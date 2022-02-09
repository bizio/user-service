// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bizio/user-service/pkg/service/v1/data (interfaces: Datastore)

// Package mock is a generated GoMock package.
package mock

import (
	v1 "github.com/bizio/user-service/pkg/api/v1"
	data "github.com/bizio/user-service/pkg/service/v1/data"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatastore is a mock of Datastore interface
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// CreateUserNotification mocks base method
func (m *MockDatastore) CreateUserNotification(arg0 *data.UserNotificationEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserNotification", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUserNotification indicates an expected call of CreateUserNotification
func (mr *MockDatastoreMockRecorder) CreateUserNotification(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserNotification", reflect.TypeOf((*MockDatastore)(nil).CreateUserNotification), arg0)
}

// GetByEmail mocks base method
func (m *MockDatastore) GetByEmail(arg0 string) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", arg0)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail
func (mr *MockDatastoreMockRecorder) GetByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockDatastore)(nil).GetByEmail), arg0)
}

// GetByEmailAndTenantId mocks base method
func (m *MockDatastore) GetByEmailAndTenantId(arg0 string, arg1 int64) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmailAndTenantId", arg0, arg1)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmailAndTenantId indicates an expected call of GetByEmailAndTenantId
func (mr *MockDatastoreMockRecorder) GetByEmailAndTenantId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmailAndTenantId", reflect.TypeOf((*MockDatastore)(nil).GetByEmailAndTenantId), arg0, arg1)
}

// GetById mocks base method
func (m *MockDatastore) GetById(arg0 int64) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockDatastoreMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockDatastore)(nil).GetById), arg0)
}

// GetByIdAndTenantId mocks base method
func (m *MockDatastore) GetByIdAndTenantId(arg0, arg1 int64) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdAndTenantId", arg0, arg1)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdAndTenantId indicates an expected call of GetByIdAndTenantId
func (mr *MockDatastoreMockRecorder) GetByIdAndTenantId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdAndTenantId", reflect.TypeOf((*MockDatastore)(nil).GetByIdAndTenantId), arg0, arg1)
}

