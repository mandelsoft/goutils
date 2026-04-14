package generics_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/generics"
)

type I interface {
	Func()
}

type II interface {
	Func()
	Other()
}

type NS struct {
}

type IS struct {
}

func (i *IS) Func() {
}

type MyInt int

type P struct {
	v string
}

func (o *P) New() *P {
	if o == nil {
		return &P{"initial"}
	}
	o.v = "set"
	return o
}

type V struct {
	v string
}

func (o V) New() V {
	o.v = "set"
	return o
}

var _ = Describe("types", func() {
	Context("implements", func() {
		It("implements", func() {
			Expect(generics.Implements[IS, I]()).To(BeFalse())
			Expect(generics.Implements[*IS, I]()).To(BeTrue())
			Expect(generics.Implements[II, I]()).To(BeTrue())
			Expect(generics.Implements[NS, I]()).To(BeFalse())
		})
	})

	Context("assign", func() {
		It("assign", func() {
			Expect(generics.CanAssign[IS, I]()).To(BeFalse())
			Expect(generics.CanAssign[*IS, I]()).To(BeTrue())
			Expect(generics.CanAssign[II, I]()).To(BeTrue())
			Expect(generics.CanAssign[NS, I]()).To(BeFalse())

			Expect(generics.CanAssign[int, int]()).To(BeTrue())
			Expect(generics.CanAssign[int8, int]()).To(BeFalse())
		})
	})

	Context("object", func() {
		It("pointer", func() {
			Expect(generics.ObjectFor[*NS]()).To(Equal(&NS{}))
		})
		It("non-pointer", func() {
			Expect(generics.ObjectFor[NS]()).To(Equal(NS{}))
		})

		It("pointer iniialized", func() {
			Expect(generics.ObjectFor[*P]()).To(Equal(&P{"set"}))
		})

		It("non-pointer iniialized", func() {
			Expect(generics.ObjectFor[V]()).To(Equal(V{"set"}))
		})
	})
})
