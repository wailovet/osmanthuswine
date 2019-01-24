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
		Host     string            `json:"host"`
		Port     string            `json:"port"`
		User     string            `json:"user"`
		Password string            `json:"password"`
		Name     string            `json:"name"`
		Params   map[string]string `json:"params"`
	} `json:"db"`
}

var instanceConfig *Config
var configFile = "./config.json"

func SetConfigFile(c string) {
	configFile = c
}

func SetConfig(c *Config) {
	instanceConfig = c
}

func GetInstanceConfig() *Config {
	if instanceConfig == nil {
		instanceConfig = &Config{
			Host:          "localhost",
			Port:          "8808",
			ApiRouter:     "/Api/*",
			CrossDomain:   "*",
			PostMaxMemory: 1024 * 1024 * 10,
			Db: struct {
				Host     string            `json:"host"`
				Port     string            `json:"port"`
				User     string            `json:"user"`
				Password string            `json:"password"`
				Name     string            `json:"name"`
				Params   map[string]string `json:"params"`
			}{
				Host:     "localhost",
				Port:     "3306",
				User:     "root",
				Password: "root",
				Name:     "test",
				Params: map[string]string{
					"charset": "utf8mb4",
				},
			},
		}

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
}
