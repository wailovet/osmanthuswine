package osmanthuswine

import (
	"github.com/wailovet/osmanthuswine/src/core"
	"github.com/wailovet/osmanthuswine/src/helper"
	"strings"
	"testing"
	"time"
)

func TestHandleFunc(t *testing.T) {

	core.GetInstanceConfig().Port = "39201"
	HandleFunc("/213", func(request core.Request, response core.Response) {
		response.DisplayByData("abcabc," + request.REQUEST["a"])
	})

	go func() {
		Run()
	}()
	t.Log("TestHandleFunc start")
	time.Sleep(time.Second * 3)
	data := helper.Get("http://127.0.0.1:39201/213", "")
	if strings.Index(data, "abcabc") == -1 {
		t.Error("HandleFunc error", data)
	}
	data = helper.Get("http://127.0.0.1:39201/213?a=dfgsdafasf", "")
	if strings.Index(data, "dfgsdafasf") == -1 {
		t.Error("HandleFunc error", data)
	}
	InstanceListener.Close()
}
