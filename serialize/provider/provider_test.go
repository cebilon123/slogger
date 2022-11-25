package provider

import (
	"github.com/cebilon123/slogger/serialize/model"
	"reflect"
	"testing"
	"time"
)

func TestExtendedLogFromLog(t *testing.T) {
	now := time.Now()
	timeNowUTCFunc = func() time.Time {
		return now
	}

	logModel := model.Log{
		Type:    1,
		Message: "Test",
		Caller:  "Test",
	}

	type args struct {
		m model.Log
	}
	tests := []struct {
		name string
		args args
		want *ExtendedLog
	}{
		{
			name: "parse model.Log into ExtendedLog",
			args: args{
				m: logModel,
			},
			want: &ExtendedLog{
				Log:       logModel,
				createdAt: now,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtendedLogFromLog(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtendedLogFromLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
