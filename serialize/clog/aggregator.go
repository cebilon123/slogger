package clog

import (
	"context"
	"github.com/cebilon123/slogger/serialize/model"
	"github.com/cebilon123/slogger/serialize/provider"
	"log"
)

// Aggregator is used to aggregate
// obtained logs from gather.
type Aggregator struct {
	logChan  chan model.Log     // logChan is a channel used to aggregate logs.
	innerCtx context.Context    // innerCtx is an inner context of aggregator used to gently close itself.
	cancel   context.CancelFunc // cancelFunc used to cancel whole aggregator.
	doneChan chan struct{}      // doneChan is used to inform if aggregator is done with the work.
}

// NewAggregator creates new aggregator that aggregates logs and send those to registered providers.
func NewAggregator(ctx context.Context, serializers ...provider.Serializer) *Aggregator {
	processSerializers(serializers...)

	innerCtx, cancel := context.WithCancel(ctx)
	logChan := make(chan model.Log)

	return &Aggregator{
		logChan:  logChan,
		innerCtx: innerCtx,
		cancel:   cancel,
		doneChan: make(chan struct{}, 1),
	}
}

// Close gently closes logChan, this way
// all the logs that were sent are safely
// serialized.
func (a *Aggregator) Close() error {
	a.cancel()
	log.Println("canceled")
	return nil
}

// Done is a channel informing if aggregator is done (i.e. it is closed, but there
// are still log messages in the channel so we should wait till all the messages
// are going to be removed from the channel).
func (a *Aggregator) Done() <-chan struct{} {
	return a.doneChan
}

// Start starts the aggregation. Logs are being sent on the
// log channel from where those can be handled during serialization process.
func (a *Aggregator) Start(ctx context.Context) {
	defer func() {
		a.doneChan <- struct{}{}
	}()

	for {
		select {
		//case l := <-a.logChan:

		case <-ctx.Done():
			// wait for all messages from log chan
			return
		case <-a.innerCtx.Done():
			// wait for all messages from log chan
			log.Println("Done")
			return
		}
	}
}

// processSerializers processes passed serializers. If there is no serializer passed
// it creates dummy serializer to basically free up logs from channel.
func processSerializers(serializers ...provider.Serializer) []provider.Serializer {
	var serializersCopy []provider.Serializer
	copy(serializersCopy, serializersCopy)
	if len(serializersCopy) == 0 {
		serializersCopy = append(serializersCopy, provider.NewDummySerializer())
	}

	return serializersCopy
}
