package game

import (
	"cabal-api/common"
	"cabal-api/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type (
	// Authorized struct of auth data
	Authorized struct {
		Code     string
		Username string
		Logger   *common.TransactionLogInfo
	}
)

// ResetSubPassword do send request reset subpassword
func (auth *Authorized) ResetSubPassword() error {
	endpoint := fmt.Sprintf("%s/game/reset-sub-password/%s/%s", config.Setting.App.GameURI, auth.Code, auth.Username)

	reqLog := *auth.Logger
	reqLog.Body = ""
	reqLog.Input = ""
	reqLog.URL = endpoint
	reqLog.Method = "GET"
	reqLog.LogWriter()

	resp, err := http.Get(endpoint)
	if err != nil {
		return errors.Wrap(err, "ResetSubPassword request GET error")
	}

	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "ResetSubPassword read response error")
	}

	respLog := *auth.Logger
	respLog.Body = ""
	respLog.Input = ""
	respLog.Output = string(bodyResp)
	respLog.Method = fmt.Sprint(resp.StatusCode)
	respLog.LogWriter()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(err, fmt.Sprintf("ResetSubPassword response httpStatus %d", resp.StatusCode))
	}

	out := OutCabal{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return errors.Wrap(err, "ResetSubPassword unmarshal fail")
	}

	if out.Result == 0 {
		return nil
	} else if out.Result == 88 {
		return errors.New("ResetSubPassword user not found")
	} else if out.Result == 99 {
		return errors.New("ResetSubPassword fail")
	}

	return errors.New("ResetSubPassword error no definition")
}
