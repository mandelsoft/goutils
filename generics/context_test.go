package generics_test

import (
	"context"

	"github.com/mandelsoft/goutils/generics"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Attribute struct {
	Value string
}

var _ = Describe("Context Test Environment", func() {
	attr := Attribute{"test"}
	var ctx context.Context

	BeforeEach(func() {
		ctx = context.TODO()
	})

	Context("value", func() {
		It("without value", func() {
			Expect(generics.FromContext[Attribute](ctx)).To(Equal(Attribute{}))
		})

		It("with value", func() {
			ctx = generics.WithValue(ctx, attr)
			Expect(generics.FromContext[Attribute](ctx)).To(Equal(attr))
		})
	})

	Context("pointer", func() {
		It("without value", func() {
			Expect(generics.FromContext[*Attribute](ctx) == nil).To(BeTrue())
		})

		It("with value", func() {
			ctx = generics.WithValue(ctx, &attr)
			Expect(generics.FromContext[*Attribute](ctx)).To(Equal(&Attribute{"test"}))
		})
	})
})
