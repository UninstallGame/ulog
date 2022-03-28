package ulog

import (
	"fmt"
	"github.com/UninstallGame/ulog/ulog/loglevel"
	"log"
	"os"
)

type ULog struct {
	logLevel int
	logger   *log.Logger
}

func New(logFile string) *ULog {
	uLog := &ULog{}
	uLog.SetLogLevel(loglevel.Debug)
	uLog.init(logFile)
	return uLog
}

func (l *ULog) SetLogLevel(logLevel int) {
	l.logLevel = logLevel
}

func (l *ULog) Debug(text string) {
	if l.logLevel != 0 {
		return
	}
	l.logger.Println(fmt.Sprintf("[DEBUG] %v", text))
}

func (l *ULog) Info(text string) {
	if l.logLevel > 1 {
		return
	}
	l.logger.Println(fmt.Sprintf("[INFO] %v", text))
}

func (l *ULog) Warning(text string) {
	if l.logLevel > 2 {
		return
	}
	l.logger.Println(fmt.Sprintf("[WARNING] %v", text))
}

func (l *ULog) Error(text string, err error) {
	if l.logLevel > 3 {
		return
	}
	l.logger.Println(fmt.Sprintf("[ERROR] %v. Error: %v", text, err.Error()))
}

func (l *ULog) Fatal(text string, err error) {
	l.logger.Println(fmt.Sprintf("[FATAL] %v. Error: %v", text, err.Error()))
}

func (l *ULog) init(logFile string) {
	f, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	l.logger = log.New(f, "", log.LstdFlags)
}
