package rabbitmq

import (
	"cabal-api/common"
	"testing"
)

func TestEmitMassageCash(t *testing.T) {
	code := common.RandomRefCode(10)
	textSucess := `{"id":1,"username":"kerkoleng","refcode":"unit-TestEmitMassageCash-` + code + `","amount":100,"created_at":"2017-11-06T15:34:41.5841895+07:00","updated_at":null}`
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				body: []byte(textSucess),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				body: []byte(``),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EmitMassageCash(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("EmitMassageCash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
