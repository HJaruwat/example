package game

import (
	"cabal-api/common"
	"cabal-api/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// LoginPassPort do send auth with game api
func LoginPassPort(username, password string, log *common.TransactionLogInfo) (bool, error) {

	token, err := GetAuthToken(log)
	if err != nil {
		return false, err
	}
	endpoint := fmt.Sprintf("%s/auth/check-old-id-and-password/%s", config.Setting.App.GameURI, token)

	body := url.Values{}
	body.Add("id", username)
	body.Add("password", password)

	reqLog := *log
	reqLog.URL = endpoint
	reqLog.Method = "POST"
	reqLog.Body = body.Encode()
	reqLog.LogWriter()

	resp, err := http.PostForm(endpoint, body)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	respLog := *log
	respLog.Output = string(bodyResp)
	respLog.Method = fmt.Sprint(resp.StatusCode)
	respLog.LogWriter()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Cabal Login Passport fail : %s", string(bodyResp))
	}

	out := OutCabal{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return false, fmt.Errorf("Cabal Login Passport unmarshal fail : %s", err.Error())
	}

	if out.Result != 0 {
		return false, nil
	}

	return true, nil
}
