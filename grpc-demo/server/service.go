package main

import (
	"context"
	"log"
	"net"

	pd "grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	PORT = "9001"
)

type SearchService struct {
	pd.UnimplementedSearchServiceServer
}

func main() {

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	server := grpc.NewServer()
	pd.RegisterSearchServiceServer(server, &SearchService{})

	server.Serve(lis)
}

func (s *SearchService) Search(ctx context.Context, req *pd.SearchRequest) (*pd.SearchResponse, error) {
	log.Printf("req:%+v", req)
	return &pd.SearchResponse{Response: req.GetRequest() + " Server"}, nil
}

func (s *SearchService) SearchIn(server pd.SearchService_SearchInServer) error {
	for {
		req, err := server.Recv()
		if err != nil {
			server.SendAndClose(&pd.SearchResponse{
				Response: "读取完成",
				Value:    0,
			})
			break
		}
		log.Printf("req: %+v", req)
	}
	return nil
}

func (s *SearchService) SearchOut(req *pd.SearchRequest, server pd.SearchService_SearchOutServer) error {
	for i := 1; i <= 10; i++ {
		server.Send(&pd.SearchResponse{
			Response: req.Request,
			Value:    req.Value,
		})
	}
	return nil
}

func (s *SearchService) SearchIO(servier pd.SearchService_SearchIOServer) error {
	return nil
}
