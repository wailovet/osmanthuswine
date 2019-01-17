package osmanthuswine

import (
	"reflect"
	"strings"
	"unicode"
)

var manage *RouterManage

type RouterManage struct {
	RegisteredData map[string]reflect.Value
}

func GetInstanceRouterManage() *RouterManage {
	if manage == nil {
		manage = new(RouterManage)
	}
	return manage
}

func (rm *RouterManage) Registered(i interface{}) {
	t := reflect.TypeOf(i)
	GetInstanceRouterManage().RegisteredData[t.String()] = reflect.ValueOf(i)
}

func (rm *RouterManage) getModuleName(name string) string {
	if name == "" {
		return "index"
	}
	return strings.ToLower(name)
}
func (rm *RouterManage) getControllerName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}
func (rm *RouterManage) getFunName(name string) string {
	if name == "" {
		return "Index"
	}
	for i, v := range name {
		return string(unicode.ToUpper(v)) + name[i+1:]
	}
	return "Index"
}

type RouterError struct {
	What string
}

func (e RouterError) Error() string {
	return e.What
}

func (rm *RouterManage) RouterSend(urlPath string, request Request, response Response) (error) {
	sar := strings.Split(urlPath, "/")
	for ; len(sar) < 5; {
		sar = append(sar, "")
	}
	//过滤非 /Api开头的
	module := rm.getModuleName(sar[2])
	controller := rm.getControllerName(sar[3])
	fun := rm.getFunName(sar[4])

	ctr := "*" + module + "." + controller

	_, ok := rm.RegisteredData[ctr]

	if !ok {
		return RouterError{
			"未注册该组件:" + ctr,
		}
	}

	f := rm.RegisteredData[ctr].MethodByName(fun)
	if !f.IsValid() {
		return RouterError{
			"组件找不到相应function:" + fun,
		}
	}
	f.Call([]reflect.Value{reflect.ValueOf(request), reflect.ValueOf(response)})

	return nil
}