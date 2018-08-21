package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type (
	// SuccessResetResponse struct of response
	SuccessResetResponse struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}
)

// ResetSubPasswordHandler do handler http request
func (a *App) ResetSubPasswordHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}
	vars := mux.Vars(r)
	username := vars["user_id"]
	if username == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBadRequest)
		return
	}
	// Log input
	writer.Logger.Input = fmt.Sprintf("username=%s", username)
	writer.Logger.URL = "/sub-password"
	writer.Logger.LogWriter()
	// Get Auth Key with Game API
	code, err := game.GetAuthToken(&writer.Logger)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorGetAuthKeyWithGameAPI)
		return
	}

	auth := &game.Authorized{
		Code:     code,
		Username: strings.ToLower(username),
		Logger:   &writer.Logger,
	}

	if err = auth.ResetSubPassword(); err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorResetSubPasswordWithGameAPI)
		return
	}

	writer.Response(SuccessResetResponse{
		Status:  true,
		Message: "reset sub-password succeed",
	})
}
