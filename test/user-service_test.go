package test

import (
	context "context"
	"errors"
	"fmt"
	"testing"

	v1 "github.com/bizio/user-service/pkg/api/v1"
	v1Service "github.com/bizio/user-service/pkg/service/v1"
	data "github.com/bizio/user-service/pkg/service/v1/data"
	"github.com/bizio/user-service/test/mock"
	"google.golang.org/grpc/metadata"

	gomock "github.com/golang/mock/gomock"
)

var mockUser *v1.User = &v1.User{
	Id:        int64(1),
	FirstName: "Test",
	LastName:  "User",
	Email:     "test.user@sosor.eu",
	Password:  "somerandompass",
	TenantId:  int64(1),
	Role:      v1.Role_USER,
	Mobile:    "3331234567",
}

var mockAdmin *v1.User = &v1.User{
	Id:        int64(2),
	FirstName: "Test",
	LastName:  "Admin",
	Email:     "test.admin@sosor.eu",
	Password:  "somerandompass",
	TenantId:  int64(1),
	Role:      v1.Role_ADMIN,
	Mobile:    "3331234567",
}

var mockTenantAdmin *v1.User = &v1.User{
	Id:        int64(3),
	FirstName: "Test",
	LastName:  "Tenant Admin",
	Email:     "test.tenant-admin@sosor.eu",
	Password:  "somerandompass",
	TenantId:  int64(1),
	Role:      v1.Role_TENANT_ADMIN,
	Mobile:    "3331234567",
}

var mockUser2 *v1.User = &v1.User{
	Id:        int64(4),
	FirstName: "Test",
	LastName:  "User",
	Email:     "test.user2@sosor.eu",
	Password:  "somerandompass",
	TenantId:  int64(2),
	Role:      v1.Role_USER,
	Mobile:    "3331234567",
}

var mockTenant *v1.Tenant = &v1.Tenant{
	Id:      1,
	Name:    "Mock tenant",
	Code:    "mocktenant",
	Email:   "mocktenant@domain.com",
	LogoUrl: "logo.png",
}

var mockTenant2 *v1.Tenant = &v1.Tenant{
	Id:      2,
	Name:    "Mock tenant2",
	Code:    "mocktenant2",
	Email:   "mocktenant2@domain.com",
	LogoUrl: "logo.png",
}

var users = []*v1.User{mockUser, mockUser2, mockTenantAdmin, mockAdmin}

func TestGet(t *testing.T) {
}

func TestGetWithEmptyContextFails(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Get(ctx, &v1.GetUserRequest{Id: mockUser.Id})
	if err == nil {
		t.Error(errors.New("Error expected"))
	} else {
		expectedErrorMessage := "rpc error: code = PermissionDenied desc = Request not authorized"
		if err.Error() != expectedErrorMessage {
			t.Error(errors.New(fmt.Sprintf("Wrong error want: %s, got: %s", expectedErrorMessage, err)))
		}
	}

}

func TestAdminCanGetUser(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockDatastore.
		EXPECT().
		GetById(mockUser.Id).
		Return(mockUser, nil)
	mockDatastore.
		EXPECT().
		GetTenantById(mockUser.TenantId).
		Return(data.ConvertTenantToEntity(mockTenant), nil)
	mockDatastore.
		EXPECT().
		GetById(mockAdmin.Id).
		Return(mockAdmin, nil)

	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	/*user*/ _, err := userServiceServer.Get(ctx, &v1.GetUserRequest{Id: mockUser.Id})
	if err != nil {
		t.Error(err)
	}

	/**
	if user.Id != mockUser.Id {
		t.Error(errors.New("User id doesn't match"))
	}
	*/

}

func TestAdminCanGetUserWithDifferentTenantId(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockDatastore.
		EXPECT().
		GetById(mockUser2.Id).
		Return(mockUser2, nil)
	mockDatastore.
		EXPECT().
		GetTenantById(mockUser2.TenantId).
		Return(data.ConvertTenantToEntity(mockTenant2), nil)
	mockDatastore.
		EXPECT().
		GetById(mockAdmin.Id).
		Return(mockAdmin, nil)

	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	user, err := userServiceServer.Get(ctx, &v1.GetUserRequest{Id: mockUser2.Id})
	if err != nil {
		t.Error(err)
	}

	if user.Id != mockUser2.Id {
		t.Error(errors.New("User id doesn't match"))
	}

}

