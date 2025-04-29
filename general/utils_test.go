package general_test

import (
	"github.com/mandelsoft/goutils/general"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Struct struct{}

var _ = Describe("Test Environment", func() {
	Context("Optional", func() {
		p := &Struct{}

		It("given", func() {
			Expect(general.Optional(0, 1)).To(Equal(0))
			Expect(general.Optional(nil, p)).To(BeIdenticalTo(p))
			Expect(general.Optional(1)).To(Equal(1))
			Expect(general.Optional(p)).To(BeIdenticalTo(p))
			Expect(general.Optional[int]()).To(Equal(0))
			Expect(general.Optional[*Struct]()).To(BeNil())
		})

		It("non-zero", func() {
			Expect(general.OptionalNonZero(0, 1)).To(Equal(1))
			Expect(general.OptionalNonZero(nil, p)).To(BeIdenticalTo(p))
			Expect(general.OptionalNonZero(1)).To(Equal(1))
			Expect(general.OptionalNonZero(p)).To(BeIdenticalTo(p))
			Expect(general.OptionalNonZero[int]()).To(Equal(0))
			Expect(general.OptionalNonZero[*Struct]()).To(BeNil())
		})
	})

	Context("OptionalDefaulted", func() {
		d := &Struct{}
		p := &Struct{}

		It("given", func() {
			Expect(general.OptionalDefaulted(5, 0, 1)).To(Equal(0))
			Expect(general.OptionalDefaulted(d, nil, p)).To(BeIdenticalTo(p))
			Expect(general.OptionalDefaulted(5, 1)).To(Equal(1))
			Expect(general.OptionalDefaulted(d, p)).To(BeIdenticalTo(p))
			Expect(general.OptionalDefaulted[int](5)).To(Equal(5))
			Expect(general.OptionalDefaulted[*Struct](d)).To(BeIdenticalTo(d))
		})

		It("non-zero", func() {
			Expect(general.OptionalNonZeroDefaulted(5, 0, 1)).To(Equal(1))
			Expect(general.OptionalNonZeroDefaulted(d, nil, p)).To(BeIdenticalTo(p))
			Expect(general.OptionalNonZeroDefaulted(5, 1)).To(Equal(1))
			Expect(general.OptionalNonZeroDefaulted(d, p)).To(BeIdenticalTo(p))
			Expect(general.OptionalNonZeroDefaulted[int](5)).To(Equal(5))
			Expect(general.OptionalNonZeroDefaulted[*Struct](d)).To(BeIdenticalTo(d))
		})
	})
})
