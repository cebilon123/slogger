package main

import (
	"context"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := proto.NewLogApiClient(conn)
	logStream, err := client.StreamLog(context.Background())
	if err != nil {
		log.Println(err)
	}

	i := 0
	for {
		time.Sleep(time.Second * 1)
		err := logStream.SendMsg(&proto.LogRequest{
			Type:    int32(i),
			Message: "",
		})
		if err != nil {
			log.Println(err)
		}
		i++
	}
}
