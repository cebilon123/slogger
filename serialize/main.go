package main

import (
	"context"
	"github.com/cebilon123/slogger/serialize/clog"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	aggregator := clog.NewAggregator(ctx)
	logger, _ := zap.NewProduction()

	go func() {
		<-exit
		logger.Info("Issued close with termination")
		_ = aggregator.Close()
	}()

	go func() {
		aggregator.Start(ctx)
	}()

	<-aggregator.Done()
}