// GetByPasswordResetToken mocks base method
func (m *MockDatastore) GetByPasswordResetToken(arg0 string) (*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPasswordResetToken", arg0)
	ret0, _ := ret[0].(*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPasswordResetToken indicates an expected call of GetByPasswordResetToken
func (mr *MockDatastoreMockRecorder) GetByPasswordResetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPasswordResetToken", reflect.TypeOf((*MockDatastore)(nil).GetByPasswordResetToken), arg0)
}

// GetPaymentById mocks base method
func (m *MockDatastore) GetPaymentById(arg0 int64) (*v1.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentById", arg0)
	ret0, _ := ret[0].(*v1.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentById indicates an expected call of GetPaymentById
func (mr *MockDatastoreMockRecorder) GetPaymentById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentById", reflect.TypeOf((*MockDatastore)(nil).GetPaymentById), arg0)
}

// GetTenantByCode mocks base method
func (m *MockDatastore) GetTenantByCode(arg0 string) (*data.TenantEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantByCode", arg0)
	ret0, _ := ret[0].(*data.TenantEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenantByCode indicates an expected call of GetTenantByCode
func (mr *MockDatastoreMockRecorder) GetTenantByCode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantByCode", reflect.TypeOf((*MockDatastore)(nil).GetTenantByCode), arg0)
}

// GetTenantByEmail mocks base method
func (m *MockDatastore) GetTenantByEmail(arg0 string) (*data.TenantEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantByEmail", arg0)
	ret0, _ := ret[0].(*data.TenantEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenantByEmail indicates an expected call of GetTenantByEmail
func (mr *MockDatastoreMockRecorder) GetTenantByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantByEmail", reflect.TypeOf((*MockDatastore)(nil).GetTenantByEmail), arg0)
}

// GetTenantById mocks base method
func (m *MockDatastore) GetTenantById(arg0 int64) (*data.TenantEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTenantById", arg0)
	ret0, _ := ret[0].(*data.TenantEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTenantById indicates an expected call of GetTenantById
func (mr *MockDatastoreMockRecorder) GetTenantById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTenantById", reflect.TypeOf((*MockDatastore)(nil).GetTenantById), arg0)
}

// GetUserNotificationById mocks base method
func (m *MockDatastore) GetUserNotificationById(arg0 int64) (*data.UserNotificationEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserNotificationById", arg0)
	ret0, _ := ret[0].(*data.UserNotificationEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserNotificationById indicates an expected call of GetUserNotificationById
func (mr *MockDatastoreMockRecorder) GetUserNotificationById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserNotificationById", reflect.TypeOf((*MockDatastore)(nil).GetUserNotificationById), arg0)
}

// GetUserNotificationByRefId mocks base method
func (m *MockDatastore) GetUserNotificationByRefId(arg0 string) (*data.UserNotificationEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserNotificationByRefId", arg0)
	ret0, _ := ret[0].(*data.UserNotificationEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserNotificationByRefId indicates an expected call of GetUserNotificationByRefId
func (mr *MockDatastoreMockRecorder) GetUserNotificationByRefId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserNotificationByRefId", reflect.TypeOf((*MockDatastore)(nil).GetUserNotificationByRefId), arg0)
}

// ListPayments mocks base method
func (m *MockDatastore) ListPayments(arg0 int64) ([]*v1.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPayments", arg0)
	ret0, _ := ret[0].([]*v1.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPayments indicates an expected call of ListPayments
func (mr *MockDatastoreMockRecorder) ListPayments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPayments", reflect.TypeOf((*MockDatastore)(nil).ListPayments), arg0)
}

// ListPendingUserNotifications mocks base method
func (m *MockDatastore) ListPendingUserNotifications() ([]*data.UserNotificationEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPendingUserNotifications")
	ret0, _ := ret[0].([]*data.UserNotificationEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPendingUserNotifications indicates an expected call of ListPendingUserNotifications
func (mr *MockDatastoreMockRecorder) ListPendingUserNotifications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPendingUserNotifications", reflect.TypeOf((*MockDatastore)(nil).ListPendingUserNotifications))
}

// ListTenants mocks base method
func (m *MockDatastore) ListTenants(arg0 *data.ListTenantsParams) ([]*data.TenantEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTenants", arg0)
	ret0, _ := ret[0].([]*data.TenantEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTenants indicates an expected call of ListTenants
func (mr *MockDatastoreMockRecorder) ListTenants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTenants", reflect.TypeOf((*MockDatastore)(nil).ListTenants), arg0)
}

// ListUserNotifications mocks base method
func (m *MockDatastore) ListUserNotifications(arg0 *data.ListUserNotificationsParams) ([]*data.UserNotificationEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserNotifications", arg0)
	ret0, _ := ret[0].([]*data.UserNotificationEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserNotifications indicates an expected call of ListUserNotifications
func (mr *MockDatastoreMockRecorder) ListUserNotifications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserNotifications", reflect.TypeOf((*MockDatastore)(nil).ListUserNotifications), arg0)
}

// ListUsers mocks base method
func (m *MockDatastore) ListUsers(arg0 *data.ListUsersParams) ([]*v1.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0)
	ret0, _ := ret[0].([]*v1.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers
func (mr *MockDatastoreMockRecorder) ListUsers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockDatastore)(nil).ListUsers), arg0)
}

// Save mocks base method
func (m *MockDatastore) Save(arg0 *v1.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockDatastoreMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockDatastore)(nil).Save), arg0)
}

// SavePayment mocks base method
func (m *MockDatastore) SavePayment(arg0 *v1.Payment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SavePayment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SavePayment indicates an expected call of SavePayment
func (mr *MockDatastoreMockRecorder) SavePayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SavePayment", reflect.TypeOf((*MockDatastore)(nil).SavePayment), arg0)
}

// SaveTenant mocks base method
func (m *MockDatastore) SaveTenant(arg0 *data.TenantEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTenant", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTenant indicates an expected call of SaveTenant
func (mr *MockDatastoreMockRecorder) SaveTenant(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTenant", reflect.TypeOf((*MockDatastore)(nil).SaveTenant), arg0)
}

// UpdateUserNotification mocks base method
func (m *MockDatastore) UpdateUserNotification(arg0 int64, arg1 *data.UserNotificationEntity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserNotification", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserNotification indicates an expected call of UpdateUserNotification
func (mr *MockDatastoreMockRecorder) UpdateUserNotification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserNotification", reflect.TypeOf((*MockDatastore)(nil).UpdateUserNotification), arg0, arg1)
}
