package alarm

import (
	"encoding/json"
	"fmt"
	"github.com/ChangQingAAS/Learing_GIN/common/function"
	"path/filepath"
	"runtime"
	"strings"
)

type errorString struct {
	s string
}

type errorInfo struct {
	Time     string `json:"time"`
	Alarm    string `json:"alarm"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
	Funcname string `json:"funcname"`
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	alarm("INFO", text)
	return &errorString{text}
}

// 发邮件
func Email(text string) error {
	alarm("EMAIL", text)
	return &errorString{text}
}

// 发短信
func Sms(text string) error {
	alarm("SMS", text)
	return &errorString{text}
}

// 发微信
func WeChat(text string) error {
	alarm("WX", text)
	return &errorString{text}
}

// function to alarm
func alarm(level string, str string) {
	currentTime := function.GetTimeStr()
	fileName, line, functionName := "?", 0, "?"
	pc, fileName, line, ok := runtime.Caller(2)

	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		functionName = filepath.Ext(functionName)
		functionName = strings.TrimPrefix(functionName, ".")
	}

	var msg = errorInfo{
		Time:     currentTime,
		Alarm:    level,
		Message:  str,
		Filename: fileName,
		Line:     line,
		Funcname: functionName,
	}

	jsons, errs := json.Marshal(msg)

	if errs != nil {
		fmt.Println("json marshal error: ", errs)
	}

	errorJsonInfo := string(jsons)

	fmt.Println(errorJsonInfo)

	if level == "EMAIL" {

	} else if level == "SMS" {
		// 执行发短信
	} else if level == "WX" {

	} else if level == "INFO" {
		// 执行记日志
	}
}
