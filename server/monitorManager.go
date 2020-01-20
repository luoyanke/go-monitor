package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

const realUrl = "http://%v:%d%v"

func monitorStart(ip string, port int, url string) {
	defer panicLogger() //panic 异常日志

	_realUrl := fmt.Sprintf(realUrl, ip, port, url)
	log.Info("start send system info to ", _realUrl)
	for {
		param := GetSystemParam(ip, port)
		resp, err := http.Post(_realUrl, "application/json;charset=utf-8", strings.NewReader(param.String()))
		if err != nil {
			log.Error(err)
			time.Sleep(time.Millisecond * 5)
			continue
		}
		if resp.StatusCode != 200 {
			log.Error(resp.Status, resp.Body)
		}
		resp.Body.Close()

		log.Debug(param.String())
		time.Sleep(time.Millisecond * 5)
	}

}

//panic 异常日志
func panicLogger() {
	if err := recover(); err != nil {
		stack := string(debug.Stack())
		log.Panic(err, "\n", stack)
	}
}
