package context

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestApp_LoginHandler(t *testing.T) {

	appTest := &App{}
	appTest.InitializeRoute()

	type args struct {
		username string
		password string
	}

	tests := []struct {
		name     string
		a        *App
		args     *args
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
			args: &args{
				username: "theblackgo1",
				password: "1C750F05A4",
			},
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := url.Values{}
			body.Add("username", tt.args.username)
			body.Add("password", tt.args.password)

			req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(body.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)
			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)
			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}
