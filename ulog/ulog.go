package ulog

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ULog struct {
	logLevel   int
	logFile    string
	tlgBotUrl  string
	tlgGroupId string
}

func New(logFile string) *ULog {
	return &ULog{logLevel: 0,
		logFile: logFile}
}

func (l *ULog) InitTlgBot(botToken, groupId string) {
	l.tlgBotUrl = fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", botToken)
	l.tlgGroupId = groupId
}

func (l *ULog) SendToTelegram(text string) {
	type postData struct {
		ChatId string `json:"chat_id"`
		Text   string `json:"text"`
	}
	x := postData{
		ChatId: l.tlgGroupId,
		Text:   text,
	}
	data := []byte(fmt.Sprintf("%v", x))
	r := bytes.NewReader(data)
	_, err := http.Post(l.tlgBotUrl, "application/json", r)
	if err != nil {
	}
}

func (l *ULog) SetLogLevel(logLevel int) {
	l.logLevel = logLevel
}

func (l *ULog) Debug(text string) {
	if l.logLevel != 0 {
		return
	}
	l.log("[DBG] ", text)
}

func (l *ULog) Info(text string) {
	if l.logLevel > 1 {
		return
	}
	l.log("[INF] ", text)
}

func (l *ULog) Warn(text string) {
	if l.logLevel > 2 {
		return
	}
	l.log("[WRN] ", text)
}

func (l *ULog) Err(text string) {
	if l.logLevel > 3 {
		return
	}
	l.log("[ERR] ", text)
}

func (l *ULog) Fatal(text string) {
	l.log("[FTL] ", text)
}

func (l *ULog) log(prefix, text string) {
	f, err := os.OpenFile(l.logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, prefix, log.LstdFlags)
	logger.Println(text)
}
