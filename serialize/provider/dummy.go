package provider

import (
	"context"
	"github.com/cebilon123/slogger/serialize/model"
)

const (
	dummySerializerName = "Dummy"
)

type dummySerializer struct {
}

var _ Serializer = (*dummySerializer)(nil)

func NewDummySerializer() *dummySerializer {
	return &dummySerializer{}
}

func (d dummySerializer) PutOnSerializationQueue(_ context.Context, _ model.Log) {
	//TODO implement me
	panic("implement me")
}

func (d dummySerializer) GetName() string {
	return dummySerializerName
}
