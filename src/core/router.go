package core

import (
	"github.com/pkg/errors"
	"log"
	"reflect"
	"strings"
	"unicode"
)

type RouterManage struct {
	RegisteredData map[string]reflect.Type
}

var instanceRouterManage *RouterManage

func GetInstanceRouterManage() *RouterManage {
	if instanceRouterManage == nil {
		instanceRouterManage = &RouterManage{} // not thread safe
		instanceRouterManage.RegisteredData = make(map[string]reflect.Type)
	}
	return instanceRouterManage
}

func (rm *RouterManage) Registered(i interface{}) {
	t := reflect.ValueOf(i)
	GetInstanceRouterManage().RegisteredData[t.Type().String()] = reflect.Indirect(t).Type()
}

func (rm *RouterManage) GetModuleName(name string) string {
	if name == "" {
		return "index"
	}
	return strings.ToLower(name)
}
func (rm *RouterManage) GetControllerName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}
func (rm *RouterManage) GetFunName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}

func (rm *RouterManage) RouterSend(urlPath string, request Request, response Response) error {
	tmp := strings.Split(urlPath, ".")
	if len(tmp) > 1 {
		urlPath = strings.Join(tmp[0:len(tmp)-1], ".")
	}

	sar := strings.Split(urlPath, "/")
	for len(sar) < 5 {
		sar = append(sar, "")
	}
	//过滤非 /Api开头的
	module := rm.GetModuleName(sar[2])
	controller := rm.GetControllerName(sar[3])
	fun := rm.GetFunName(sar[4])

	ctr := "*" + module + "." + controller

	_, ok := rm.RegisteredData[ctr]
	if !ok {
		return errors.New("未注册该组件:" + ctr)
	}

	vc := reflect.New(rm.RegisteredData[ctr])
	f := vc.MethodByName(fun)
	if !f.IsValid() {
		return errors.New("组件找不到相应function:" + fun)
	}

	init := vc.MethodByName("ControllerInit")
	if init.IsValid() {
		init.Call([]reflect.Value{reflect.ValueOf(request), reflect.ValueOf(response)})
		f.Call(nil)
	} else {
		//兼容模式
		log.Println("兼容模式")
		f.Call([]reflect.Value{reflect.ValueOf(request), reflect.ValueOf(response)})
	}

	return nil
}
