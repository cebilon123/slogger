package publisher

import (
	"context"
	"github.com/cebilon123/slogger/gather/clog"
	"github.com/cebilon123/slogger/gather/model"
	"go.uber.org/zap"
)

// logs is a logs' publisher.
type logs struct {
	aggregator *clog.Aggregator
	logger     *zap.Logger

	workerPool chan any
}

// New creates new sender that handles sending
// logs to desired address.
func New(a *clog.Aggregator, l *zap.Logger) *logs {
	return &logs{
		aggregator: a,
		logger:     l,
	}
}

// StartPublishing starts to listen for aggregator
// logs. If it gets the logs, the logs sends them to
// the desired address.
func (s *logs) StartPublishing(ctx context.Context) error {
	logsChan := s.aggregator.GetLogsChannel()
	go startHandlingLogsInChannel(ctx, logsChan, s.logger)
	return nil
}

func startHandlingLogsInChannel(ctx context.Context, logsChan <-chan model.Log, lg *zap.Logger) {
	for {
		select {
		case l := <-logsChan:
			lg.Info(l.Message)
		case <-ctx.Done():
			return
		}
	}
}
