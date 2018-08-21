package context

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_CharacterHandler(t *testing.T) {

	appTest := &App{}
	appTest.InitializeRoute()

	type args struct {
		username string
	}
	tests := []struct {
		name     string
		a        *App
		args     args
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
			args: args{
				username: "hjaruwat",
			},
			wantCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/characters/"+tt.args.username, nil)
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)
			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)
			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}
