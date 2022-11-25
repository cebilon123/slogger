// Package provider handles serialization providers logic.
package provider

import (
	"context"
	"github.com/cebilon123/slogger/serialize/model"
	"time"
)

// timeNowUTCFunc is used mainly to replace time.now in tests.
var timeNowUTCFunc = func() time.Time {
	return time.Now().UTC() // for now by default we are using UTC, but we need to add possibility to use other timezones
}

// Serializer must be implemented by all serialization
// providers. It shares logic for enqueuing logs, which then
// should be serialized in queue.
type Serializer interface {
	// PutOnSerializationQueue puts logs on serialization queue to ensure order of logs.
	PutOnSerializationQueue(ctx context.Context, log model.Log)
	// GetName returns name of given serializer.
	GetName() string
}

type ExtendedLog struct {
	model.Log

	createdAt time.Time
}

// ExtendedLogFromLog creates ExtendedLog from model.Log.
func ExtendedLogFromLog(m model.Log) *ExtendedLog {
	return &ExtendedLog{
		Log:       m,
		createdAt: timeNowUTCFunc(),
	}
}
