package common

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriter_ResponseError(t *testing.T) {
	writer := Writer{
		Resp: httptest.NewRecorder(),
		Logger: TransactionLogInfo{
			RefCode: RandomRefCode(5),
		},
	}
	type args struct {
		status int
		err    *ErrorResponse
	}
	tests := []struct {
		name string
		w    Writer
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			w:    writer,
			args: args{
				err: &ErrorResponse{
					Error:            "internal_server_error",
					ErrorDescription: "some thing wrong",
				},
				status: http.StatusInternalServerError,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.ResponseError(tt.args.status, tt.args.err)
		})
	}
}

func TestWriter_Response(t *testing.T) {
	writer := Writer{
		Resp: httptest.NewRecorder(),
		Logger: TransactionLogInfo{
			RefCode: RandomRefCode(5),
		},
	}

	type args struct {
		data interface{}
	}
	tests := []struct {
		name string
		w    Writer
		args args
	}{
		// TODO: Add test cases.
		{
			name: "success",
			w:    writer,
			args: args{
				data: "test pass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.w.Response(tt.args.data)
		})
	}
}
