package main

import (
	"context"
	"fmt"
	"grpc-lesson/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// establish connection with server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	// always finish connection
	defer conn.Close()

	// get file service client
	client := pb.NewFilesServiceClient(conn)
	callListFiles(client)

}

func callListFiles(client pb.FilesServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res.GetFilenames())
}
