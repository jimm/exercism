package paasio

import (
	"io"
	"sync"
)

const testVersion = 3

type readCounter struct {
	sync.Mutex
	r    io.Reader
	n    int64
	nops int
}

type writeCounter struct {
	sync.Mutex
	w    io.Writer
	n    int64
	nops int
}

type readWriteCounter struct {
	r *readCounter
	w *writeCounter
}

// **************** read ****************

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{r: reader}
}

func (c *readCounter) Read(buf []byte) (int, error) {
	c.Lock()
	defer c.Unlock()

	var n int
	var err error
	if n, err = c.r.Read(buf); err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

func (c *readCounter) ReadCount() (int64, int) {
	return c.n, c.nops
}

// **************** write ****************

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{w: writer}
}

func (c *writeCounter) Write(buf []byte) (int, error) {
	c.Lock()
	defer c.Unlock()

	var n int
	var err error
	if n, err = c.w.Write(buf); err == nil {
		c.n += int64(n)
		c.nops++
	}
	return n, err
}

func (c *writeCounter) WriteCount() (int64, int) {
	return c.n, c.nops
}

// **************** read/write ****************

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	rc := NewReadCounter(rw).(*readCounter)
	wc := NewWriteCounter(rw).(*writeCounter)
	return &readWriteCounter{r: rc, w: wc}
}

func (c *readWriteCounter) Read(buf []byte) (int, error) {
	return c.r.Read(buf)
}

func (c *readWriteCounter) Write(buf []byte) (int, error) {
	return c.w.Write(buf)
}

func (c *readWriteCounter) ReadCount() (int64, int) {
	return c.r.n, c.r.nops
}

func (c *readWriteCounter) WriteCount() (int64, int) {
	return c.w.n, c.w.nops
}
