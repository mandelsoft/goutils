package errors_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/errors"
)

type ERR struct {
}

func (e ERR) Error() string {
	return "err"
}

var _ = Describe("errors", func() {
	Context("ErrReadOnly", func() {
		It("identifies kind error", func() {
			uerr := errors.ErrReadOnly("KIND", "obj")

			Expect(errors.IsErrReadOnlyKind(uerr, "KIND")).To(BeTrue())
			Expect(errors.IsErrReadOnlyKind(uerr, "other")).To(BeFalse())

		})
		It("message with elem", func() {
			uerr := errors.ErrReadOnly("KIND", "obj")

			Expect(uerr.Error()).To(Equal("KIND \"obj\" is readonly"))
		})
		It("message without elem", func() {
			uerr := errors.ErrReadOnly()

			Expect(uerr.Error()).To(Equal("readonly"))
		})
	})
	Context("ErrUnkown", func() {
		It("identifies kind error", func() {
			uerr := errors.ErrUnknown("KIND", "obj")

			Expect(errors.IsErrUnknownKind(uerr, "KIND")).To(BeTrue())
			Expect(errors.IsErrUnknownKind(uerr, "other")).To(BeFalse())

		})
		It("finds error in history", func() {
			uerr := errors.ErrUnknown("KIND", "obj")
			werr := errors.Wrapf(uerr, "wrapped")

			Expect(errors.IsErrUnknownKind(werr, "KIND")).To(BeTrue())
			Expect(errors.IsErrUnknownKind(werr, "other")).To(BeFalse())
		})

		It("finds error in history with list", func() {
			uerr := errors.ErrUnknown("KIND", "obj")
			werr := errors.Wrapf(uerr, "wrapped")
			lerr := errors.ErrList().Add(fmt.Errorf("some error"), werr)
			Expect(errors.IsErrUnknownKind(lerr, "KIND")).To(BeTrue())
			Expect(errors.IsErrUnknownKind(lerr, "other")).To(BeFalse())

			Expect(errors.IsA(lerr, &errors.UnknownError{})).To(BeTrue())
			Expect(errors.IsOfType[*errors.UnknownError](lerr)).To(BeTrue())
		})
	})

	Context("ERR", func() {
		It("finds value error in history with list", func() {
			uerr := ERR{}
			werr := errors.Wrapf(uerr, "wrapped")
			lerr := errors.ErrList().Add(fmt.Errorf("some error"), werr)

			Expect(errors.IsA(lerr, ERR{})).To(BeTrue())
			Expect(errors.IsA(lerr, &ERR{})).To(BeTrue())
			Expect(errors.IsOfType[*ERR](lerr)).To(BeTrue())
			Expect(errors.IsOfType[ERR](lerr)).To(BeTrue())
		})

		It("finds pointer error in history with list", func() {
			uerr := &ERR{}
			werr := errors.Wrapf(uerr, "wrapped")
			lerr := errors.ErrList().Add(fmt.Errorf("some error"), werr)

			Expect(errors.IsA(lerr, ERR{})).To(BeTrue())
			Expect(errors.IsA(lerr, &ERR{})).To(BeTrue())
			Expect(errors.IsOfType[*ERR](lerr)).To(BeTrue())
			Expect(errors.IsOfType[ERR](lerr)).To(BeTrue())
		})
	})
})
