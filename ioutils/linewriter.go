package ioutils

import (
	"bytes"
	"io"
	"sync"

	"github.com/mandelsoft/goutils/general"
)

// LineWriter only forwards a sequence of complete lines
// as single call to the underlying writer. The rest of a Write
// will be cached until a further Write with at least one
// line end will be called.
// Pending content will be written as line
// when Close is called.
// Optionally, an EOL scanner can be given. It detects
// the next line break in a buffer. This can be useful
// if escape sequences should
// be handled correctly. The default scanner just looks
// for the next newline rune.
type LineWriter struct {
	lock   sync.Mutex
	scan   func([]byte) int
	writer io.Writer
	buf    bytes.Buffer
}

var _ io.WriteCloser = (*LineWriter)(nil)

func defaultScanner(data []byte) int {
	return bytes.IndexByte(data, '\n')
}

func NewLineWriter(writer io.Writer, scan ...func([]byte) int) *LineWriter {
	return &LineWriter{
		writer: writer,
		scan:   general.OptionalDefaulted(defaultScanner, scan...),
	}
}

func (w *LineWriter) Write(p []byte) (n int, err error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	c := false
	s := w.buf.Len()
	for {
		nl := w.scan(p)
		if nl < 0 {
			break
		}
		w.buf.Write(p[:nl+1])
		c = true
		p = p[nl+1:]
	}

	if c {
		n, err = w.writer.Write(w.buf.Bytes())
		n -= s
		if err != nil {
			if n < 0 {
				return 0, err
			}
			return n, err
		}
		w.buf.Reset()
	}
	w.buf.Write(p)
	return len(p) + n, nil
}

func (w *LineWriter) Close() error {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.buf.Len() == 0 {
		return nil
	}
	w.buf.Write([]byte("\n"))
	_, err := w.writer.Write(w.buf.Bytes())
	w.buf.Reset()
	return err
}
