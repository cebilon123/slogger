package clog

import (
	"context"
	"errors"
	"github.com/cebilon123/slogger/gather/clog/model"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

var (
	ErrClientDisconnected = errors.New("client disconnected")
)

type logApiServer struct {
	proto.UnimplementedLogApiServer

	aggregator *Aggregator
	logger     *zap.Logger
}

var _ proto.LogApiServer = (*logApiServer)(nil)

func NewLogApiServer(_ context.Context, a *Aggregator, l *zap.Logger) *logApiServer {
	return &logApiServer{
		aggregator: a,
		logger:     l,
	}
}

func (l *logApiServer) StreamLog(stream proto.LogApi_StreamLogServer) error {
	for {
		lr, err := stream.Recv()

		if err == io.EOF {
			return ErrClientDisconnected
		}

		if err != nil {
			err = stream.SendAndClose(&proto.LogResponse{Code: -1})
			return err
		}

		l.aggregator.EnqueueLog(model.Log{
			Type:    lr.Type,
			Message: lr.Message,
			Caller:  lr.Caller,
		})
	}
}
func (l *logApiServer) Log(ctx context.Context, lr *proto.LogRequest) (*proto.LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
