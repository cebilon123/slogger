package clog

import (
	"context"
	"github.com/cebilon123/slogger/gather/clog/model"
	"github.com/cebilon123/slogger/gather/pkg/gen/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

type logApiServer struct {
	proto.UnimplementedLogApiServer
	aggregator *Aggregator
}

var _ proto.LogApiServer = (*logApiServer)(nil)

func NewLogApiServer(_ context.Context, a *Aggregator) *logApiServer {
	return &logApiServer{
		aggregator: a,
	}
}

func (l *logApiServer) StreamLog(stream proto.LogApi_StreamLogServer) error {
	for {
		lr, err := stream.Recv()
		if err == io.EOF {
			log.Println("DONE")
			return nil
		}
		if err != nil {
			err = stream.SendAndClose(&proto.LogResponse{Type: -1})
			return err
		}

		l.aggregator.EnqueueLog(model.Log{
			Type:    lr.Type,
			Message: lr.Message,
		})
	}
}
func (l *logApiServer) Log(context.Context, *proto.LogRequest) (*proto.LogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}
