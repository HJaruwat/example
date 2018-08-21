package game

import (
	"cabal-api/common"
	"cabal-api/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	key = "23efwefwpgjeter3ffe"
)

type (
	// OutCabal struct output
	OutCabal struct {
		Result  int    `json:"result"`
		AuthKey string `json:"authkey"`
	}
)

// GetAuthToken do send request get token
func GetAuthToken(log *common.TransactionLogInfo) (string, error) {
	endpoint := fmt.Sprintf("%s/auth/get-key/%s", config.Setting.App.GameURI, key)

	reqLog := *log
	reqLog.Body = ""
	reqLog.Input = ""
	reqLog.URL = endpoint
	reqLog.Method = "GET"
	reqLog.LogWriter()

	resp, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	respLog := *log
	respLog.Body = ""
	respLog.Input = ""
	respLog.Output = string(bodyResp)
	respLog.Method = fmt.Sprint(resp.StatusCode)
	respLog.LogWriter()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Cabal getToken fail : %s", string(bodyResp))
	}

	out := OutCabal{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return "", fmt.Errorf("Cabal getToken unmarshal fail : %s", err.Error())
	}

	return out.AuthKey, nil
}
