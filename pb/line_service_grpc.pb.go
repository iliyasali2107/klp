// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: line_service.proto

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

// SubscribeServiceClient is the client API for SubscribeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscribeServiceClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (SubscribeService_SubscribeClient, error)
}

type subscribeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribeServiceClient(cc grpc.ClientConnInterface) SubscribeServiceClient {
	return &subscribeServiceClient{cc}
}

func (c *subscribeServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (SubscribeService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscribeService_ServiceDesc.Streams[0], "/klp.SubscribeService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscribeServiceSubscribeClient{stream}
	return x, nil
}

type SubscribeService_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type subscribeServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *subscribeServiceSubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *subscribeServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscribeServiceServer is the server API for SubscribeService service.
// All implementations must embed UnimplementedSubscribeServiceServer
// for forward compatibility
type SubscribeServiceServer interface {
	Subscribe(SubscribeService_SubscribeServer) error
	mustEmbedUnimplementedSubscribeServiceServer()
}

// UnimplementedSubscribeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscribeServiceServer struct {
}

func (UnimplementedSubscribeServiceServer) Subscribe(SubscribeService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscribeServiceServer) mustEmbedUnimplementedSubscribeServiceServer() {}

// UnsafeSubscribeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscribeServiceServer will
// result in compilation errors.
type UnsafeSubscribeServiceServer interface {
	mustEmbedUnimplementedSubscribeServiceServer()
}

func RegisterSubscribeServiceServer(s grpc.ServiceRegistrar, srv SubscribeServiceServer) {
	s.RegisterService(&SubscribeService_ServiceDesc, srv)
}

func _SubscribeService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SubscribeServiceServer).Subscribe(&subscribeServiceSubscribeServer{stream})
}

type SubscribeService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type subscribeServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *subscribeServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *subscribeServiceSubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SubscribeService_ServiceDesc is the grpc.ServiceDesc for SubscribeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscribeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "klp.SubscribeService",
	HandlerType: (*SubscribeServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _SubscribeService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "line_service.proto",
}
