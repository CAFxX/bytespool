package internal

import "bytes"

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
