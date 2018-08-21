package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"cabal-api/game"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type (
	// InputMigrateCabal struct input migrate info
	InputMigrateCabal struct {
		ExeID    string `json:"exe_id"`
		EstID    string `json:"est_id"`
		ServerID string `json:"server_id"`
	}
	// SuccessMigrateResponse struct of success response
	SuccessMigrateResponse struct {
		Status  bool               `json:"status"`
		Code    int                `json:"code"`
		Message string             `json:"message"`
		Data    *InputMigrateCabal `json:"data"`
	}
)

// MigrateHandler do handler http request
func (a *App) MigrateHandler(w http.ResponseWriter, r *http.Request) {
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

	input := InputMigrateCabal{}
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBodyDecode)
		return
	}

	// validate exe_id
	if input.ExeID == "" {
		writer.Logger.Error = fmt.Errorf("ExeID is invalid")
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	// validate est_id
	if input.EstID == "" {
		writer.Logger.Error = fmt.Errorf("EstID is invalid")
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorEstID)
		return
	}

	// validate server_id
	if input.ServerID == "" {
		writer.Logger.Error = fmt.Errorf("ServerID is invalid")
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorServerID)
		return
	}

	result, code, err := game.MigrateAccount(input.ExeID, input.EstID, input.ServerID, &writer.Logger)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorMigrate)
		return
	}

	if !result {
		var message string
		switch code {
		case 88:
			message = "EST account not found"
			break
		case 94:
			message = "estID account has no character"
			break
		case 95:
			message = "character already exists in EXE account"
			break
		case 96:
			message = "please logout before migrate"
			break
		case 97:
			message = "already migrated account"
			break
		case 98:
			message = "migrating account has completed the migration"
			break
		case 99:
			message = "account is blocked account"
			break
		default:
			message = "out of conditions"
			break
		}

		writer.Response(SuccessMigrateResponse{
			Status:  false,
			Message: message,
			Code:    code,
			Data:    &input,
		})
		return
	}

	writer.Response(SuccessMigrateResponse{
		Status:  true,
		Code:    http.StatusOK,
		Message: "success",
		Data:    &input,
	})
}