func TestTenantAdminCanGetUserWithSameTenantId(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockTenantAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockDatastore.
		EXPECT().
		GetById(mockTenantAdmin.Id).
		Return(mockTenantAdmin, nil)
	mockDatastore.
		EXPECT().
		GetTenantById(mockTenantAdmin.TenantId).
		Return(data.ConvertTenantToEntity(mockTenant), nil)
	mockDatastore.
		EXPECT().
		GetById(mockUser.Id).
		Return(mockUser, nil)

	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	user, err := userServiceServer.Get(ctx, &v1.GetUserRequest{Id: mockUser.Id, TenantId: mockUser.TenantId})
	if err != nil {
		t.Error(err)
	}

	if user.Id != mockUser.Id {
		t.Error(errors.New("User id doesn't match"))
	}

}

func TestTenantAdminCannotGetUserWithDifferentTenantId(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockTenantAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockDatastore.
		EXPECT().
		GetById(mockTenantAdmin.Id).
		Return(mockTenantAdmin, nil)

	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Get(ctx, &v1.GetUserRequest{Id: mockUser2.Id, TenantId: mockUser2.TenantId})
	if err == nil {
		t.Error(errors.New("Error expected"))
	} else {
		expectedErrorMessage := "rpc error: code = PermissionDenied desc = Request not authorized"
		if err.Error() != expectedErrorMessage {
			t.Error(errors.New(fmt.Sprintf("Wrong error want: %s, got: %s", expectedErrorMessage, err)))
		}
	}

}
func TestCreateWithEmptyContextFails(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(context.Background(), &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		Mobile:    mockUser.Mobile,
	})

	if err == nil {
		t.Error(errors.New("Error expected"))
	} else {
		expectedErrorMessage := "rpc error: code = PermissionDenied desc = Request not authorized"
		if err.Error() != expectedErrorMessage {
			t.Error(errors.New(fmt.Sprintf("Wrong error want: %s, got: %s", expectedErrorMessage, err)))
		}
	}
}

// context with no metadata
func TestCreateWithoutUserContextFails(t *testing.T) {
	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("random", "somedata"))
	mockDatastore := mock.NewMockDatastore(mockCtrl)
	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		Mobile:    mockUser.Mobile,
	})

	if err == nil {
		t.Error(errors.New("Error expected"))
	} else {
		expectedErrorMessage := "rpc error: code = PermissionDenied desc = Invalid user id"
		if err.Error() != expectedErrorMessage {
			t.Error(errors.New(fmt.Sprintf("Wrong error want: %s, got: %s", expectedErrorMessage, err)))
		}
	}
}

func TestAdminCanCreateTenantAdminWithDifferentTenantId(t *testing.T) {

	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)

	tenantId := int64(2)
	mockDatastore.
		EXPECT().
		GetById(mockAdmin.Id).
		Return(mockAdmin, nil)

	// check for duplicates
	mockDatastore.
		EXPECT().
		GetByEmailAndTenantId(mockTenantAdmin.Email, tenantId).
		Return(nil, errors.New(data.ErrNotFound)).
		AnyTimes()

	mockDatastore.
		EXPECT().
		Save(gomock.Any()).
		Return(nil).
		AnyTimes()

	mockDatastore.
		EXPECT().
		ListUsers(gomock.Any()).
		Return(users, nil).
		AnyTimes()

	var userNotifications []*data.UserNotificationEntity
	mockDatastore.
		EXPECT().
		ListUserNotifications(gomock.Any()).
		Return(userNotifications, nil).
		AnyTimes()

	mockQueue := mock.NewMockQueue(mockCtrl)

	mockCacheService := mock.NewMockCacheService(mockCtrl)
	mockCacheService.
		EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockTenantAdmin.FirstName,
		LastName:  mockTenantAdmin.LastName,
		Email:     mockTenantAdmin.Email,
		Password:  "randompass",
		TenantId:  tenantId,
		Role:      mockTenantAdmin.Role,
		Mobile:    mockTenantAdmin.Mobile,
	})

	if err != nil {
		t.Error(err)
	}

}
func TestAdminCanCreateUserWithDifferentTenantId(t *testing.T) {

	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockAdmin.Id)))

	tenantId := int64(222)
	mockDatastore := mock.NewMockDatastore(mockCtrl)
	// check for duplicates
	mockDatastore.
		EXPECT().
		GetByEmailAndTenantId(mockUser.Email, tenantId).
		Return(nil, errors.New(data.ErrNotFound)).
		AnyTimes()

	mockDatastore.
		EXPECT().
		GetById(mockAdmin.Id).
		Return(mockAdmin, nil)
	mockDatastore.
		EXPECT().
		Save(gomock.Any()).
		Return(nil).
		AnyTimes()
	mockDatastore.
		EXPECT().
		ListUsers(gomock.Any()).
		Return(users, nil).
		AnyTimes()

	var userNotifications []*data.UserNotificationEntity
	mockDatastore.
		EXPECT().
		ListUserNotifications(gomock.Any()).
		Return(userNotifications, nil).
		AnyTimes()

	mockQueue := mock.NewMockQueue(mockCtrl)

	mockCacheService := mock.NewMockCacheService(mockCtrl)
	mockCacheService.
		EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		TenantId:  tenantId,
		Role:      mockUser.Role,
		Mobile:    mockUser.Mobile,
	})

	if err != nil {
		t.Error(err)
	}

}

