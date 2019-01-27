package helper

import (
	"testing"
	"github.com/wailovet/osmanthuswine/src/core"
)

func TestGetStructByDb(t *testing.T) {

	instanceConfig := &core.Config{
		Host:          "localhost",
		Port:          "8808",
		ApiRouter:     "/Api/*",
		CrossDomain:   "*",
		PostMaxMemory: 1024 * 1024 * 10,
		Db: struct {
			Host        string            `json:"host"`
			Port        string            `json:"port"`
			User        string            `json:"user"`
			Password    string            `json:"password"`
			Name        string            `json:"name"`
			MaxOpenConn int               `json:"max_open_conn"`
			Params      map[string]string `json:"params"`
		}{
			MaxOpenConn: 500,
			Params: map[string]string{
				"charset":   "utf8mb4",
				"parseTime": "true",
			},
		},
	}
	core.SetConfig(instanceConfig)

	s, _ := GetStructByDb("ox_system", "test", "system")
	println(s)
}
