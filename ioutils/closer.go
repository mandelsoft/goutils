package ioutils

import (
	"fmt"
	"io"

	"github.com/mandelsoft/goutils/errors"
)

////////////////////////////////////////////////////////////////////////////////

// NopWriteCloser returns a WriteCloser with a no-op Close method wrapping
// the provided Writer w.
func NopWriteCloser(w io.Writer) io.WriteCloser {
	return nopCloser{w}
}

type nopCloser struct {
	io.Writer
}

func (nopCloser) Close() error { return nil }

// NopReadCloser returns a ReadCloser with a no-op Close method wrapping
// the provided Reader r.
func NopReadCloser(r io.Reader) io.ReadCloser {
	return io.NopCloser(r)

}

////////////////////////////////////////////////////////////////////////////////

type once struct {
	callbacks []CloserCallback
	closer    io.Closer
}

// CloserCallback is a function called when Close is called.
type CloserCallback func()

// OnceCloser assures that an underlying io.Closer
// is called only once, even if called multiple times.
// Additionally, callback functions can be added.
func OnceCloser(c io.Closer, callbacks ...CloserCallback) io.Closer {
	return &once{callbacks, c}
}

func (c *once) Close() error {
	if c.closer == nil {
		return nil
	}

	t := c.closer
	c.closer = nil
	err := t.Close()

	for _, cb := range c.callbacks {
		cb()
	}

	if err != nil {
		return fmt.Errorf("unable to close: %w", err)
	}

	return nil
}

// Close calls Close on a sequence of io.Closer and aggregates
// potential error.
func Close(closer ...io.Closer) error {
	if len(closer) == 0 {
		return nil
	}
	list := errors.ErrList()
	for _, c := range closer {
		if c != nil {
			list.Add(c.Close())
		}
	}
	return list.Result()
}

// CloseFunc is function usable as io.Closer.
type CloseFunc func() error

func (c CloseFunc) Close() error {
	return c()
}
