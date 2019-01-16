package core

import "reflect"

var manage *RouterManage

type RouterManage struct {
	RegisteredData map[string]reflect.Value
}

func GetInstance() *RouterManage {
	if manage == nil {
		manage = new(RouterManage)
	}
	return manage
}

func (rm *RouterManage) Registered(i interface{}) {
	t := reflect.TypeOf(i)
	GetInstance().RegisteredData[t.String()] = reflect.ValueOf(i)
}
