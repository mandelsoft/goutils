package reflectutils_test

import (
	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/reflectutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type N struct {
	next *N
}

var _ reflectutils.Unwrappable[*N] = (*N)(nil)

func (n *N) Unwrap() *N {
	return n.next
}

type M struct {
	next *M
}

func (n *M) Unwrap(in *N) *M {
	return n.next
}

type O struct {
	next *O
}

func (n *O) Unwrap() (*O, *N) {
	return n.next, nil
}

var _ = Describe("Unwrap Test Environment", func() {
	It("unwrap any", func() {
		n := &N{&N{&N{}}}
		i := 0
		for n != nil {
			n = generics.Cast[*N](reflectutils.UnwrapAny(n))
			i++
		}
		Expect(i).To(Equal(3))
	})

	It("unwrap typed", func() {
		n := &N{&N{&N{}}}
		i := 0
		for n != nil {
			n = reflectutils.Unwrap[*N](n)
			i++
		}
		Expect(i).To(Equal(3))
	})

	It("no unwrap for arg count mismatch", func() {
		n := &M{&M{&M{}}}
		i := 0
		for n != nil {
			n = generics.Cast[*M](reflectutils.UnwrapAny(n))
			i++
		}
		Expect(i).To(Equal(1))
	})

	It("no unwrap for result count mismatch", func() {
		n := &O{&O{&O{}}}
		i := 0
		for n != nil {
			n = generics.Cast[*O](reflectutils.UnwrapAny(n))
			i++
		}
		Expect(i).To(Equal(1))
	})
})
