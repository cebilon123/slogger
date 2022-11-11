package clog

import (
	"context"
	"github.com/cebilon123/slogger/gather/model"
	"sync/atomic"
)

// Aggregator is logs aggregator. It is called
// by a lot of different goroutines, so all operations
// need to be handled without data races.
type Aggregator struct {
	logsChan chan model.Log

	enqueuedLogsCount uint64
}

func (a *Aggregator) Close() error {
	close(a.logsChan)
	return nil
}

func NewAggregator(_ context.Context) *Aggregator {
	logsChan := make(chan model.Log)
	return &Aggregator{
		logsChan: logsChan,
	}
}

// EnqueueLog puts logs to the logs channel which
// then is used to handle sending of logs from different
// goroutines.
func (a *Aggregator) EnqueueLog(l model.Log) {
	a.logsChan <- l
	atomic.AddUint64(&a.enqueuedLogsCount, 1)
}

// GetLogsChannel gets readonly logs channel.
func (a *Aggregator) GetLogsChannel() <-chan model.Log {
	return a.logsChan
}
