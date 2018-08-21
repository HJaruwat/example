package game

import (
	"cabal-api/common"
	"testing"
)

func TestGetAuthToken(t *testing.T) {
	type args struct {
		log *common.TransactionLogInfo
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
				log: &common.TransactionLogInfo{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAuthToken(tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAuthToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != "" {
				t.Logf("GetAuthToken() = %v", got)
			}
		})
	}
}
