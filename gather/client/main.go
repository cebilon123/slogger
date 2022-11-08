package main

import (
	"context"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := proto.NewLogApiClient(conn)
	for i := 0; i < 25; i++ {
		res, err := client.Log(context.Background(), &proto.LogRequest{
			Type:    int32(i),
			Message: "asd",
		})
		if err != nil {
			log.Println(err)
		}
		log.Printf("Got response: %d", res.Type)
	}
}
