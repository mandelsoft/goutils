package reflectutils

import (
	"fmt"
	"reflect"

	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/sliceutils"
)

// CallMethodByInterfaceVA calls a void method on object o with
// one argument a. The methid is specified by the interface
// M, which should implement exactly one appropriate method.
func CallMethodByInterfaceVA[M, B any](o B, a interface{}) {
	t := generics.TypeOf[M]()
	if t.NumMethod() != 1 {
		panic(fmt.Sprintf("invalid setter type %s", t))
	}
	m := t.Method(0)
	reflect.ValueOf(o).MethodByName(m.Name).Call([]reflect.Value{reflect.ValueOf(a)})
}

func CallMethodByInterface[M, B any](o B, args ...interface{}) []interface{} {
	t := generics.TypeOf[M]()
	if t.NumMethod() != 1 {
		panic(fmt.Sprintf("invalid setter type %s", t))
	}
	m := t.Method(0)
	v := sliceutils.Transform(args, reflect.ValueOf)
	r := reflect.ValueOf(o).MethodByName(m.Name).Call(v)
	return sliceutils.Transform(r, MapValueToInterface)
}
