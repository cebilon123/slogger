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

	aggregator := clog.NewAggregator(ctx)
	apiServer := clog.NewLogApiServer(ctx, aggregator, logger)
	sender := clog.NewSender(aggregator, logger)
	if err := sender.StartSendingProcess(ctx); err != nil {
		return err
	}

	proto.RegisterLogApiServer(grpcServer, apiServer)

	logger.Info(fmt.Sprintf("Starting TCP server, address: %s", listener.Addr().String()))

	err = grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}

// initApiServers initialize api servers and all dependencies used in application.
func initApiServers(ctx context.Context, grpcServer *grpc.Server, l *zap.Logger) {

}
