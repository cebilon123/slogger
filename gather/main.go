package main

import (
	"context"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"github.com/cebilon123/slogger/gather/server"
	"log"
)

type testApiServer struct {
	proto.UnimplementedLogApiServer
}

func (t testApiServer) Log(ctx context.Context, in *proto.LogRequest) (*proto.LogResponse, error) {
	return &proto.LogResponse{Type: in.Type}, nil
}

func main() {
	handler := server.NewHandler()
	if err := handler.Start(); err != nil {
		log.Fatalln(err)
	}
}
