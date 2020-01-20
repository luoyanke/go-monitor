package main

import (
	"encoding/json"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	log "github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"time"
)

type SystemParam struct {
	HostName   string    `json:"hostName"`
	NodeIp     string    `json:"nodeIp"`
	CpuTotal   float64   `json:"cpuTotal"`
	CpuPercent []float64 `json:"cpuPercent"`
	MemInfo    MemInfo   `json:"memInfo"`
	Time       int64     `json:"time"`
}

//内存数据
type MemInfo struct {
	Total       uint64  `json:"total"`       //系统总的可用物理内存大小
	Available   uint64  `json:"available"`   //还可以被 应用程序 使用的物理内存大小
	Free        uint64  `json:"free"`        //还有多少物理内存可用
	Used        uint64  `json:"used"`        //已被使用的物理内存大小
	UsedPercent float64 `json:"usedPercent"` //已被使用的物理内存大小 百分比
}

func GetSystemParam(ip string, port int) *SystemParam {

	systemParam := new(SystemParam)
	systemParam.Time = time.Now().UTC().UnixNano()
	systemParam.getHostName()
	systemParam.getIp(ip, port)
	systemParam.getMem()
	systemParam.getCpu()
	log.Debug(systemParam)
	return systemParam
}

//通过udp检测对外通讯所使用的IP
func getOutboundIP(ip string, port int) string {
	url := ip + ":" + strconv.Itoa(port)
	conn, err := net.Dial("udp", url)
	if err != nil {
		log.Error(err)
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

func (systemParam *SystemParam) getIp(ip string, port int) {
	if GetSetting().CurrentIp == "" {
		GetSetting().CurrentIp = getOutboundIP(ip, port)
	}
	systemParam.NodeIp = GetSetting().CurrentIp
}

func (systemParam *SystemParam) getHostName() {
	if GetSetting().NodeName != "" {
		systemParam.HostName = GetSetting().NodeName
		return
	}
	h, err := os.Hostname()
	if err != nil {
		log.Error(err)
		return
	}
	systemParam.HostName = h
}

func (systemParam *SystemParam) getMem() {
	stat, err := mem.VirtualMemory()
	info := new(MemInfo)
	if err != nil {
		log.Error(err)
		return
	}
	info.Available = stat.Available
	info.Free = stat.Free
	info.Total = stat.Total
	info.Used = stat.Used
	info.UsedPercent = stat.UsedPercent
	systemParam.MemInfo = *info

}

func (systemParam *SystemParam) getCpu() {
	percent, _ := cpu.Percent(time.Second, true)
	systemParam.CpuPercent = make([]float64, len(percent))
	for i := range percent {
		systemParam.CpuPercent[i] = percent[i]
	}

}

func (systemParam SystemParam) String() string {

	s, err := json.Marshal(systemParam)
	if err != nil {
		log.Error(err)
		nilResult, _ := json.Marshal(SystemParam{})
		return string(nilResult)

	}
	return string(s)
}
