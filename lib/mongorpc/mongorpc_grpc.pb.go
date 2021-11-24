// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mongorpc

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

// MongoRPCClient is the client API for MongoRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongoRPCClient interface {
	GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Value, error)
	InsertDocument(ctx context.Context, in *InsertDocumentRequest, opts ...grpc.CallOption) (*ObjectId, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Empty, error)
}

type mongoRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoRPCClient(cc grpc.ClientConnInterface) MongoRPCClient {
	return &mongoRPCClient{cc}
}

func (c *mongoRPCClient) GetDocument(ctx context.Context, in *GetDocumentRequest, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/GetDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) InsertDocument(ctx context.Context, in *InsertDocumentRequest, opts ...grpc.CallOption) (*ObjectId, error) {
	out := new(ObjectId)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/InsertDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoRPCClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/mongorpc.MongoRPC/DeleteDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoRPCServer is the server API for MongoRPC service.
// All implementations must embed UnimplementedMongoRPCServer
// for forward compatibility
type MongoRPCServer interface {
	GetDocument(context.Context, *GetDocumentRequest) (*Value, error)
	InsertDocument(context.Context, *InsertDocumentRequest) (*ObjectId, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*Empty, error)
	mustEmbedUnimplementedMongoRPCServer()
}

// UnimplementedMongoRPCServer must be embedded to have forward compatible implementations.
type UnimplementedMongoRPCServer struct {
}

func (UnimplementedMongoRPCServer) GetDocument(context.Context, *GetDocumentRequest) (*Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDocument not implemented")
}
func (UnimplementedMongoRPCServer) InsertDocument(context.Context, *InsertDocumentRequest) (*ObjectId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertDocument not implemented")
}
func (UnimplementedMongoRPCServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDocument not implemented")
}
func (UnimplementedMongoRPCServer) mustEmbedUnimplementedMongoRPCServer() {}

// UnsafeMongoRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongoRPCServer will
// result in compilation errors.
type UnsafeMongoRPCServer interface {
	mustEmbedUnimplementedMongoRPCServer()
}

func RegisterMongoRPCServer(s grpc.ServiceRegistrar, srv MongoRPCServer) {
	s.RegisterService(&MongoRPC_ServiceDesc, srv)
}

func _MongoRPC_GetDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).GetDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/GetDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).GetDocument(ctx, req.(*GetDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_InsertDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).InsertDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/InsertDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).InsertDocument(ctx, req.(*InsertDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoRPC_DeleteDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoRPCServer).DeleteDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongorpc.MongoRPC/DeleteDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoRPCServer).DeleteDocument(ctx, req.(*DeleteDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MongoRPC_ServiceDesc is the grpc.ServiceDesc for MongoRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongoRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongorpc.MongoRPC",
	HandlerType: (*MongoRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDocument",
			Handler:    _MongoRPC_GetDocument_Handler,
		},
		{
			MethodName: "InsertDocument",
			Handler:    _MongoRPC_InsertDocument_Handler,
		},
		{
			MethodName: "DeleteDocument",
			Handler:    _MongoRPC_DeleteDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mongorpc/mongorpc.proto",
}
