package main

import (
	"context"
	"log"
	"net"

	pd "grpc-demo/proto"

	"google.golang.org/grpc"
)

type SearchService struct {
	pd.UnsafeSearchServiceServer
}

func (s *SearchService) Search(ctx context.Context, req *pd.SearchRequest) (*pd.SearchResponse, error) {
	log.Printf("req:%+v", req)
	return &pd.SearchResponse{Response: req.GetRequest() + " Server"}, nil
}

const PORT = "9001"

func main() {

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server := grpc.NewServer()
	pd.RegisterSearchServiceServer(server, &SearchService{})

	server.Serve(lis)
}
