package context

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_MigrateHandler(t *testing.T) {

	appTest := &App{}
	appTest.InitializeRoute()

	tests := []struct {
		name     string
		a        *App
		args     *InputMigrateCabal
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			a:    appTest,
			args: &InputMigrateCabal{
				ExeID:    "kerkoleng",
				EstID:    "ker54933945",
				ServerID: "18",
			},
			wantCode: http.StatusOK,
		},
		{
			name: "fail",
			a:    appTest,
			args: &InputMigrateCabal{
				ExeID:    "",
				EstID:    "ker54933945",
				ServerID: "18",
			},
			wantCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.args)
			t.Log(string(b))
			req, _ := http.NewRequest("POST", "/migrate", bytes.NewBuffer(b))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			tt.a.Router.ServeHTTP(rr, req)
			resp, _ := ioutil.ReadAll(rr.Result().Body)
			t.Logf("Response : %s", resp)
			checkResponseCode(t, tt.wantCode, rr.Code)
		})
	}
}
