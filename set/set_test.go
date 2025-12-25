package set_test

import (
	"github.com/mandelsoft/goutils/set"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set Test Environment", func() {
	var s1 set.Set[string]

	BeforeEach(func() {
		s1 = set.New[string]()
	})

	Context("one set", func() {
		It("should be empty", func() {
			Expect(s1.Len()).To(Equal(0))
			Expect(s1.IsEmpty()).To(BeTrue())

			s1.Add("a")

			Expect(s1.Len()).To(Equal(1))
			Expect(s1.IsEmpty()).To(BeFalse())
		})

		It("adds", func() {
			Expect(s1.Add("a")).To(Equal(set.New("a")))
		})

		It("equals", func() {
			Expect(s1.Add("a").Equal(set.New("a"))).To(BeTrue())
			Expect(s1.Add("a").Equal(set.New[string]())).To(BeFalse())
			Expect(s1.Add("a", "b").Equal(set.New("a"))).To(BeFalse())
			Expect(s1.Add("a", "b").Equal(set.New("a", "b"))).To(BeTrue())
		})

		It("with added", func() {
			n := s1.Add("a").WithAdded("c", "b")
			Expect(n.Equal(s1)).To(BeFalse())
			Expect(n).To(Equal(set.New("a", "b", "c")))
			Expect(s1).To(Equal(set.New("a")))
		})

		It("ordered", func() {
			s1.Add("b", "c", "a")
			Expect(s1.AsArray()).To(ConsistOf("a", "b", "c"))
			Expect(set.AsSortedArray(s1)).To(Equal([]string{"a", "b", "c"}))
		})

		It("delete", func() {
			s1.Add("a", "b", "c")
			Expect(s1.Delete("a", "b")).To(Equal(set.New("c")))
		})

		It("get", func() {
			s1.Add("a", "b", "c")
			cmp := s1.Clone()

			e, ok := s1.GetAny()
			Expect(ok).To(BeTrue())
			Expect(e).To(BeElementOf(cmp.AsArray()))
			Expect(s1).To(Equal(cmp.Delete(e)))

			e, ok = s1.GetAny()
			Expect(ok).To(BeTrue())
			Expect(e).To(BeElementOf(cmp.AsArray()))
			Expect(s1).To(Equal(cmp.Delete(e)))

			e, ok = s1.GetAny()
			Expect(ok).To(BeTrue())
			Expect(e).To(BeElementOf(cmp.AsArray()))
			Expect(s1).To(Equal(cmp.Delete(e)))

			e, ok = s1.GetAny()
			Expect(ok).To(BeFalse())
		})

		It("iterator", func() {
			s1.Add("a", "b", "c")

			var r []string
			for e := range s1.Elements {
				r = append(r, e)
			}
			Expect(r).To(ConsistOf("a", "b", "c"))
		})
	})

	Context("two sets", func() {
		var s2 set.Set[string]

		BeforeEach(func() {
			s2 = set.New[string]("x", "y", "z")
			s1.Add("x", "y", "a", "b")
		})

		It("intersection", func() {
			Expect(s1.Intersection(s2)).To(Equal(set.New[string]("x", "y")))
		})

		It("union", func() {
			Expect(s1.Union(s2)).To(Equal(set.New[string]("a", "b", "x", "y", "z")))
		})

		It("difference", func() {
			Expect(s1.Difference(s2)).To(Equal(set.New[string]("a", "b")))
			Expect(s2.Difference(s1)).To(Equal(set.New[string]("z")))
			Expect(s1.Difference(s1)).To(Equal(set.New[string]()))
		})

		It("symmetric diff", func() {
			Expect(s1.SymmetricDifference(s2)).To(Equal(set.New[string]("a", "b", "z")))
			Expect(s2.SymmetricDifference(s1)).To(Equal(set.New[string]("a", "b", "z")))
		})

		It("superset", func() {
			Expect(s1.IsSuperset(s2)).To(BeFalse())
			Expect(s2.IsSuperset(s1)).To(BeFalse())

			Expect(s1.IsSubset(s2)).To(BeFalse())
			Expect(s2.IsSubset(s1)).To(BeFalse())

			s1.Add("z")

			Expect(s2.IsSuperset(s1)).To(BeFalse())
			Expect(s1.IsSuperset(s2)).To(BeTrue())

			Expect(s2.IsSubset(s1)).To(BeTrue())
			Expect(s1.IsSubset(s2)).To(BeFalse())
		})
	})
})
