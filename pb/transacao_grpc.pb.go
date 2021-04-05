// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// QixClient is the client API for Qix service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QixClient interface {
	EnviaQix(ctx context.Context, in *QixRequest, opts ...grpc.CallOption) (*QixReply, error)
}

type qixClient struct {
	cc grpc.ClientConnInterface
}

func NewQixClient(cc grpc.ClientConnInterface) QixClient {
	return &qixClient{cc}
}

func (c *qixClient) EnviaQix(ctx context.Context, in *QixRequest, opts ...grpc.CallOption) (*QixReply, error) {
	out := new(QixReply)
	err := c.cc.Invoke(ctx, "/qix2.Qix/EnviaQix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QixServer is the server API for Qix service.
// All implementations must embed UnimplementedQixServer
// for forward compatibility
type QixServer interface {
	EnviaQix(context.Context, *QixRequest) (*QixReply, error)
	mustEmbedUnimplementedQixServer()
}

// UnimplementedQixServer must be embedded to have forward compatible implementations.
type UnimplementedQixServer struct {
}

func (UnimplementedQixServer) EnviaQix(context.Context, *QixRequest) (*QixReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviaQix not implemented")
}
func (UnimplementedQixServer) mustEmbedUnimplementedQixServer() {}

// UnsafeQixServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QixServer will
// result in compilation errors.
type UnsafeQixServer interface {
	mustEmbedUnimplementedQixServer()
}

func RegisterQixServer(s grpc.ServiceRegistrar, srv QixServer) {
	s.RegisterService(&Qix_ServiceDesc, srv)
}

func _Qix_EnviaQix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QixServer).EnviaQix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qix2.Qix/EnviaQix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QixServer).EnviaQix(ctx, req.(*QixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Qix_ServiceDesc is the grpc.ServiceDesc for Qix service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Qix_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qix2.Qix",
	HandlerType: (*QixServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviaQix",
			Handler:    _Qix_EnviaQix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transacao.proto",
}

// LoadClient is the client API for Load service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoadClient interface {
	EnviaLoad(ctx context.Context, in *LoadInfo, opts ...grpc.CallOption) (*LoadRecebido, error)
}

type loadClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadClient(cc grpc.ClientConnInterface) LoadClient {
	return &loadClient{cc}
}

func (c *loadClient) EnviaLoad(ctx context.Context, in *LoadInfo, opts ...grpc.CallOption) (*LoadRecebido, error) {
	out := new(LoadRecebido)
	err := c.cc.Invoke(ctx, "/qix2.Load/EnviaLoad", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadServer is the server API for Load service.
// All implementations must embed UnimplementedLoadServer
// for forward compatibility
type LoadServer interface {
	EnviaLoad(context.Context, *LoadInfo) (*LoadRecebido, error)
	mustEmbedUnimplementedLoadServer()
}

// UnimplementedLoadServer must be embedded to have forward compatible implementations.
type UnimplementedLoadServer struct {
}

func (UnimplementedLoadServer) EnviaLoad(context.Context, *LoadInfo) (*LoadRecebido, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviaLoad not implemented")
}
func (UnimplementedLoadServer) mustEmbedUnimplementedLoadServer() {}

// UnsafeLoadServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoadServer will
// result in compilation errors.
type UnsafeLoadServer interface {
	mustEmbedUnimplementedLoadServer()
}

func RegisterLoadServer(s grpc.ServiceRegistrar, srv LoadServer) {
	s.RegisterService(&Load_ServiceDesc, srv)
}

func _Load_EnviaLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadServer).EnviaLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qix2.Load/EnviaLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadServer).EnviaLoad(ctx, req.(*LoadInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Load_ServiceDesc is the grpc.ServiceDesc for Load service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Load_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qix2.Load",
	HandlerType: (*LoadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviaLoad",
			Handler:    _Load_EnviaLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/transacao.proto",
}