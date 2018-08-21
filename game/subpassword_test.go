package game

import (
	"cabal-api/common"
	"testing"
)

func TestAuthorized_ResetSubPassword(t *testing.T) {
	logger := &common.TransactionLogInfo{}

	code, err := GetAuthToken(logger)
	if err != nil {
		t.Error(err)
		return
	}

	auth := &Authorized{
		Code:     code,
		Logger:   logger,
		Username: "kerkoleng",
	}

	tests := []struct {
		name    string
		auth    *Authorized
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "success",
			auth:    auth,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.auth.ResetSubPassword(); (err != nil) != tt.wantErr {
				t.Errorf("Authorized.ResetSubPassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
