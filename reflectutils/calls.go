package reflectutils

import (
	"fmt"
	"reflect"

	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/sliceutils"
)

// GetInterfaceMethod gets the method of an interface
// with one method.
func GetInterfaceMethod[M any]() reflect.Method {
	t := generics.TypeOf[M]()
	if t.NumMethod() != 1 {
		panic(fmt.Sprintf("invalid setter type %s", t))
	}
	return t.Method(0)
}

// CallMethodByInterfaceVA calls a void method on object o with
// one argument a. The method is specified by the interface
// M, which should implement exactly one appropriate method.
func CallMethodByInterfaceVA[M, B any](o B, a interface{}) {
	CallMethodByNameVA[B](GetInterfaceMethod[M]().Name, o, a)
}

func CallMethodByInterfaceV[M, B any](o B, args ...interface{}) {
	CallMethodByNameV[B](GetInterfaceMethod[M]().Name, o, args...)
}

func CallMethodByInterface[M, B any](o B, args ...interface{}) []interface{} {
	return CallMethodByName[B](GetInterfaceMethod[M]().Name, o, args...)
}

func CallMethodByNameVA[B any](n string, o B, a interface{}) {
	reflect.ValueOf(o).MethodByName(n).Call([]reflect.Value{reflect.ValueOf(a)})
}

func CallMethodByNameV[B any](n string, o B, args ...interface{}) {
	v := sliceutils.Transform(args, reflect.ValueOf)
	reflect.ValueOf(o).MethodByName(n).Call(v)
}

func CallMethodByName[B any](n string, o B, args ...interface{}) []interface{} {
	v := sliceutils.Transform(args, reflect.ValueOf)
	r := reflect.ValueOf(o).MethodByName(n).Call(v)
	return sliceutils.Transform(r, MapValueToInterface)
}

// GetInterfaceMethodFor used an interface with one method to determine
// the method to look for. If the given object implements this method
// the appropriate object method is returned, nil otherwise.
// The interface must contain a single method.
func GetInterfaceMethodFor[M any](o any) reflect.Value {
	if o == nil {
		return reflect.Value{}
	}

	if _, ok := generics.TryCast[M](o); !ok {
		return reflect.Value{}
	}
	return reflect.ValueOf(o).MethodByName(GetInterfaceMethod[M]().Name)
}

// CallOptionalInterfaceMethodOn calls an interface method on an object if
// it implements the interface.
func CallOptionalInterfaceMethodOn[M any](o any, args ...interface{}) []interface{} {
	m := GetInterfaceMethodFor[M](o)
	if !m.IsValid() {
		return nil
	}
	r := m.Call(sliceutils.Transform(args, reflect.ValueOf))
	return sliceutils.Transform(r, MapValueToInterface)
}

func CallOptionalInterfaceMethodOnV[M any](o any, args ...interface{}) bool {
	m := GetInterfaceMethodFor[M](o)
	if !m.IsValid() {
		return false
	}
	m.Call(sliceutils.Transform(args, reflect.ValueOf))
	return true
}

// CallOptionalInterfaceMethodOnE is like CallOptionalInterfaceMethodOn
// but for a method returning an error.
func CallOptionalInterfaceMethodOnE[M any](o any, args ...interface{}) error {
	m := GetInterfaceMethodFor[M](o)
	if !m.IsValid() {
		return nil
	}
	r := m.Call(sliceutils.Transform(args, reflect.ValueOf))
	return generics.Cast[error](r[0].Interface())
}
