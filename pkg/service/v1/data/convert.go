package data

import (
	"time"

	v1 "github.com/bizio/user-service/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func ConvertToEntity(user *v1.User) (*UserEntity, error) {

	createdAt, err := ptypes.Timestamp(user.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	updatedAt, err := ptypes.Timestamp(user.GetUpdatedAt())
	if err != nil {
		return nil, err
	}

	var notifications []*NotificationEntity
	for _, notification := range user.GetNotifications() {
		notifications = append(notifications, &NotificationEntity{Type: notification.GetType(), Priority: notification.GetPriority()})
	}

	entity := &UserEntity{
		Id:                 user.GetId(),
		FirstName:          user.GetFirstName(),
		LastName:           user.GetLastName(),
		Email:              user.GetEmail(),
		Password:           user.GetPassword(),
		TenantId:           user.GetTenantId(),
		Role:               user.GetRole(),
		Mobile:             user.GetMobile(),
		Notifications:      notifications,
		PasswordResetToken: user.GetPasswordResetToken(),
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}

	return entity, nil
}

func ConvertToStruct(entity *UserEntity) (*v1.User, error) {

	createdAt, err := ptypes.TimestampProto(entity.CreatedAt)
	if err != nil {
		return nil, err
	}

	updatedAt, err := ptypes.TimestampProto(entity.UpdatedAt)
	if err != nil {
		return nil, err
	}

	var notifications []*v1.NotificationSettings
	for _, notification := range entity.Notifications {
		notifications = append(notifications, &v1.NotificationSettings{Type: notification.Type, Priority: notification.Priority})
	}

	user := &v1.User{
		Id:                 entity.Id,
		FirstName:          entity.FirstName,
		LastName:           entity.LastName,
		Email:              entity.Email,
		Password:           entity.Password,
		TenantId:           entity.TenantId,
		Role:               entity.Role,
		Mobile:             entity.Mobile,
		Notifications:      notifications,
		PasswordResetToken: entity.PasswordResetToken,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}

	return user, nil
}

func ConvertUserNotificationToStruct(entity *UserNotificationEntity) (*v1.UserNotification, error) {

	var sentAt *timestamp.Timestamp

	createdAt, err := ptypes.TimestampProto(entity.CreatedAt)
	if err != nil {
		return nil, err
	}

	sentAt, err = ptypes.TimestampProto(entity.SentAt)
	if err != nil {
		return nil, err
	}

	userNotification := &v1.UserNotification{
		Id:           entity.Id,
		UserId:       entity.UserId,
		Email:        entity.Email,
		Mobile:       entity.Mobile,
		AlertId:      entity.AlertId,
		ShortMessage: entity.ShortMessage,
		LongMessage:  entity.LongMessage,
		Type:         v1.NotificationType(entity.Type),
		CreatedAt:    createdAt,
		SentAt:       sentAt,
		RefId:        entity.RefId,
		Status:       entity.Status,
		Reason:       entity.Reason,
		Price:        entity.Price,
		TenantCode:   entity.TenantCode,
		TenantName:   entity.TenantName,
	}

	return userNotification, nil
}

func ConvertUserNotificationToEntity(userNotification *v1.UserNotification) (*UserNotificationEntity, error) {

	var sentAt time.Time
	createdAt, err := ptypes.Timestamp(userNotification.CreatedAt)
	if err != nil {
		return nil, err
	}

	if userNotification.SentAt != nil {
		sentAt, err = ptypes.Timestamp(userNotification.SentAt)
		if err != nil {
			return nil, err
		}
	}

	entity := &UserNotificationEntity{
		Id:           userNotification.GetId(),
		UserId:       userNotification.GetUserId(),
		Email:        userNotification.GetEmail(),
		Mobile:       userNotification.GetMobile(),
		AlertId:      userNotification.GetAlertId(),
		ShortMessage: userNotification.GetShortMessage(),
		LongMessage:  userNotification.GetLongMessage(),
		Type:         userNotification.GetType(),
		CreatedAt:    createdAt,
		SentAt:       sentAt,
		RefId:        userNotification.GetRefId(),
		Status:       userNotification.GetStatus(),
		Reason:       userNotification.GetReason(),
		Price:        userNotification.GetPrice(),
		TenantCode:   userNotification.GetTenantCode(),
		TenantName:   userNotification.GetTenantName(),
	}

	return entity, nil
}

func ConvertPaymentToEntity(payment *v1.Payment) *PaymentEntity {

	entity := &PaymentEntity{
		Id:          payment.GetId(),
		TenantId:    payment.GetTenantId(),
		Ref:         payment.GetRef(),
		Amount:      payment.GetAmount(),
		PaymentDate: ConvertTimestamp(payment.GetPaymentDate()),
		PaidBy:      payment.GetPaidBy(),
		CreatedAt:   ConvertTimestamp(payment.GetCreatedAt()),
		CreatedBy:   payment.GetCreatedBy(),
	}

	return entity
}

func ConvertPaymentEntityToStruct(entity *PaymentEntity) *v1.Payment {

	payment := &v1.Payment{
		Id:          entity.Id,
		TenantId:    entity.TenantId,
		Ref:         entity.Ref,
		Amount:      entity.Amount,
		PaymentDate: ConvertTime(entity.PaymentDate),
		PaidBy:      entity.PaidBy,
		CreatedAt:   ConvertTime(entity.CreatedAt),
		CreatedBy:   entity.CreatedBy,
	}

	return payment
}

func ConvertTenantEntityToStruct(entity *TenantEntity) *v1.Tenant {
	return &v1.Tenant{
		Id:        entity.Id,
		Name:      entity.Name,
		Code:      entity.Code,
		Email:     entity.Email,
		Credit:    entity.Credit,
		CreatedAt: ConvertTime(entity.CreatedAt),
		UpdatedAt: ConvertTime(entity.UpdatedAt),
	}
}

func ConvertTenantToEntity(tenant *v1.Tenant) *TenantEntity {
	return &TenantEntity{
		Id:        tenant.Id,
		Name:      tenant.Name,
		Code:      tenant.Code,
		Email:     tenant.Email,
		Credit:    tenant.Credit,
		CreatedAt: ConvertTimestamp(tenant.CreatedAt),
		UpdatedAt: ConvertTimestamp(tenant.UpdatedAt),
	}
}

func ConvertTime(t time.Time) *timestamp.Timestamp {
	ts, _ := ptypes.TimestampProto(t)

	return ts
}

func ConvertTimestamp(ts *timestamp.Timestamp) time.Time {

	t, _ := ptypes.Timestamp(ts)
	return t
}
