// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/v1/api.proto

package api

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

// SmartCalendarClient is the client API for SmartCalendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmartCalendarClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskID, error)
	CreateTaskStage(ctx context.Context, in *CreateTaskStageRequest, opts ...grpc.CallOption) (*TaskStageID, error)
	GetTasksList(ctx context.Context, in *GetTasksListRequest, opts ...grpc.CallOption) (*TasksList, error)
	SubscribeUser(ctx context.Context, in *SubscribeUserRequest, opts ...grpc.CallOption) (*SubscribeUserResponse, error)
	UnsubscribeUser(ctx context.Context, in *UnsubscribeUserRequest, opts ...grpc.CallOption) (*UnsubscribeUserResponse, error)
	GetUserTasksByTime(ctx context.Context, in *GetUserTasksByTimeRequest, opts ...grpc.CallOption) (*TasksList, error)
}

type smartCalendarClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartCalendarClient(cc grpc.ClientConnInterface) SmartCalendarClient {
	return &smartCalendarClient{cc}
}

func (c *smartCalendarClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/SmartCalendar/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) CreateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*TaskID, error) {
	out := new(TaskID)
	err := c.cc.Invoke(ctx, "/SmartCalendar/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) CreateTaskStage(ctx context.Context, in *CreateTaskStageRequest, opts ...grpc.CallOption) (*TaskStageID, error) {
	out := new(TaskStageID)
	err := c.cc.Invoke(ctx, "/SmartCalendar/CreateTaskStage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) GetTasksList(ctx context.Context, in *GetTasksListRequest, opts ...grpc.CallOption) (*TasksList, error) {
	out := new(TasksList)
	err := c.cc.Invoke(ctx, "/SmartCalendar/GetTasksList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) SubscribeUser(ctx context.Context, in *SubscribeUserRequest, opts ...grpc.CallOption) (*SubscribeUserResponse, error) {
	out := new(SubscribeUserResponse)
	err := c.cc.Invoke(ctx, "/SmartCalendar/SubscribeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) UnsubscribeUser(ctx context.Context, in *UnsubscribeUserRequest, opts ...grpc.CallOption) (*UnsubscribeUserResponse, error) {
	out := new(UnsubscribeUserResponse)
	err := c.cc.Invoke(ctx, "/SmartCalendar/UnsubscribeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartCalendarClient) GetUserTasksByTime(ctx context.Context, in *GetUserTasksByTimeRequest, opts ...grpc.CallOption) (*TasksList, error) {
	out := new(TasksList)
	err := c.cc.Invoke(ctx, "/SmartCalendar/GetUserTasksByTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmartCalendarServer is the server API for SmartCalendar service.
// All implementations must embed UnimplementedSmartCalendarServer
// for forward compatibility
type SmartCalendarServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	CreateTask(context.Context, *Task) (*TaskID, error)
	CreateTaskStage(context.Context, *CreateTaskStageRequest) (*TaskStageID, error)
	GetTasksList(context.Context, *GetTasksListRequest) (*TasksList, error)
	SubscribeUser(context.Context, *SubscribeUserRequest) (*SubscribeUserResponse, error)
	UnsubscribeUser(context.Context, *UnsubscribeUserRequest) (*UnsubscribeUserResponse, error)
	GetUserTasksByTime(context.Context, *GetUserTasksByTimeRequest) (*TasksList, error)
	mustEmbedUnimplementedSmartCalendarServer()
}

// UnimplementedSmartCalendarServer must be embedded to have forward compatible implementations.
type UnimplementedSmartCalendarServer struct {
}

func (UnimplementedSmartCalendarServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSmartCalendarServer) CreateTask(context.Context, *Task) (*TaskID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedSmartCalendarServer) CreateTaskStage(context.Context, *CreateTaskStageRequest) (*TaskStageID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskStage not implemented")
}
func (UnimplementedSmartCalendarServer) GetTasksList(context.Context, *GetTasksListRequest) (*TasksList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTasksList not implemented")
}
func (UnimplementedSmartCalendarServer) SubscribeUser(context.Context, *SubscribeUserRequest) (*SubscribeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeUser not implemented")
}
func (UnimplementedSmartCalendarServer) UnsubscribeUser(context.Context, *UnsubscribeUserRequest) (*UnsubscribeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsubscribeUser not implemented")
}
func (UnimplementedSmartCalendarServer) GetUserTasksByTime(context.Context, *GetUserTasksByTimeRequest) (*TasksList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTasksByTime not implemented")
}
func (UnimplementedSmartCalendarServer) mustEmbedUnimplementedSmartCalendarServer() {}

// UnsafeSmartCalendarServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmartCalendarServer will
// result in compilation errors.
type UnsafeSmartCalendarServer interface {
	mustEmbedUnimplementedSmartCalendarServer()
}

func RegisterSmartCalendarServer(s grpc.ServiceRegistrar, srv SmartCalendarServer) {
	s.RegisterService(&SmartCalendar_ServiceDesc, srv)
}

func _SmartCalendar_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).CreateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_CreateTaskStage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).CreateTaskStage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/CreateTaskStage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).CreateTaskStage(ctx, req.(*CreateTaskStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_GetTasksList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTasksListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).GetTasksList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/GetTasksList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).GetTasksList(ctx, req.(*GetTasksListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_SubscribeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).SubscribeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/SubscribeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).SubscribeUser(ctx, req.(*SubscribeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_UnsubscribeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).UnsubscribeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/UnsubscribeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).UnsubscribeUser(ctx, req.(*UnsubscribeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartCalendar_GetUserTasksByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTasksByTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartCalendarServer).GetUserTasksByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SmartCalendar/GetUserTasksByTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartCalendarServer).GetUserTasksByTime(ctx, req.(*GetUserTasksByTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmartCalendar_ServiceDesc is the grpc.ServiceDesc for SmartCalendar service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmartCalendar_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SmartCalendar",
	HandlerType: (*SmartCalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _SmartCalendar_CreateUser_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _SmartCalendar_CreateTask_Handler,
		},
		{
			MethodName: "CreateTaskStage",
			Handler:    _SmartCalendar_CreateTaskStage_Handler,
		},
		{
			MethodName: "GetTasksList",
			Handler:    _SmartCalendar_GetTasksList_Handler,
		},
		{
			MethodName: "SubscribeUser",
			Handler:    _SmartCalendar_SubscribeUser_Handler,
		},
		{
			MethodName: "UnsubscribeUser",
			Handler:    _SmartCalendar_UnsubscribeUser_Handler,
		},
		{
			MethodName: "GetUserTasksByTime",
			Handler:    _SmartCalendar_GetUserTasksByTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}
