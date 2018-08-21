package context

import (
	"bytes"
	"cabal-api/common"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestApp_CashHandler(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	mock := appTest.InitializeDatabaseTest()
	refcode := "unit-refcode-" + common.RandomRefCode(5)
	mock.ExpectPrepare("INSERT INTO `topup`").ExpectExec().WithArgs("kerkoleng", refcode, 100, 1, 0).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectPrepare("UPDATE `topup`").ExpectExec().WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	type args struct {
		input InputCash
	}

	args1 := args{
		input: InputCash{
			Username:      "kerkoleng",
			Amount:        100,
			Bonus:         1,
			ReferenceCode: refcode,
		},
	}

	args2 := args{
		input: InputCash{
			Username:      "kerkoleng",
			Amount:        100,
			ReferenceCode: "",
		},
	}

	tests := []struct {
		name     string
		a        *App
		args     args
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name:     "success case",
			a:        appTest,
			args:     args1,
			wantCode: http.StatusOK,
		},
		{
			name:     "error case",
			a:        appTest,
			args:     args2,
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.args.input)
			t.Log(string(b))
			req, _ := http.NewRequest("POST", "/topup", bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)

			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)
			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}

func TestApp_InsertTopup(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	mock := appTest.InitializeDatabaseTest()
	refcode := "unit-refcode-" + common.RandomRefCode(5)
	mock.ExpectPrepare("INSERT INTO `topup`").ExpectExec().WithArgs("kerkoleng", refcode, 100, 1, 0).WillReturnResult(sqlmock.NewResult(1, 1))

	type args struct {
		input *InputCash
	}
	tests := []struct {
		name    string
		a       *App
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
			args: args{
				input: &InputCash{
					Amount:        100,
					ReferenceCode: refcode,
					Username:      "kerkoleng",
					Bonus:         1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.InsertTopup(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("App.InsertTopup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_UpdateStatusPutMessage(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	mock := appTest.InitializeDatabaseTest()
	mock.ExpectPrepare("UPDATE `topup`").ExpectExec().WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	refcode := "unit-refcode-" + common.RandomRefCode(5)
	type args struct {
		input *InputCash
	}
	tests := []struct {
		name    string
		a       *App
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success case",
			a:    appTest,
			args: args{
				input: &InputCash{
					ID:            1,
					Amount:        100,
					ReferenceCode: refcode,
					Username:      "kerkoleng",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UpdateStatusPutMessage(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("App.UpdateStatusPutMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
