package context

import (
	"cabal-api/common"
	"cabal-api/errors"
	"net/http"
	"time"
)

type (

	// TopupCallBack struct of data model topup table
	TopupCallBack struct {
		ID            int64     `json:"id"`
		Username      string    `json:"username"`
		Amount        int       `json:"amount"`
		Bonus         int       `json:"bonus"`
		ReferenceID   string    `json:"reference_id"`
		FlagQueue     bool      `json:"flag_queue"`
		FlagSent      bool      `json:"flag_sent"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		CreatedString string    `json:"-"`
		UpdatedString string    `json:"-"`
	}
	// SuccessTopupCallbackResponse struct of response when success case
	SuccessTopupCallbackResponse struct {
		Status  bool           `json:"status"`
		Message string         `json:"message"`
		Data    *TopupCallBack `json:"data"`
	}
)

// CashCallBackHandler do handler http request
func (a *App) CashCallBackHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	r.ParseForm()

	writer := common.Writer{
		Resp: w,
		Logger: common.TransactionLogInfo{
			RefCode: common.RandomRefCode(5),
		},
	}
	input := &TopupCallBack{}
	// validate ReferenceID
	if input.ReferenceID = r.FormValue("refcode"); input.ReferenceID == "" {
		writer.ResponseError(http.StatusBadRequest, errors.ErrorReferenceNotFound)
		return
	}

	err = a.GetTopupWithReferenceCode(input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorGetTopupWithReferenceCode)
		return
	}

	input.FlagSent = true

	err = a.UpdateTopupSentWithID(input)
	if err != nil {
		writer.Logger.Error = err
		writer.Logger.LogWriter()
		writer.ResponseError(http.StatusBadRequest, errors.ErrorUpdateTransaction)
		return
	}

	writer.Response(SuccessTopupCallbackResponse{
		Status:  true,
		Message: "success",
		Data:    input,
	})

	return
}

// GetTopupWithReferenceCode do select topup
func (a *App) GetTopupWithReferenceCode(input *TopupCallBack) error {
	var err error
	err = a.CheckAndRetryConnection()
	if err != nil {
		return err
	}

	err = a.DB.QueryRow("select * from topup where reference_id = ?", input.ReferenceID).Scan(&input.ID, &input.Username, &input.ReferenceID, &input.Amount, &input.Bonus, &input.FlagQueue, &input.FlagSent, &input.CreatedString, &input.UpdatedString)
	if err != nil {
		return err
	}

	if input.CreatedAt, err = time.Parse(time.RFC3339, input.CreatedString); err != nil {
		return err
	}
	if input.UpdatedAt, _ = time.Parse(time.RFC3339, input.UpdatedString); err != nil {
		return err
	}

	return nil
}

// UpdateTopupSentWithID do select topup
func (a *App) UpdateTopupSentWithID(input *TopupCallBack) error {
	var err error
	err = a.CheckAndRetryConnection()
	if err != nil {
		return err
	}

	stmt, err := a.DB.Prepare("UPDATE `topup` SET flag_sent=? WHERE id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(input.FlagSent, input.ID)
	if err != nil {
		return err
	}

	return nil
}
