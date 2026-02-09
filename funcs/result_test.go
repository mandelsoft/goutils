package funcs_test

import (
	. "github.com/mandelsoft/goutils/funcs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func multiresult() (int, string, int, bool) {
	return 1, "test", 2, true
}

var _ = Describe("result", func() {
	Context("result", func() {

		It("first", func() {
			Expect(First(multiresult())).To(Equal(1))
		})

		It("second", func() {
			Expect(Second(multiresult())).To(Equal("test"))
		})

		It("third", func() {
			Expect(Third(multiresult())).To(Equal(2))
		})

		It("fourth", func() {
			Expect(Fourth(multiresult())).To(Equal(true))
		})

		It("last", func() {
			Expect(Last[bool](multiresult())).To(Equal(true))
		})
	})
})
