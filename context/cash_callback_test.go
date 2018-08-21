package context

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestApp_UpdateTopupSentWithID(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	mock := appTest.InitializeDatabaseTest()
	mock.ExpectPrepare("UPDATE `topup`").ExpectExec().WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(1, 1))
	type args struct {
		input *TopupCallBack
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
				input: &TopupCallBack{
					ID:       1,
					FlagSent: true,
				},
			},
			wantErr: false,
		},
		{
			name: "fail case",
			a:    appTest,
			args: args{
				input: &TopupCallBack{
					ID:       0,
					FlagSent: false,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.UpdateTopupSentWithID(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("App.UpdateTopupSentWithID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_GetTopupWithReferenceCode(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	columns := []string{"id", "username", "reference_id", "amount", "bonus", "flag_queue", "flag_sent", "created_at", "updated_at"}

	mock := appTest.InitializeDatabaseTest()
	mock.ExpectQuery("select (.+) from topup").WithArgs("1").WillReturnRows(sqlmock.NewRows(columns).FromCSVString(`1,kerkoleng,1,10,0,1,0,2006-01-02T15:04:05Z,2006-01-02T15:04:05Z`))

	type args struct {
		input *TopupCallBack
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
				input: &TopupCallBack{
					ReferenceID: "1",
					ID:          1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.a.GetTopupWithReferenceCode(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("App.GetTopupWithReferenceCode() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.args.input.ID <= 0 {
				t.Errorf("App.GetTopupWithReferenceCode() error = topup id not found")
			}
		})
	}
}

func TestApp_CashCallBackHandler(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()
	columns := []string{"id", "username", "reference_id", "amount", "bonus", "flag_queue", "flag_sent", "created_at", "updated_at"}

	mock := appTest.InitializeDatabaseTest()
	mock.ExpectQuery("select (.+) from topup").WithArgs("1").WillReturnRows(sqlmock.NewRows(columns).FromCSVString(`1,kerkoleng,1,10,0,1,0,2006-01-02T15:04:05Z,2006-01-02T15:04:05Z`))
	mock.ExpectPrepare("UPDATE `topup`").ExpectExec().WithArgs(true, 1).WillReturnResult(sqlmock.NewResult(1, 1))

	tests := []struct {
		name     string
		a        *App
		refCode  string
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name:     "success",
			a:        appTest,
			refCode:  "1",
			wantCode: 200,
		},
		{
			name:     "fail",
			a:        appTest,
			refCode:  "",
			wantCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := url.Values{}
			data.Add("refcode", tt.refCode)

			req, _ := http.NewRequest("POST", "/topup/callback", bytes.NewBufferString(data.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)

			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)

			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}
