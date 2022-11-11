package publisher

import (
	"context"
	"github.com/cebilon123/slogger/gather/clog"
	"github.com/cebilon123/slogger/gather/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"sync"
	"testing"
	"time"
)

func Test_logsPublisher_StartPublishing_ReturnsNil(t *testing.T) {
	ctx := context.Background()
	l, _ := zap.NewDevelopment()
	a := clog.NewAggregator(ctx)

	lPublisher := New(a, l)
	_ = lPublisher.StartPublishing(ctx)

	a.EnqueueLog(model.Log{
		Type:    0,
		Message: "test",
		Caller:  "test.caller",
	})

	time.Sleep(time.Second * 1)

	assert.Equal(t, true, lPublisher.so.isHandlingLogs)
}

func Test_logsPublisher_CloseCalled_StopsPublishing(t *testing.T) {
	ctx := context.Background()
	l, _ := zap.NewDevelopment()
	a := clog.NewAggregator(ctx)

	lPublisher := New(a, l)

	_ = lPublisher.StartPublishing(ctx)

	a.EnqueueLog(model.Log{
		Type:    0,
		Message: "test",
		Caller:  "test.caller",
	})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond * 500)
		_ = lPublisher.Close()
		wg.Done()
	}()

	wg.Wait()
	time.Sleep(time.Millisecond * 500)
	assert.Equal(t, true, lPublisher.so.isDone)
}
