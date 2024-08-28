package maputils_test

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/maputils"
)

func TransformValue(i int) string {
	return fmt.Sprintf("%d", i)
}

func TransformKey(n string) string {
	return strings.ToUpper(n)
}

var _ = Describe("MapUtils Test Environment", func() {
	Context("transform", func() {
		It("transforms values", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.TransformValues(in, TransformValue)).To(Equal(map[string]string{
				"alice": "24",
				"bob":   "25",
			}))
		})

		It("transforms values into slice", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.TransformedValues(in, TransformValue)).To(ConsistOf([]string{
				"24",
				"25",
			}))
		})

		It("transforms values into ordered slice", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.OrderedTransformedValues(in, TransformValue)).To(Equal([]string{
				"24",
				"25",
			}))
		})

		It("transforms keys", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.TransformKeys(in, TransformKey)).To(Equal(map[string]int{
				"ALICE": 24,
				"BOB":   25,
			}))
		})

		It("transforms keys into slice", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.TransformedKeys(in, TransformKey)).To(ConsistOf([]string{
				"ALICE",
				"BOB",
			}))
		})

		It("transforms keys into ordered slice", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.TransformedKeys(in, TransformKey)).To(Equal([]string{
				"ALICE",
				"BOB",
			}))
		})

		It("transforms maps", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.Transform(in, maputils.KeyValueTransformer(TransformKey, TransformValue))).To(Equal(map[string]string{
				"ALICE": "24",
				"BOB":   "25",
			}))
		})
	})
})
