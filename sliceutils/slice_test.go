package sliceutils_test

import (
	"strconv"

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

	Context("aggregate", func() {
		It("aggregates empty", func() {
			data := []int{}

			r := sliceutils.Aggregate(data, "x", func(s string, i int) string { return s + strconv.Itoa(i) })
			Expect(r).To(Equal("x"))
		})

		It("aggregates", func() {
			data := []int{1, 2, 3, 4, 5}

			r := sliceutils.Aggregate(data, "x", func(s string, i int) string { return s + strconv.Itoa(i) })
			Expect(r).To(Equal("x12345"))
		})
	})

	Context("prefix", func() {
		It("accepts prefix", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasPrefix(data, 1)).To(BeTrue())
			Expect(sliceutils.HasPrefix(data, 1, 2)).To(BeTrue())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3)).To(BeTrue())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3, 4)).To(BeTrue())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3, 4, 5)).To(BeTrue())
		})

		It("accepts prefix by func", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasPrefixFunc(data, equals, 1)).To(BeTrue())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2)).To(BeTrue())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3)).To(BeTrue())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3, 4)).To(BeTrue())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3, 4, 5)).To(BeTrue())
		})

		It("rejects prefix", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasPrefix(data, 2)).To(BeFalse())
			Expect(sliceutils.HasPrefix(data, 1, 3)).To(BeFalse())
			Expect(sliceutils.HasPrefix(data, 1, 2, 4)).To(BeFalse())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3, 5)).To(BeFalse())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasPrefix(data, 1, 2, 3, 4, 5, 6)).To(BeFalse())
		})

		It("rejects prefix by func", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasPrefixFunc(data, equals, 2)).To(BeFalse())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 3)).To(BeFalse())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 4)).To(BeFalse())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3, 5)).To(BeFalse())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasPrefixFunc(data, equals, 1, 2, 3, 4, 5, 6)).To(BeFalse())
		})
	})

	Context("suffix", func() {

		It("accepts suffix", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasSuffix(data, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffix(data, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffix(data, 3, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffix(data, 2, 3, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffix(data, 1, 2, 3, 4, 5)).To(BeTrue())
		})

		It("rejects suffix", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasSuffix(data, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 5, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 4, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 3, 3, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 2, 2, 3, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 1, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffix(data, 1, 2, 3, 4, 5, 6)).To(BeFalse())
		})

		It("accepts suffix by func", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasSuffixFunc(data, equals, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffixFunc(data, equals, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffixFunc(data, equals, 3, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffixFunc(data, equals, 2, 3, 4, 5)).To(BeTrue())
			Expect(sliceutils.HasSuffixFunc(data, equals, 1, 2, 3, 4, 5)).To(BeTrue())
		})

		It("rejects suffix", func() {
			data := []int{1, 2, 3, 4, 5}

			Expect(sliceutils.HasSuffixFunc(data, equals, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 5, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 4, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 3, 3, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 2, 2, 3, 4, 5)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 1, 2, 3, 4, 6)).To(BeFalse())
			Expect(sliceutils.HasSuffixFunc(data, equals, 1, 2, 3, 4, 5, 6)).To(BeFalse())
		})
	})
})

func equals(a, b int) bool {
	return a == b
}
