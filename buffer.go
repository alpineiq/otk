package otk

import (
	"bytes"
	"io"

	"golang.org/x/xerrors"
)

var (
	_ io.ReaderAt = (*Buffer)(nil)

	ErrOffsetOOB = xerrors.New("offset out of bounds")
)

func NewBuffer(buf []byte) *Buffer {
	var b Buffer
	b.Write(buf)
	return &b
}

func NewBufferString(s string) *Buffer {
	var b Buffer
	b.WriteString(s)
	return &b
}

// Buffer is a drop-in replacement for bytes.Buffer that implements the io.ReaderAt interface.
type Buffer struct {
	bytes.Buffer
	ref []byte
}

func (b *Buffer) WriteByte(c byte) (err error) {
	err = b.Buffer.WriteByte(c)
	b.ref = b.Buffer.Bytes()
	return
}

func (b *Buffer) WriteRune(r rune) (n int, err error) {
	n, err = b.Buffer.WriteRune(r)
	b.ref = b.Buffer.Bytes()
	return
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	n, err = b.Buffer.Write(p)
	b.ref = b.Buffer.Bytes()
	return
}

func (b *Buffer) WriteString(s string) (n int, err error) {
	n, err = b.Buffer.WriteString(s)
	b.ref = b.Buffer.Bytes()
	return
}

func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = b.Buffer.ReadFrom(r)
	b.ref = b.Buffer.Bytes()
	return
}

func (b *Buffer) ReadAt(p []byte, off int64) (n int, err error) {
	if int(off) > len(b.ref) {
		err = ErrOffsetOOB
		return
	}
	n = copy(p, b.ref[off:])
	return
}
