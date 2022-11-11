package publisher

import (
	"context"
	"errors"
	"github.com/cebilon123/slogger/gather/clog"
	"go.uber.org/zap"
)

var (
	ErrLogsPublisherNotClosed = errors.New("logsPublisher closing failed")
)

// logsPublisher is a logsPublisher' publisher.
type logsPublisher struct {
	aggregator *clog.Aggregator
	logger     *zap.Logger

	doneChan chan struct{}
	so       stateObserve
}

func (s *logsPublisher) Close() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			s.logger.Error("logsPublisher - Close(): publisher cannot be closed.")
			err = ErrLogsPublisherNotClosed
		}
	}()
	s.doneChan <- struct{}{}
	return err
}

// stateObserve is used to observe state of logsPublisher.
// It can be used in the tests.
type stateObserve struct {
	isDone         bool
	isHandlingLogs bool
}

// New creates new sender that handles sending
// logsPublisher to desired address.
func New(a *clog.Aggregator, l *zap.Logger) *logsPublisher {
	doneChan := make(chan struct{}, 1)
	return &logsPublisher{
		aggregator: a,
		logger:     l,
		doneChan:   doneChan,
	}
}

// StartPublishing starts to listen for aggregator
// logsPublisher. If it gets the logsPublisher, the logsPublisher sends them to
// the desired address.
func (s *logsPublisher) StartPublishing(ctx context.Context) error {
	s.so.isHandlingLogs = true
	go s.startHandlingLogsInChannel(ctx, s.logger)
	return nil
}

func (s *logsPublisher) startHandlingLogsInChannel(ctx context.Context, lg *zap.Logger) {
	defer func() {
		s.so.isDone = true
		s.so.isHandlingLogs = false
	}()

	logsChan := s.aggregator.GetLogsChannel()
	for {
		select {
		case l := <-logsChan:
			lg.Info(l.Message)
		case <-ctx.Done():
			return
		case <-s.doneChan:
			return
		}
	}
}
