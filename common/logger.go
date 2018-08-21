package common

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

var (
	//LogChannel of
	LogChannel chan string

	//LogRequestChannel of
	LogRequestChannel chan string

	//LogErrorChannel of
	LogErrorChannel chan string

	filename string
)

func init() {
	LogChannel = make(chan string)
	LogRequestChannel = make(chan string)
	LogErrorChannel = make(chan string)

	go LogLogger(LogChannel)
	go LogRequestLogger(LogRequestChannel)
	go LogErrorLogger(LogErrorChannel)

	_, filename, _, _ = runtime.Caller(0)
	filename = strings.Split(filename, "common")[0]
	filename = path.Dir(filename)

}

//LogLogger is log file
func LogLogger(ch chan string) {
	for {
		logMsg := <-ch
		currentTime := time.Now()
		filename := fmt.Sprintf("%v/%04d-%02d.log", filename+"/logs/transaction", currentTime.Year(), currentTime.Month())
		outfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			log.Println("can't open file", filename)
			log.Println(logMsg)
			continue
		}
		logger := log.New(outfile, "", log.LstdFlags)
		logger.Println(logMsg)
	}
}

//LogRequestLogger is log file
func LogRequestLogger(ch chan string) {
	for {
		logMsg := <-ch
		currentTime := time.Now()
		filename := fmt.Sprintf("%v/%04d-%02d.log", filename+"/logs/request", currentTime.Year(), currentTime.Month())
		outfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			log.Println("can't open file", filename)
			log.Println(logMsg)
			continue
		}
		logger := log.New(outfile, "", log.LstdFlags)
		logger.Println(logMsg)
	}
}

//LogErrorLogger is log file
func LogErrorLogger(ch chan string) {
	for {
		logMsg := <-ch
		currentTime := time.Now()
		filename := fmt.Sprintf("%v/%04d-%02d.log", filename+"/logs/error", currentTime.Year(), currentTime.Month())
		outfile, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			log.Println("can't open file", filename)
			log.Println(logMsg)
			continue
		}
		logger := log.New(outfile, "", log.LstdFlags)
		logger.Println(logMsg)
	}
}
