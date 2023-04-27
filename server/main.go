package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	// this struct allows us to use methods such as ListFiles
	// add methods and fields that we define to server struct
	pb.UnimplementedFilesServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked")
	dir := "/Users/user/workspace/grpc-lesson/storage/"

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

func (*server) Download(req *pb.DownloadRequest, stream pb.FilesService_DownloadServer) error {
	fmt.Println("Download was invoked")

	filename := req.GetFilename()
	path := "/Users/user/workspace/grpc-lesson/storage/" + filename

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// prepare for byte type whose length is 5
	buf := make([]byte, 5)
	for {
		// read file content
		// n = byte value
		n, err := file.Read(buf)
		// if nothing was read(0) or buf was read till the end
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// send data to client as response
		res := &pb.DownloadResponse{Data: buf[:n]}
		sendErr := stream.Send(res)
		if sendErr != nil {
			return sendErr
		}

		// solution for process not to finish 1 sec
		time.Sleep(1 * time.Second)

	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// get structure of grpc
	s := grpc.NewServer()

	// register structure content at grpc, which means grpc can provide methods we define
	// this method is created automatically by protocol buffer's compile
	pb.RegisterFilesServiceServer(s, &server{})

	fmt.Println("server is running")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
