package common

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

type (
	//TransactionLogInfo struct of log
	TransactionLogInfo struct {
		RefCode    string
		URL        string
		Method     string
		Input      string
		Output     string
		Error      error
		Body       string
		StatusCode int
	}
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//FailOnError is print log and save log
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return
	}
}

//LogWriter is center write
func (logInfo *TransactionLogInfo) LogWriter() {

	if logInfo.Error != nil {
		LogErrorChannel <- LogStringerError(logInfo)
	}

	LogRequestChannel <- LogStringer(logInfo)
}

//LogStringer is format log
func LogStringer(logInfo *TransactionLogInfo) string {

	logString := fmt.Sprintf("RefCode:%s", logInfo.RefCode)
	if logInfo.URL != "" {
		logString += fmt.Sprintf(" URL:%s", logInfo.URL)
	}

	if logInfo.Method != "" {
		logString += fmt.Sprintf(" Method:%s", logInfo.Method)
	}

	if logInfo.Body != "" {
		logString += fmt.Sprintf(" Body:%s", logInfo.Body)
	}

	if logInfo.Input != "" {
		logString += fmt.Sprintf(" Request:%s", logInfo.Input)
	}

	if logInfo.StatusCode != 0 {
		logString += fmt.Sprintf(" StatusCode:%d", logInfo.StatusCode)
	}

	if logInfo.Output != "" {
		logString += fmt.Sprintf(" Response:%s", logInfo.Output)
	}

	return logString
}

//LogStringerError is format log error
func LogStringerError(logInfo *TransactionLogInfo) string {
	_, path, lineNumber, _ := runtime.Caller(3)
	paths := strings.Split(path, "/")
	filename := fmt.Sprintf("%v:%v", paths[len(paths)-1], lineNumber)

	logString := fmt.Sprintf("RefCode:%s URL:%s Request:%s Error: %s %s", logInfo.RefCode, logInfo.URL, logInfo.Input, filename, logInfo.Error.Error())

	return logString
}

//RandomRefCode get random string
func RandomRefCode(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
