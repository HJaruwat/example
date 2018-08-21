package common

import (
	"encoding/json"
	"net/http"
)

// Writer instace of
type Writer struct {
	Resp   http.ResponseWriter
	Logger TransactionLogInfo
}

// ErrorResponse instace of
type ErrorResponse struct {
	Error            string      `json:"error"`
	ErrorDescription string      `json:"error_description"`
	ErrorData        interface{} `json:"error_data,omitempty"`
}

// ResponseError is
func (w Writer) ResponseError(status int, err *ErrorResponse) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(status)

	byteString, _ := json.Marshal(&err)
	w.Resp.Write(byteString)
	log := &TransactionLogInfo{
		RefCode: w.Logger.RefCode,
	}

	log.Output = string(byteString)
	log.LogWriter()
	return
}

// Response is
func (w Writer) Response(data interface{}) {
	w.Resp.Header().Set("Content-Type", "application/json")
	w.Resp.WriteHeader(200)
	log := &TransactionLogInfo{
		RefCode: w.Logger.RefCode,
	}

	if data != nil {
		byteString, _ := json.Marshal(&data)
		w.Resp.Write(byteString)

		log.Output = string(byteString)
		log.LogWriter()
	} else {
		log.Output = "OK"
		log.LogWriter()
	}
	return
}
