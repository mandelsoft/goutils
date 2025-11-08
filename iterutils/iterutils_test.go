package iterutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/iterutils"
)

var _ = Describe("Iterutils", func() {
	Context("Get", func() {
		It("should return element at index", func() {
			items := []string{"a", "b", "c"}
			element := iterutils.Get(iterutils.For(items...))
			Expect(element).To(Equal(items))
		})
	})

	Context("For/Get", func() {
		It("should iterate over slice elements", func() {
			items := []string{"a", "b", "c"}
			var result []string
			for e := range iterutils.For(items...) {
				result = append(result, e)
			}
			Expect(result).To(Equal(items))
		})

		It("should handle empty slice", func() {
			var empty []int
			count := 0
			for range iterutils.For(empty...) {
				count++
			}
			Expect(count).To(BeZero())
		})
	})

	Context("Reverse", func() {
		It("should reverse slice elements", func() {
			original := []string{"a", "b", "c"}
			reversed := iterutils.Reverse(iterutils.For(original...))
			Expect(iterutils.Get(reversed)).To(Equal([]string{"c", "b", "a"}))
		})
	})
})
