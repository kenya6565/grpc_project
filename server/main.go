package main

import (
	"context"
	"grpc-lesson/pb"
)

type server struct {
	pb.UnimplementedFilesServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {

}
