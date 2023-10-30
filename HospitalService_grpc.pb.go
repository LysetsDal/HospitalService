// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: HospitalService/HospitalService.proto

package proto

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

const (
	HospitalService_RegisterClient_FullMethodName      = "/HospitalService/RegisterClient"
	HospitalService_GetClients_FullMethodName          = "/HospitalService/GetClients"
	HospitalService_SendMessageToServer_FullMethodName = "/HospitalService/SendMessageToServer"
	HospitalService_SendSecretToServer_FullMethodName  = "/HospitalService/SendSecretToServer"
)

// HospitalServiceClient is the client API for HospitalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HospitalServiceClient interface {
	RegisterClient(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error)
	GetClients(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*ClientIDsList, error)
	SendMessageToServer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error)
	SendSecretToServer(ctx context.Context, in *ShareSecretRes, opts ...grpc.CallOption) (*MessageRes, error)
}

type hospitalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHospitalServiceClient(cc grpc.ClientConnInterface) HospitalServiceClient {
	return &hospitalServiceClient{cc}
}

func (c *hospitalServiceClient) RegisterClient(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRes, error) {
	out := new(RegisterRes)
	err := c.cc.Invoke(ctx, HospitalService_RegisterClient_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) GetClients(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*ClientIDsList, error) {
	out := new(ClientIDsList)
	err := c.cc.Invoke(ctx, HospitalService_GetClients_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) SendMessageToServer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error) {
	out := new(MessageRes)
	err := c.cc.Invoke(ctx, HospitalService_SendMessageToServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hospitalServiceClient) SendSecretToServer(ctx context.Context, in *ShareSecretRes, opts ...grpc.CallOption) (*MessageRes, error) {
	out := new(MessageRes)
	err := c.cc.Invoke(ctx, HospitalService_SendSecretToServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HospitalServiceServer is the server API for HospitalService service.
// All implementations must embed UnimplementedHospitalServiceServer
// for forward compatibility
type HospitalServiceServer interface {
	RegisterClient(context.Context, *RegisterReq) (*RegisterRes, error)
	GetClients(context.Context, *MessageReq) (*ClientIDsList, error)
	SendMessageToServer(context.Context, *MessageReq) (*MessageRes, error)
	SendSecretToServer(context.Context, *ShareSecretRes) (*MessageRes, error)
	mustEmbedUnimplementedHospitalServiceServer()
}

// UnimplementedHospitalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHospitalServiceServer struct {
}

func (UnimplementedHospitalServiceServer) RegisterClient(context.Context, *RegisterReq) (*RegisterRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedHospitalServiceServer) GetClients(context.Context, *MessageReq) (*ClientIDsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClients not implemented")
}
func (UnimplementedHospitalServiceServer) SendMessageToServer(context.Context, *MessageReq) (*MessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToServer not implemented")
}
func (UnimplementedHospitalServiceServer) SendSecretToServer(context.Context, *ShareSecretRes) (*MessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSecretToServer not implemented")
}
func (UnimplementedHospitalServiceServer) mustEmbedUnimplementedHospitalServiceServer() {}

// UnsafeHospitalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HospitalServiceServer will
// result in compilation errors.
type UnsafeHospitalServiceServer interface {
	mustEmbedUnimplementedHospitalServiceServer()
}

func RegisterHospitalServiceServer(s grpc.ServiceRegistrar, srv HospitalServiceServer) {
	s.RegisterService(&HospitalService_ServiceDesc, srv)
}

func _HospitalService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_RegisterClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).RegisterClient(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_GetClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).GetClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_GetClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).GetClients(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_SendMessageToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).SendMessageToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_SendMessageToServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).SendMessageToServer(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HospitalService_SendSecretToServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareSecretRes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HospitalServiceServer).SendSecretToServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HospitalService_SendSecretToServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HospitalServiceServer).SendSecretToServer(ctx, req.(*ShareSecretRes))
	}
	return interceptor(ctx, in, info, handler)
}

// HospitalService_ServiceDesc is the grpc.ServiceDesc for HospitalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HospitalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "HospitalService",
	HandlerType: (*HospitalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterClient",
			Handler:    _HospitalService_RegisterClient_Handler,
		},
		{
			MethodName: "GetClients",
			Handler:    _HospitalService_GetClients_Handler,
		},
		{
			MethodName: "SendMessageToServer",
			Handler:    _HospitalService_SendMessageToServer_Handler,
		},
		{
			MethodName: "SendSecretToServer",
			Handler:    _HospitalService_SendSecretToServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HospitalService/HospitalService.proto",
}

const (
	ClientService_SendMessageToPeer_FullMethodName = "/ClientService/SendMessageToPeer"
	ClientService_ShareSecret_FullMethodName       = "/ClientService/ShareSecret"
	ClientService_SendShareToPeer_FullMethodName   = "/ClientService/SendShareToPeer"
)

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	SendMessageToPeer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error)
	ShareSecret(ctx context.Context, in *ShareSecretReq, opts ...grpc.CallOption) (*ShareSecretRes, error)
	SendShareToPeer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*ShareMessage, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) SendMessageToPeer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*MessageRes, error) {
	out := new(MessageRes)
	err := c.cc.Invoke(ctx, ClientService_SendMessageToPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) ShareSecret(ctx context.Context, in *ShareSecretReq, opts ...grpc.CallOption) (*ShareSecretRes, error) {
	out := new(ShareSecretRes)
	err := c.cc.Invoke(ctx, ClientService_ShareSecret_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) SendShareToPeer(ctx context.Context, in *MessageReq, opts ...grpc.CallOption) (*ShareMessage, error) {
	out := new(ShareMessage)
	err := c.cc.Invoke(ctx, ClientService_SendShareToPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	SendMessageToPeer(context.Context, *MessageReq) (*MessageRes, error)
	ShareSecret(context.Context, *ShareSecretReq) (*ShareSecretRes, error)
	SendShareToPeer(context.Context, *MessageReq) (*ShareMessage, error)
	mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) SendMessageToPeer(context.Context, *MessageReq) (*MessageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToPeer not implemented")
}
func (UnimplementedClientServiceServer) ShareSecret(context.Context, *ShareSecretReq) (*ShareSecretRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShareSecret not implemented")
}
func (UnimplementedClientServiceServer) SendShareToPeer(context.Context, *MessageReq) (*ShareMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendShareToPeer not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_SendMessageToPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).SendMessageToPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_SendMessageToPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).SendMessageToPeer(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_ShareSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShareSecretReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).ShareSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_ShareSecret_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).ShareSecret(ctx, req.(*ShareSecretReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_SendShareToPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).SendShareToPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientService_SendShareToPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).SendShareToPeer(ctx, req.(*MessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessageToPeer",
			Handler:    _ClientService_SendMessageToPeer_Handler,
		},
		{
			MethodName: "ShareSecret",
			Handler:    _ClientService_ShareSecret_Handler,
		},
		{
			MethodName: "SendShareToPeer",
			Handler:    _ClientService_SendShareToPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "HospitalService/HospitalService.proto",
}
