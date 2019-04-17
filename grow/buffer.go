package grow

import (
	"bytes"

	"github.com/CAFxX/bytespool"
)

func BytesBuffer(b *bytes.Buffer, n int) *bytes.Buffer {
	t := (b.Len() + n) * 2
	if t <= b.Cap() {
		return b
	}
	nb := bytespool.GetBytesBuffer(t)
	nb.Write(b.Bytes())
	bytespool.PutBytesBuffer(b)
	return nb
}
