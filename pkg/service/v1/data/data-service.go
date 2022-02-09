package data

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/bizio/wa-srv-base/cache"
	v1 "github.com/bizio/user-service/pkg/api/v1"
)

const (
	KIND                     = "User"
	TENANT                   = "Tenant"
	USER_NOTIFICATION        = "UserNotification"
	PAYMENT                  = "Payment"
	ErrNotFound              = "Not found"
	ErrInvalidEntityType     = "Invalid entity type"
	ErrContextCancelled      = "Context cancelled"
	UserCacheKey             = "user"
	UserNotificationCacheKey = "user_notification"
)

type Datastore interface {
	Save(user *v1.User) error
	GetById(id int64) (*v1.User, error)
	GetByIdAndTenantId(id, tenantId int64) (*v1.User, error)
	GetByEmail(email string) (*v1.User, error)
	GetByEmailAndTenantId(email string, tenantId int64) (*v1.User, error)
	GetByPasswordResetToken(token string) (*v1.User, error)
	ListUsers(params *ListUsersParams) ([]*v1.User, error)

	CreateUserNotification(userNotification *UserNotificationEntity) error
	UpdateUserNotification(id int64, entity *UserNotificationEntity) error
	ListUserNotifications(params *ListUserNotificationsParams) ([]*UserNotificationEntity, error)
	ListPendingUserNotifications() ([]*UserNotificationEntity, error)
	GetUserNotificationById(id int64) (*UserNotificationEntity, error)
	GetUserNotificationByRefId(refId string) (*UserNotificationEntity, error)

	GetTenantById(id int64) (*TenantEntity, error)
	GetTenantByCode(code string) (*TenantEntity, error)
	GetTenantByEmail(email string) (*TenantEntity, error)
	ListTenants(params *ListTenantsParams) ([]*TenantEntity, error)
	SaveTenant(tenant *TenantEntity) error

	GetPaymentById(paymentId int64) (*v1.Payment, error)
	SavePayment(payment *v1.Payment) error
	ListPayments(tenantId int64) ([]*v1.Payment, error)
}

type NotificationEntity struct {
	Type     v1.NotificationType
	Priority int32
}

type UserEntity struct {
	Id                 int64
	FirstName          string
	LastName           string
	Email              string
	Password           string
	TenantId           int64
	Role               v1.Role
	Mobile             string
	Notifications      []*NotificationEntity
	PasswordResetToken string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type UserNotificationEntity struct {
	Id           int64
	UserId       int64
	Email        string
	Mobile       string
	AlertId      int64
	ShortMessage string
	LongMessage  string
	Type         v1.NotificationType
	CreatedAt    time.Time
	SentAt       time.Time
	RefId        string
	Status       v1.UserNotificationStatus
	Reason       string
	Price        float32
	TenantName   string
	TenantCode   string
}

type ListUserNotificationsParams struct {
	UserId         int64
	Email          string
	Mobile         string
	AlertId        int64
	CreatedAtStart *time.Time
	CreatedAtEnd   *time.Time
	SentAtStart    *time.Time
	SentAtEnd      *time.Time
	Status         *v1.UserNotificationStatus
	Type           *v1.NotificationType
	Limit          int
	Offset         int
}

type ListUsersParams struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	TenantId  int64
	Role      int
	Limit     int
	Offset    int
}

type PaymentEntity struct {
	Id          int64
	TenantId    int64
	Ref         string
	Amount      float32
	PaymentDate time.Time
	PaidBy      int64
	CreatedAt   time.Time
	CreatedBy   int64
}