func TestAdminCanCreateUserWithSameTenantId(t *testing.T) {

	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	// check for duplicates
	mockDatastore.
		EXPECT().
		GetByEmailAndTenantId(mockUser.Email, mockUser.TenantId).
		Return(nil, errors.New(data.ErrNotFound)).
		AnyTimes()

	mockDatastore.
		EXPECT().
		GetById(mockAdmin.Id).
		Return(mockAdmin, nil)
	mockDatastore.
		EXPECT().
		Save(gomock.Any()).
		Return(nil).
		AnyTimes()
	mockDatastore.
		EXPECT().
		ListUsers(gomock.Any()).
		Return(users, nil).
		AnyTimes()

	var userNotifications []*data.UserNotificationEntity
	mockDatastore.
		EXPECT().
		ListUserNotifications(gomock.Any()).
		Return(userNotifications, nil).
		AnyTimes()

	mockQueue := mock.NewMockQueue(mockCtrl)

	mockCacheService := mock.NewMockCacheService(mockCtrl)
	mockCacheService.
		EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		TenantId:  mockAdmin.TenantId,
		Role:      mockUser.Role,
		Mobile:    mockUser.Mobile,
	})

	if err != nil {
		t.Error(err)
	}

}

func TestTenantAdminCanCreateUserWithSameTenantId(t *testing.T) {

	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockTenantAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	// check for duplicates
	mockDatastore.
		EXPECT().
		GetByEmailAndTenantId(mockUser.Email, mockUser.TenantId).
		Return(nil, errors.New(data.ErrNotFound)).
		AnyTimes()

	mockDatastore.
		EXPECT().
		GetById(mockTenantAdmin.Id).
		Return(mockAdmin, nil)
	mockDatastore.
		EXPECT().
		Save(gomock.Any()).
		Return(nil).
		AnyTimes()
	mockDatastore.
		EXPECT().
		ListUsers(gomock.Any()).
		Return(users, nil).
		AnyTimes()

	var userNotifications []*data.UserNotificationEntity
	mockDatastore.
		EXPECT().
		ListUserNotifications(gomock.Any()).
		Return(userNotifications, nil).
		AnyTimes()

	mockQueue := mock.NewMockQueue(mockCtrl)

	mockCacheService := mock.NewMockCacheService(mockCtrl)
	mockCacheService.
		EXPECT().
		Set(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		TenantId:  mockUser.TenantId,
		Role:      mockUser.Role,
		Mobile:    mockUser.Mobile,
	})

	if err != nil {
		t.Error(err)
	}

}

func TestTenantAdminCannotCreateUserWithDifferentTenantId(t *testing.T) {

	mockCtrl, ctx := gomock.WithContext(context.Background(), t)
	defer mockCtrl.Finish()

	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("user", fmt.Sprintf("%d", mockTenantAdmin.Id)))

	mockDatastore := mock.NewMockDatastore(mockCtrl)
	// check for duplicates
	mockDatastore.
		EXPECT().
		GetByEmailAndTenantId(mockUser.Email, mockUser.TenantId).
		Return(nil, errors.New(data.ErrNotFound)).
		AnyTimes()

	mockDatastore.
		EXPECT().
		GetById(mockTenantAdmin.Id).
		Return(mockTenantAdmin, nil)
	mockDatastore.
		EXPECT().
		Save(gomock.Any()).
		Return(nil).
		AnyTimes()

	mockQueue := mock.NewMockQueue(mockCtrl)
	mockCacheService := mock.NewMockCacheService(mockCtrl)

	userServiceServer := v1Service.NewUserServiceServer(mockDatastore, mockQueue, mockCacheService)
	_, err := userServiceServer.Create(ctx, &v1.CreateUserRequest{
		FirstName: mockUser.FirstName,
		LastName:  mockUser.LastName,
		Email:     mockUser.Email,
		Password:  "randompass",
		TenantId:  2,
		Role:      mockUser.Role,
		Mobile:    mockUser.Mobile,
	})

	if err == nil {
		t.Error(errors.New("Error expected"))
	} else {
		expectedErrorMessage := "rpc error: code = PermissionDenied desc = Request not authorized"
		if err.Error() != expectedErrorMessage {
			t.Error(errors.New(fmt.Sprintf("Wrong error want: %s, got: %s", expectedErrorMessage, err)))
		}
	}

}
