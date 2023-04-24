package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func callListFiles(client pb.FilesServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())
}

func callDownload(client pb.FilesServiceClient) {
	req := &pb.DownloadRequest{Filename: "name.txt"}
	stream, err := client.Download(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// always accept it every time server returns response
		res, err := stream.Recv()
		// when server reaching EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("Response from Download(Bytes): %v", res.GetData())
		log.Printf("Response from Download(string): %v", string(res.GetData()))
	}

}

func main() {
	// establish connection with localhost:50051
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	// always finish connection
	defer conn.Close()

	// get file service client
	// this method is created automatically by protocol buffer's compile
	// this func name includes service in proto file

	// func NewFilesServiceClient(cc grpc.ClientConnInterface) FilesServiceClient {
	// 	return &filesServiceClient{cc}
	// }
	// this returns interface FilesServiceClient
	// this client is FilesServiceClient type
	client := pb.NewFilesServiceClient(conn)
	// callListFiles(client)
	callDownload(client)

}
