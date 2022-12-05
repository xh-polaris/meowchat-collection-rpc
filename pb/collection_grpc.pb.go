// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: collection.proto

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

// CatClient is the client API for Cat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatClient interface {
	GetManyCat(ctx context.Context, in *GetManyCatReq, opts ...grpc.CallOption) (*GetManyCatResp, error)
	GetCat(ctx context.Context, in *GetCatReq, opts ...grpc.CallOption) (*GetCatResp, error)
	AddCat(ctx context.Context, in *AddCatReq, opts ...grpc.CallOption) (*AddCatResp, error)
	UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error)
	DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error)
}

type catClient struct {
	cc grpc.ClientConnInterface
}

func NewCatClient(cc grpc.ClientConnInterface) CatClient {
	return &catClient{cc}
}

func (c *catClient) GetManyCat(ctx context.Context, in *GetManyCatReq, opts ...grpc.CallOption) (*GetManyCatResp, error) {
	out := new(GetManyCatResp)
	err := c.cc.Invoke(ctx, "/cat.cat/GetManyCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catClient) GetCat(ctx context.Context, in *GetCatReq, opts ...grpc.CallOption) (*GetCatResp, error) {
	out := new(GetCatResp)
	err := c.cc.Invoke(ctx, "/cat.cat/GetCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catClient) AddCat(ctx context.Context, in *AddCatReq, opts ...grpc.CallOption) (*AddCatResp, error) {
	out := new(AddCatResp)
	err := c.cc.Invoke(ctx, "/cat.cat/AddCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catClient) UpdateCat(ctx context.Context, in *UpdateCatReq, opts ...grpc.CallOption) (*UpdateCatResp, error) {
	out := new(UpdateCatResp)
	err := c.cc.Invoke(ctx, "/cat.cat/UpdateCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *catClient) DeleteCat(ctx context.Context, in *DeleteCatReq, opts ...grpc.CallOption) (*DeleteCatResp, error) {
	out := new(DeleteCatResp)
	err := c.cc.Invoke(ctx, "/cat.cat/DeleteCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServer is the server API for Cat service.
// All implementations must embed UnimplementedCatServer
// for forward compatibility
type CatServer interface {
	GetManyCat(context.Context, *GetManyCatReq) (*GetManyCatResp, error)
	GetCat(context.Context, *GetCatReq) (*GetCatResp, error)
	AddCat(context.Context, *AddCatReq) (*AddCatResp, error)
	UpdateCat(context.Context, *UpdateCatReq) (*UpdateCatResp, error)
	DeleteCat(context.Context, *DeleteCatReq) (*DeleteCatResp, error)
	mustEmbedUnimplementedCatServer()
}

// UnimplementedCatServer must be embedded to have forward compatible implementations.
type UnimplementedCatServer struct {
}

func (UnimplementedCatServer) GetManyCat(context.Context, *GetManyCatReq) (*GetManyCatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyCat not implemented")
}
func (UnimplementedCatServer) GetCat(context.Context, *GetCatReq) (*GetCatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCat not implemented")
}
func (UnimplementedCatServer) AddCat(context.Context, *AddCatReq) (*AddCatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCat not implemented")
}
func (UnimplementedCatServer) UpdateCat(context.Context, *UpdateCatReq) (*UpdateCatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCat not implemented")
}
func (UnimplementedCatServer) DeleteCat(context.Context, *DeleteCatReq) (*DeleteCatResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCat not implemented")
}
func (UnimplementedCatServer) mustEmbedUnimplementedCatServer() {}

// UnsafeCatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatServer will
// result in compilation errors.
type UnsafeCatServer interface {
	mustEmbedUnimplementedCatServer()
}

func RegisterCatServer(s grpc.ServiceRegistrar, srv CatServer) {
	s.RegisterService(&Cat_ServiceDesc, srv)
}

func _Cat_GetManyCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetManyCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cat.cat/GetManyCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetManyCat(ctx, req.(*GetManyCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cat_GetCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cat.cat/GetCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetCat(ctx, req.(*GetCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cat_AddCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).AddCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cat.cat/AddCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).AddCat(ctx, req.(*AddCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cat_UpdateCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).UpdateCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cat.cat/UpdateCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).UpdateCat(ctx, req.(*UpdateCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cat_DeleteCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).DeleteCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cat.cat/DeleteCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).DeleteCat(ctx, req.(*DeleteCatReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Cat_ServiceDesc is the grpc.ServiceDesc for Cat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cat.cat",
	HandlerType: (*CatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetManyCat",
			Handler:    _Cat_GetManyCat_Handler,
		},
		{
			MethodName: "GetCat",
			Handler:    _Cat_GetCat_Handler,
		},
		{
			MethodName: "AddCat",
			Handler:    _Cat_AddCat_Handler,
		},
		{
			MethodName: "UpdateCat",
			Handler:    _Cat_UpdateCat_Handler,
		},
		{
			MethodName: "DeleteCat",
			Handler:    _Cat_DeleteCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection.proto",
}