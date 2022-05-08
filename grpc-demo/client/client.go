package main

import (
	"context"
	"log"
	"sync"

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

	//第三种，出参为流
	clientSearchOut(client)

	//第四种，出入为流
	clientSearchIO(client)

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
		log.Printf("client.SearchIn err: %v", err)
	}
}

func clientSearchOut(client pd.SearchServiceClient) {
	req := &pd.SearchRequest{
		Request: "clientSearchOut",
	}
	resp, err := client.SearchOut(context.Background(), req)
	if err != nil {
		log.Fatalf("resp.Recv() err: %v", err)
	}

	for {
		r, err := resp.Recv()
		if err != nil {
			log.Printf("client.SearchOut err: %v", err)
			break
		}
		log.Printf("SearchOut req: %+v", r)
	}
}

func clientSearchIO(client pd.SearchServiceClient) {
	resp, err := client.SearchIO(context.Background())
	if err != nil {
		log.Fatalf("resp.Recv() err: %v", err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			err := resp.Send(&pd.SearchRequest{
				Request: "SearchIO",
			})
			if err != nil {
				wg.Done()
				break
			}
		}
	}()

	go func() {
		for {
			r, err := resp.Recv()
			if err != nil {
				wg.Done()
				log.Fatalf("resp.Recv() err: %v", err)
			}
			log.Printf("SearchIO req: %+v", r)
		}
	}()
	wg.Wait()
}
