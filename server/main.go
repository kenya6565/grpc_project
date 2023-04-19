package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFilesServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked")
	dir := "/Users/user/workspace/grpc-lesson/storage"

	paths, err := ioutil.ReadDir(dir)

	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range paths {
		// if path is not a directory, which is actually a file name
		if !path.IsDir() {
			// append file name to slice
			filenames = append(filenames, path.Name())
		}
	}

	// create ListFilesResponse as return val
	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// get structure of grpc
	s := grpc.NewServer()

	// register structure content at grpc, which means grpc can provide methods we define
	pb.RegisterFilesServiceServer(s, &server{})

	fmt.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
