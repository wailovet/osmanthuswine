package core

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"github.com/wailovet/osmanthuswine/src/helper"
)

type Config struct {
	Port          string `json:"port"`
	Host          string `json:"host"`
	CrossDomain   string `json:"cross_domain"`
	PostMaxMemory int64  `json:"post_max_memory"`
}

func (c *Config) ReadConfig(file string) {
	configText, err := ioutil.ReadFile("./config/main.json")
	if err != nil {
		helper.GetInstanceLog().Out("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}
	json.Unmarshal(configText, *c)

	if c.PostMaxMemory <= 0 {
		c.PostMaxMemory = 1024 * 1024 * 10
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "8808"
	}

}
