package main

import (
	pb "github.com/go-projects/tag-service/proto"
	"google.golang.org/grpc"
	"log"
	"context"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, "localhost:50051",  []grpc.DialOption{grpc.WithBlock()})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}