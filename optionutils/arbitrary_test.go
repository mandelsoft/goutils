package optionutils_test

import (
	"github.com/mandelsoft/goutils/optionutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type ArbitraryOptions struct {
	A string
	B string
}

type ArbitraryOption interface {
	ApplyToArbirary(opts *ArbitraryOptions)
}

type A string

func (o A) ApplyToArbirary(opts *ArbitraryOptions) {
	opts.A = string(o)
}

type B string

func (o B) ApplyToArbirary(opts *ArbitraryOptions) {
	opts.B = string(o)
}

var _ = Describe("Arbitrary Option interface Test Environment", func() {
	It("", func() {

		opts := optionutils.EvalArbitraryOptions[ArbitraryOption, ArbitraryOptions](A("a"), B("b"))
		Expect(opts.A).To(Equal("a"))
		Expect(opts.B).To(Equal("b"))
	})
})
