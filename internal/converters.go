package internal

import (
	"bufio"
	"bytes"
	"io"
)

// Bb2bs extracts the backing slice from a bytes.Buffer.
// The bytes.Buffer can not be used and must be discarded when this function returns.
func Bb2bs(b *bytes.Buffer) []byte {
	var zeros [256]byte
	b.Reset()
	c, r := b.Cap()/len(zeros), b.Cap()%len(zeros)
	for i := 0; i < c; i++ {
		b.Write(zeros[:])
	}
	b.Write(zeros[:r])
	return b.Bytes()
}

// Br2bs extracts the backing slice from a bufio.Reader
// The bufio.Reader can not be used and must be discarded when this function returns.
func Br2bs(r *bufio.Reader) []byte {
	d := doroboReader{}
	r.Reset(&d)
	r.ReadByte()
	r.Reset(nil)
	return d.b
}

type doroboReader struct {
	io.Reader
	b []byte
}

func (d *doroboReader) Read(buf []byte) (int, error) {
	d.b = buf
	return len(buf), nil
}

// Bw2bs extracts the backing slice from a bufio.Writer
// The bufio.Writer can not be used and must be discarded when this function returns.
func Bw2bs(w *bufio.Writer) []byte {
	d := doroboWriter{}
	w.Reset(&d)
	zeros := _zeros[:]
	if w.Size() < len(zeros) {
		zeros = zeros[:w.Size()]
	}
	for d.b == nil {
		w.Write(zeros)
	}
	w.Reset(nil)
	return d.b
}

type doroboWriter struct {
	io.Writer
	b []byte
}

func (d *doroboWriter) Write(buf []byte) (int, error) {
	d.b = buf
	return len(buf), nil
}

var _zeros = [256]byte{}
