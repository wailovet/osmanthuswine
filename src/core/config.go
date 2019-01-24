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
	ApiRouter     string `json:"api_router"`
	PostMaxMemory int64  `json:"post_max_memory"`
	Db struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Name     string `json:"name"`
		User     string `json:"user"`
		Password string `json:"password"`
		Charset  string `json:"charset"`
	} `json:"db"`
}

var instanceConfig *Config
var configFile = "./config/main.json"

func SetConfigFile(c string) {
	configFile = c
}

func SetConfig(c *Config) {
	instanceConfig = c
}

func GetInstanceConfig() *Config {
	if instanceConfig == nil {
		instanceConfig = &Config{} // not thread safe
		instanceConfig.ReadConfig(configFile)
	}
	return instanceConfig
}

func (c *Config) ReadConfig(file string) {
	configText, err := ioutil.ReadFile(file)
	if err != nil {
		helper.GetInstanceLog().Out("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}
	err = json.Unmarshal(configText, c)
	if err != nil {
		helper.GetInstanceLog().Out("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}
	if c.PostMaxMemory <= 0 {
		c.PostMaxMemory = 1024 * 1024 * 10
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "8808"
	}
	if c.ApiRouter == "" {
		c.ApiRouter = "/Api/*"
	}
	if c.CrossDomain == "" {
		c.CrossDomain = "*"
	}

}
