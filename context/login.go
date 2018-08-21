package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"net/http"
)

type (
	// SuccessLoginResponse struct of response
	SuccessLoginResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
)

// LoginHandler do handler http request
func (a *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}
	// check body
	err = r.ParseForm()
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBodyDecode)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	if password == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorPasswordIsRequired)
		return
	}

	success, err := game.LoginPassPort(username, password, &writer.Logger)
	if err != nil {
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorLoginPassPort)
		return
	}

	if !success {
		writer.Response(SuccessLoginResponse{
			Status:  success,
			Message: "username or password is incorrect",
		})
		return

	}

	writer.Response(SuccessLoginResponse{
		Status:  success,
		Message: "login success",
	})
	return
}
