package atomic_test

import (
	"github.com/mandelsoft/goutils/atomic"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type S struct {
	value string
}

var _ = Describe("Atomic Test Environment", func() {
	Context("nilable", func() {
		It("initial", func() {
			var a atomic.Value[*S]
			Expect(a.Load()).To(BeNil())
		})
		It("set/get", func() {
			var a atomic.Value[*S]
			v := &S{"alice"}
			a.Store(v)
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("swap", func() {
			var a atomic.Value[*S]

			v := &S{"alice"}
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Swap(nil)).To(BeIdenticalTo(v))
			Expect(a.Load()).To(BeNil())
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("compareandswap", func() {
			var a atomic.Value[*S]

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
		It("initial", func() {
			var a atomic.Value[int]
			Expect(a.Load()).To(Equal(0))
		})

		It("set/get", func() {
			var a atomic.Value[int]
			v := 5
			a.Store(v)
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("swap", func() {
			var a atomic.Value[*S]

			v := &S{"alice"}
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Swap(nil)).To(BeIdenticalTo(v))
			Expect(a.Load()).To(BeNil())
			Expect(a.Swap(v)).To(BeNil())
			Expect(a.Load()).To(BeIdenticalTo(v))
		})

		It("compareandswap", func() {
			var a atomic.Value[int]

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
