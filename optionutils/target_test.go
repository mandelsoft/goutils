package optionutils_test

import (
	"github.com/mandelsoft/goutils/optionutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type OrigOptions struct {
	Orig string
}

type DerivedOptions struct {
	OrigOptions
	Derived string
}

type orig string

func (o orig) ApplyTo(opts *OrigOptions) {
	opts.Orig = string(o)
}

func WithOrig(s string) optionutils.Option[*OrigOptions] {
	return orig(s)
}

type derived string

func (o derived) ApplyTo(opts *DerivedOptions) {
	opts.Derived = string(o)
}

func WithDerived(s string) optionutils.Option[*DerivedOptions] {
	return derived(s)
}

func WithOrigDerived(s string) optionutils.Option[*DerivedOptions] {
	return optionutils.MapOptionTarget[*DerivedOptions, *OrigOptions](WithOrig(s))
}

var _ = Describe("Mapped Options Test Environment", func() {
	It("", func() {
		var opts DerivedOptions

		list := []optionutils.Option[*DerivedOptions]{
			WithDerived("d"),
			WithOrigDerived("o"),
		}

		optionutils.ApplyOptions(&opts, list...)

		Expect(opts.Derived).To(Equal("d"))
		Expect(opts.Orig).To(Equal(""))

		optionutils.ApplyOptions(&opts.OrigOptions, optionutils.FilterMappedOptions[*OrigOptions](list...)...)
		Expect(opts.Orig).To(Equal("o"))
	})
})
