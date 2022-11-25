package clog

import (
	"github.com/cebilon123/slogger/serialize/provider"
	"reflect"
	"testing"
)

func Test_processSerializers(t *testing.T) {
	type args struct {
		serializers []provider.Serializer
	}
	tests := []struct {
		name string
		args args
		want []provider.Serializer
	}{
		{
			name: "empty serializers, creates dummy serializer",
			args: args{
				serializers: nil,
			},
			want: []provider.Serializer{
				provider.NewDummySerializer(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processSerializers(tt.args.serializers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processSerializers() = %v, want %v", got, tt.want)
			}
		})
	}
}
