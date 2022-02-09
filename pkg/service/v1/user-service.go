package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	AlertV1 "github.com/bizio/wa-srv-alert/pkg/api/v1"
	"github.com/bizio/wa-srv-base/cache"
	v1 "github.com/bizio/user-service/pkg/api/v1"
	"github.com/bizio/user-service/pkg/service/v1/cloudpubsub"
	"github.com/bizio/user-service/pkg/service/v1/data"
	"github.com/bizio/user-service/pkg/service/v1/utils"
	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/nyaruka/phonenumbers"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ErrInvalidToken     = "Invalid token"
	ErrInvalidFirstName = "Invalid first name"
	ErrInvalidLastName  = "Invalid last name"
	ErrInvalidEmail     = "Invalid email"
	ErrInvalidPassword  = "Invalid password"
	ErrPasswordTooShort = "Password is too short"
	ErrInvalidTelephone = "Invalid telephone"
	ErrInvalidMobile    = "Invalid mobile"
	ErrInvalidUserId    = "Invalid user id"
	ErrInvalidRole      = "Role not defined or invalid"
	ErrUserExists       = "A user with the same email address already exists"
	ErrUserNotFound     = "User not found"

	ErrUserNotificationNotFound      = "User notification not found"
	ErrInvalidUserNotificationRefId  = "Invalid message reference ID"
	ErrInvalidUserNotificationStatus = "Invalid message notification status"

	ErrInvalidTenantName = "Invalid name"
	ErrInvalidTenantCode = "Invalid code"
	ErrTenantExists      = "A tenant with the same email or code already exists"

	ErrPermissionDenied   = "Request not authorized"
	ErrInvalidCredentials = "Invalid username or password"
	ErrInvalidTenantId    = "Invalid tenant id"

	ErrInvalidPaymentAmount = "Invalid payment amount"
	ErrInvalidPaymentRef    = "Invalid payment reference"

	UsersNotificationsReportCacheKey = "users_notifications_report"
	PaymentsListCacheKey             = "tenant_payments"

	DateLayoutISO = "2006-01-02"
	DateLayoutIT  = "January 2, 2006"

	MinPasswordLength = 8
)

type UserServiceServer struct {
	datastore data.Datastore
	queue     cloudpubsub.Queue
	cache     cache.CacheService
	logger    *log.Logger
}

