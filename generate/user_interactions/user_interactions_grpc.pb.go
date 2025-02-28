// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: user_interactions.proto

package user_interactions

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

// UserInteractionsServiceClient is the client API for UserInteractionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserInteractionsServiceClient interface {
	Create(ctx context.Context, in *UserInteraction, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserInteractionSRes, error)
	GetById(ctx context.Context, in *UserInteractionId, opts ...grpc.CallOption) (*UserInteractionRes, error)
	Update(ctx context.Context, in *UserInteractionRes, opts ...grpc.CallOption) (*Void, error)
	Delete(ctx context.Context, in *UserInteractionId, opts ...grpc.CallOption) (*Void, error)
}

type userInteractionsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserInteractionsServiceClient(cc grpc.ClientConnInterface) UserInteractionsServiceClient {
	return &userInteractionsServiceClient{cc}
}

func (c *userInteractionsServiceClient) Create(ctx context.Context, in *UserInteraction, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/userInteractions.UserInteractionsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInteractionsServiceClient) Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*UserInteractionSRes, error) {
	out := new(UserInteractionSRes)
	err := c.cc.Invoke(ctx, "/userInteractions.UserInteractionsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInteractionsServiceClient) GetById(ctx context.Context, in *UserInteractionId, opts ...grpc.CallOption) (*UserInteractionRes, error) {
	out := new(UserInteractionRes)
	err := c.cc.Invoke(ctx, "/userInteractions.UserInteractionsService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInteractionsServiceClient) Update(ctx context.Context, in *UserInteractionRes, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/userInteractions.UserInteractionsService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInteractionsServiceClient) Delete(ctx context.Context, in *UserInteractionId, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/userInteractions.UserInteractionsService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInteractionsServiceServer is the server API for UserInteractionsService service.
// All implementations must embed UnimplementedUserInteractionsServiceServer
// for forward compatibility
type UserInteractionsServiceServer interface {
	Create(context.Context, *UserInteraction) (*Void, error)
	Get(context.Context, *Void) (*UserInteractionSRes, error)
	GetById(context.Context, *UserInteractionId) (*UserInteractionRes, error)
	Update(context.Context, *UserInteractionRes) (*Void, error)
	Delete(context.Context, *UserInteractionId) (*Void, error)
	mustEmbedUnimplementedUserInteractionsServiceServer()
}

// UnimplementedUserInteractionsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserInteractionsServiceServer struct {
}

func (UnimplementedUserInteractionsServiceServer) Create(context.Context, *UserInteraction) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserInteractionsServiceServer) Get(context.Context, *Void) (*UserInteractionSRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserInteractionsServiceServer) GetById(context.Context, *UserInteractionId) (*UserInteractionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserInteractionsServiceServer) Update(context.Context, *UserInteractionRes) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserInteractionsServiceServer) Delete(context.Context, *UserInteractionId) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserInteractionsServiceServer) mustEmbedUnimplementedUserInteractionsServiceServer() {
}

// UnsafeUserInteractionsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserInteractionsServiceServer will
// result in compilation errors.
type UnsafeUserInteractionsServiceServer interface {
	mustEmbedUnimplementedUserInteractionsServiceServer()
}

func RegisterUserInteractionsServiceServer(s grpc.ServiceRegistrar, srv UserInteractionsServiceServer) {
	s.RegisterService(&UserInteractionsService_ServiceDesc, srv)
}

func _UserInteractionsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInteraction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInteractionsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userInteractions.UserInteractionsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInteractionsServiceServer).Create(ctx, req.(*UserInteraction))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInteractionsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInteractionsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userInteractions.UserInteractionsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInteractionsServiceServer).Get(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInteractionsService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInteractionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInteractionsServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userInteractions.UserInteractionsService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInteractionsServiceServer).GetById(ctx, req.(*UserInteractionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInteractionsService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInteractionRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInteractionsServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userInteractions.UserInteractionsService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInteractionsServiceServer).Update(ctx, req.(*UserInteractionRes))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInteractionsService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInteractionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInteractionsServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userInteractions.UserInteractionsService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInteractionsServiceServer).Delete(ctx, req.(*UserInteractionId))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInteractionsService_ServiceDesc is the grpc.ServiceDesc for UserInteractionsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInteractionsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userInteractions.UserInteractionsService",
	HandlerType: (*UserInteractionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserInteractionsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _UserInteractionsService_Get_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UserInteractionsService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserInteractionsService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserInteractionsService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_interactions.proto",
}
