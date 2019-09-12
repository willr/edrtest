package action

import (
	"io"
	"sync/atomic"
)

type CountWriter struct {
	io.Writer
	count uint64
}

func NewCountWriter(w io.Writer) *CountWriter {
	return &CountWriter{
		Writer: w,
	}
}

func (counter *CountWriter) Write(buf []byte) (int, error) {
	n, err := counter.Writer.Write(buf)
	atomic.AddUint64(&counter.count, uint64(n))
	return n, err
}

func (counter *CountWriter) Count() uint64 {
	return atomic.LoadUint64(&counter.count)
}
