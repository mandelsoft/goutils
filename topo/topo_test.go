package topo_test

import (
	"iter"

	"github.com/mandelsoft/goutils/maputils"
	"github.com/mandelsoft/goutils/topo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func dep(graph map[string][]string) topo.Dependencies[string, string] {
	return func(e string) iter.Seq2[string, string] {
		return func(yield func(string, string) bool) {
			for _, d := range graph[e] {
				if !yield(d, d) {
					return
				}
			}
		}
	}
}

func iterate(graph map[string][]string) iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, k := range maputils.OrderedKeys(graph) {
			if !yield(k, k) {
				return
			}
		}
	}
}

var _ = Describe("Topo Sort Test", func() {
	It("order", func() {
		graph := map[string][]string{
			"A": {"B", "C"},
			"B": {"D", "E"},
			"C": nil,
			"D": {"C"},
			"E": nil,
		}

		o, c := topo.Sort[string, string](iterate(graph), dep(graph))
		Expect(c).To(BeNil())
		Expect(o).To(Equal([]string{"C", "D", "E", "B", "A"}))
	})

	It("cycle", func() {
		graph := map[string][]string{
			"A": {"B", "C"},
			"B": {"D", "E"},
			"C": nil,
			"D": {"C", "A"},
			"E": nil,
		}

		o, c := topo.Sort[string, string](iterate(graph), dep(graph))
		Expect(o).To(BeNil())
		Expect(c).To(Equal([]string{"A", "B", "D", "A"}))
	})
})
