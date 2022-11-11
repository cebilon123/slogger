package main

import (
	"context"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"strconv"
	"sync"
)

// just used to test things out
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

	var wg sync.WaitGroup
	for j := 0; j < 1000000; j++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			err := logStream.SendMsg(&proto.LogRequest{
				Type:    int32(j),
				Message: strconv.Itoa(j),
			})
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
			}
		}(j)
	}

	wg.Wait()
}
