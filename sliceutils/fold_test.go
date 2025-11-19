package sliceutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/sliceutils"
)

func AddV(e int, sum float32) float32 {
	return float32(e) + sum
}

func MulIV(i, e int, aggr float32) float32 {
	return float32(i*e) + aggr
}

var _ = Describe("Folding", func() {
	var slice = []int{1, 2, 3, 4, 5}
	Context("V", func() {
		It("folds empty slice", func() {
			Expect(sliceutils.FoldV([]int{}, 0, AddV)).To(Equal(float32(0)))
		})

		It("folds slice", func() {
			Expect(sliceutils.FoldV(slice, 0, AddV)).To(Equal(float32((5 + 1) * 5 / 2)))
		})
	})

	Context("IV", func() {
		It("folds empty slice", func() {
			Expect(sliceutils.FoldIV([]int{}, 0, MulIV)).To(Equal(float32(0)))
		})

		It("folds slice", func() {
			Expect(sliceutils.FoldIV(slice, 0, MulIV)).To(Equal(float32(2 + 3*2 + 4*3 + 5*4)))
		})
	})
})
