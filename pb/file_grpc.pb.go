// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/file.proto

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

// FilesServiceClient is the client API for FilesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FilesServiceClient interface {
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FilesService_DownloadClient, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (FilesService_UploadClient, error)
}

type filesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFilesServiceClient(cc grpc.ClientConnInterface) FilesServiceClient {
	return &filesServiceClient{cc}
}

func (c *filesServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, "/file.FilesService/ListFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *filesServiceClient) Download(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (FilesService_DownloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FilesService_ServiceDesc.Streams[0], "/file.FilesService/Download", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesServiceDownloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FilesService_DownloadClient interface {
	Recv() (*DownloadResponse, error)
	grpc.ClientStream
}

type filesServiceDownloadClient struct {
	grpc.ClientStream
}

func (x *filesServiceDownloadClient) Recv() (*DownloadResponse, error) {
	m := new(DownloadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *filesServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FilesService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &FilesService_ServiceDesc.Streams[1], "/file.FilesService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &filesServiceUploadClient{stream}
	return x, nil
}

type FilesService_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type filesServiceUploadClient struct {
	grpc.ClientStream
}

func (x *filesServiceUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *filesServiceUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilesServiceServer is the server API for FilesService service.
// All implementations must embed UnimplementedFilesServiceServer
// for forward compatibility
type FilesServiceServer interface {
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	Download(*DownloadRequest, FilesService_DownloadServer) error
	Upload(FilesService_UploadServer) error
	mustEmbedUnimplementedFilesServiceServer()
}

// UnimplementedFilesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFilesServiceServer struct {
}

func (UnimplementedFilesServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedFilesServiceServer) Download(*DownloadRequest, FilesService_DownloadServer) error {
	return status.Errorf(codes.Unimplemented, "method Download not implemented")
}
func (UnimplementedFilesServiceServer) Upload(FilesService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedFilesServiceServer) mustEmbedUnimplementedFilesServiceServer() {}

// UnsafeFilesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FilesServiceServer will
// result in compilation errors.
type UnsafeFilesServiceServer interface {
	mustEmbedUnimplementedFilesServiceServer()
}

func RegisterFilesServiceServer(s grpc.ServiceRegistrar, srv FilesServiceServer) {
	s.RegisterService(&FilesService_ServiceDesc, srv)
}

func _FilesService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FilesServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/file.FilesService/ListFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FilesServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FilesService_Download_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DownloadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FilesServiceServer).Download(m, &filesServiceDownloadServer{stream})
}

type FilesService_DownloadServer interface {
	Send(*DownloadResponse) error
	grpc.ServerStream
}

type filesServiceDownloadServer struct {
	grpc.ServerStream
}

func (x *filesServiceDownloadServer) Send(m *DownloadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FilesService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FilesServiceServer).Upload(&filesServiceUploadServer{stream})
}

type FilesService_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type filesServiceUploadServer struct {
	grpc.ServerStream
}

func (x *filesServiceUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *filesServiceUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FilesService_ServiceDesc is the grpc.ServiceDesc for FilesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FilesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "file.FilesService",
	HandlerType: (*FilesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFiles",
			Handler:    _FilesService_ListFiles_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Download",
			Handler:       _FilesService_Download_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upload",
			Handler:       _FilesService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/file.proto",
}
