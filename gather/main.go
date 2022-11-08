package main

import (
	"context"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type testApiServer struct {
	proto.UnimplementedLogApiServer
}

func (t testApiServer) Log(ctx context.Context, in *proto.LogRequest) (*proto.LogResponse, error) {
	return &proto.LogResponse{Type: in.Type}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterLogApiServer(grpcServer, &testApiServer{})

	log.Println("Started TCP server")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
