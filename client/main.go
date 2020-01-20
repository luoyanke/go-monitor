package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"
)

/**
命令行启动server等操作
gomcli -start --serverpath server/go_monitor_server
gomcli -stop
gomcli -restart --serverpath server/go_monitor_server
*/
var (
	go_monitor_serverd_path string // 监听服务路径
	log_level               string //服务器日志级别

	help  bool
	start bool
	stop  bool
)

func main() {
	dir, _ := os.Getwd() //获取client当前路径
	flag.BoolVar(&help, "help", false, "get help")
	flag.BoolVar(&start, "start", false, "start monitor server")
	flag.BoolVar(&stop, "stop", false, "stop monitor server")
	flag.StringVar(&go_monitor_serverd_path, "serverpath", dir, "server absolute path")
	flag.StringVar(&log_level, "log_level", "info", "logger level. default: `info`.  debug > info > error")

	flag.Parse()

	if help {
		flag.Usage()
		return
	} else if start {
		stopOp()
		startOp()
		return
	} else if stop {
		stopOp()
		return
	} else {
		flag.Usage()
		return
	}

}

func startOp() {

	if log_level != "info" && log_level != "debug" && log_level != "error" {
		fmt.Println("error log level %s", log_level)
		flag.Usage()
		os.Exit(0)
	}

	//给go_monitor_server操作权限
	syscall.Chmod(go_monitor_serverd_path+"/go_monitor_server", 755)

	//直接执行，不要通过shell 方式运行
	sh_command := exec.Command(go_monitor_serverd_path+"/go_monitor_server", "-config_dir",
		go_monitor_serverd_path+"/go-monitor-config.json", "-log_level", log_level, "&")
	//sh_command.Stdout = os.Stdout //查看服务端的日志
	sh_err := sh_command.Start()
	if sh_err != nil {
		log.Fatal(sh_err)
	}
	//这里只是调试用，没有实际意义，因为Goland中调试会直接退出。这里休眠几秒是为了看看日志
	fmt.Println("waiting server start ...")
	time.Sleep(time.Second * 10)
	fmt.Println("server start success.")
}

func stopOp() {
	//复杂的命令通过shell脚本执行，要不然会出错
	command := exec.Command("sh", "-c", "ps -ef|grep go_monitor_server |grep -v grep |awk '{print $2}' |xargs kill -9")
	err := command.Run()
	if err != nil {
		log.Println(err)
	}
}
