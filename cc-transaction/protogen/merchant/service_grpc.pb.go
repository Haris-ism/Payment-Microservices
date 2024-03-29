// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/merchant/service.proto

package merchant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MerchantServicesClient is the client API for MerchantServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MerchantServicesClient interface {
	InquiryItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InquiryMerchantItemsModel, error)
	InquiryDiscounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InquiryMerchantDiscountsModel, error)
	TransItems(ctx context.Context, in *ReqTransItemsModel, opts ...grpc.CallOption) (*ResMerchantTransModel, error)
}

type merchantServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewMerchantServicesClient(cc grpc.ClientConnInterface) MerchantServicesClient {
	return &merchantServicesClient{cc}
}

func (c *merchantServicesClient) InquiryItems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InquiryMerchantItemsModel, error) {
	out := new(InquiryMerchantItemsModel)
	err := c.cc.Invoke(ctx, "/MerchantServices/InquiryItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServicesClient) InquiryDiscounts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*InquiryMerchantDiscountsModel, error) {
	out := new(InquiryMerchantDiscountsModel)
	err := c.cc.Invoke(ctx, "/MerchantServices/InquiryDiscounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *merchantServicesClient) TransItems(ctx context.Context, in *ReqTransItemsModel, opts ...grpc.CallOption) (*ResMerchantTransModel, error) {
	out := new(ResMerchantTransModel)
	err := c.cc.Invoke(ctx, "/MerchantServices/TransItems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MerchantServicesServer is the server API for MerchantServices service.
// All implementations must embed UnimplementedMerchantServicesServer
// for forward compatibility
type MerchantServicesServer interface {
	InquiryItems(context.Context, *emptypb.Empty) (*InquiryMerchantItemsModel, error)
	InquiryDiscounts(context.Context, *emptypb.Empty) (*InquiryMerchantDiscountsModel, error)
	TransItems(context.Context, *ReqTransItemsModel) (*ResMerchantTransModel, error)
	mustEmbedUnimplementedMerchantServicesServer()
}

// UnimplementedMerchantServicesServer must be embedded to have forward compatible implementations.
type UnimplementedMerchantServicesServer struct {
}

func (UnimplementedMerchantServicesServer) InquiryItems(context.Context, *emptypb.Empty) (*InquiryMerchantItemsModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryItems not implemented")
}
func (UnimplementedMerchantServicesServer) InquiryDiscounts(context.Context, *emptypb.Empty) (*InquiryMerchantDiscountsModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InquiryDiscounts not implemented")
}
func (UnimplementedMerchantServicesServer) TransItems(context.Context, *ReqTransItemsModel) (*ResMerchantTransModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransItems not implemented")
}
func (UnimplementedMerchantServicesServer) mustEmbedUnimplementedMerchantServicesServer() {}

// UnsafeMerchantServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MerchantServicesServer will
// result in compilation errors.
type UnsafeMerchantServicesServer interface {
	mustEmbedUnimplementedMerchantServicesServer()
}

func RegisterMerchantServicesServer(s grpc.ServiceRegistrar, srv MerchantServicesServer) {
	s.RegisterService(&MerchantServices_ServiceDesc, srv)
}

func _MerchantServices_InquiryItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServicesServer).InquiryItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantServices/InquiryItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServicesServer).InquiryItems(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantServices_InquiryDiscounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServicesServer).InquiryDiscounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantServices/InquiryDiscounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServicesServer).InquiryDiscounts(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MerchantServices_TransItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqTransItemsModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MerchantServicesServer).TransItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MerchantServices/TransItems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MerchantServicesServer).TransItems(ctx, req.(*ReqTransItemsModel))
	}
	return interceptor(ctx, in, info, handler)
}

// MerchantServices_ServiceDesc is the grpc.ServiceDesc for MerchantServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MerchantServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MerchantServices",
	HandlerType: (*MerchantServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InquiryItems",
			Handler:    _MerchantServices_InquiryItems_Handler,
		},
		{
			MethodName: "InquiryDiscounts",
			Handler:    _MerchantServices_InquiryDiscounts_Handler,
		},
		{
			MethodName: "TransItems",
			Handler:    _MerchantServices_TransItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/merchant/service.proto",
}
