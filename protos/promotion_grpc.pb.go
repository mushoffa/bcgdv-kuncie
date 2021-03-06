// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PromotionServiceClient is the client API for PromotionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PromotionServiceClient interface {
	CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error)
	GetAllPromotions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllPromotionsResponse, error)
	UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*UpdatePromotionResponse, error)
	DeletePromotion(ctx context.Context, in *DeletePromotionRequest, opts ...grpc.CallOption) (*DeletePromotionResponse, error)
	CheckPromotion(ctx context.Context, in *CheckPromotionRequest, opts ...grpc.CallOption) (*CheckPromotionResponse, error)
}

type promotionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPromotionServiceClient(cc grpc.ClientConnInterface) PromotionServiceClient {
	return &promotionServiceClient{cc}
}

func (c *promotionServiceClient) CreatePromotion(ctx context.Context, in *CreatePromotionRequest, opts ...grpc.CallOption) (*CreatePromotionResponse, error) {
	out := new(CreatePromotionResponse)
	err := c.cc.Invoke(ctx, "/PromotionService/CreatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) GetAllPromotions(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllPromotionsResponse, error) {
	out := new(GetAllPromotionsResponse)
	err := c.cc.Invoke(ctx, "/PromotionService/GetAllPromotions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) UpdatePromotion(ctx context.Context, in *UpdatePromotionRequest, opts ...grpc.CallOption) (*UpdatePromotionResponse, error) {
	out := new(UpdatePromotionResponse)
	err := c.cc.Invoke(ctx, "/PromotionService/UpdatePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) DeletePromotion(ctx context.Context, in *DeletePromotionRequest, opts ...grpc.CallOption) (*DeletePromotionResponse, error) {
	out := new(DeletePromotionResponse)
	err := c.cc.Invoke(ctx, "/PromotionService/DeletePromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *promotionServiceClient) CheckPromotion(ctx context.Context, in *CheckPromotionRequest, opts ...grpc.CallOption) (*CheckPromotionResponse, error) {
	out := new(CheckPromotionResponse)
	err := c.cc.Invoke(ctx, "/PromotionService/CheckPromotion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PromotionServiceServer is the server API for PromotionService service.
// All implementations must embed UnimplementedPromotionServiceServer
// for forward compatibility
type PromotionServiceServer interface {
	CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error)
	GetAllPromotions(context.Context, *empty.Empty) (*GetAllPromotionsResponse, error)
	UpdatePromotion(context.Context, *UpdatePromotionRequest) (*UpdatePromotionResponse, error)
	DeletePromotion(context.Context, *DeletePromotionRequest) (*DeletePromotionResponse, error)
	CheckPromotion(context.Context, *CheckPromotionRequest) (*CheckPromotionResponse, error)
	mustEmbedUnimplementedPromotionServiceServer()
}

// UnimplementedPromotionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPromotionServiceServer struct {
}

func (UnimplementedPromotionServiceServer) CreatePromotion(context.Context, *CreatePromotionRequest) (*CreatePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePromotion not implemented")
}
func (UnimplementedPromotionServiceServer) GetAllPromotions(context.Context, *empty.Empty) (*GetAllPromotionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPromotions not implemented")
}
func (UnimplementedPromotionServiceServer) UpdatePromotion(context.Context, *UpdatePromotionRequest) (*UpdatePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePromotion not implemented")
}
func (UnimplementedPromotionServiceServer) DeletePromotion(context.Context, *DeletePromotionRequest) (*DeletePromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePromotion not implemented")
}
func (UnimplementedPromotionServiceServer) CheckPromotion(context.Context, *CheckPromotionRequest) (*CheckPromotionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPromotion not implemented")
}
func (UnimplementedPromotionServiceServer) mustEmbedUnimplementedPromotionServiceServer() {}

// UnsafePromotionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PromotionServiceServer will
// result in compilation errors.
type UnsafePromotionServiceServer interface {
	mustEmbedUnimplementedPromotionServiceServer()
}

func RegisterPromotionServiceServer(s grpc.ServiceRegistrar, srv PromotionServiceServer) {
	s.RegisterService(&PromotionService_ServiceDesc, srv)
}

func _PromotionService_CreatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).CreatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PromotionService/CreatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).CreatePromotion(ctx, req.(*CreatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_GetAllPromotions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).GetAllPromotions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PromotionService/GetAllPromotions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).GetAllPromotions(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_UpdatePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).UpdatePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PromotionService/UpdatePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).UpdatePromotion(ctx, req.(*UpdatePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_DeletePromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).DeletePromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PromotionService/DeletePromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).DeletePromotion(ctx, req.(*DeletePromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PromotionService_CheckPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPromotionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PromotionServiceServer).CheckPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PromotionService/CheckPromotion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PromotionServiceServer).CheckPromotion(ctx, req.(*CheckPromotionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PromotionService_ServiceDesc is the grpc.ServiceDesc for PromotionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PromotionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PromotionService",
	HandlerType: (*PromotionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePromotion",
			Handler:    _PromotionService_CreatePromotion_Handler,
		},
		{
			MethodName: "GetAllPromotions",
			Handler:    _PromotionService_GetAllPromotions_Handler,
		},
		{
			MethodName: "UpdatePromotion",
			Handler:    _PromotionService_UpdatePromotion_Handler,
		},
		{
			MethodName: "DeletePromotion",
			Handler:    _PromotionService_DeletePromotion_Handler,
		},
		{
			MethodName: "CheckPromotion",
			Handler:    _PromotionService_CheckPromotion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "promotion.proto",
}
