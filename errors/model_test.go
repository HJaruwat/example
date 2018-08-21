package errors

import (
	"cabal-api/common"
	"testing"
)

func TestErrorData(t *testing.T) {
	var v interface{}

	type args struct {
		err  string
		desc string
		data interface{}
	}
	tests := []struct {
		name string
		args args
		want *common.ErrorResponse
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				err:  "unit testing",
				desc: "unit testing description",
				data: "data",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v = ErrorData(tt.args.err, tt.args.desc, tt.args.data)
			switch v.(type) {
			case *common.ErrorResponse:
				t.Log("test passed")
			default:
				t.Errorf("ErrorData() = %v, want %v", v, tt.want)
			}

		})
	}
}
