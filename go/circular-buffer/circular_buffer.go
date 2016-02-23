package circular

import "errors"

const testVersion = 3

type Buffer struct {
	size       int
	numEntries int
	readIndex  int
	writeIndex int
	data       []byte
}

func NewBuffer(size int) *Buffer {
	return &Buffer{size, 0, 0, 0, make([]byte, size)}
}

func (buf *Buffer) ReadByte() (byte, error) {
	if buf.isEmpty() {
		return 0, errors.New("buffer is empty")
	}
	c := buf.data[buf.readIndex]
	buf.moveReadHead()
	buf.numEntries--
	return c, nil
}

func (buf *Buffer) WriteByte(c byte) error {
	if buf.isFull() {
		return errors.New("buffer is full")
	}
	buf.data[buf.writeIndex] = c
	buf.moveWriteHead()
	buf.numEntries++
	return nil
}

func (buf *Buffer) Overwrite(c byte) {
	if buf.isFull() {
		index := (buf.writeIndex - buf.numEntries + buf.size) % buf.size
		buf.data[index] = c
		buf.moveReadHead()
		buf.moveWriteHead()
	} else {
		buf.WriteByte(c)
	}
}

func (buf *Buffer) Reset() {
	buf.numEntries = 0
	buf.readIndex = 0
	buf.writeIndex = 0
	buf.numEntries = 0
}

func isEmpty() bool {
	return buf.numEntries == 0
}

func isFull() bool {
	return buf.numEntries == buf.size
}

func (buf *Buffer) moveReadHead() {
	buf.readIndex = (buf.readIndex + 1) % buf.size
}

func (buf *Buffer) moveWriteHead() {
	buf.writeIndex = (buf.writeIndex + 1) % buf.size
}
