package registered

import "reflect"

var RegisteredData = make(map[string]reflect.Value)

func Registered(i interface{}) {
	t := reflect.TypeOf(i)
	RegisteredData[t.String()] = reflect.ValueOf(i)
	println("registered:", t.String())
}
