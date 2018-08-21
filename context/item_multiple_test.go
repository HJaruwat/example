package context

import (
	"bytes"
	"cabal-api/common"
	"cabal-api/rabbitmq"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_ItemMultipleHandler(t *testing.T) {

	appTest := &App{}
	appTest.InitializeRoute()
	code := common.RandomRefCode(10)

	// array reward
	r1 := rabbitmq.Rewards{
		ItemID: "3588",
		Amount: 1,
		Name:   "Superioi Odd Clircle",
	}

	r2 := rabbitmq.Rewards{
		ItemID: "33555596",
		Amount: 1,
		Name:   "GM's Blessing (LV.1) Holy Water",
	}

	var arrRewards []rabbitmq.Rewards
	arrRewards = append(arrRewards, r1)
	arrRewards = append(arrRewards, r2)

	tests := []struct {
		name     string
		a        *App
		args     *rabbitmq.InputItemMultiple
		wantCode int
	}{
		// TODO: Add test cases.
		{
			name: "succecss case",
			a:    appTest,
			args: &rabbitmq.InputItemMultiple{
				Username:      "kerkoleng",
				ReferenceCode: "unit-testing-api-" + code,
				Rewards:       arrRewards,
				ServerID:      "1",
			},
			wantCode: http.StatusOK,
		},
		{
			name: "fail case",
			a:    appTest,
			args: &rabbitmq.InputItemMultiple{
				Username:      "",
				ReferenceCode: "unit-testing-api-" + code,
				Rewards:       nil,
			},
			wantCode: http.StatusBadRequest,
		},
		{
			name: "fail case 2 ",
			a:    appTest,
			args: &rabbitmq.InputItemMultiple{
				Username:      "kerkoleng",
				ReferenceCode: "unit-testing-api-" + code,
				Rewards:       nil,
			},
			wantCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				b, _ := json.Marshal(tt.args)
				t.Log(string(b))
				req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(b))
				req.Header.Set("Content-Type", "application/json")
				rr := httptest.NewRecorder()
				tt.a.Router.ServeHTTP(rr, req)
				resp, _ := ioutil.ReadAll(rr.Result().Body)
				t.Logf("Response : %s", resp)
				checkResponseCode(t, tt.wantCode, rr.Code)
			})
		})
	}
}
