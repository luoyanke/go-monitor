package main

import (
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
)

type Setting struct {
	NodeName  string    `json:"nodeName"`
	CurrentIp string    `json:"currentIp"`
	Rec       []Receive `json:"receive"`
}

type Receive struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Url  string `json:"url"`
}

var setting *Setting

func GetSetting() *Setting {
	return setting
}
func InitSetting(config_dir string) *Setting {

	setting = new(Setting)
	file, err := os.Open(config_dir)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	decoder.Decode(setting)

	log.Info("init setting ...")
	return setting
}

func Set(settingObj Setting) {

}
