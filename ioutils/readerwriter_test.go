package ioutils_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"

	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/goutils/ioutils"
)

var data = "this is some test data"

var Err = fmt.Errorf("closed")

type TestReader struct {
	lock sync.Mutex
	r    io.Reader
}

var _ io.ReadCloser = (*TestReader)(nil)

func (r *TestReader) Read(buf []byte) (int, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.r == nil {
		return 0, Err
	}
	return r.r.Read(buf)
}

func (r *TestReader) Close() error {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.r == nil {
		return Err
	}
	r.r = nil
	return nil
}

func (r *TestReader) IsClosed() bool {
	r.lock.Lock()
	defer r.lock.Unlock()
	return r.r == nil
}

var _ = Describe("Test Environment", func() {
	var reader *TestReader

	BeforeEach(func() {
		reader = &TestReader{r: bytes.NewBufferString(data)}
	})

	Context("dup reader", func() {
		It("closed single view", func() {
			r := Must(ioutils.NewDupReadCloser(reader))
			MustBeSuccessful(r.Close())
			Expect(reader.IsClosed()).To(BeTrue())

			Expect(r.Close()).To(BeIdenticalTo(os.ErrClosed))
		})

		It("delayed close for multiple views", func() {
			r := Must(ioutils.NewDupReadCloser(reader))
			r2 := Must(r.Dup())
			MustBeSuccessful(r.Close())
			Expect(reader.IsClosed()).To(BeFalse())
			Expect(r.Close()).To(BeIdenticalTo(os.ErrClosed))
			MustBeSuccessful(r2.Close())
			Expect(reader.IsClosed()).To(BeTrue())
			Expect(r2.Close()).To(BeIdenticalTo(os.ErrClosed))
		})

		It("reads using multiple views", func() {
			r := Must(ioutils.NewDupReadCloser(reader))
			r2 := Must(r.Dup())

			var buf [5]byte
			Expect(r.Read(buf[:])).To(Equal(5))
			Expect(string(buf[:])).To(Equal("this "))
			MustBeSuccessful(r.Close())
			Expect(reader.IsClosed()).To(BeFalse())

			Expect(r2.Read(buf[:])).To(Equal(5))
			Expect(string(buf[:])).To(Equal("is so"))
			MustBeSuccessful(r2.Close())
			Expect(reader.IsClosed()).To(BeTrue())
		})

		It("handles closed readers", func() {
			MustBeSuccessful(reader.Close())

			r := Must(ioutils.NewDupReadCloser(reader))
			r2 := Must(r.Dup())

			MustBeSuccessful(r.Close())

			var buf [5]byte
			ExpectError(r2.Read(buf[:])).To(BeIdenticalTo(Err))
			ExpectError(r2.Close()).To(BeIdenticalTo(Err))
		})
	})
})
