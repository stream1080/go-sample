package main

import (
	"context"
	"log"

	pd "grpc-demo/proto"

	"google.golang.org/grpc"
)

const PORT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	client := pd.NewSearchServiceClient(conn)

	// 第一种，常规调用
	clientSearch(client)

	// 第二种，入参为流
	clientSearchIn(client)

}

func clientSearch(client pd.SearchServiceClient) {
	resp, err := client.Search(context.Background(), &pd.SearchRequest{
		Request: "gRPC",
		Value:   1,
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}
	log.Printf("resp: %s", resp.GetResponse())
}

func clientSearchIn(client pd.SearchServiceClient) {
	resp, err := client.SearchIn(context.Background())
	for i := 1; i <= 10; i++ {
		resp.Send(&pd.SearchRequest{
			Request: "SearchIn",
			Value:   int32(i),
		})
		if i == 10 {
			c, err := resp.CloseAndRecv()
			if err != nil {
				log.Fatalf("client.CloseAndRecv() err: %v", err)
			}
			log.Printf("CloseAndRecv(): %+v", c)
		}
	}
	if err != nil {
		log.Fatalf("client.SearchIn err: %v", err)
	}
}
