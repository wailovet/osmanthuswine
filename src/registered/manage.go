package registeredmanage

import "reflect"

var manage *RegisteredManage

type RegisteredManage struct {
	RegisteredData map[string]reflect.Value
}

func GetInstance() *RegisteredManage {
	if manage == nil {
		manage = new(RegisteredManage)
	}
	return manage
}

func (rm *RegisteredManage) Registered(i interface{}) {
	t := reflect.TypeOf(i)
	GetInstance().RegisteredData[t.String()] = reflect.ValueOf(i)
}
