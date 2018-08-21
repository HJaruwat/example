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
	// InputCash struct of input message
	InputCash struct {
		ID            int64      `json:"id"`
		Username      string     `json:"username"`
		ReferenceCode string     `json:"refcode"`
		Amount        int        `json:"amount"`
		Bonus         int        `json:"bonus"`
		CreatedAt     time.Time  `json:"created_at"`
		UpdatedAt     *time.Time `json:"updated_at"`
	}
	// SuccessResponse struct of response when success case
	SuccessResponse struct {
		Status  bool      `json:"status"`
		Message string    `json:"message"`
		Data    InputCash `json:"data"`
	}
)

// CashHandler do handler http request
func (a *App) CashHandler(w http.ResponseWriter, r *http.Request) {
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

	// init input data
	input := InputCash{}
	if errDecode := json.NewDecoder(r.Body).Decode(&input); errDecode != nil {
		writer.Logger.Error = errDecode
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBodyDecode)
		return
	}

	input.CreatedAt = time.Now()
	// Log input value
	inputStr, err := json.Marshal(&input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorBodyEncode)
		return
	}
	writer.Logger.Input = string(inputStr)
	writer.Logger.LogWriter()

	// validate username
	if input.Username == "" {
		writer.Logger.Error = fmt.Errorf("Username is invalid")
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUsername)
		return
	}

	if input.ReferenceCode == "" {
		writer.Logger.Error = fmt.Errorf("ReferenceCode is invalid")
		writer.Logger.LogWriter()
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

	// validate amount
	if input.Amount < 1 {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorAmount)
		return
	}

	// insert record in database
	if errInsert := a.InsertTopup(&input); errInsert != nil {
		writer.Logger.Error = errInsert
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorCreateTransaction)
		return
	}

	byteJSON, err := json.Marshal(input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorInternalServer)
		return
	}

	err = rabbitmq.EmitMassageCash(byteJSON)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusInternalServerError, errors.ErrorPutMessage)
		return
	}

	if errUpdate := a.UpdateStatusPutMessage(&input); errUpdate != nil {
		writer.Logger.Error = errUpdate
		writer.Logger.LogWriter()
		// writer.ResponseError(http.StatusInternalServerError, errors.ErrorUpdateTransaction)
		// return
	}

	writer.Response(SuccessResponse{
		Status:  true,
		Message: "success",
		Data:    input,
	})
}

// InsertTopup do database insert process
func (a *App) InsertTopup(input *InputCash) error {
	var err error
	err = a.CheckAndRetryConnection()
	if err != nil {
		return err
	}

	stmt, err := a.DB.Prepare("INSERT INTO `topup` (`username`, `reference_id`, `amount` , `bonus`,  `flag_queue`) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	res, err := a.Exec(stmt, input.Username, input.ReferenceCode, input.Amount, input.Bonus, 0)
	if err != nil {
		return err
	}
	defer stmt.Close()

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	input.ID = lastID
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt < 1 {
		return fmt.Errorf("0 row effect")
	}

	return nil
}

// UpdateStatusPutMessage do database update status
func (a *App) UpdateStatusPutMessage(input *InputCash) error {
	var err error
	err = a.CheckAndRetryConnection()
	if err != nil {
		return err
	}

	stmt, err := a.DB.Prepare("UPDATE `topup` SET flag_queue=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	updateTime := time.Now()
	input.UpdatedAt = &updateTime

	res, err := a.Exec(stmt, 1, input.ID)
	if err != nil {
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowCnt < 1 {
		return fmt.Errorf("0 row effect")
	}

	return err
}
