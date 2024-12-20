// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.0
// source: proto/mail_proto/mail_server.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MailService_SendVerificationCodeMail_FullMethodName = "/proto.MailService/SendVerificationCodeMail"
	MailService_SendResetPasswordMail_FullMethodName    = "/proto.MailService/SendResetPasswordMail"
	MailService_SendPlainTextMail_FullMethodName        = "/proto.MailService/SendPlainTextMail"
)

// MailServiceClient is the client API for MailService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailServiceClient interface {
	SendVerificationCodeMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
	SendResetPasswordMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
	SendPlainTextMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error)
}

type mailServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMailServiceClient(cc grpc.ClientConnInterface) MailServiceClient {
	return &mailServiceClient{cc}
}

func (c *mailServiceClient) SendVerificationCodeMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MailResponse)
	err := c.cc.Invoke(ctx, MailService_SendVerificationCodeMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailServiceClient) SendResetPasswordMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MailResponse)
	err := c.cc.Invoke(ctx, MailService_SendResetPasswordMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailServiceClient) SendPlainTextMail(ctx context.Context, in *MailRequest, opts ...grpc.CallOption) (*MailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MailResponse)
	err := c.cc.Invoke(ctx, MailService_SendPlainTextMail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailServiceServer is the server API for MailService service.
// All implementations must embed UnimplementedMailServiceServer
// for forward compatibility.
type MailServiceServer interface {
	SendVerificationCodeMail(context.Context, *MailRequest) (*MailResponse, error)
	SendResetPasswordMail(context.Context, *MailRequest) (*MailResponse, error)
	SendPlainTextMail(context.Context, *MailRequest) (*MailResponse, error)
	mustEmbedUnimplementedMailServiceServer()
}

// UnimplementedMailServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMailServiceServer struct{}

func (UnimplementedMailServiceServer) SendVerificationCodeMail(context.Context, *MailRequest) (*MailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendVerificationCodeMail not implemented")
}
func (UnimplementedMailServiceServer) SendResetPasswordMail(context.Context, *MailRequest) (*MailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendResetPasswordMail not implemented")
}
func (UnimplementedMailServiceServer) SendPlainTextMail(context.Context, *MailRequest) (*MailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPlainTextMail not implemented")
}
func (UnimplementedMailServiceServer) mustEmbedUnimplementedMailServiceServer() {}
func (UnimplementedMailServiceServer) testEmbeddedByValue()                     {}

// UnsafeMailServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailServiceServer will
// result in compilation errors.
type UnsafeMailServiceServer interface {
	mustEmbedUnimplementedMailServiceServer()
}

func RegisterMailServiceServer(s grpc.ServiceRegistrar, srv MailServiceServer) {
	// If the following call pancis, it indicates UnimplementedMailServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MailService_ServiceDesc, srv)
}

func _MailService_SendVerificationCodeMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendVerificationCodeMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailService_SendVerificationCodeMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendVerificationCodeMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailService_SendResetPasswordMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendResetPasswordMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailService_SendResetPasswordMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendResetPasswordMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MailService_SendPlainTextMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailServiceServer).SendPlainTextMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MailService_SendPlainTextMail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailServiceServer).SendPlainTextMail(ctx, req.(*MailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailService_ServiceDesc is the grpc.ServiceDesc for MailService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MailService",
	HandlerType: (*MailServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendVerificationCodeMail",
			Handler:    _MailService_SendVerificationCodeMail_Handler,
		},
		{
			MethodName: "SendResetPasswordMail",
			Handler:    _MailService_SendResetPasswordMail_Handler,
		},
		{
			MethodName: "SendPlainTextMail",
			Handler:    _MailService_SendPlainTextMail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/mail_proto/mail_server.proto",
}
