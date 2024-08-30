package reflectutils_test

import (
	"github.com/mandelsoft/goutils/reflectutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Method interface {
	Method(a, b string) string
}

type MethodVA interface {
	MethodVA(a string)
}

type Target struct {
	Result string
}

func (t *Target) Method(a, b string) string {
	t.Result = a + b
	return t.Result
}

func (t *Target) MethodVA(a string) {
	t.Result = a
}

var _ = Describe("Calls Test Environment", func() {
	It("calls method based on interface", func() {
		t := &Target{}
		r := reflectutils.CallMethodByInterface[Method](t, "alice", "bob")
		Expect(r).To(Equal([]interface{}{"alicebob"}))
	})

	It("calls method VA based on interface", func() {
		t := &Target{}
		reflectutils.CallMethodByInterfaceVA[MethodVA](t, "alice")
		Expect(t.Result).To(Equal("alice"))
	})
})
