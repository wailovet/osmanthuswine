package owstruct

import (
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
)

type Config struct {
	Port        string `json:"port"`
	Host        string `json:"host"`
	CrossDomain string `json:"cross_domain"`
	MaxMemory   int64  `json:"post_max_memory"`
}

func (c *Config) ReadConfig(file string) {
	configText, err := ioutil.ReadFile("./config/main.json")
	if err != nil {
		log.Println("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}
	json.Unmarshal(configText, *c)
}
