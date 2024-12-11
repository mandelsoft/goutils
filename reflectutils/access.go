package reflectutils

import (
	"reflect"
	"unsafe"
)

func UnexportedFieldByName(p any, name string) reflect.Value {
	v := reflect.ValueOf(p).Elem().FieldByName(name)
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func UnexportedFieldByIndex(p any, n int) reflect.Value {
	v := reflect.ValueOf(p).Elem().Field(n)
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}
