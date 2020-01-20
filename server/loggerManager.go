package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

func log_init(log_level string) {
	if log_level == "info" {
		log.SetLevel(log.InfoLevel)
	}
	if log_level == "debug" {
		log.SetLevel(log.DebugLevel)
	}
	log.AddHook(newMonitorHook())
	go loggerFileTask() //监控日志总体大小，不要写满磁盘
}

func newMonitorHook() *MonitorHook {

	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	os.Mkdir("log", os.ModePerm) //有error不用管
	logFile, _ := os.OpenFile("log/"+getCurrentDate()+".server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	return &MonitorHook{getCurrentDate(), logFile}
}

type MonitorHook struct {
	LogDate string
	LogFile *os.File
}

func (mHook *MonitorHook) Levels() []log.Level {
	return log.AllLevels
}

func (mHook *MonitorHook) Fire(entry *log.Entry) error {
	var lock sync.Mutex
	lock.Lock()
	if getCurrentDate() != mHook.LogDate {
		mHook.LogFile.Close() //关闭之前的日志文件
		mHook.LogFile, _ = os.OpenFile("log/"+getCurrentDate()+".server.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		mHook.LogDate = getCurrentDate()
	}
	msg, _ := entry.String()
	mHook.LogFile.WriteString(msg)
	lock.Unlock()
	return nil
}

func getCurrentDate() string {
	format := time.Now().Format("2006-01-02")
	return format
}

func loggerFileTask() {
	defer func() { //panic 异常日志
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			log.Panic(err, "\n", stack)
		}
	}()

	for {
		fileInfos, _ := ioutil.ReadDir("log")
		for _, file := range fileInfos {
			if time.Now().Unix()-file.ModTime().Unix() > 86400*2 {
				os.Remove("log/" + file.Name())
			}
		}
		time.Sleep(time.Hour * 4)
	}
}
