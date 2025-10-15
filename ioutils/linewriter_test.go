package ioutils_test

import (
	"bytes"
	"io"

	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/ioutils"
)

var _ = Describe("Line Writer Test Environment", func() {
	var buf bytes.Buffer
	var w io.WriteCloser

	BeforeEach(func() {
		buf.Reset()
		w = ioutils.NewLineWriter(&buf)
	})

	It("write line", func() {
		Expect(w.Write([]byte("line\n"))).To(Equal(5))
		Expect(buf.String()).To(Equal("line\n"))
	})

	It("write multiple lines", func() {
		Expect(w.Write([]byte("line1\nline2\n"))).To(Equal(12))
		Expect(buf.String()).To(Equal("line1\nline2\n"))
	})

	It("write partial line", func() {
		Expect(w.Write([]byte("line"))).To(Equal(4))
		Expect(buf.String()).To(Equal(""))

		Expect(w.Write([]byte("1\n"))).To(Equal(2))
		Expect(buf.String()).To(Equal("line1\n"))
	})

	It("write multiple lines plus partial", func() {
		Expect(w.Write([]byte("line1\nline2\nline"))).To(Equal(16))
		Expect(buf.String()).To(Equal("line1\nline2\n"))

		Expect(w.Write([]byte("3\n"))).To(Equal(2))
		Expect(buf.String()).To(Equal("line1\nline2\nline3\n"))
	})

	It("write mixed line", func() {
		Expect(w.Write([]byte("line1\nline"))).To(Equal(10))
		Expect(buf.String()).To(Equal("line1\n"))

		Expect(w.Write([]byte("2\n"))).To(Equal(2))
		Expect(buf.String()).To(Equal("line1\nline2\n"))
	})

	It("multiple partial lines", func() {
		Expect(w.Write([]byte("li"))).To(Equal(2))
		Expect(buf.String()).To(Equal(""))
		Expect(w.Write([]byte("ne1"))).To(Equal(3))
		Expect(buf.String()).To(Equal(""))
		Expect(w.Write([]byte("\nline"))).To(Equal(5))
		Expect(buf.String()).To(Equal("line1\n"))

		Expect(w.Write([]byte("2\n"))).To(Equal(2))
		Expect(buf.String()).To(Equal("line1\nline2\n"))
	})

	It("write repeated mixed line", func() {
		Expect(w.Write([]byte("line1\nline"))).To(Equal(10))
		Expect(buf.String()).To(Equal("line1\n"))

		Expect(w.Write([]byte("2\nline"))).To(Equal(6))
		Expect(buf.String()).To(Equal("line1\nline2\n"))

		Expect(w.Write([]byte("3\n"))).To(Equal(2))
		Expect(buf.String()).To(Equal("line1\nline2\nline3\n"))
	})

	It("close write pending", func() {
		Expect(w.Write([]byte("line1\nline"))).To(Equal(10))
		Expect(buf.String()).To(Equal("line1\n"))

		MustBeSuccessful(w.Close())

		Expect(buf.String()).To(Equal("line1\nline\n"))
	})

})
