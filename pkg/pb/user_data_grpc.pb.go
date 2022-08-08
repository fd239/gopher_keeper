// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: user_data.proto

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

// UserDataServiceClient is the client API for UserDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserDataServiceClient interface {
	SaveText(ctx context.Context, in *SaveTextRequest, opts ...grpc.CallOption) (*SaveTextResponse, error)
	GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error)
	SaveCard(ctx context.Context, in *SaveCardRequest, opts ...grpc.CallOption) (*SaveCardResponse, error)
	GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error)
	SaveFile(ctx context.Context, opts ...grpc.CallOption) (UserDataService_SaveFileClient, error)
}

type userDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserDataServiceClient(cc grpc.ClientConnInterface) UserDataServiceClient {
	return &userDataServiceClient{cc}
}

func (c *userDataServiceClient) SaveText(ctx context.Context, in *SaveTextRequest, opts ...grpc.CallOption) (*SaveTextResponse, error) {
	out := new(SaveTextResponse)
	err := c.cc.Invoke(ctx, "/api.UserDataService/SaveText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) GetText(ctx context.Context, in *GetTextRequest, opts ...grpc.CallOption) (*GetTextResponse, error) {
	out := new(GetTextResponse)
	err := c.cc.Invoke(ctx, "/api.UserDataService/GetText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) SaveCard(ctx context.Context, in *SaveCardRequest, opts ...grpc.CallOption) (*SaveCardResponse, error) {
	out := new(SaveCardResponse)
	err := c.cc.Invoke(ctx, "/api.UserDataService/SaveCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) GetCard(ctx context.Context, in *GetCardRequest, opts ...grpc.CallOption) (*GetCardResponse, error) {
	out := new(GetCardResponse)
	err := c.cc.Invoke(ctx, "/api.UserDataService/GetCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userDataServiceClient) SaveFile(ctx context.Context, opts ...grpc.CallOption) (UserDataService_SaveFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserDataService_ServiceDesc.Streams[0], "/api.UserDataService/SaveFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &userDataServiceSaveFileClient{stream}
	return x, nil
}

type UserDataService_SaveFileClient interface {
	Send(*FileRequest) error
	CloseAndRecv() (*FileResponse, error)
	grpc.ClientStream
}

type userDataServiceSaveFileClient struct {
	grpc.ClientStream
}

func (x *userDataServiceSaveFileClient) Send(m *FileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userDataServiceSaveFileClient) CloseAndRecv() (*FileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserDataServiceServer is the server API for UserDataService service.
// All implementations must embed UnimplementedUserDataServiceServer
// for forward compatibility
type UserDataServiceServer interface {
	SaveText(context.Context, *SaveTextRequest) (*SaveTextResponse, error)
	GetText(context.Context, *GetTextRequest) (*GetTextResponse, error)
	SaveCard(context.Context, *SaveCardRequest) (*SaveCardResponse, error)
	GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error)
	SaveFile(UserDataService_SaveFileServer) error
	mustEmbedUnimplementedUserDataServiceServer()
}

// UnimplementedUserDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserDataServiceServer struct {
}

func (UnimplementedUserDataServiceServer) SaveText(context.Context, *SaveTextRequest) (*SaveTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveText not implemented")
}
func (UnimplementedUserDataServiceServer) GetText(context.Context, *GetTextRequest) (*GetTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetText not implemented")
}
func (UnimplementedUserDataServiceServer) SaveCard(context.Context, *SaveCardRequest) (*SaveCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveCard not implemented")
}
func (UnimplementedUserDataServiceServer) GetCard(context.Context, *GetCardRequest) (*GetCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCard not implemented")
}
func (UnimplementedUserDataServiceServer) SaveFile(UserDataService_SaveFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveFile not implemented")
}
func (UnimplementedUserDataServiceServer) mustEmbedUnimplementedUserDataServiceServer() {}

// UnsafeUserDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserDataServiceServer will
// result in compilation errors.
type UnsafeUserDataServiceServer interface {
	mustEmbedUnimplementedUserDataServiceServer()
}

func RegisterUserDataServiceServer(s grpc.ServiceRegistrar, srv UserDataServiceServer) {
	s.RegisterService(&UserDataService_ServiceDesc, srv)
}

func _UserDataService_SaveText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).SaveText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserDataService/SaveText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).SaveText(ctx, req.(*SaveTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_GetText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).GetText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserDataService/GetText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).GetText(ctx, req.(*GetTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_SaveCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).SaveCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserDataService/SaveCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).SaveCard(ctx, req.(*SaveCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_GetCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserDataServiceServer).GetCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.UserDataService/GetCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserDataServiceServer).GetCard(ctx, req.(*GetCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserDataService_SaveFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserDataServiceServer).SaveFile(&userDataServiceSaveFileServer{stream})
}

type UserDataService_SaveFileServer interface {
	SendAndClose(*FileResponse) error
	Recv() (*FileRequest, error)
	grpc.ServerStream
}

type userDataServiceSaveFileServer struct {
	grpc.ServerStream
}

func (x *userDataServiceSaveFileServer) SendAndClose(m *FileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userDataServiceSaveFileServer) Recv() (*FileRequest, error) {
	m := new(FileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UserDataService_ServiceDesc is the grpc.ServiceDesc for UserDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UserDataService",
	HandlerType: (*UserDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveText",
			Handler:    _UserDataService_SaveText_Handler,
		},
		{
			MethodName: "GetText",
			Handler:    _UserDataService_GetText_Handler,
		},
		{
			MethodName: "SaveCard",
			Handler:    _UserDataService_SaveCard_Handler,
		},
		{
			MethodName: "GetCard",
			Handler:    _UserDataService_GetCard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveFile",
			Handler:       _UserDataService_SaveFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user_data.proto",
}
