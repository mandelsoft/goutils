package sliceutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/sliceutils"
)

func Compare(a, b int) int {
	return a - b
}

func Match(a int) bool {
	return a == 3
}

var _ = Describe("SliceUtils Test Environment", func() {
	Context("insert", func() {
		It("ordered", func() {
			Expect(sliceutils.InsertAscending([]int{1, 3, 5}, 0)).To(Equal([]int{0, 1, 3, 5}))
			Expect(sliceutils.InsertAscending([]int{1, 3, 5}, 2)).To(Equal([]int{1, 2, 3, 5}))
			Expect(sliceutils.InsertAscending([]int{1, 3, 5}, 6)).To(Equal([]int{1, 3, 5, 6}))
			Expect(sliceutils.InsertAscending([]int{}, 1)).To(Equal([]int{1}))
		})
		It("ordered reverse", func() {
			Expect(sliceutils.InsertDescending([]int{5, 3, 1}, 0)).To(Equal([]int{5, 3, 1, 0}))
			Expect(sliceutils.InsertDescending([]int{5, 3, 1}, 2)).To(Equal([]int{5, 3, 2, 1}))
			Expect(sliceutils.InsertDescending([]int{5, 3, 1}, 6)).To(Equal([]int{6, 5, 3, 1}))
			Expect(sliceutils.InsertDescending([]int{}, 1)).To(Equal([]int{1}))
		})

		It("ordered func", func() {
			Expect(sliceutils.InsertAscendingFunc([]int{1, 3, 5}, 0, Compare)).To(Equal([]int{0, 1, 3, 5}))
			Expect(sliceutils.InsertAscendingFunc([]int{1, 3, 5}, 2, Compare)).To(Equal([]int{1, 2, 3, 5}))
			Expect(sliceutils.InsertAscendingFunc([]int{1, 3, 5}, 6, Compare)).To(Equal([]int{1, 3, 5, 6}))
			Expect(sliceutils.InsertAscendingFunc([]int{}, 1, Compare)).To(Equal([]int{1}))
		})
		It("ordered reverse func", func() {
			Expect(sliceutils.InsertDescendingFunc([]int{5, 3, 1}, 0, Compare)).To(Equal([]int{5, 3, 1, 0}))
			Expect(sliceutils.InsertDescendingFunc([]int{5, 3, 1}, 2, Compare)).To(Equal([]int{5, 3, 2, 1}))
			Expect(sliceutils.InsertDescendingFunc([]int{5, 3, 1}, 6, Compare)).To(Equal([]int{6, 5, 3, 1}))
			Expect(sliceutils.InsertDescendingFunc([]int{}, 1, Compare)).To(Equal([]int{1}))
		})

		It("before first func", func() {
			Expect(sliceutils.InsertBeforeFirstFunc([]int{1, 3, 5}, 0, Match)).To(Equal([]int{1, 0, 3, 5}))
			Expect(sliceutils.InsertBeforeFirstFunc([]int{1, 3, 3, 5}, 0, Match)).To(Equal([]int{1, 0, 3, 3, 5}))
		})
		It("after first func", func() {
			Expect(sliceutils.InsertAfterFirstFunc([]int{1, 3, 5}, 0, Match)).To(Equal([]int{1, 3, 0, 5}))
			Expect(sliceutils.InsertAfterFirstFunc([]int{1, 3, 3, 3, 5}, 0, Match)).To(Equal([]int{1, 3, 0, 3, 3, 5}))
		})
		It("after last func", func() {
			Expect(sliceutils.InsertAfterLastFunc([]int{1, 3, 5}, 0, Match)).To(Equal([]int{1, 3, 0, 5}))
			Expect(sliceutils.InsertAfterLastFunc([]int{1, 3, 3, 5}, 0, Match)).To(Equal([]int{1, 3, 3, 0, 5}))
		})
		It("before last func", func() {
			Expect(sliceutils.InsertBeforeLastFunc([]int{1, 3, 5}, 0, Match)).To(Equal([]int{1, 0, 3, 5}))
			Expect(sliceutils.InsertBeforeLastFunc([]int{1, 3, 3, 3, 5}, 0, Match)).To(Equal([]int{1, 3, 3, 0, 3, 5}))
		})
	})
})
