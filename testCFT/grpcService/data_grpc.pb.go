// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: data.proto

package grpcService

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

// CftServiceClient is the client API for CftService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CftServiceClient interface {
	GetFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	GetFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error)
	CreateFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error)
	UpdateFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error)
	DeleteFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error)
}

type cftServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCftServiceClient(cc grpc.ClientConnInterface) CftServiceClient {
	return &cftServiceClient{cc}
}

func (c *cftServiceClient) GetFiles(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcService.CftService/GetFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cftServiceClient) GetFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcService.CftService/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cftServiceClient) CreateFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcService.CftService/CreateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cftServiceClient) UpdateFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcService.CftService/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cftServiceClient) DeleteFile(ctx context.Context, in *FileWithData, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/grpcService.CftService/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CftServiceServer is the server API for CftService service.
// All implementations must embed UnimplementedCftServiceServer
// for forward compatibility
type CftServiceServer interface {
	GetFiles(context.Context, *Empty) (*Response, error)
	GetFile(context.Context, *FileWithData) (*Response, error)
	CreateFile(context.Context, *FileWithData) (*Response, error)
	UpdateFile(context.Context, *FileWithData) (*Response, error)
	DeleteFile(context.Context, *FileWithData) (*Response, error)
	mustEmbedUnimplementedCftServiceServer()
}

// UnimplementedCftServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCftServiceServer struct {
}

func (UnimplementedCftServiceServer) GetFiles(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFiles not implemented")
}
func (UnimplementedCftServiceServer) GetFile(context.Context, *FileWithData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedCftServiceServer) CreateFile(context.Context, *FileWithData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFile not implemented")
}
func (UnimplementedCftServiceServer) UpdateFile(context.Context, *FileWithData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedCftServiceServer) DeleteFile(context.Context, *FileWithData) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedCftServiceServer) mustEmbedUnimplementedCftServiceServer() {}

// UnsafeCftServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CftServiceServer will
// result in compilation errors.
type UnsafeCftServiceServer interface {
	mustEmbedUnimplementedCftServiceServer()
}

func RegisterCftServiceServer(s grpc.ServiceRegistrar, srv CftServiceServer) {
	s.RegisterService(&CftService_ServiceDesc, srv)
}

func _CftService_GetFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServiceServer).GetFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.CftService/GetFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServiceServer).GetFiles(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CftService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileWithData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.CftService/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServiceServer).GetFile(ctx, req.(*FileWithData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CftService_CreateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileWithData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServiceServer).CreateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.CftService/CreateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServiceServer).CreateFile(ctx, req.(*FileWithData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CftService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileWithData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.CftService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServiceServer).UpdateFile(ctx, req.(*FileWithData))
	}
	return interceptor(ctx, in, info, handler)
}

func _CftService_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileWithData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CftServiceServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcService.CftService/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CftServiceServer).DeleteFile(ctx, req.(*FileWithData))
	}
	return interceptor(ctx, in, info, handler)
}

// CftService_ServiceDesc is the grpc.ServiceDesc for CftService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CftService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcService.CftService",
	HandlerType: (*CftServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFiles",
			Handler:    _CftService_GetFiles_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _CftService_GetFile_Handler,
		},
		{
			MethodName: "CreateFile",
			Handler:    _CftService_CreateFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _CftService_UpdateFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _CftService_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}