func (u *UserServiceServer) Create(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {

	var user, newUser *v1.User
	var err error

	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[Create::Error] %s", err)
		return nil, err
	}

	user = &v1.User{
		FirstName:     req.GetFirstName(),
		LastName:      req.GetLastName(),
		Email:         req.GetEmail(),
		Password:      req.GetPassword(),
		Mobile:        req.GetMobile(),
		Notifications: req.GetNotifications(),
	}
	newUser, err = u.validateUser(user, true)
	if err != nil {
		u.logger.Printf("[Create::Error] %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	switch v1.Role(authUser.Role) {
	case v1.Role_ADMIN:
		// admin user can send tenantId and role in the request payload
		if req.GetTenantId() > 0 {
			newUser.TenantId = req.GetTenantId()
		} else {
			newUser.TenantId = authUser.TenantId
		}
		newUser.Role = req.GetRole()
		break
	case v1.Role_TENANT_ADMIN:
		// tenant admins can only create users within the same tenant
		// @todo maybe we can just ignore the tenandId in the request
		if req.GetTenantId() > 0 && authUser.TenantId != req.GetTenantId() {
			u.logger.Printf("[Create::Error] %s", ErrPermissionDenied)
			return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
		}

		newUser.TenantId = authUser.TenantId
		newUser.Role = v1.Role_USER
	case v1.Role_USER:
		// users cannot create other users
		u.logger.Printf("[Create::Error] %s", ErrPermissionDenied)
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}

	u.logger.Printf("[Create::Debug] %v", newUser)
	if newUser.GetTenantId() == 0 {
		u.logger.Printf("[Create::Error] empty tenant id %v", newUser)
		return nil, status.Error(codes.Internal, ErrInvalidTenantId)
	}

	// check if user exists
	_, err = u.datastore.GetByEmailAndTenantId(newUser.Email, newUser.GetTenantId())
	if err != nil {
		if err.Error() == data.ErrNotFound {
			// do nothing, continue execution
		} else {
			u.logger.Printf("[Create::Error] %s", err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	} else {
		// if user exists return error
		u.logger.Printf("[Create::Error] %s", ErrUserExists)
		return nil, status.Error(codes.AlreadyExists, ErrUserExists)

	}

	// set created at date
	newUser.CreatedAt = ptypes.TimestampNow()
	newUser.UpdatedAt = ptypes.TimestampNow()
	err = u.datastore.Save(newUser)
	if err != nil {
		u.logger.Printf("[Create::Error] %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	// regenerate cache
	_, err = u.generateReport(newUser.TenantId)
	if err != nil {
		u.logger.Printf("[Create::Error] %s", err)
	}

	return &v1.CreateUserResponse{Id: newUser.Id}, nil

}

func (u *UserServiceServer) Update(ctx context.Context, req *v1.UpdateUserRequest) (*v1.GetUserResponse, error) {

	var existingUser, currentUser, updatedUser *v1.User
	var err error

	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[Update::Error] %s", err)
		return nil, err
	}

	currentUser, err = u.datastore.GetById(req.GetUser().GetId())
	if err != nil {
		u.logger.Printf("[Update::Error] %s", err)
		return nil, err
	}

	updateUserRequest := req.GetUser()
	// try to update user with request data
	proto.Merge(currentUser, updateUserRequest)

	// validate request data
	passwordChange := len(req.GetUser().GetPassword()) > 0
	updatedUser, err = u.validateUser(currentUser, passwordChange)
	if err != nil {
		u.logger.Printf("[Update::Error] %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// try to update user with validated data
	proto.Merge(currentUser, updatedUser)

	// if notifications have changed we just replace the map
	// and remove duplicates
	if len(updateUserRequest.Notifications) > 0 {
		currentUser.Notifications = updateUserRequest.Notifications

		// remove duplicates from notifications after the merge
		var notifications []*v1.NotificationSettings
		var exists map[string]bool = make(map[string]bool)
		for _, notification := range currentUser.GetNotifications() {
			stringType := notification.GetType().String()
			ok := exists[stringType]
			if !ok {
				exists[stringType] = true
				notifications = append(notifications, notification)
			}
		}
		currentUser.Notifications = notifications
	}

	// check if email has changed and if there's a duplicate
	if len(updateUserRequest.GetEmail()) > 0 {
		existingUser, err = u.datastore.GetByEmail(updateUserRequest.GetEmail())
		if err != nil {
			if err.Error() == data.ErrNotFound {
				// do nothing, continue execution
			} else {
				u.logger.Printf("[Update::Error] %s", err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		} else if existingUser.GetId() != currentUser.GetId() {
			// if user exists return error
			u.logger.Printf("[Update::Error] %s", ErrUserExists)
			return nil, status.Error(codes.AlreadyExists, ErrUserExists)

		}
	}

	switch v1.Role(authUser.Role) {
	case v1.Role_ADMIN:
		// if tenantId and role have been passed in the request
		// they have been already merged
		break
	case v1.Role_TENANT_ADMIN:
		// tenant admins can only update users within the same tenant
		// @todo maybe we can just ignore the tenandId in the request
		if updateUserRequest.GetTenantId() > 0 && authUser.GetTenantId() != updateUserRequest.GetTenantId() {
			u.logger.Printf("[Update::Error] %s", ErrPermissionDenied)
			return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
		}

		currentUser.TenantId = authUser.TenantId
	case v1.Role_USER:
		// users can only update their own profile
		if authUser.GetId() != currentUser.GetId() {
			u.logger.Printf("[Update::Error] %s", ErrPermissionDenied)
			return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
		}

		// in case a user tries to update his/her own tenantId or role
		currentUser.TenantId = authUser.TenantId
		currentUser.Role = v1.Role_USER
	}

	// set updated at date
	currentUser.UpdatedAt = ptypes.TimestampNow()

	err = u.datastore.Save(currentUser)
	if err != nil {
		u.logger.Printf("[Update::Error] %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return u.userResponse(currentUser), nil

}
func (u *UserServiceServer) Get(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserResponse, error) {

	var user *v1.User
	var err error

	// authorized user
	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[Get::Error] %s", err)
		return nil, err
	}

	switch v1.Role(authUser.Role) {
	case v1.Role_ADMIN:
		// admin can get users from all tenants
		break
	case v1.Role_TENANT_ADMIN:
		// tenant admins can only get users within the same tenant
		if req.GetTenantId() > 0 && authUser.TenantId != req.GetTenantId() {
			u.logger.Printf("[Get::Error] %s", ErrPermissionDenied)
			return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
		}
		break
	case v1.Role_USER:
		// users can only get their profile
		if req.GetId() != authUser.GetId() {
			u.logger.Printf("[Get::Error] %s", ErrPermissionDenied)
			return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
		}
	}

	user, err = u.datastore.GetById(req.GetId())
	// check if the datastore returned any error
	if err != nil {
		u.logger.Printf("[Get::Error] %s", err)
		return nil, err
	}

	return u.userResponse(user), nil
}

func (u *UserServiceServer) ListUsers(ctx context.Context, req *v1.ListUsersRequest) (*v1.ListUsersResponse, error) {

	var response *v1.ListUsersResponse
	var list []*v1.User
	var err error

	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[ListUsers::Error] %s", err)
		return nil, err
	}

	response = &v1.ListUsersResponse{}
	params := &data.ListUsersParams{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Limit:     int(req.GetLimit()),
		Offset:    int(req.GetOffset()),
	}

	switch v1.Role(authUser.Role) {
	case v1.Role_ADMIN:
		// admins can list all users
		params.TenantId = req.GetTenantId()
		params.Role = int(req.GetRole())
	case v1.Role_TENANT_ADMIN:
		// tenant admins can only list users within the same tenant
		params.TenantId = authUser.TenantId
		params.Role = int(v1.Role_USER)
	case v1.Role_USER:
		// users cannot list users
		u.logger.Printf("[ListUsers::Error] %s", ErrPermissionDenied)
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}

	list, err = u.datastore.ListUsers(params)
	if err != nil {
		u.logger.Printf("[ListUsers::Error] %s", err)
		return nil, err
	}

	response.Users = list

	return response, nil
}

func (u *UserServiceServer) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (*v1.ResetPasswordResponse, error) {

	if len(req.GetEmail()) == 0 {
		return nil, status.Error(codes.InvalidArgument, ErrInvalidEmail)
	}

	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[ResetPassword::Error] %s", err)
		return nil, err
	}

	// we expect this to be a request from a logged out user
	// so the call should be internal
	if v1.Role(authUser.Role) != v1.Role_TENANT_ADMIN {
		u.logger.Printf("[ResetPassword::Error] %s", ErrPermissionDenied)
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}

	// lookup user
	user, err := u.datastore.GetByEmailAndTenantId(req.GetEmail(), authUser.TenantId)
	if err != nil {
		u.logger.Printf("[ResetPassword::Error] %s", err)
		if err.Error() == data.ErrNotFound {
			return nil, status.Error(codes.NotFound, ErrUserNotFound)
		} else {
			return nil, err
		}
	}

	// 30 minutes expiration
	minutes, _ := time.ParseDuration("30m")
	token := utils.GenerateToken(time.Now().Add(minutes))

	// update user with password reset token
	user.PasswordResetToken = token

	err = u.datastore.Save(user)
	if err != nil {
		u.logger.Printf("[ResetPassword::Error] %s", err)
		return nil, err
	}

	tenant, err := u.datastore.GetTenantById(user.TenantId)
	if err != nil {
		u.logger.Printf("[ResetPassword::Error] %s", err)
		return nil, err
	}

	// send notification
	userNotification := &v1.UserNotification{
		UserId:      user.GetId(),
		Email:       user.GetEmail(),
		LongMessage: token,
		Type:        v1.NotificationType_EMAIL,
		TenantName:  tenant.Name,
		TenantCode:  tenant.Code,
		CreatedAt:   ptypes.TimestampNow(),
	}

	err = u.CreateUserNotification(ctx, userNotification)
	if err != nil {
		u.logger.Printf("[ResetPassword::Error] %s", err)
		return nil, err
	}

	return &v1.ResetPasswordResponse{Token: token}, nil
}

// we are not checking any authorization context here, shoud we?
func (u *UserServiceServer) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordRequest) (*v1.UpdatePasswordResponse, error) {

	var err error

	if len(req.GetToken()) == 0 {
		u.logger.Printf("[UpdatePassword::Error] %s", ErrInvalidToken)
		return nil, status.Error(codes.InvalidArgument, ErrInvalidToken)
	}

	// validate token
	err = utils.ValidateToken(req.GetToken())
	if err != nil {
		u.logger.Printf("[UpdatePassword::Error] %s", err)
		return nil, err
	}

	// lookup user by password reset token
	user, err := u.datastore.GetByPasswordResetToken(req.GetToken())
	if err != nil {
		u.logger.Printf("[UpdatePassword::Error] %s", err)
		return nil, err
	}

	// check password length and make sure is different from previous one
	newPassword := req.GetPassword()

	if len(newPassword) < MinPasswordLength {
		u.logger.Printf("[UpdatePassword::Error] %s", errors.New(ErrPasswordTooShort))
		return nil, status.Error(codes.InvalidArgument, ErrPasswordTooShort)
	}

	encryptedPassword, err := u.encryptPassword(newPassword)
	if err != nil {
		u.logger.Printf("[UpdatePassword::Error] %s", err)
		return nil, err
	}

	// trying to reset the same password
	if user.Password == encryptedPassword {
		u.logger.Printf("[UpdatePassword::Error] trying to reset password to previous one %s", errors.New(ErrPasswordTooShort))
		return nil, status.Error(codes.InvalidArgument, ErrInvalidPassword)
	}

	// update user with new password and clear password reset token
	user.Password = encryptedPassword
	user.PasswordResetToken = ""

	err = u.datastore.Save(user)
	if err != nil {
		u.logger.Printf("[UpdatePassword::Error] %s", err)
		return nil, err
	}

	return &v1.UpdatePasswordResponse{}, nil
}

// Gets the cached report for user notifications
func (u *UserServiceServer) GetReport(ctx context.Context, req *v1.GetReportRequest) (*v1.Report, error) {

	var err error
	var report *v1.Report
	var tenantId int64

	authUser, err := u.CheckAuthorization(ctx, &empty.Empty{})
	if err != nil {
		u.logger.Printf("[ListUsers::Error] %s", err)
		return nil, err
	}
	switch v1.Role(authUser.Role) {
	case v1.Role_ADMIN:
		tenantId = req.GetTenantId()
	case v1.Role_TENANT_ADMIN:
		tenantId = authUser.TenantId
	case v1.Role_USER:
		u.logger.Printf("[GetReport::Error] %s", ErrPermissionDenied)
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}

	cacheKey := fmt.Sprintf("%s_%d", UsersNotificationsReportCacheKey, tenantId)
	cachedReport, cacheErr := u.cache.Get(cacheKey)
	if cacheErr != nil {
		report, err = u.generateReport(tenantId)
		if err != nil {
			u.logger.Printf("[GetReport::Error] %s", err)
			return nil, err
		}

		if cacheErr == cache.ErrNotFound {
			bytes, err := json.Marshal(report)
			if err != nil {
				u.logger.Printf("[GetReport::Error] %s", err)
				return nil, err
			}

			err = u.cache.Set(cacheKey, bytes, 0)
			if err != nil {
				u.logger.Printf("[GetReport::Error] %s", err)
			}
		}

	} else {

		err = json.Unmarshal([]byte(cachedReport), &report)
		if err != nil {
			u.logger.Printf("[GetReport::Error] %s", err)
			return nil, err
		}
	}

	return report, nil
}

func (u *UserServiceServer) GetByEmailAndPassword(ctx context.Context, req *v1.GetUserByEmailAndPasswordRequest) (*v1.GetUserResponse, error) {
	var user *v1.User
	var err error

	user, err = u.datastore.GetByEmailAndTenantId(req.GetEmail(), req.GetTenantId())
	if err != nil {
		u.logger.Printf("[GetByEmail::Error] %s", err)
		if err.Error() == data.ErrNotFound {
			return nil, status.Error(codes.NotFound, ErrUserNotFound)
		}
		return nil, err
	}

	byteHash := []byte(user.GetPassword())
	err = bcrypt.CompareHashAndPassword(byteHash, []byte(req.GetPassword()))
	if err != nil {
		u.logger.Printf("[GetByEmail::Error] %s", err)
		if err == bcrypt.ErrMismatchedHashAndPassword || err == bcrypt.ErrHashTooShort {
			u.logger.Printf("[GetByEmail::Error::InvalidPassword] %s", err)
			return nil, status.Error(codes.InvalidArgument, ErrInvalidCredentials)
		} else {
			return nil, err
		}

	}

	return u.userResponse(user), nil
}

func (u *UserServiceServer) ListUserNotifications(ctx context.Context, req *v1.ListUserNotificationsRequest) (*v1.ListUserNotificationsResponse, error) {

	var err error
	var email, mobile, formattedMobile string
	//var notificationStatus *v1.UserNotificationStatus
	//var notificationType *v1.NotificationType

	params := &data.ListUserNotificationsParams{
		UserId:  req.GetUserId(),
		AlertId: req.GetAlertId(),
	}

	if len(req.GetEmail()) > 0 {
		email = strings.TrimSpace(req.GetEmail())
		if len(email) > 0 {
			_, err = mail.ParseAddress(email)
			if err != nil {
				return nil, errors.New(ErrInvalidEmail)
			}
		}
		params.Email = email
	}

	if len(req.GetMobile()) > 0 {
		mobile = strings.TrimSpace(req.GetMobile())
		if len(mobile) != 0 {
			parsedMobile, err := phonenumbers.Parse(mobile, "IT")
			isValidNumber := phonenumbers.IsValidNumberForRegion(parsedMobile, "IT")
			if err != nil || isValidNumber == false {
				u.logger.Printf("Error parsing phone number: %s", err)
				return nil, errors.New(ErrInvalidMobile)
			} else {
				formattedMobile = phonenumbers.Format(parsedMobile, phonenumbers.INTERNATIONAL)
			}

			params.Mobile = formattedMobile
		}
	}

	// even if the status is not passed as part of the user request
	// this value is defaulted to 0 (PENDING) so we are not able to
	// query without passing this parameter unless we ignore it completely
	// @todo: fix NotificationStatus
	notificationStatus := req.GetStatus()
	params.Status = &notificationStatus

	if req.GetType() > 0 {
		notificationType := req.GetType()
		params.Type = &notificationType
	}

	if req.GetLimit() > 0 {
		params.Limit = int(req.GetLimit())
		if req.GetOffset() > 0 {
			params.Offset = int(req.GetOffset())
		}
	}

	list, err := u.datastore.ListUserNotifications(params)
	if err != nil {
		u.logger.Printf("[ListUserNotifications::Error] %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	var notifications []*v1.UserNotification
	for _, entity := range list {
		obj, _ := data.ConvertUserNotificationToStruct(entity)
		notifications = append(notifications, obj)
	}

	return &v1.ListUserNotificationsResponse{Total: int32(len(notifications)), Notifications: notifications}, nil

}

func (u *UserServiceServer) validateUser(user *v1.User, generateEncryptedPassword bool) (*v1.User, error) {

	var err error
	var newUser *v1.User
	var formattedMobile string

	firstName := strings.Title(strings.ToLower(strings.TrimSpace(user.GetFirstName())))
	lastName := strings.Title(strings.ToLower(strings.TrimSpace(user.GetLastName())))
	email := strings.ToLower(strings.TrimSpace(user.GetEmail()))
	password := strings.TrimSpace(user.GetPassword())
	mobile := strings.TrimSpace(user.GetMobile())

	if len(firstName) == 0 {
		return nil, errors.New(ErrInvalidFirstName)
	}

	if len(lastName) == 0 {
		return nil, errors.New(ErrInvalidLastName)
	}

	if len(email) == 0 {
		return nil, errors.New(ErrInvalidEmail)
	}

	if len(mobile) == 0 {
		return nil, errors.New(ErrInvalidMobile)
	} else {
		parsedMobile, err := phonenumbers.Parse(mobile, "IT")
		isValidNumber := phonenumbers.IsValidNumberForRegion(parsedMobile, "IT")
		if err != nil || isValidNumber == false {
			u.logger.Printf("Error parsing phone number: %s", err)
			return nil, errors.New(ErrInvalidMobile)
		} else {
			formattedMobile = phonenumbers.Format(parsedMobile, phonenumbers.INTERNATIONAL)
		}
	}

	_, err = mail.ParseAddress(email)
	if err != nil {
		return nil, errors.New(ErrInvalidEmail)
	}

	if len(password) == 0 {
		return nil, errors.New(ErrInvalidPassword)
	}

	if len(password) < MinPasswordLength {
		return nil, errors.New(ErrPasswordTooShort)
	}

	newUser = &v1.User{
		FirstName:     strings.Title(firstName),
		LastName:      strings.Title(lastName),
		Email:         strings.ToLower(email),
		Password:      password,
		Mobile:        formattedMobile,
		Notifications: user.GetNotifications(),
	}

	// we only encrypt passwords for new users and leave unchanged for updates
	if generateEncryptedPassword {
		encryptedPassword, err := u.encryptPassword(password)
		if err != nil {
			u.logger.Printf("[validateUser::Error] %s", err)
			return nil, err
		}

		newUser.Password = encryptedPassword
	}

	return newUser, nil
}

func (u *UserServiceServer) CheckAuthorization(ctx context.Context, req *empty.Empty) (*v1.User, error) {

	var user *v1.User
	var userId int64
	var err error

	// check permissions
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		u.logger.Printf("[CheckAuthorization::Error] Missing authorization credentials from context")
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}

	// validate user
	userFromContext, ok := md["user"]
	if !ok {
		u.logger.Printf("[CheckAuthorization::Error] %s", ErrInvalidUserId)
		return nil, status.Error(codes.PermissionDenied, ErrInvalidUserId)
	}

	userIdFromContext := userFromContext[0]
	if len(userIdFromContext) == 0 {
		u.logger.Printf("[CheckAuthorization::Error] %s", ErrInvalidUserId)
		return nil, status.Error(codes.PermissionDenied, ErrInvalidUserId)
	}

	userId, err = strconv.ParseInt(userIdFromContext, 10, 64)
	if err != nil {
		u.logger.Printf("[CheckAuthorization::Error] %s", ErrInvalidUserId)
		return nil, status.Error(codes.PermissionDenied, ErrInvalidUserId)
	}

	// get the information of the user that initiated the request
	user, err = u.datastore.GetById(userId)
	if err != nil {
		u.logger.Printf("[CheckAuthorization::Error] %s", err)
		return nil, status.Error(codes.PermissionDenied, ErrPermissionDenied)
	}
	user = &v1.User{
		Email:    tenant.Email,
		TenantId: tenant.Id,
		Role:     v1.Role_ADMIN,
	}

	return user, nil
}

// Counts users by type and caches the result
func (u *UserServiceServer) generateReport(tenantId int64) (*v1.Report, error) {

	// report includes a map of all notification types groupped by status
	// and total cost
	var notificationsReport map[string]*v1.NotificationsReport = make(map[string]*v1.NotificationsReport)
	var report *v1.Report = &v1.Report{
		Total: 0,
		Cost:  0.00,
		Data:  notificationsReport,
	}

	cacheKey := fmt.Sprintf("%s_%d", UsersNotificationsReportCacheKey, tenantId)

	params := &data.ListUsersParams{TenantId: tenantId}

	// fetch all users for the given tenant
	users, err := u.datastore.ListUsers(params)
	if err != nil {
		u.logger.Printf("[generateReport::Error] %s", err)
		return nil, err
	}

	// loop through users list and get notifications
	for _, user := range users {

		userNotifications, err := u.datastore.ListUserNotifications(&data.ListUserNotificationsParams{UserId: user.GetId()})
		if err != nil {
			u.logger.Printf("[generateReport::Error] %s", err)
			return nil, err
		}

		report.Total += int32(len(userNotifications))

		// loop through the current user's notifications and
		// updates totals by notification type
		for _, notification := range userNotifications {
			notificationType := strings.ToLower(notification.Type.String())

			// sets or adds to the notification type total
			notificationsReport[notificationType] = &v1.NotificationsReport{
				Sent:    0,
				Failed:  0,
				Pending: 0,
				Cost:    0.00,
			}

			notificationsReport[notificationType].Cost += notification.Price
			report.Cost += notification.Price

			switch notification.Status {
			case v1.UserNotificationStatus_PENDING:
				notificationsReport[notificationType].Pending += 1
			case v1.UserNotificationStatus_SENT:
				notificationsReport[notificationType].Sent += 1
			case v1.UserNotificationStatus_FAILED:
				notificationsReport[notificationType].Failed += 1
			default:
				err = errors.New("Invalid status")
				u.logger.Printf("[generateReport::Error] %s", err)
				return nil, err

			}
		}

	}

	report.Data = notificationsReport

	bytes, err := json.Marshal(report)
	if err != nil {
		u.logger.Printf("[generateReport::Error] %s", err)
		return nil, err
	}

	err = u.cache.Set(cacheKey, bytes, 0)
	if err != nil {
		u.logger.Printf("[generateReport::Error] %s", err)
	}

	return report, nil

}
func (u *UserServiceServer) CheckAlerts() {

	messages := make(chan *pubsub.Message)
	go func() {
		err := u.queue.GetMessages(messages)
		if err != nil {
			u.logger.Printf("Error checking new alerts from queue: %s", err)
			return
		}
	}()

	for {

		select {
		case message := <-messages:

			ctx := context.Background()
			alert := AlertV1.Alert{}

			err := proto.Unmarshal(message.Data, &alert)
			if err != nil {
				u.logger.Printf("[CheckAlerts::Error] Error unmarshaling message: %s", err)
				continue
			}

			md := map[string]string{
				"user":   "api",
				"tenant": strconv.FormatInt(alert.GetTenantId(), 10),
			}
			ctx = metadata.NewIncomingContext(ctx, metadata.New(md))
			usersList, err := u.ListUsers(ctx, &v1.ListUsersRequest{TenantId: alert.GetTenantId(), Role: v1.Role_USER})
			if err != nil {
				u.logger.Printf("[CheckAlerts::Error] Error getting users list: %s", err)
				continue
			}

			u.logger.Printf("[CheckAlerts::Debug] found %d users with tenantId %d", len(usersList.GetUsers()), alert.GetTenantId())

			for _, user := range usersList.GetUsers() {
				userNotificationPreferences := user.GetNotifications()
				if len(userNotificationPreferences) == 0 {
					u.logger.Printf("[CheckAlerts] user %d has not set any notification preferences", user.GetId())
					continue
				}

				tenant, err := u.GetTenant(ctx, &v1.GetTenantRequest{Id: user.GetTenantId()})
				if err != nil {
					u.logger.Printf("[CheckAlerts::Error] Error getting tenant: %s", err)
					continue
				}

				for _, userNotificationPreference := range userNotificationPreferences {

					userNotification := &v1.UserNotification{
						UserId:       user.GetId(),
						Email:        user.GetEmail(),
						Mobile:       user.GetMobile(),
						AlertId:      alert.GetId(),
						ShortMessage: alert.GetShortMessage(),
						LongMessage:  alert.GetLongMessage(),
						Type:         userNotificationPreference.GetType(),
						TenantName:   tenant.GetName(),
						TenantCode:   tenant.GetCode(),
						CreatedAt:    ptypes.TimestampNow(),
					}

					err := u.CreateUserNotification(ctx, userNotification)
					if err != nil {
						u.logger.Printf("[CheckAlerts::Error] %s", err)
						u.logger.Printf("[CheckAlerts::debug] %v", userNotification)
						continue
					}
				}

			}
			message.Ack()
		}
	}
}

func (u *UserServiceServer) UpdateUserNotification(ctx context.Context, req *v1.UserNotificationUpdateRequest) (*v1.UserNotification, error) {

	u.logger.Printf("[UpdateUserNotification::req] %v", req)
	var entity *data.UserNotificationEntity
	var refId, reason string
	var notificationStatus v1.UserNotificationStatus
	var price float32
	var id int64
	var err error
	var creditUpdateRequired bool

	id = req.GetId()
	refId = strings.TrimSpace(req.GetRefId())
	reason = strings.TrimSpace(req.GetReason())
	notificationStatus = req.GetStatus()
	price = req.GetPrice()

	// webhook and callback functions only have a refId
	if id > 0 {
		u.logger.Printf("[UpdateUserNotification::GetUserNotificationById] %d", id)
		entity, err = u.datastore.GetUserNotificationById(id)
	} else {
		if len(refId) == 0 {
			u.logger.Printf("[UpdateUserNotification::Error] %s", ErrInvalidUserNotificationRefId)
			return nil, status.Error(codes.InvalidArgument, ErrInvalidUserNotificationRefId)
		}
		u.logger.Printf("[UpdateUserNotification::GetUserNotificationByRefId] %s", refId)
		entity, err = u.datastore.GetUserNotificationByRefId(refId)
	}

	if err != nil {
		u.logger.Printf("[UpdateUserNotification::Error] %s", err)
		if err.Error() == data.ErrNotFound {
			return nil, status.Error(codes.NotFound, ErrUserNotificationNotFound)
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	if id > 0 {
		entity.RefId = refId
	}

	if notificationStatus != entity.Status {
		entity.Status = notificationStatus
		if notificationStatus == v1.UserNotificationStatus_SENT {
			sentAt, err := ptypes.Timestamp(req.GetSentAt())
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			entity.SentAt = sentAt
		}
	}

	if len(reason) > 0 {
		entity.Reason = reason
	}

	if price > 0 {
		entity.Price = price
		// we need to update tenant's credit
		creditUpdateRequired = true
	}

	err = u.datastore.UpdateUserNotification(entity.Id, entity)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userNotification, err := data.ConvertUserNotificationToStruct(entity)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	u.logger.Printf("[UpdateUserNotification::debug] entity price: %v, struct price: %v", entity.Price, userNotification.GetPrice())
	if creditUpdateRequired {
		u.logger.Printf("[UpdateUserNotification::debug] credit update required for %s", entity.TenantCode)
		tenant, err := u.datastore.GetTenantByCode(entity.TenantCode)
		if err != nil {
			u.logger.Printf("[UpdateUserNotification::Error] %s", err)
		} else {
			tenant.Credit -= userNotification.GetPrice()
			err = u.datastore.SaveTenant(tenant)
			if err != nil {
				u.logger.Printf("[UpdateUserNotification:::Error] %s", err)
			}
		}
	}

	return userNotification, nil

}

/**
 * Search for pending notifications and pushes them to the sending queue
 */
func (u *UserServiceServer) QueueUserNotifications(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {

	res := &empty.Empty{}

	list, err := u.datastore.ListPendingUserNotifications()
	if err != nil {
		u.logger.Printf("[QueueUserNotification::Error] %s", err)
		return res, err
	}

	if list == nil {
		u.logger.Print("[QueueUserNotifications] List empty, nothing to do\n")
		return res, nil
	}

	for _, pendingMessage := range list {

		userNotification, _ := data.ConvertUserNotificationToStruct(pendingMessage)
		bytes, err := proto.Marshal(userNotification)
		if err != nil {
			u.logger.Printf("[QueueUserNotifications::Error] %s", err)
			return res, err
		}

		// publish user_notification message to queue
		writeTopic, err := u.queue.GetTopic(cloudpubsub.WriteTopic)
		if err != nil {
			u.logger.Printf("[QueueUserNotifications::Error] %s", err)
			return res, err
		}

		publishResult := writeTopic.Publish(ctx, &pubsub.Message{Data: bytes})
		if err != nil {
			u.logger.Printf("[QueueUserNotifications::Error] %s", err)
			return res, err
		} else {
			_, err = publishResult.Get(ctx)
			if err != nil {
				u.logger.Printf("[QueueUserNotifications::Error] %s", err)
				return res, err
			}
		}
	}

	return res, nil

}

func (u *UserServiceServer) CreateUserNotification(ctx context.Context, userNotification *v1.UserNotification) error {

	entity, err := data.ConvertUserNotificationToEntity(userNotification)
	if err != nil {
		return err
	}

	err = u.datastore.CreateUserNotification(entity)
	if err != nil {
		u.logger.Printf("[CreateUserNotification::Error] %s", err)
		return err
	}

	userNotification, _ = data.ConvertUserNotificationToStruct(entity)
	bytes, err := proto.Marshal(userNotification)
	if err != nil {
		u.logger.Printf("[CreateUserNotification::Error] %s", err)
		return err
	}

	// publish user_notification message to queue
	writeTopic, err := u.queue.GetTopic(cloudpubsub.WriteTopic)
	if err != nil {
		u.logger.Printf("[CreateUserNotification::Error] %s", err)
		return err
	}

	publishResult := writeTopic.Publish(ctx, &pubsub.Message{Data: bytes})
	if err != nil {
		u.logger.Printf("[CreateUserNotification::Error] %s", err)
		return err
	} else {
		_, err = publishResult.Get(ctx)
		if err != nil {
			u.logger.Printf("[CreateUserNotification::Error] %s", err)
			return err
		}
	}

	return nil

}

// removes fields not required in the response like passwords
func (u *UserServiceServer) userResponse(user *v1.User) *v1.GetUserResponse {

	var res *v1.GetUserResponse

	tenant, err := u.datastore.GetTenantById(user.TenantId)
	if err != nil {
		u.logger.Printf("[userResponse::Error] %s", err)
		return nil
	}

	res = &v1.GetUserResponse{
		Id:            user.GetId(),
		FirstName:     user.GetFirstName(),
		LastName:      user.GetLastName(),
		Email:         user.GetEmail(),
		TenantId:      user.GetTenantId(),
		Tenant:        data.ConvertTenantEntityToStruct(tenant),
		Role:          user.GetRole(),
		Mobile:        user.GetMobile(),
		Notifications: user.GetNotifications(),
		CreatedAt:     user.GetCreatedAt(),
		UpdatedAt:     user.GetUpdatedAt(),
	}

	return res
}

func (u *UserServiceServer) encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Printf("[encryptPassword::Error] %s", err)
		return "", err
	}

	return string(hash), nil
}

func NewUserServiceServer(datastoreService data.Datastore, queue cloudpubsub.Queue, cache cache.CacheService) *UserServiceServer {
	return &UserServiceServer{
		datastoreService,
		queue,
		cache,
		log.New(os.Stderr, "[UserServiceServer] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
