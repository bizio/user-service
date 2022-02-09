// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-service.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("user-service.proto", fileDescriptor_2a3086c73a75cdba) }

var fileDescriptor_2a3086c73a75cdba = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x4e, 0x13, 0x41,
	0x14, 0xc6, 0x11, 0x4d, 0x23, 0x53, 0xa8, 0x70, 0x4a, 0xff, 0x42, 0x80, 0x4e, 0x88, 0x26, 0x24,
	0x74, 0x03, 0xc6, 0xc4, 0x70, 0x87, 0x04, 0x1b, 0x83, 0x1a, 0xe5, 0x4f, 0xbc, 0x31, 0x31, 0x4b,
	0x3b, 0xc0, 0x2a, 0xdd, 0x5d, 0x76, 0x66, 0x4b, 0x90, 0x70, 0xe3, 0x2b, 0xf8, 0x50, 0x3e, 0x80,
	0xaf, 0xe0, 0x83, 0x98, 0x39, 0x67, 0x66, 0x77, 0xdb, 0x2e, 0x35, 0xde, 0xcd, 0x7e, 0x33, 0xf3,
	0xfb, 0x66, 0xce, 0x7e, 0x67, 0x18, 0xc4, 0x52, 0x44, 0x9b, 0x52, 0x44, 0x03, 0xaf, 0x2b, 0xda,
	0x61, 0x14, 0xa8, 0x00, 0xa6, 0x07, 0x5b, 0xcd, 0xe5, 0xf3, 0x20, 0x38, 0xbf, 0x14, 0x8e, 0x1b,
	0x7a, 0x8e, 0xeb, 0xfb, 0x81, 0x72, 0x95, 0x17, 0xf8, 0x92, 0x56, 0x34, 0x97, 0xcc, 0x2c, 0x7e,
	0x9d, 0xc6, 0x67, 0x8e, 0xe8, 0x87, 0xea, 0xc6, 0x4c, 0x32, 0x8d, 0xa4, 0xf1, 0xf6, 0xaf, 0x12,
	0x2b, 0x9e, 0x48, 0x11, 0x1d, 0x91, 0x01, 0xbc, 0x63, 0x85, 0xbd, 0x48, 0xb8, 0x4a, 0x40, 0xa5,
	0x3d, 0xd8, 0x6a, 0xd3, 0x58, 0x2f, 0x38, 0x14, 0x57, 0xb1, 0x90, 0xaa, 0x59, 0x1d, 0x95, 0x65,
	0x18, 0xf8, 0x52, 0xf0, 0xc5, 0x1f, 0xbf, 0xff, 0xfc, 0x9c, 0x2e, 0xf1, 0x19, 0x67, 0xb0, 0xe5,
	0x68, 0x03, 0xb9, 0xf3, 0x60, 0x03, 0x3e, 0xb1, 0xc2, 0x49, 0xd8, 0x4b, 0x70, 0x34, 0xce, 0xe2,
	0xca, 0x5a, 0xee, 0x08, 0x35, 0xc4, 0x6a, 0x21, 0x6b, 0x69, 0xbb, 0x9c, 0xb0, 0x9c, 0x5b, 0x3c,
	0xb3, 0xd7, 0xbb, 0xdb, 0x79, 0xa4, 0x47, 0xf0, 0x9a, 0x3d, 0xec, 0x08, 0x05, 0x30, 0xb4, 0x7d,
	0x02, 0xb2, 0x8a, 0xc8, 0x79, 0x28, 0x65, 0x90, 0x5e, 0xef, 0x0e, 0x02, 0x56, 0x4d, 0x0f, 0xf5,
	0x3e, 0x50, 0xde, 0x99, 0xd7, 0xc5, 0x4a, 0x42, 0x0b, 0x0f, 0x3c, 0xa2, 0xd2, 0x5a, 0xeb, 0xb4,
	0x98, 0xb7, 0x84, 0x2f, 0xa3, 0x55, 0x95, 0x2f, 0x68, 0x2b, 0x3f, 0x33, 0x83, 0x15, 0xb9, 0x66,
	0x95, 0xb7, 0x9e, 0x54, 0xa3, 0xbb, 0x24, 0xac, 0x69, 0x58, 0xee, 0x94, 0xb5, 0x6b, 0x4d, 0x58,
	0x61, 0xae, 0xd9, 0x40, 0xef, 0x32, 0x8c, 0x7b, 0xc3, 0x25, 0xab, 0x7e, 0x8c, 0x45, 0x2c, 0xc6,
	0x9d, 0xab, 0x6d, 0x4a, 0x4b, 0xdb, 0xa6, 0xa5, 0xbd, 0xaf, 0xd3, 0xd2, 0xbc, 0x47, 0xe7, 0x1c,
	0x4d, 0x96, 0x79, 0x6d, 0xcc, 0xc4, 0xb9, 0xd2, 0x0e, 0xfa, 0x9a, 0x27, 0x6c, 0xa6, 0x23, 0xd4,
	0xa1, 0x08, 0x83, 0x48, 0xc1, 0xa2, 0xf9, 0x23, 0xf4, 0x69, 0xaf, 0xc3, 0xb4, 0x4a, 0x12, 0x7f,
	0x8a, 0xc8, 0x35, 0x58, 0x49, 0x7f, 0xcf, 0x30, 0x38, 0x22, 0xd2, 0x01, 0x9b, 0xb1, 0x05, 0x90,
	0x84, 0x4d, 0x3e, 0x2d, 0xb6, 0x32, 0xa2, 0x9a, 0xca, 0x2c, 0xa0, 0x43, 0x11, 0xd2, 0x7c, 0xc2,
	0x19, 0x9b, 0x3b, 0x14, 0x52, 0xa8, 0x0f, 0xae, 0x94, 0xd7, 0x41, 0xd4, 0x83, 0x3a, 0x9d, 0x28,
	0x23, 0x59, 0x68, 0x23, 0x67, 0xc6, 0x80, 0xd7, 0x10, 0xdc, 0x84, 0x7a, 0x7a, 0xf4, 0xd0, 0xac,
	0x71, 0x22, 0xbd, 0x03, 0xfa, 0xac, 0x44, 0xb9, 0x49, 0x8c, 0x1a, 0x69, 0x33, 0x8c, 0x3a, 0x35,
	0xf3, 0xa6, 0x8c, 0xd5, 0x3a, 0x5a, 0xad, 0xf0, 0x46, 0x8e, 0x55, 0x8c, 0x5b, 0x74, 0xe9, 0x8f,
	0x59, 0xa5, 0x23, 0xd4, 0xab, 0x9b, 0xfd, 0xbe, 0xeb, 0x5d, 0xee, 0xfa, 0xbd, 0xc4, 0x75, 0x3d,
	0xd3, 0x18, 0xe3, 0xd3, 0x13, 0xdb, 0x67, 0x0a, 0x5e, 0x32, 0xd8, 0xbb, 0x10, 0xdd, 0x6f, 0xbb,
	0xb1, 0xba, 0x08, 0x22, 0xef, 0x3b, 0x35, 0xc9, 0x7d, 0xd1, 0x79, 0x6c, 0x3b, 0x83, 0x4f, 0xc1,
	0x01, 0x9b, 0xa5, 0xf7, 0xe2, 0x58, 0xf8, 0xae, 0xaf, 0xa0, 0x96, 0xbe, 0x20, 0xa4, 0x0c, 0x05,
	0x82, 0x24, 0xdb, 0xaf, 0xbc, 0xa8, 0xaf, 0xaa, 0x50, 0xc3, 0xf6, 0xf9, 0xcc, 0x66, 0xa9, 0x38,
	0x59, 0x58, 0x56, 0xc9, 0x83, 0x3d, 0x43, 0x58, 0x6b, 0xbb, 0x96, 0x81, 0x39, 0xb7, 0x34, 0xc0,
	0x37, 0xa5, 0x40, 0x63, 0xe8, 0x60, 0x6a, 0x0d, 0xda, 0xa6, 0xf6, 0x7e, 0x6e, 0x1d, 0xb9, 0x00,
	0xf3, 0x43, 0x5c, 0xfd, 0xac, 0x74, 0xd9, 0x93, 0x64, 0xe7, 0x6e, 0xaf, 0xef, 0x61, 0x97, 0x99,
	0x5c, 0x92, 0x9a, 0xe4, 0xb5, 0x36, 0xa6, 0x9b, 0x9a, 0xaf, 0x22, 0xbd, 0x01, 0xb5, 0x51, 0xba,
	0xe3, 0x12, 0xf1, 0x88, 0x15, 0x33, 0xfb, 0xfe, 0xdf, 0xa0, 0x8c, 0x06, 0x73, 0x90, 0xad, 0x31,
	0x9c, 0x33, 0xd6, 0xd1, 0x29, 0xbf, 0xe9, 0x0b, 0x5f, 0xd1, 0xab, 0x9d, 0x7e, 0x5b, 0x64, 0x51,
	0xcb, 0x46, 0xe3, 0x2f, 0x10, 0xe3, 0xc0, 0x66, 0x4e, 0x75, 0xbf, 0xe8, 0xe3, 0x86, 0xb4, 0x50,
	0x3a, 0xb7, 0x66, 0xa4, 0x55, 0x38, 0x65, 0x73, 0x14, 0x02, 0xeb, 0x55, 0x4f, 0x73, 0x31, 0xc9,
	0x6e, 0x03, 0xed, 0xd6, 0xf9, 0x6a, 0x8e, 0xdd, 0x9b, 0x8c, 0x9b, 0x4e, 0xcb, 0x57, 0x36, 0xab,
	0x2f, 0x6e, 0xb6, 0x4a, 0x48, 0x4a, 0x61, 0x15, 0xeb, 0x50, 0x1f, 0x9f, 0x30, 0x45, 0x32, 0xd9,
	0x81, 0x7f, 0xd9, 0x9d, 0x16, 0xb0, 0x05, 0x9e, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x83, 0x5a,
	0x90, 0x20, 0xb1, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUserNotification(ctx context.Context, in *UserNotificationUpdateRequest, opts ...grpc.CallOption) (*UserNotification, error)
	ListUserNotifications(ctx context.Context, in *ListUserNotificationsRequest, opts ...grpc.CallOption) (*ListUserNotificationsResponse, error)
	QueueUserNotifications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error)
	ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error)
	// used for login
	GetByEmailAndPassword(ctx context.Context, in *GetUserByEmailAndPasswordRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CheckAuthorization(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error)
	// Tenant service
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// Tenant service
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	GetTenantAdmins(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error)
	GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*Payment, error)
	CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Payment, error)
	ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserNotification(ctx context.Context, in *UserNotificationUpdateRequest, opts ...grpc.CallOption) (*UserNotification, error) {
	out := new(UserNotification)
	err := c.cc.Invoke(ctx, "/v1.UserService/UpdateUserNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUserNotifications(ctx context.Context, in *ListUserNotificationsRequest, opts ...grpc.CallOption) (*ListUserNotificationsResponse, error) {
	out := new(ListUserNotificationsResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ListUserNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueueUserNotifications(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/v1.UserService/QueueUserNotifications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetReport(ctx context.Context, in *GetReportRequest, opts ...grpc.CallOption) (*Report, error) {
	out := new(Report)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	out := new(ListUsersResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...grpc.CallOption) (*UpdatePasswordResponse, error) {
	out := new(UpdatePasswordResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/UpdatePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByEmailAndPassword(ctx context.Context, in *GetUserByEmailAndPasswordRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetByEmailAndPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckAuthorization(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/v1.UserService/CheckAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/v1.UserService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/v1.UserService/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetTenantAdmins(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetTenantAdmins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListTenants(ctx context.Context, in *ListTenantsRequest, opts ...grpc.CallOption) (*ListTenantsResponse, error) {
	out := new(ListTenantsResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ListTenants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetPayment(ctx context.Context, in *GetPaymentRequest, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/v1.UserService/GetPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreatePayment(ctx context.Context, in *CreatePaymentRequest, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/v1.UserService/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListPayments(ctx context.Context, in *ListPaymentsRequest, opts ...grpc.CallOption) (*ListPaymentsResponse, error) {
	out := new(ListPaymentsResponse)
	err := c.cc.Invoke(ctx, "/v1.UserService/ListPayments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Create(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Update(context.Context, *UpdateUserRequest) (*GetUserResponse, error)
	Get(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUserNotification(context.Context, *UserNotificationUpdateRequest) (*UserNotification, error)
	ListUserNotifications(context.Context, *ListUserNotificationsRequest) (*ListUserNotificationsResponse, error)
	QueueUserNotifications(context.Context, *empty.Empty) (*empty.Empty, error)
	GetReport(context.Context, *GetReportRequest) (*Report, error)
	ListUsers(context.Context, *ListUsersRequest) (*ListUsersResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*UpdatePasswordResponse, error)
	// used for login
	GetByEmailAndPassword(context.Context, *GetUserByEmailAndPasswordRequest) (*GetUserResponse, error)
	CheckAuthorization(context.Context, *empty.Empty) (*User, error)
	// Tenant service
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	// Tenant service
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	GetTenantAdmins(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	ListTenants(context.Context, *ListTenantsRequest) (*ListTenantsResponse, error)
	GetPayment(context.Context, *GetPaymentRequest) (*Payment, error)
	CreatePayment(context.Context, *CreatePaymentRequest) (*Payment, error)
	ListPayments(context.Context, *ListPaymentsRequest) (*ListPaymentsResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserNotificationUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/UpdateUserNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserNotification(ctx, req.(*UserNotificationUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUserNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUserNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ListUserNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUserNotifications(ctx, req.(*ListUserNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueueUserNotifications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueueUserNotifications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/QueueUserNotifications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueueUserNotifications(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetReport(ctx, req.(*GetReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*ListUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/UpdatePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdatePassword(ctx, req.(*UpdatePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByEmailAndPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByEmailAndPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByEmailAndPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetByEmailAndPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByEmailAndPassword(ctx, req.(*GetUserByEmailAndPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CheckAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckAuthorization(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetTenantAdmins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetTenantAdmins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetTenantAdmins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetTenantAdmins(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListTenants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListTenants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ListTenants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListTenants(ctx, req.(*ListTenantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/GetPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPayment(ctx, req.(*GetPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreatePayment(ctx, req.(*CreatePaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.UserService/ListPayments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListPayments(ctx, req.(*ListPaymentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "UpdateUserNotification",
			Handler:    _UserService_UpdateUserNotification_Handler,
		},
		{
			MethodName: "ListUserNotifications",
			Handler:    _UserService_ListUserNotifications_Handler,
		},
		{
			MethodName: "QueueUserNotifications",
			Handler:    _UserService_QueueUserNotifications_Handler,
		},
		{
			MethodName: "GetReport",
			Handler:    _UserService_GetReport_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _UserService_ResetPassword_Handler,
		},
		{
			MethodName: "UpdatePassword",
			Handler:    _UserService_UpdatePassword_Handler,
		},
		{
			MethodName: "GetByEmailAndPassword",
			Handler:    _UserService_GetByEmailAndPassword_Handler,
		},
		{
			MethodName: "CheckAuthorization",
			Handler:    _UserService_CheckAuthorization_Handler,
		},
		{
			MethodName: "CreateTenant",
			Handler:    _UserService_CreateTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _UserService_UpdateTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _UserService_GetTenant_Handler,
		},
		{
			MethodName: "GetTenantAdmins",
			Handler:    _UserService_GetTenantAdmins_Handler,
		},
		{
			MethodName: "ListTenants",
			Handler:    _UserService_ListTenants_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _UserService_GetPayment_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _UserService_CreatePayment_Handler,
		},
		{
			MethodName: "ListPayments",
			Handler:    _UserService_ListPayments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-service.proto",
}
