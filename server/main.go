package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
)

func main() {
	// 捕捉 Ctrl+c 和 kill 信号，写入signalChan
	signalChan := make(chan os.Signal, 1)
	//signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(signalChan)
	// 此处执行处理逻辑
	go func() {
		log.Info("begin monitor ...")

		for _, receive := range GetSetting().Rec {
			monitorStart(receive.Ip, receive.Port, receive.Url)
		}
	}()
	// signalChan阻塞进程
	//log.Warn(<-signalChan)
	// 捕捉信号后在Exit函数中处理信息，例如内存持久化等信息防止丢失
	log.Fatal(" monitor exit  ,", <-signalChan)
}

func init() {
	var config_dir string
	var log_level string
	flag.StringVar(&config_dir, "config_dir", "monitor-config.json", "config file path")
	flag.StringVar(&log_level, "log_level", "info", "logger level . default: `info`.  debug > info > error")
	flag.Parse()
	flag.Usage()        //打印使用帮助
	log_init(log_level) //初始化日志框架
	banner()            //启动打印

	dir, _ := os.Getwd()
	log.Info("working dir:", dir)
	log.Info("starting server ...")

	InitSetting(config_dir)
}
