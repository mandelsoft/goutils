package testutils

import (
	"fmt"
	"reflect"

	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/types"
)

// ContainInOrder checks a slice to contain some elements in order.
func ContainInOrder(elems ...any) types.GomegaMatcher {
	return &orderMatcher{Expected: elems}
}

type orderMatcher struct {
	Expected []any
	last     int
	elem     int
}

func (m *orderMatcher) Match(actual interface{}) (success bool, err error) {
	v := reflect.ValueOf(actual)
	if v.Kind() != reflect.Slice {
		return false, fmt.Errorf("After matcher expects a slice")
	}

	m.elem = 0
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		if reflect.DeepEqual(e.Interface(), m.Expected[m.elem]) {
			m.last = i
			m.elem++
			if m.elem == len(m.Expected) {
				return true, nil
			}
		}
	}
	return false, nil
}

func (m *orderMatcher) FailureMessage(actual interface{}) (message string) {
	if m.elem == 0 {
		return fmt.Sprintf("first element (%s) not found in %s", format.Object(m.Expected[m.elem], 0), format.Object(actual, 0))
	}
	return fmt.Sprintf("element %d (%s) not found after position %d in %s", m.elem, format.Object(m.Expected[m.elem], 0), m.last, format.Object(actual, 0))
}

func (m *orderMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("elements %s found in order in %s", format.Object(m.Expected, 0), format.Object(actual, 0))

}
