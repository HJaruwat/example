package context

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_ResetSubPasswordHandler(t *testing.T) {
	appTest := &App{}
	appTest.InitializeRoute()

	type args struct {
		username string
	}
	tests := []struct {
		name     string
		a        *App
		username string
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name:     "success",
			a:        appTest,
			username: "kerkoleng",
			wantCode: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodPut, fmt.Sprintf("/sub-password/%s", tt.username), nil)
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)
			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)
			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}
