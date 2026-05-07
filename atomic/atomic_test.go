package atomic_test

import (
	"github.com/mandelsoft/goutils/atomic"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type S struct {
	value string
}

func (s *S) String() string {
	return s.value
}

var _ I = (*S)(nil)

type O struct {
	S
}

type I interface {
	String() string
}

var _ = Describe("Atomic Test Environment", func() {
	Context("interface", func() {
		var a atomic.InterfaceValue[I]

		BeforeEach(func() {
			a = atomic.InterfaceValue[I]{}
		})

		It("initial", func() {
			Expect(a.Load()).To(BeNil())
		})

		It("set nil", func() {
			a.Store(&S{"alice"})
			Expect(a.Load().String()).To(Equal("alice"))
			a.Store(nil)
			Expect(a.Load()).To(BeNil())
			Expect(a.Load() == nil).To(BeTrue())
			Expect(a.Load() == nil).To(BeTrue())
		})
		It("set/get", func() {
			v := &S{"alice"}
			a.Store(v)
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("swap", func() {
			v := &S{"alice"}
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Swap(nil)).To(BeIdenticalTo(v))
			Expect(a.Load()).To(BeNil())
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("compareandswap", func() {
			v := &S{"alice"}
			Expect(a.CompareAndSwap(v, nil)).To(BeFalse())
			Expect(a.Load()).To(BeNil())
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
			Expect(a.CompareAndSwap(nil, nil)).To(BeFalse())
			Expect(a.CompareAndSwap(v, nil)).To(BeTrue())
			Expect(a.Load()).To(BeNil())
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

	})

	Context("pointer", func() {
		var a atomic.Value[*S]

		BeforeEach(func() {
			a = atomic.Value[*S]{}
		})

		It("initial", func() {
			Expect(a.Load()).To(BeNil())
		})
		It("set initial nil", func() {
			a.Store(nil)
			Expect(a.Load()).To(BeNil())
			Expect(a.Load() == nil).To(BeTrue())
		})
		It("set nil", func() {
			a.Store(&S{})
			a.Store(nil)
			Expect(a.Load()).To(BeNil())
			Expect(a.Load() == nil).To(BeTrue())
		})
		It("set/get", func() {
			v := &S{"alice"}
			a.Store(v)
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("swap", func() {
			v := &S{"alice"}
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Swap(nil)).To(BeIdenticalTo(v))
			Expect(a.Load()).To(BeNil())
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("compareandswap", func() {
			v := &S{"alice"}
			Expect(a.CompareAndSwap(v, nil)).To(BeFalse())
			Expect(a.Load()).To(BeNil())
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
			Expect(a.CompareAndSwap(nil, nil)).To(BeFalse())
			Expect(a.CompareAndSwap(v, nil)).To(BeTrue())
			Expect(a.Load()).To(BeNil())
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

	})

	Context("not nilable", func() {
		var a atomic.Value[int]

		BeforeEach(func() {
			a = atomic.Value[int]{}
		})

		It("initial", func() {
			Expect(a.Load()).To(Equal(0))
		})

		It("set/get", func() {
			v := 5
			a.Store(v)
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("swap", func() {
			Expect(a.Swap(1)).To(Equal(0))
			Expect(a.Swap(2)).To(BeIdenticalTo(1))
			Expect(a.Load()).To(BeIdenticalTo(2))
		})

		It("compareandswap", func() {
			v := 5
			Expect(a.CompareAndSwap(v, 0)).To(BeFalse())
			Expect(a.Load()).To(Equal(0))
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
			Expect(a.CompareAndSwap(nil, 0)).To(BeFalse())
			Expect(a.CompareAndSwap(v, 0)).To(BeTrue())
			Expect(a.Load()).To(Equal(0))
			Expect(a.CompareAndSwap(nil, v)).To(BeTrue())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})
	})
})
