package paasio

import (
	"io"
)

const testVersion = 3

type counter struct {
	n    int64
	nops int
	ch   chan int64
}

type readCounter struct {
	r io.Reader
	c *counter
}

type writeCounter struct {
	w io.Writer
	c *counter
}

type readWriteCounter struct {
	r *readCounter
	w *writeCounter
}

// **************** read ****************

func NewReadCounter(reader io.Reader) ReadCounter {
	rc := readCounter{r: reader, c: newCounter()}
	return &rc
}

func (rc *readCounter) Read(buf []byte) (int, error) {
	var n int
	var err error
	if n, err = rc.r.Read(buf); err == nil {
		rc.c.add(int64(n))
	}
	return n, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	return rc.c.n, rc.c.nops
}

// **************** write ****************

func NewWriteCounter(writer io.Writer) WriteCounter {
	wc := writeCounter{w: writer, c: newCounter()}
	return &wc
}

func (wc *writeCounter) Write(buf []byte) (int, error) {
	var n int
	var err error
	if n, err = wc.w.Write(buf); err == nil {
		wc.c.add(int64(n))
	}
	return n, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	return wc.c.n, wc.c.nops
}

// **************** read/write ****************

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	rc := NewReadCounter(rw).(*readCounter)
	wc := NewWriteCounter(rw).(*writeCounter)
	return &readWriteCounter{r: rc, w: wc}
}

func (rwc *readWriteCounter) Read(buf []byte) (int, error) {
	return rwc.r.Read(buf)
}

func (rwc *readWriteCounter) Write(buf []byte) (int, error) {
	return rwc.w.Write(buf)
}

func (rwc *readWriteCounter) ReadCount() (int64, int) {
	return rwc.r.ReadCount()
}

func (rwc *readWriteCounter) WriteCount() (int64, int) {
	return rwc.w.WriteCount()
}

// **************** counter ****************

func newCounter() *counter {
	c := counter{ch: make(chan int64)}
	go func() {
		for {
			c.n += <-c.ch
			c.nops++
		}
	}()
	return &c
}

func (c *counter) add(n int64) {
	c.ch <- n
}
