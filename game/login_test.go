package game

import (
	"cabal-api/common"
	"testing"
)

func TestLoginPassPort(t *testing.T) {
	type args struct {
		username string
		password string
		log      *common.TransactionLogInfo
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				username: "theblackgo1",
				password: "1C750F05A4",
				log:      &common.TransactionLogInfo{},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoginPassPort(tt.args.username, tt.args.password, tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginPassPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LoginPassPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
