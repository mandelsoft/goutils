package stringutils_test

import (
	"github.com/mandelsoft/goutils/stringutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Padding Test Environment", func() {
	Context("single", func() {
		It("pad right", func() {
			Expect(stringutils.PadRight("test", 6, ' ')).To(Equal("test  "))
			Expect(stringutils.PadRight("test", 4, ' ')).To(Equal("test"))
			Expect(stringutils.PadRight("test", 3, ' ')).To(Equal("test"))
		})

		It("pad left", func() {
			Expect(stringutils.PadLeft("test", 6, ' ')).To(Equal("  test"))
			Expect(stringutils.PadLeft("test", 4, ' ')).To(Equal("test"))
			Expect(stringutils.PadLeft("test", 3, ' ')).To(Equal("test"))
		})
	})
	Context("align", func() {
		It("align right", func() {
			Expect(stringutils.AlignRight([]string{"alice", "bob", "charly"}, ' ')).To(Equal(
				[]string{" alice", "   bob", "charly"}))
		})

		It("align left", func() {
			Expect(stringutils.AlignLeft([]string{"alice", "bob", "charly"}, ' ')).To(Equal(
				[]string{"alice ", "bob   ", "charly"}))
		})
	})
})
