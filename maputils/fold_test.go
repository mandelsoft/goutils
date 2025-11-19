package maputils_test

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"strings"

	"github.com/mandelsoft/goutils/maputils"
)

func AddV(e int, sum float32) float32 {
	return float32(e) + sum
}

func MulIV(k string, e int, aggr float32) float32 {
	return float32(len(k)*e) + aggr
}

func AggrV(v int, aggr string) string {
	sub := fmt.Sprintf("%d", v)
	if len(aggr) == 0 {
		return sub
	}
	return aggr + ", " + sub
}

func AggrKV(k string, v int, aggr string) string {
	sub := fmt.Sprintf("%s:%d", k, v)
	if len(aggr) == 0 {
		return sub
	}
	return aggr + ", " + sub
}

var _ = Describe("Folding", func() {
	var m = map[string]int{"a": 1, "bb": 2, "c": 3, "dd": 4, "e": 5}

	Context("V", func() {
		It("folds empty map", func() {
			Expect(maputils.FoldV(map[string]int{}, 0, AddV)).To(Equal(float32(0)))
		})

		It("folds map", func() {
			Expect(maputils.FoldV(m, 0, AddV)).To(Equal(float32((5 + 1) * 5 / 2)))
		})
	})

	Context("KV", func() {
		It("folds empty map", func() {
			Expect(maputils.FoldKV(map[string]int{}, 0, MulIV)).To(Equal(float32(0)))
		})

		It("folds map", func() {
			Expect(maputils.FoldKV(m, 0, MulIV)).To(Equal(float32(1*1 + 2*2 + 1*3 + 2*4 + 1*5)))
		})
	})

	Context("ordered SV", func() {
		It("folds empty map", func() {
			Expect(maputils.FoldSV(map[string]int{}, "", AggrV)).To(Equal(""))
		})

		It("folds map", func() {
			Expect(maputils.FoldSV(m, "", AggrV)).To(Equal("1, 2, 3, 4, 5"))
		})
	})

	Context("ordered SKV", func() {
		It("folds empty map", func() {
			Expect(maputils.FoldSKV(map[string]int{}, "", AggrKV)).To(Equal(""))
		})

		It("folds map", func() {
			Expect(maputils.FoldSKV(m, "", AggrKV)).To(Equal("a:1, bb:2, c:3, dd:4, e:5"))
		})
	})

	Context("SVFunc", func() {
		cmp := func(k1, k2 string) int { return strings.Compare(k1, k2) }

		It("folds empty map", func() {
			Expect(maputils.FoldSVFunc(map[string]int{}, "", AggrV, cmp)).To(Equal(""))
		})

		It("folds map", func() {
			Expect(maputils.FoldSVFunc(m, "", AggrV, cmp)).To(Equal("1, 2, 3, 4, 5"))
		})
	})

	Context("SKVFunc", func() {
		cmp := func(k1, k2 string) int { return strings.Compare(k1, k2) }

		It("folds empty map", func() {
			Expect(maputils.FoldSKVFunc(map[string]int{}, "", AggrKV, cmp)).To(Equal(""))
		})

		It("folds ordered map", func() {
			Expect(maputils.FoldSKVFunc(m, "", AggrKV, cmp)).To(Equal("a:1, bb:2, c:3, dd:4, e:5"))
		})
	})
})
