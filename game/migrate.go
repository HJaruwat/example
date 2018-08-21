package game

import (
	"cabal-api/common"
	"cabal-api/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MigrateAccount do send auth with game api
func MigrateAccount(username, estID, serverID string, log *common.TransactionLogInfo) (bool, int, error) {
	token, err := GetAuthToken(log)
	if err != nil {
		return false, 0, err
	}

	endpoint := config.Setting.App.GameURI
	endpoint = fmt.Sprintf("%s/game/move-character/%s/%s/%s/%s", endpoint, token, username, estID, serverID)

	reqLog := *log
	reqLog.URL = endpoint
	reqLog.Method = "GET"
	reqLog.LogWriter()

	resp, err := http.Get(endpoint)
	if err != nil {
		return false, 0, err
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, 0, err
	}

	respLog := *log
	respLog.Output = string(bodyResp)
	respLog.StatusCode = resp.StatusCode
	respLog.LogWriter()

	if resp.StatusCode != http.StatusOK {
		return false, 0, fmt.Errorf("Cabal CabalMigrate fail : %s", string(bodyResp))
	}

	out := OutCabal{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return false, 0, fmt.Errorf("Cabal CabalMigrate unmarshal fail : %s", err.Error())
	}

	if out.Result != 0 {
		return false, out.Result, nil
	}

	return true, out.Result, nil
}
