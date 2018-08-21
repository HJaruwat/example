package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

type (
	// SuccessInfoResponse struct of response when success case
	SuccessInfoResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
)

// InfoHandler do handler http request
func (a *App) InfoHandler(w http.ResponseWriter, r *http.Request) {
	// var err error
	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}

	vars := mux.Vars(r)
	values := url.Values{}

	username := vars["user_id"]
	values.Add("user_id", username)

	writer.Logger.URL = r.URL.String()
	writer.Logger.Input = values.Encode()
	writer.Logger.LogWriter()

	if username == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	created, err := game.CheckUser(username, &writer.Logger)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorCannotConnectGameAPI)
		return
	}

	if !created {
		//ErrorGameUserNotFound
		writer.Logger.Error = fmt.Errorf(errors.ErrorGameUserNotFound.ErrorDescription)
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorGameUserNotFound)
		return
	}

	writer.Response(&SuccessInfoResponse{
		Status:  true,
		Message: "username already have account in game",
	})

}
