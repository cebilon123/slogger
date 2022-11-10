package clog

import (
	"context"
	"go.uber.org/zap"
)

type Sender struct {
	aggregator *Aggregator
	logger     *zap.Logger
}

func NewSender(a *Aggregator, l *zap.Logger) *Sender {
	return &Sender{
		aggregator: a,
		logger:     l,
	}
}

// StartSendingProcess starts to listen for aggregator
// logs. If it gets the logs, the Sender sends them to
// the desired address.
func (s *Sender) StartSendingProcess(ctx context.Context) error {
	logsChan := s.aggregator.GetLogsChannel()
	go func() {
		for {
			select {
			case l := <-logsChan:
				s.logger.Info(l.Message)
			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}
