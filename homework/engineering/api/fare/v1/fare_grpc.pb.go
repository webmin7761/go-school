// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package fare

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

// FareServiceClient is the client API for FareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FareServiceClient interface {
	//新建运价
	CreateFare(ctx context.Context, in *CreateFareRequest, opts ...grpc.CallOption) (*CreateFareReply, error)
	//删除运价
	UpdateFare(ctx context.Context, in *UpdateFareRequest, opts ...grpc.CallOption) (*UpdateFareReply, error)
	//删除运价
	DeleteFare(ctx context.Context, in *DeleteFareRequest, opts ...grpc.CallOption) (*DeleteFareReply, error)
	//删除运价
	GetFare(ctx context.Context, in *GetFareRequest, opts ...grpc.CallOption) (*GetFareReply, error)
	//机票销售价格计算
	Pricing(ctx context.Context, in *PricingRequest, opts ...grpc.CallOption) (*PricingResponse, error)
}

type fareServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFareServiceClient(cc grpc.ClientConnInterface) FareServiceClient {
	return &fareServiceClient{cc}
}

func (c *fareServiceClient) CreateFare(ctx context.Context, in *CreateFareRequest, opts ...grpc.CallOption) (*CreateFareReply, error) {
	out := new(CreateFareReply)
	err := c.cc.Invoke(ctx, "/fare.FareService/CreateFare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) UpdateFare(ctx context.Context, in *UpdateFareRequest, opts ...grpc.CallOption) (*UpdateFareReply, error) {
	out := new(UpdateFareReply)
	err := c.cc.Invoke(ctx, "/fare.FareService/UpdateFare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) DeleteFare(ctx context.Context, in *DeleteFareRequest, opts ...grpc.CallOption) (*DeleteFareReply, error) {
	out := new(DeleteFareReply)
	err := c.cc.Invoke(ctx, "/fare.FareService/DeleteFare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) GetFare(ctx context.Context, in *GetFareRequest, opts ...grpc.CallOption) (*GetFareReply, error) {
	out := new(GetFareReply)
	err := c.cc.Invoke(ctx, "/fare.FareService/GetFare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fareServiceClient) Pricing(ctx context.Context, in *PricingRequest, opts ...grpc.CallOption) (*PricingResponse, error) {
	out := new(PricingResponse)
	err := c.cc.Invoke(ctx, "/fare.FareService/Pricing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FareServiceServer is the server API for FareService service.
// All implementations must embed UnimplementedFareServiceServer
// for forward compatibility
type FareServiceServer interface {
	//新建运价
	CreateFare(context.Context, *CreateFareRequest) (*CreateFareReply, error)
	//删除运价
	UpdateFare(context.Context, *UpdateFareRequest) (*UpdateFareReply, error)
	//删除运价
	DeleteFare(context.Context, *DeleteFareRequest) (*DeleteFareReply, error)
	//删除运价
	GetFare(context.Context, *GetFareRequest) (*GetFareReply, error)
	//机票销售价格计算
	Pricing(context.Context, *PricingRequest) (*PricingResponse, error)
	mustEmbedUnimplementedFareServiceServer()
}

// UnimplementedFareServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFareServiceServer struct {
}

func (UnimplementedFareServiceServer) CreateFare(context.Context, *CreateFareRequest) (*CreateFareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFare not implemented")
}
func (UnimplementedFareServiceServer) UpdateFare(context.Context, *UpdateFareRequest) (*UpdateFareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFare not implemented")
}
func (UnimplementedFareServiceServer) DeleteFare(context.Context, *DeleteFareRequest) (*DeleteFareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFare not implemented")
}
func (UnimplementedFareServiceServer) GetFare(context.Context, *GetFareRequest) (*GetFareReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFare not implemented")
}
func (UnimplementedFareServiceServer) Pricing(context.Context, *PricingRequest) (*PricingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pricing not implemented")
}
func (UnimplementedFareServiceServer) mustEmbedUnimplementedFareServiceServer() {}

// UnsafeFareServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FareServiceServer will
// result in compilation errors.
type UnsafeFareServiceServer interface {
	mustEmbedUnimplementedFareServiceServer()
}

func RegisterFareServiceServer(s grpc.ServiceRegistrar, srv FareServiceServer) {
	s.RegisterService(&FareService_ServiceDesc, srv)
}

func _FareService_CreateFare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).CreateFare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fare.FareService/CreateFare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).CreateFare(ctx, req.(*CreateFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_UpdateFare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).UpdateFare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fare.FareService/UpdateFare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).UpdateFare(ctx, req.(*UpdateFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_DeleteFare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).DeleteFare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fare.FareService/DeleteFare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).DeleteFare(ctx, req.(*DeleteFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_GetFare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).GetFare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fare.FareService/GetFare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).GetFare(ctx, req.(*GetFareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FareService_Pricing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PricingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FareServiceServer).Pricing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fare.FareService/Pricing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FareServiceServer).Pricing(ctx, req.(*PricingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FareService_ServiceDesc is the grpc.ServiceDesc for FareService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FareService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fare.FareService",
	HandlerType: (*FareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFare",
			Handler:    _FareService_CreateFare_Handler,
		},
		{
			MethodName: "UpdateFare",
			Handler:    _FareService_UpdateFare_Handler,
		},
		{
			MethodName: "DeleteFare",
			Handler:    _FareService_DeleteFare_Handler,
		},
		{
			MethodName: "GetFare",
			Handler:    _FareService_GetFare_Handler,
		},
		{
			MethodName: "Pricing",
			Handler:    _FareService_Pricing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fare.proto",
}