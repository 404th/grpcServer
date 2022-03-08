// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user.proto

package user_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementClient interface {
	CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error)
	GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersList, error)
	DeleteUser(ctx context.Context, in *IDTracker, opts ...grpc.CallOption) (*Deleted, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) CreateUser(ctx context.Context, in *NewUser, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user_service.UserManagement/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetUsers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UsersList, error) {
	out := new(UsersList)
	err := c.cc.Invoke(ctx, "/user_service.UserManagement/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteUser(ctx context.Context, in *IDTracker, opts ...grpc.CallOption) (*Deleted, error) {
	out := new(Deleted)
	err := c.cc.Invoke(ctx, "/user_service.UserManagement/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility
type UserManagementServer interface {
	CreateUser(context.Context, *NewUser) (*User, error)
	GetUsers(context.Context, *Empty) (*UsersList, error)
	DeleteUser(context.Context, *IDTracker) (*Deleted, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServer struct {
}

func (UnimplementedUserManagementServer) CreateUser(context.Context, *NewUser) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserManagementServer) GetUsers(context.Context, *Empty) (*UsersList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserManagementServer) DeleteUser(context.Context, *IDTracker) (*Deleted, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserManagement/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateUser(ctx, req.(*NewUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserManagement/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUsers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDTracker)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_service.UserManagement/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteUser(ctx, req.(*IDTracker))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_service.UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserManagement_CreateUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserManagement_GetUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagement_DeleteUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
