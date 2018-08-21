package game

import (
	"cabal-api/common"
	"cabal-api/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type (
	//Response Est Api Get CharacterList
	Response struct {
		Result int  `json:"result"`
		Data   Data `json:"data"`
	}
	//Data User and Servers.
	Data struct {
		User    User      `json:"user"`
		Servers []Servers `json:"servers"`
	}

	//User Detail.
	User struct {
		UserNumber int    `json:"userNum"`
		ID         string `json:"id"`
		LoginTime  string `json:"loginTime"`
		CreateDate string `json:"createDate"`
		Cash       int    `json:"cash"`
	}

	//Servers Detail.
	Servers struct {
		Server     Server            `json:"server"`
		Characters []CharacterDetail `json:"characters"`
	}

	//Server ServerID and ServerName.
	Server struct {
		ServerID   int    `json:"serverIdx"`
		ServerName string `json:"serverName"`
	}
	//CharacterDetail detail of character.
	CharacterDetail struct {
		Name       string `json:"name"`
		Level      int    `json:"lev"`
		Style      string `json:"style"`
		LoginTime  string `json:"logintime"`
		CreateDate string `json:"createDate"`
	}
)

// CheckUser do request to check created user
func CheckUser(username string, log *common.TransactionLogInfo) (bool, error) {

	token, err := GetAuthToken(log)
	if err != nil {
		return false, err
	}

	endpoint := config.Setting.App.GameURI
	endpoint = fmt.Sprintf("%s/account/check-created-account/%s/%s", endpoint, token, username)

	requestLog := *log
	requestLog.Body = ""
	requestLog.Input = ""
	requestLog.URL = endpoint
	requestLog.Method = "GET"
	requestLog.LogWriter()

	resp, err := http.Get(endpoint)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	respLog := *log
	requestLog.Body = ""
	requestLog.Input = ""
	respLog.Output = string(bodyResp)
	respLog.Method = fmt.Sprint(resp.StatusCode)
	respLog.LogWriter()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("check user fail : %s", string(bodyResp))
	}

	out := OutCabal{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return false, fmt.Errorf("check user unmarshal fail : %s", err.Error())
	}

	if out.Result != 0 {
		return false, nil
	}

	return true, nil
}

func CharacterInfo(username string, log *common.TransactionLogInfo) (*Response, error) {
	token, err := GetAuthToken(log)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/game/get-all-characters/%s/%s", config.Setting.App.GameURI, token, username)

	reqLog := *log
	reqLog.URL = endpoint
	reqLog.Method = "GET"
	reqLog.LogWriter()

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error Request Est Api %s", endpoint)
	}

	bodyResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respLog := *log
	respLog.Output = string(bodyResp)
	respLog.StatusCode = resp.StatusCode
	respLog.Method = resp.Request.Method
	respLog.LogWriter()

	out := &Response{}
	err = json.Unmarshal(bodyResp, &out)
	if err != nil {
		return nil, err
	}

	return out, nil

}
