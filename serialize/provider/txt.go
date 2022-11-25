package provider

import (
	"context"
	"github.com/cebilon123/slogger/serialize/model"
)

// TxtSerializer is a simple txt
// serializer. It serializes it
// to text documents of .txt format.
type TxtSerializer struct {
	logsQueueChan chan model.Log
}

// NewTxtProvider creates new text serializer
// provider. It enables to log into .txt format
// files.
func NewTxtProvider() *TxtSerializer {
	logsQueueChan := make(chan model.Log)
	return &TxtSerializer{
		logsQueueChan: logsQueueChan,
	}
}

func (t TxtSerializer) PutOnSerializationQueue(ctx context.Context, log model.Log) {
	t.logsQueueChan <- log
}
