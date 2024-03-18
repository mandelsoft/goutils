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

		It("transforms keys", func() {
			in := map[string]int{
				"alice": 24,
				"bob":   25,
			}

			Expect(maputils.MapKeys(in, TransformKey)).To(Equal(map[string]int{
				"ALICE": 24,
				"BOB":   25,
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
