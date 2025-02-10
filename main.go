package main

import (
	"reflect"
	"runtime"
)

func main() {
	t := reflect.ValueOf(runtime.GOROOT()).MethodByName("SplitList").Call([]reflect.Value{})
	
	println(t.Len())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		println(method.Name)
	}
}