package server

import (
	"context"
	"fmt"
	"github.com/cebilon123/slogger/gather/clog"
	"github.com/cebilon123/slogger/gather/config"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

// Start starts handler which handles initialization
// and TCP server. If error is not nil, application
// should log fatal.
func (h *handler) Start() error {
	ctx := context.Background()
	if err := config.NewEnvironment(); err != nil {
		return err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Environment.Port))
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	initApiServers(ctx, grpcServer)

	logger.Info(fmt.Sprintf("Starting TCP server on port: %s", config.Environment.Port))
	logger.Info(fmt.Sprintf("Server address: %s", listener.Addr().String()))

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

// initApiServers initialize api servers and all dependencies used in application.
func initApiServers(ctx context.Context, grpcServer *grpc.Server) {
	aggregator := clog.NewAggregator(ctx)
	apiServer := clog.NewLogApiServer(ctx, aggregator)
	proto.RegisterLogApiServer(grpcServer, apiServer)
}
