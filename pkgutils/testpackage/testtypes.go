package testpackage

import (
	"github.com/mandelsoft/goutils/sliceutils"

	"github.com/mandelsoft/goutils/pkgutils"
)

type (
	MyStruct struct{}

	MyList     []int
	MyArray    [3]int
	MyMap      map[int]int
	MyChan     chan int
	MyFuncType func()
)

func MyFunc(i ...int) (string, error) {
	return pkgutils.GetPackageName(sliceutils.Convert[interface{}](i)...)
}
