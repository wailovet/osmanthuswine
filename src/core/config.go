package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port                string          `json:"port"`
	Host                string          `json:"host"`
	CrossDomain         string          `json:"cross_domain"`
	ApiRouter           string          `json:"api_router"`
	StaticRouter        string          `json:"static_router"`
	IsStaticStripPrefix bool            `json:"is_static_strip_prefix"`
	StaticFileSystem    http.FileSystem `json:"static_file_system"`
	PostMaxMemory       int64           `json:"post_max_memory"`
	Db                  ConfigDb        `json:"db"`
	Redis               ConfigRedis     `json:"redis"`
	UpdateDir           string          `json:"update_dir"`
	UpdatePath          string          `json:"update_path"`
}

type ConfigDb struct {
	Host        string            `json:"host"`
	Port        string            `json:"port"`
	User        string            `json:"user"`
	Password    string            `json:"password"`
	Name        string            `json:"name"`
	Prefix      string            `json:"prefix"`
	MaxOpenConn int               `json:"max_open_conn"`
	Params      map[string]string `json:"params"`
	Debug       bool              `json:"debug"`
}
type ConfigRedis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

var instanceConfig *Config
var configFile = "./config.json"
var privateConfigFile = "./private.json"

func SetConfigFile(c string) {
	configFile = c
}

func SetConfig(c *Config) {
	instanceConfig = c
}

func GetInstanceConfig() *Config {
	if instanceConfig == nil {
		instanceConfig = &Config{
			Host:             "localhost",
			Port:             "8808",
			ApiRouter:        "/Api/*",
			StaticRouter:     "/*",
			StaticFileSystem: http.Dir("static"),
			CrossDomain:      "*",
			PostMaxMemory:    1024 * 1024 * 10,
			Db: ConfigDb{
				Host:        "localhost",
				Port:        "3306",
				User:        "root",
				Password:    "root",
				Name:        "test",
				Prefix:      "",
				MaxOpenConn: 500,
				Params: map[string]string{
					"charset":   "utf8mb4",
					"parseTime": "true",
				},
				Debug: true,
			},
			Redis: ConfigRedis{
				Addr:     "localhost:6379",
				Password: "",
				Db:       0,
			},
			UpdateDir:  "",
			UpdatePath: "",
		}

		instanceConfig.ReadConfig(configFile)
		instanceConfig.ReadPrivateConfig(privateConfigFile)
	}
	return instanceConfig
}

func (c *Config) ReadConfig(file string) {
	configText, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("配置文件读取错误,启动默认配置:", err.Error())
		return
	}
	err = json.Unmarshal(configText, c)
	if err != nil {
		log.Println("配置文件错误,启动失败:", err.Error())
		os.Exit(0)
	}
}

func (c *Config) ReadPrivateConfig(file string) {
	configText, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("未加载", privateConfigFile, ":", err.Error())
		return
	}
	err = json.Unmarshal(configText, c)
	if err != nil {
		log.Println("未加载", privateConfigFile, ":", err.Error())
	}
}
