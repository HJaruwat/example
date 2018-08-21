package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

// CharacterHandler do handler http request character Info.
func (a *App) CharacterHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}
	// check body
	vars := mux.Vars(r)
	values := url.Values{}

	username := vars["user_id"]
	values.Add("user_id", username)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBodyDecode)
		return
	}

	if username == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	resp, err := game.CharacterInfo(username, &writer.Logger)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorCharacterInfo)
		return
	}

	writer.Response(resp)
	return
}
