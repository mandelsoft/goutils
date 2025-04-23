package generics_test

import (
	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/generics"
)

type S struct {
}

type W1 struct {
}

func (*W1) Unwrap(int) any {
	return nil
}

type W2 struct {
}

func (*W2) Unwrap() {
}

type T1 struct {
}

func (t T1) Unwrap() *T0 {
	return &T0{}
}

type T0 struct {
}

func (t *T0) Unwrap() any {
	return &S{}
}

type TA struct {
	ref *TA
}

func (t *TA) Unwrap() *TA {
	return t.ref
}

var _ = Describe("generic unwrap", func() {
	Context("unwrap any", func() {
		It("unwraps any type", func() {
			Expect(generics.UnwrapAny(&T1{})).To(DeepEqual(&T0{}))
		})
		It("fails on missing or wrong Unwrap method", func() {
			Expect(generics.UnwrapAny(&S{})).To(BeNil())
			Expect(generics.UnwrapAny(&W1{})).To(BeNil())
			Expect(generics.UnwrapAny(&W2{})).To(BeNil())
		})
	})

	Context("unwrap until", func() {
		It("unwraps any type", func() {
			r, ok := generics.UnwrapUntil[*S](&T0{})
			Expect(ok).To(BeTrue())
			Expect(r).To(DeepEqual(&S{}))
		})

		It("deep unwraps any type", func() {
			r, ok := generics.UnwrapUntil[*S](&T1{})
			Expect(ok).To(BeTrue())
			Expect(r).To(DeepEqual(&S{}))
		})

		It("fails on wrong unwrapped type", func() {
			_, ok := generics.UnwrapUntil[S](&T1{})
			Expect(ok).To(BeFalse())
		})
	})

	Context("typed", func() {
		It("unwraps nil", func() {
			Expect(generics.UnwrapWith[*T0](nil)).To(BeNil())
		})

		It("unwraps no match", func() {
			Expect(generics.UnwrapWith[*T0](&T0{})).To(BeNil())
		})

		It("unwraps any type", func() {
			t := &T1{}

			Expect(generics.UnwrapWith[*T0](t)).NotTo(BeNil())
		})

		It("unwraps nil ref", func() {
			t := &TA{}

			Expect(generics.UnwrapWith[*TA](t)).To(BeNil())
		})

	})

	Context("all", func() {
		It("unwraps nil", func() {
			Expect(generics.UnwrapAllWith[*TA](nil)).To(BeNil())
		})
		It("unwraps", func() {
			b := &TA{nil}
			r1 := &TA{b}
			r2 := &TA{r1}

			Expect(generics.UnwrapAllWith[*TA](r2)).To(BeIdenticalTo(b))
			Expect(generics.UnwrapWith[*TA](generics.UnwrapAllWith[*TA](r2))).To(BeNil())
		})
	})
})
