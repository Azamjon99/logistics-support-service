// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: service.proto

package logistic_support

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	GetRatingByOrder(ctx context.Context, in *GetRatingByOrderRequest, opts ...grpc.CallOption) (*RatingResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/RatingService/CreateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/RatingService/UpdateRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/RatingService/GetRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatingByOrder(ctx context.Context, in *GetRatingByOrderRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, "/RatingService/GetRatingByOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	CreateRating(context.Context, *CreateRatingRequest) (*RatingResponse, error)
	UpdateRating(context.Context, *UpdateRatingRequest) (*RatingResponse, error)
	GetRating(context.Context, *GetRatingRequest) (*RatingResponse, error)
	GetRatingByOrder(context.Context, *GetRatingByOrderRequest) (*RatingResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) CreateRating(context.Context, *CreateRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRating not implemented")
}
func (UnimplementedRatingServiceServer) UpdateRating(context.Context, *UpdateRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (UnimplementedRatingServiceServer) GetRating(context.Context, *GetRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRating not implemented")
}
func (UnimplementedRatingServiceServer) GetRatingByOrder(context.Context, *GetRatingByOrderRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingByOrder not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_CreateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).CreateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RatingService/CreateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).CreateRating(ctx, req.(*CreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RatingService/UpdateRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).UpdateRating(ctx, req.(*UpdateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RatingService/GetRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRating(ctx, req.(*GetRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatingByOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingByOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingByOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RatingService/GetRatingByOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingByOrder(ctx, req.(*GetRatingByOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRating",
			Handler:    _RatingService_CreateRating_Handler,
		},
		{
			MethodName: "UpdateRating",
			Handler:    _RatingService_UpdateRating_Handler,
		},
		{
			MethodName: "GetRating",
			Handler:    _RatingService_GetRating_Handler,
		},
		{
			MethodName: "GetRatingByOrder",
			Handler:    _RatingService_GetRatingByOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}