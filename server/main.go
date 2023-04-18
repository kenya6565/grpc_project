package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io/ioutil"
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