type TenantEntity struct {
	Id        int64
	Name      string
	Code      string
	Email     string
	Credit    float32
	LogoUrl   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ListTenantsParams struct {
	Name   string
	Email  string
	Code   string
	Limit  int
	Offset int
}

type DatastoreService struct {
	client *datastore.Client
	cache  cache.CacheService
	logger *log.Logger
}

func NewDatastoreService(projectId string, cache cache.CacheService) Datastore {

	logger := log.New(os.Stderr, "[DatastoreService] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Printf("Connecting datastore for project id %s", projectId)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := datastore.NewClient(ctx, projectId)
	if err != nil {
		logger.Fatalf("Cannot create new datastore client %s", err)
		return nil
	}

	// Verify that we can communicate and authenticate with the datastore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		logger.Fatalf("datastoredb: could not connect: %v", err)
		return nil
	}
	if err := t.Rollback(); err != nil {
		logger.Fatalf("datastoredb: could not connect: %v", err)
		return nil
	}

	return &DatastoreService{client, cache, logger}

}

func (d *DatastoreService) Save(newUser *v1.User) error {

	var k *datastore.Key

	// new user
	if newUser.GetId() == 0 {
		k = datastore.IncompleteKey(KIND, nil)
		d.logger.Printf("Creating key for new user: %v", k)
	} else {
		// update user
		k = datastore.IDKey(KIND, newUser.GetId(), nil)
		d.logger.Printf("User key: %v", k)
	}
	entity, err := ConvertToEntity(newUser)
	if err != nil {
		d.logger.Printf("Cannot save user: %s", err)
		return err
	}

	newKey, err := d.client.Put(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("Cannot save user: %s", err)
		return err
	}

	// if we are creating a new user we have to update Id field
	if entity.Id == 0 {
		entity.Id = newKey.ID
		newUser.Id = newKey.ID
		_, err = d.client.Put(context.Background(), newKey, entity)
		if err != nil {
			d.logger.Printf("Cannot update new user id: %s", err)
			return err
		}
	}

	d.logger.Printf("User successfully saved with id %d", entity.Id)
	return nil
}

func (d *DatastoreService) GetById(id int64) (*v1.User, error) {

	var err error
	var user *v1.User

	entity := &UserEntity{}
	k := datastore.IDKey(KIND, id, nil)
	err = d.client.Get(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("[GetById::Error] %s", err)
		return nil, handleError(err)
	}

	user, err = ConvertToStruct(entity)
	if err != nil {
		return nil, err
	}
	/*
		cacheKey := fmt.Sprintf("%s_%d", UserCacheKey, id)

		cachedUser, cacheErr := d.cache.Get(cacheKey)
		if cacheErr != nil {

			err := d.client.Get(context.Background(), k, entity)
			if err != nil {
				d.logger.Printf("[GetById::Error] %s", err)
				return nil, handleError(err)
			}

			user, err = ConvertToStruct(entity)
			if err != nil {
				return nil, err
			}

			if cacheErr == cache.ErrNotFound {

				bytes, err := proto.Marshal(user)
				if err != nil {
					d.logger.Printf("[GetById::Error] %s", err)
					return nil, err
				}

				err = d.cache.Set(cacheKey, bytes, 0)
				if err != nil {
					d.logger.Printf("[GetById::Cache::Error] %s", err)
				}

			} else {
				d.logger.Printf("[GetById::Error] %s", cacheErr)
			}

		} else {

			user = &v1.User{}
			err = proto.Unmarshal([]byte(cachedUser), user)
			if err != nil {
				d.logger.Printf("[GetById::Error] %s", err)
				return nil, status.Error(codes.Internal, err.Error())
			}

		}
	*/
	return user, nil
}

func (d *DatastoreService) GetByIdAndTenantId(id, tenantId int64) (*v1.User, error) {
	var res []*UserEntity
	var err error

	query := datastore.NewQuery(KIND).
		Filter("Id = ", id).
		Filter("TenantId = ", tenantId).
		Limit(1)

	_, err = d.client.GetAll(context.Background(), query, &res)
	if err != nil {
		d.logger.Printf("GetByIdAndTenantId::Error] %s", err)
		return nil, err
	}

	return ConvertToStruct(res[0])
}

func (d *DatastoreService) GetByEmail(email string) (*v1.User, error) {

	var users []*UserEntity
	var err error

	query := datastore.NewQuery(KIND).
		Filter("Email = ", email).
		Limit(1)

	_, err = d.client.GetAll(context.Background(), query, &users)
	if err != nil {
		d.logger.Printf("GetByEmail::Error] %s", err)
		return nil, handleError(err)
	}

	if len(users) == 0 {
		d.logger.Printf("GetByEmail::Error] %s", ErrNotFound)
		return nil, errors.New(ErrNotFound)
	}

	return ConvertToStruct(users[0])
}

func (d *DatastoreService) GetByEmailAndTenantId(email string, tenantId int64) (*v1.User, error) {

	var users []*UserEntity
	var err error

	query := datastore.NewQuery(KIND).
		Filter("Email = ", email).
		Filter("TenantId = ", tenantId).
		Limit(1)

	_, err = d.client.GetAll(context.Background(), query, &users)
	if err != nil {
		d.logger.Printf("[GetByEmailAndTenantId::Error] %s", err)
		return nil, handleError(err)
	}

	if len(users) == 0 {
		d.logger.Printf("[GetByEmailAndTenantId::Error] %s %d %s", email, tenantId, ErrNotFound)
		return nil, errors.New(ErrNotFound)
	}

	return ConvertToStruct(users[0])
}

func (d *DatastoreService) GetByPasswordResetToken(token string) (*v1.User, error) {

	var users []*UserEntity
	var err error

	query := datastore.NewQuery(KIND).
		Filter("PasswordResetToken = ", token).
		Limit(1)

	_, err = d.client.GetAll(context.Background(), query, &users)
	if err != nil {
		d.logger.Printf("[GetByPasswordResetToken::Error] %s", err)
		return nil, handleError(err)
	}

	if len(users) == 0 {
		d.logger.Printf("[GetByPasswordResetToken::Error] %s", ErrNotFound)
		return nil, errors.New(ErrNotFound)
	}

	return ConvertToStruct(users[0])
}

func (d *DatastoreService) GetTenantById(id int64) (*TenantEntity, error) {

	tenant := &TenantEntity{}
	k := datastore.IDKey(TENANT, id, nil)
	err := d.client.Get(context.Background(), k, tenant)
	if err != nil {
		d.logger.Printf("[GetTenantById::Error] %s", err)
		return nil, handleError(err)
	}
	return tenant, err
}

func (d *DatastoreService) GetTenantByCode(code string) (*TenantEntity, error) {

	d.logger.Printf("[GetTenantByCode::Debug] code: %s", code)
	list, err := d.ListTenants(&ListTenantsParams{Code: code})
	if err != nil {
		d.logger.Printf("[GetTenantByCode::Error] %s", err)
		return nil, handleError(err)
	}

	if len(list) == 0 {
		d.logger.Printf("[GetTenantByCode::Error] %s", ErrNotFound)
		return nil, errors.New(ErrNotFound)
	}

	return list[0], nil

}

func (d *DatastoreService) GetTenantByEmail(email string) (*TenantEntity, error) {

	list, err := d.ListTenants(&ListTenantsParams{Email: email})
	if err != nil {
		d.logger.Printf("[GetTenantByEmail::Error] %s", err)
		return nil, handleError(err)
	}

	if len(list) == 0 {
		return nil, errors.New(ErrNotFound)
	}

	return list[0], nil

}

// code and email should be unique across db so we force the limit to 1
// @todo: implement limit + offset
func (d *DatastoreService) ListTenants(params *ListTenantsParams) ([]*TenantEntity, error) {

	d.logger.Printf("[ListTenants::Debug] ListTenantsParams: %v", params)
	var tenants []*TenantEntity
	var err error
	var code, email, name string

	code = strings.TrimSpace(params.Code)
	email = strings.TrimSpace(params.Email)
	name = strings.TrimSpace(params.Name)

	query := datastore.NewQuery(TENANT)
	if len(code) > 0 {
		query = query.Filter("Code =", code).Limit(1)
	}

	if len(email) > 0 {
		query = query.Filter("Email =", email).Limit(1)
	}

	if len(name) > 0 {
		query = query.Filter("Name =", name)
	}

	_, err = d.client.GetAll(context.Background(), query, &tenants)
	if err != nil {
		d.logger.Printf("[ListTenants::Error] %s", err)
		return nil, handleError(err)
	}

	return tenants, err
}

func (d *DatastoreService) SaveTenant(tenant *TenantEntity) error {

	var k *datastore.Key

	if tenant.Id == 0 {
		k = datastore.IncompleteKey(TENANT, nil)
		d.logger.Printf("Creating key for new tenant: %v", k)
	} else {
		k = datastore.IDKey(TENANT, tenant.Id, nil)
		d.logger.Printf("Tenant key: %v", k)
	}

	newKey, err := d.client.Put(context.Background(), k, tenant)
	if err != nil {
		d.logger.Printf("Cannot save tenant: %s", err)
		return err
	}

	// if we are creating a new tenant we have to update Id field
	if tenant.Id == 0 {
		tenant.Id = newKey.ID
		_, err = d.client.Put(context.Background(), newKey, tenant)
		if err != nil {
			d.logger.Printf("Cannot update new tenant id: %s", err)
			return err
		}
	}

	d.logger.Printf("Tenant successfully saved with id %d", tenant.Id)
	return nil
}

func (d *DatastoreService) ListUsers(params *ListUsersParams) ([]*v1.User, error) {

	var users []*v1.User
	var err error
	var keys []*datastore.Key

	query := datastore.NewQuery(KIND).KeysOnly() // individual entities will be retrieved from cache
	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	if len(params.FirstName) > 0 {
		query = query.Filter("FirstName =", params.FirstName)
	}

	if len(params.LastName) > 0 {
		query = query.Filter("LastName =", params.LastName)
	}

	if len(params.Email) > 0 {
		query = query.Filter("Email =", params.Email)
	}

	if params.TenantId > 0 {
		query = query.Filter("TenantId =", params.TenantId)
	}

	if params.Role >= 0 {
		query = query.Filter("Role =", params.Role)
	}

	keys, err = d.client.GetAll(context.Background(), query, nil)
	if err != nil {
		d.logger.Printf("[ListUsers::Error] %s", err)
		return nil, handleError(err)
	}

	for _, key := range keys {
		user, err := d.GetById(key.ID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (d *DatastoreService) ListUserNotifications(params *ListUserNotificationsParams) ([]*UserNotificationEntity, error) {
	var list []*UserNotificationEntity
	var err error
	var keys []*datastore.Key

	query := datastore.NewQuery(USER_NOTIFICATION).KeysOnly() // individual entities will be retrieved from cache if needed

	if params.Limit > 0 {
		query = query.Limit(params.Limit)
	}

	if params.Offset > 0 {
		query = query.Offset(params.Offset)
	}

	if params.UserId > 0 {
		query = query.Filter("UserId = ", params.UserId)
	}

	if len(params.Email) > 0 {
		query = query.Filter("Email = ", params.Email)
	}

	if len(params.Mobile) > 0 {
		query = query.Filter("Mobile = ", params.Mobile)
	}

	if params.AlertId > 0 {
		query = query.Filter("AlertId = ", params.AlertId)
	}

	if params.CreatedAtStart != nil {
		query = query.Filter("CreatedAt >= ", params.CreatedAtStart)
	}

	if params.CreatedAtEnd != nil {
		query = query.Filter("CreatedAt <= ", params.CreatedAtEnd)
	}

	if params.SentAtStart != nil {
		query = query.Filter("SentAt >= ", params.SentAtStart)
	}

	if params.SentAtEnd != nil {
		query = query.Filter("SentAt <= ", params.SentAtEnd)
	}

	if params.Status != nil {
		query = query.Filter("Status = ", int32(*params.Status))
	}

	if params.Type != nil {
		query = query.Filter("Type = ", int32(*params.Type))
	}

	keys, err = d.client.GetAll(context.Background(), query, nil)
	if err != nil {
		d.logger.Printf("[ListUserNotifications::Error] %s", err)
		return nil, handleError(err)
	}

	for _, key := range keys {
		userNotification, err := d.GetUserNotificationById(key.ID)
		if err != nil {
			d.logger.Printf("[ListUserNotifications::Error] %s", err)
			return nil, err
		}
		list = append(list, userNotification)
	}

	return list, nil
}

func (d *DatastoreService) ListPendingUserNotifications() ([]*UserNotificationEntity, error) {
	var err error
	var list []*UserNotificationEntity

	query := datastore.NewQuery(USER_NOTIFICATION).
		Filter("Status = ", 0). //v1.UserNotificationStatus_PENDING
		Filter("RefId = ", "")

	_, err = d.client.GetAll(context.Background(), query, &list)
	if err != nil {
		d.logger.Printf("[ListPendingUserNotifications::Error] %s", err)
		return nil, handleError(err)
	}

	return list, err
}

func (d *DatastoreService) GetUserNotificationById(id int64) (*UserNotificationEntity, error) {

	var err error
	var entity *UserNotificationEntity = &UserNotificationEntity{}

	k := datastore.IDKey(USER_NOTIFICATION, id, nil)
	err = d.client.Get(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("[GetUserNotificationById::Error] %s", err)
		return nil, handleError(err)
	}

	return entity, nil
}

func (d *DatastoreService) GetUserNotificationByRefId(refId string) (*UserNotificationEntity, error) {

	var err error
	var list []*UserNotificationEntity

	query := datastore.NewQuery(USER_NOTIFICATION).Filter("RefId =", refId).Limit(1)
	_, err = d.client.GetAll(context.Background(), query, &list)
	if err != nil {
		d.logger.Printf("[GetUserNotificationByRefId::Error] %s", err)
		return nil, handleError(err)
	}

	if len(list) == 0 {
		return nil, errors.New(ErrNotFound)
	}

	return list[0], nil
}

func (d *DatastoreService) CreateUserNotification(entity *UserNotificationEntity) error {

	k := datastore.IncompleteKey(USER_NOTIFICATION, nil)
	newKey, err := d.client.Put(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("Cannot create new user notification: %s", err)
		return err
	}

	// update Id field
	entity.Id = newKey.ID
	_, err = d.client.Put(context.Background(), newKey, entity)
	if err != nil {
		d.logger.Printf("Cannot update new user notification id: %s", err)
		return err
	}

	d.logger.Printf("User notification successfully created with id %d", entity.Id)
	return nil
}

func (d *DatastoreService) UpdateUserNotification(id int64, entity *UserNotificationEntity) error {

	var err error

	d.logger.Printf("UpdateUserNotification: entity price: %v", entity.Price)
	k := datastore.IDKey(USER_NOTIFICATION, entity.Id, nil)
	_, err = d.client.Put(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("Cannot update user notification: %s", err)
		return err
	}

	d.logger.Printf("User notification successfully updated with id %d", entity.Id)

	return nil

}

func (d *DatastoreService) GetPaymentById(paymentId int64) (*v1.Payment, error) {

	var entity *PaymentEntity
	var err error

	entity = &PaymentEntity{}
	k := datastore.IDKey(PAYMENT, paymentId, nil)
	err = d.client.Get(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("[GetPaymentById::Error] %s", err)
		return nil, handleError(err)
	}

	return ConvertPaymentEntityToStruct(entity), nil
}

func (d *DatastoreService) SavePayment(payment *v1.Payment) error {

	k := datastore.IncompleteKey(PAYMENT, nil)
	entity := ConvertPaymentToEntity(payment)

	newKey, err := d.client.Put(context.Background(), k, entity)
	if err != nil {
		d.logger.Printf("Cannot create new payment: %s", err)
		return err
	}

	// update Id field
	entity.Id = newKey.ID
	payment.Id = newKey.ID
	_, err = d.client.Put(context.Background(), newKey, entity)
	if err != nil {
		d.logger.Printf("Cannot update new entity id: %s", err)
		return err
	}

	d.logger.Printf("Payment successfully created with id %d", entity.Id)
	return nil

}

func (d *DatastoreService) ListPayments(tenantId int64) ([]*v1.Payment, error) {

	var payments []*PaymentEntity
	var list []*v1.Payment
	var err error

	query := datastore.NewQuery(PAYMENT).Filter("TenantId =", tenantId)

	_, err = d.client.GetAll(context.Background(), query, &payments)
	if err != nil {
		d.logger.Printf("[ListPayments::Error] %s", err)
		return nil, handleError(err)
	}

	for _, payment := range payments {
		list = append(list, ConvertPaymentEntityToStruct(payment))
	}

	return list, err
}

func handleError(datastoreError error) error {
	var err error

	switch datastoreError.Error() {
	case "datastore: no such entity":
		err = errors.New(ErrNotFound)
	case "datastore: invalid entity type":
		err = errors.New(ErrInvalidEntityType)
	case "code = Unauthenticated desc = transport: context canceled":
		err = errors.New(ErrContextCancelled)
	default:
		log.Printf("Error not implemented: %s", datastoreError)
		err = errors.New("Datastore error")
	}

	return err
}
