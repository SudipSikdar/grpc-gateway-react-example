// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: service.proto

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

// GrpcServiceClient is the client API for GrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcServiceClient interface {
	GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error)
}

type grpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcServiceClient(cc grpc.ClientConnInterface) GrpcServiceClient {
	return &grpcServiceClient{cc}
}

func (c *grpcServiceClient) GetMessage(ctx context.Context, in *GetMessageRequest, opts ...grpc.CallOption) (*GetMessageResponse, error) {
	out := new(GetMessageResponse)
	err := c.cc.Invoke(ctx, "/reactgrpc.GrpcService/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcServiceServer is the server API for GrpcService service.
// All implementations must embed UnimplementedGrpcServiceServer
// for forward compatibility
type GrpcServiceServer interface {
	GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error)
	mustEmbedUnimplementedGrpcServiceServer()
}

// UnimplementedGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcServiceServer struct {
}

func (UnimplementedGrpcServiceServer) GetMessage(context.Context, *GetMessageRequest) (*GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedGrpcServiceServer) mustEmbedUnimplementedGrpcServiceServer() {}

// UnsafeGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcServiceServer will
// result in compilation errors.
type UnsafeGrpcServiceServer interface {
	mustEmbedUnimplementedGrpcServiceServer()
}

func RegisterGrpcServiceServer(s grpc.ServiceRegistrar, srv GrpcServiceServer) {
	s.RegisterService(&GrpcService_ServiceDesc, srv)
}

func _GrpcService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reactgrpc.GrpcService/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcServiceServer).GetMessage(ctx, req.(*GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GrpcService_ServiceDesc is the grpc.ServiceDesc for GrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reactgrpc.GrpcService",
	HandlerType: (*GrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _GrpcService_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
