package ioutils

import (
	"github.com/mandelsoft/goutils/general"
	"io"
	"os"
	"sync"

	"github.com/mandelsoft/goutils/errors"
	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/sliceutils"
)

////////////////////////////////////////////////////////////////////////////////

type additionalCloser[T any] struct {
	msg              []interface{}
	wrapped          T
	additionalCloser io.Closer
}

func (c *additionalCloser[T]) Close() error {
	var list *errors.ErrorList
	if len(c.msg) == 0 {
		list = errors.ErrListf("close")
	} else {
		if s, ok := c.msg[0].(string); ok && len(c.msg) > 1 {
			list = errors.ErrListf(s, c.msg[1:]...)
		} else {
			list = errors.ErrList(c.msg...)
		}
	}
	if cl, ok := generics.TryCast[io.Closer](c.wrapped); ok {
		list.Add(cl.Close())
	}
	if c.additionalCloser != nil {
		list.Add(c.additionalCloser.Close())
	}
	return list.Result()
}

func newAdditionalCloser[T any](w T, closer io.Closer, msg ...interface{}) additionalCloser[T] {
	return additionalCloser[T]{
		wrapped:          w,
		msg:              msg,
		additionalCloser: closer,
	}
}

////////////////////////////////////////////////////////////////////////////////

type readCloser struct {
	additionalCloser[io.Reader]
}

var _ io.ReadCloser = (*readCloser)(nil)

// Deprecated: use AddReaderCloser .
func AddCloser(reader io.ReadCloser, closer io.Closer, msg ...string) io.ReadCloser {
	return AddReaderCloser(reader, closer, sliceutils.AsAny(msg)...)
}

func ReadCloser(r io.Reader) io.ReadCloser {
	return AddReaderCloser(r, nil)
}

func AddReaderCloser(reader io.Reader, closer io.Closer, msg ...interface{}) io.ReadCloser {
	return &readCloser{
		additionalCloser: newAdditionalCloser[io.Reader](reader, closer, msg...),
	}
}

func (c *readCloser) Read(p []byte) (n int, err error) {
	return c.wrapped.Read(p)
}

type writeCloser struct {
	additionalCloser[io.Writer]
}

var _ io.WriteCloser = (*writeCloser)(nil)

func WriteCloser(w io.Writer) io.WriteCloser {
	return AddWriterCloser(w, nil)
}

func AddWriterCloser(writer io.Writer, closer io.Closer, msg ...interface{}) io.WriteCloser {
	return &writeCloser{
		additionalCloser: newAdditionalCloser[io.Writer](writer, closer, msg...),
	}
}

func (c *writeCloser) Write(p []byte) (n int, err error) {
	return c.wrapped.Write(p)
}

////////////////////////////////////////////////////////////////////////////////

type DupReadCloser interface {
	io.ReadCloser
	Dup() (DupReadCloser, error)
}

type dupReadCloser struct {
	lock  sync.Mutex
	rc    io.ReadCloser
	count int
}

func (d *dupReadCloser) Read(p []byte) (n int, err error) {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.rc.Read(p)
}

func (d *dupReadCloser) Close() error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.count == 0 {
		return os.ErrClosed
	}
	d.count--
	if d.count == 0 {
		return d.rc.Close()
	}
	return nil
}

func (d *dupReadCloser) Dup() (DupReadCloser, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if d.count == 0 {
		return nil, os.ErrClosed
	}
	d.count++
	return d, nil
}

func NewDupReadCloser(rc io.ReadCloser, errs ...error) (DupReadCloser, error) {
	if err := general.Optional(errs...); err != nil {
		return nil, err
	}
	if d, ok := rc.(*dupReadCloser); ok {
		return d.Dup()
	}
	return &dupReadCloser{
		rc:    rc,
		count: 1,
	}, nil
}
