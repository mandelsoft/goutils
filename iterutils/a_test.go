package iterutils_test

import (
	"iter"

	"github.com/mandelsoft/goutils/iterutils"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type A struct {
	M string
}

func (a A) String() string {
	return a.M
}

type I interface {
	String() string
}

var _ I = (*A)(nil)

func iterSet[A comparable](set map[A]struct{}) iter.Seq[A] {
	return func(yield func(A) bool) {
		for e := range set {
			if !yield(e) {
				return
			}
		}
	}
}

var _ = Describe("Iterutils", func() {

	Context("when using iterators", func() {
		setA := map[A]struct{}{
			A{"a"}: struct{}{}, A{"b"}: struct{}{},
		}

		setI := map[I]struct{}{
			A{"a"}: struct{}{}, A{"b"}: struct{}{},
		}

		It("iter A", func() {
			list := []string{}
			for e := range iterSet(setA) {
				list = append(list, e.M)
			}
			Expect(list).To(ConsistOf("a", "b"))
		})

		It("iter I", func() {
			list := []string{}
			for e := range iterSet(setI) {
				list = append(list, e.String())
			}
			Expect(list).To(ConsistOf("a", "b"))
		})

		It("converted iter A", func() {
			list := []string{}
			for e := range iterutils.Convert[A](iterSet(setI)) {
				list = append(list, e.M)
			}
			Expect(list).To(ConsistOf("a", "b"))
		})

		It("converted iter I", func() {
			list := []string{}
			for e := range iterutils.Convert[I](iterSet(setA)) {
				list = append(list, e.String())
			}
			Expect(list).To(ConsistOf("a", "b"))
		})
	})
})
