package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"cabal-api/rabbitmq"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type (

	// SuccessItemMultipleResponse struct success response
	SuccessItemMultipleResponse struct {
		Status  bool                        `json:"status"`
		Message string                      `json:"message"`
		Data    *rabbitmq.InputItemMultiple `json:"data"`
	}
)

// ItemMultipleHandler do handler http request
func (a *App) ItemMultipleHandler(w http.ResponseWriter, r *http.Request) {
	var err error

	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}

	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "json") {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorContentType)
		return
	}

	input := &rabbitmq.InputItemMultiple{}

	if err = json.NewDecoder(r.Body).Decode(input); err != nil {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorInternalServer)
		return
	}

	if input.Username == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	input.CreatedAt = time.Now()

	if input.ReferenceCode == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorReferenceNotFound)
		return
	}

	// check user
	created, err := game.CheckUser(input.Username, &writer.Logger)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorCannotConnectGameAPI)
		return
	}

	if !created {
		writer.Logger.Error = fmt.Errorf(errors.ErrorGameUserNotFound.ErrorDescription)
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorGameUserNotFound)
		return
	}

	if len(input.Rewards) < 1 {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorRewardsList)
		return
	}

	// save message before push to queue
	byteBody, _ := json.Marshal(input)
	writer.Logger.Body = string(byteBody)
	writer.Logger.LogWriter()

	err = rabbitmq.EmitMassageItem(input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorPutMessage)
		return
	}

	writer.Response(SuccessItemMultipleResponse{
		Status:  true,
		Message: "success",
		Data:    input,
	})
}
