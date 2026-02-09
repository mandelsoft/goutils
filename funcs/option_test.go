package funcs_test

import (
	"fmt"
	
	. "github.com/mandelsoft/goutils/funcs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func F(i int) Option[float64] {
	if i == 0 {
		return None[float64](fmt.Errorf("division by zero"))
	}
	return Some(1.0 / float64(i))
}

func G(d float64) Option[string] {
	return Some(fmt.Sprintf("%f", d))
}

var _ = Describe("option", func() {
	Context("option", func() {

		It("succeed", func() {
			Expect(AndThen(F(2), G).Value()).To(Equal("0.500000"))
		})

		It("fail early", func() {
			Expect(AndThen(F(0), G).IsNone()).To(BeTrue())
			Expect(AndThen(F(0), G).Error()).To(MatchError("division by zero"))
		})

	})
})
