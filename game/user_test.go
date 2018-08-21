package game

import (
	"cabal-api/common"
	"testing"
)

func TestCheckUser(t *testing.T) {
	type args struct {
		username string
		log      *common.TransactionLogInfo
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
				username: "kerkoleng",
				log:      &common.TransactionLogInfo{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckUser(tt.args.username, tt.args.log)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			t.Logf("CheckUser() = %v", got)

		})
	}
}
