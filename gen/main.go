package main

import (
	"fmt"
	"os"
	"text/template"
)

// TODO: get/getb should try also bigger pools if the correct pool is empty

const tmpl = `// Code generated by bytespool-gen. DO NOT EDIT.

// Package bytespool implements a series of sync.Pools tailored to safely recycle
// byte slices and bytes.Buffers. It works by defining a number of buckets, each for
// slices or Buffers of a certain range of capacities, and directing the Put and Get
// requests to the appropriate bucket.
// This package is autogenerated to allow to customize the granularity of the buckets
// and the range of supported capacities. This version was generated with the following
// buckets:{{range .}} {{.Short}}B{{end}}.
package bytespool

import (
	"bytes"
	"bufio"
	"io"
	"net/http/httputil"
	"sync"	
)

var buckets = [...]int{ {{range .}}{{.Bytes}},{{end}} }

{{range .}}
var pool{{.Short}}  sync.Pool // *[]byte, {{.Bytes}} <= cap < {{.BytesHigh}}
var poolb{{.Short}} sync.Pool // *bytes.Buffer, {{.Bytes}} <= Cap < {{.BytesHigh}}
var poolr{{.Short}} sync.Pool // *bufio.Reader, {{.Bytes}} <= Size < {{.BytesHigh}}
var poolw{{.Short}} sync.Pool // *bufio.Writer, {{.Bytes}} <= Size < {{.BytesHigh}}

func get{{.Short}}() *[]byte {
    p, _ := pool{{.Short}}.Get().(*[]byte)
    return p
}

func getb{{.Short}}() *bytes.Buffer {
    b, _ := poolb{{.Short}}.Get().(*bytes.Buffer)
    return b
} 

func getr{{.Short}}() *bufio.Reader {
    r, _ := poolr{{.Short}}.Get().(*bufio.Reader)
    return r
} 

func getw{{.Short}}() *bufio.Writer {
    w, _ := poolw{{.Short}}.Get().(*bufio.Writer)
    return w
} 

func put{{.Short}}(b *[]byte) {
	pool{{.Short}}.Put(b)
}

func putb{{.Short}}(b *bytes.Buffer) {
	poolb{{.Short}}.Put(b)
}

func putr{{.Short}}(r *bufio.Reader) {
	poolr{{.Short}}.Put(r)
}

func putw{{.Short}}(w *bufio.Writer) {
	poolw{{.Short}}.Put(w)
}

// GetBytesBuffer{{.Short}} gets a bytes.Buffer with a capacity of at least {{.Short}} bytes.
func GetBytesBuffer{{.Short}}() *bytes.Buffer {
    if b := getb{{.Short}}(); b != nil {
        return b
	}
	if p := get{{.Short}}(); p != nil {
        return bytes.NewBuffer(*p)
	}
    return bytes.NewBuffer(make([]byte, {{.Bytes}}))
}

// GetBytesSlice{{.Short}} gets a byte slice with a capacity of at least {{.Short}} bytes and length of {{.Short}} bytes.
func GetBytesSlice{{.Short}}() []byte {
    if p := get{{.Short}}(); p != nil {
        return *p
	}
	if b := getb{{.Short}}(); b != nil {
		return bb2bs(b)
	}
    p := make([]byte, {{.Bytes}})
    return p
}

// GetBytesSlicePtr{{.Short}} is like GetBytesSlice{{.Short}} but returns a pointer to the slice instead. This is needed
// as PutBytesSlice{{.Short}} requires a pointer-sized allocation per call, whereas PutBytesSlicePtr{{.Short}} does not
// allocate.
func GetBytesSlicePtr{{.Short}}() *[]byte {
    if p := get{{.Short}}(); p != nil {
        return p
	}
	if b := getb{{.Short}}(); b != nil {
		p := bb2bs(b)
		return &p
	}
    p := make([]byte, {{.Bytes}})
    return &p
}

func GetBufioReader{{.Short}}(pr io.Reader) *bufio.Reader {
	if r := getr{{.Short}}(); r != nil {
		r.Reset(pr)
        return r
	}
	return bufio.NewReaderSize(pr, {{.Bytes}})
}

func GetBufioWriter{{.Short}}(pw io.Writer) *bufio.Writer {
	if w := getw{{.Short}}(); w != nil {
		w.Reset(pw)
        return w
	}
	return bufio.NewWriterSize(pw, {{.Bytes}})
}

// PutBytesBuffer{{.Short}} recycles the passed bytes.Buffer. If the bytes.Buffer can not be recycled (e.g. because its capacity
// is too small or too big) false is returned and the bytes.Buffer is unmodified, otherwise the bytes.Buffer is
// recycled and true is returned. In the latter case, the caller should never use again the passed bytes.Buffer.
// PutBytesBuffer{{.Short}} is optimized for bytes.Buffer of capacity [{{.Bytes}}, {{.BytesHigh}}) but will accepts other
// sizes as well.
func PutBytesBuffer{{.Short}}(b *bytes.Buffer) bool {
    if b == nil {
        return false
    }
    if l := b.Cap(); l < {{.Bytes}} || l >= {{.BytesHigh}} {
        return PutBytesBuffer(b)
    }
	b.Reset()
    putb{{.Short}}(b)
    return true
}

// PutBytesSlice{{.Short}} recycles the passed byte slice. If the byte slice can not be recycled (e.g. because its capacity
// is too small or too big) false is returned and the bytes.Buffer is returned unmodified, otherwise the byte slice is
// recycled and true is returned. In the latter case, the caller should never use again the passed byte slice.
// PutBytesSlice{{.Short}} is optimized for byte slice of capacity [{{.Bytes}}, {{.BytesHigh}}) but will accepts other
// sizes as well.
// PutBytesSlice{{.Short}}, contrary to PutBytesSlicePtr{{.Short}}, will perform a pointer-sized allocation for each call.
func PutBytesSlice{{.Short}}(p []byte) bool {
    if l := cap(p); l < {{.Bytes}} || l >= {{.BytesHigh}} {
        return PutBytesSlice(p)
    }
    p = p[0:{{.Bytes}}]
    put{{.Short}}(&p)
    return true
}

// PutBytesSlicePtr{{.Short}} is like PutBytesSlice{{.Short}}, but it accepts a pointer to the byte slice and does not perform 
// a pointer-sized allocation for each call.
func PutBytesSlicePtr{{.Short}}(p *[]byte) bool {
    if p == nil {
        return false
    }
    if l := cap(*p); l < {{.Bytes}} || l >= {{.BytesHigh}} {
        return PutBytesSlicePtr(p)
    }
    *p = (*p)[0:{{.Bytes}}]
    put{{.Short}}(p)
    return true
}

func PutBufioReader{{.Short}}(r *bufio.Reader) bool {
    if r == nil {
        return false
    }
    if l := r.Size(); l < {{.Bytes}} || l >= {{.BytesHigh}} {
        return PutBufioReader(r)
    }
	r.Reset(nil) // to not keep the parent reader alive
    putr{{.Short}}(r)
    return true
}

func PutBufioWriter{{.Short}}(w *bufio.Writer) bool {
    if w == nil {
        return false
    }
    if l := w.Size(); l < {{.Bytes}} || l >= {{.BytesHigh}} {
        return PutBufioWriter(w)
    }
	w.Reset(nil) // to not keep the parent writer alive
    putw{{.Short}}(w)
    return true
}

// BufferPool{{.Short}} is a httputil.BufferPool that provides byte slices of {{.Short}} bytes. 
type BufferPool{{.Short}} struct {
	httputil.BufferPool
}

// Get implements httputil.BufferPool.Get
func (_ BufferPool{{.Short}}) Get() []byte {
	return GetBytesSlice{{.Short}}()
}

// Put implements httputil.BufferPool.Put
func (_ BufferPool{{.Short}}) Put(b []byte) {
	PutBytesSlice{{.Short}}(b)
}

// BufferPtrPool{{.Short}} is like BufferPool{{.Short}}, but using pointer to byte slices instead (to avoid allocations during Put).
// For this reason it is not compatible with httputil.BufferPool.
type BufferPtrPool{{.Short}} struct {}

// Get gets a byte slice from the pool. See GetBytesSlicePtr{{.Short}} for details.
func (_ BufferPtrPool{{.Short}}) Get() *[]byte {
	return GetBytesSlicePtr{{.Short}}()
}

// Put inserts a byte slice in the pool. See PutBytesSlicePtr{{.Short}} for details.
func (_ BufferPtrPool{{.Short}}) Put(b *[]byte) {
	PutBytesSlicePtr{{.Short}}(b)
}
{{end}}

// GetBytesBuffer returns a bytes.Buffer with at least size bytes of capacity.
// If your code uses buffers of static size, it is more performant to call one of the GetBytesBufferXxx functions instead.
// Calling GetBytesBuffer with a negative size panics.
func GetBytesBuffer(size int) *bytes.Buffer {
    switch {
    {{range .}}
    case size > {{.BytesLow}} && size <= {{.Bytes}}:
        return GetBytesBuffer{{.Short}}()
    {{end}}
	default:
		return bytes.NewBuffer(make([]byte, size))
	}
}

// GetBytesSlice returns a bytes.Buffer with at least size bytes of capacity.
// If your code uses buffers of static size, it is more performant to call one of the GetBytesSliceXxx functions instead.
// GetBytesSlice, contrary to GetBytesSlicePtr, performs a pointer-sized allocation per call.
// Calling GetBytesSlice with a negative size panics.
func GetBytesSlice(size int) []byte {
    switch {
    {{range .}}
    case size > {{.BytesLow}} && size <= {{.Bytes}}:
        return GetBytesSlice{{.Short}}()
    {{end}}
	default:
		return make([]byte, size)
	}
}

// GetBytesSlicePtr is like GetBytesSlice but returns a pointer to the byte slice.
// Contrary to GetBytesSlice, it does not perform a pointer-sized allocation per call.
func GetBytesSlicePtr(size int) *[]byte {
    switch {
    {{range .}}
    case size > {{.BytesLow}} && size <= {{.Bytes}}:
        return GetBytesSlicePtr{{.Short}}()
    {{end}}
	default:
		p := make([]byte, size)
		return &p
	}
}

// PutBytesBuffer recycles the passed bytes.Buffer. If the bytes.Buffer can not be recycled (e.g. because its capacity
// is too small or too big) false is returned and the bytes.Buffer is unmodified, otherwise the bytes.Buffer is
// recycled and true is returned. In the latter case, the caller should never use again the passed bytes.Buffer.
func PutBytesBuffer(b *bytes.Buffer) bool {
    if b == nil {
        return false
	}
	size := b.Cap()
    switch {
    {{range .}}
	case size >= {{.Bytes}} && size < {{.BytesHigh}}:
		b.Reset()
		putb{{.Short}}(b)
    {{end}}
	default:
		return false
	}
	return true
}

func PutBytesSlice(b []byte) bool {
	size := cap(b)
    switch {
    {{range .}}
    case size >= {{.Bytes}} && size < {{.BytesHigh}}:
		b = b[0:{{.Bytes}}]
		put{{.Short}}(&b)
    {{end}}
	default:
		return false
	}
	return true
}

func PutBytesSlicePtr(b *[]byte) bool {
    if b == nil {
        return false
	}
	size := cap(*b)
    switch {
    {{range .}}
    case size >= {{.Bytes}} && size < {{.BytesHigh}}:
		*b = (*b)[0:{{.Bytes}}]
		put{{.Short}}(b)
    {{end}}
	default:
		return false
	}
	return true
}

func PutBufioReader(r *bufio.Reader) bool {
    if r == nil {
        return false
	}
	size := r.Size()
    switch {
    {{range .}}
	case size >= {{.Bytes}} && size < {{.BytesHigh}}:
		r.Reset(nil) // to not keep the parent reader alive
		putr{{.Short}}(r)
    {{end}}
	default:
		return false
	}
	return true
}

func PutBufioWriter(w *bufio.Writer) bool {
    if w == nil {
        return false
	}
	size := w.Size()
    switch {
    {{range .}}
	case size >= {{.Bytes}} && size < {{.BytesHigh}}:
		w.Reset(nil) // to not keep the parent writer alive
		putw{{.Short}}(w)
    {{end}}
	default:
		return false
	}
	return true
}

func bb2bs(b *bytes.Buffer) []byte {
	var zeros [256]byte
	b.Reset()
	c, r := b.Cap() / len(zeros), b.Cap() % len(zeros)
	for i := 0; i < c; i++ {
		b.Write(zeros[:])
	} 
	b.Write(zeros[:r])
	return b.Bytes()
}
`

type tmpldata struct {
	Short     string
	Bytes     int
	BytesHigh int
	BytesLow  int
}

func main() {
	t, _ := template.New("Code").Parse(tmpl)
	td := []tmpldata{}
	for i := 256; i <= 16*1024*1024; i = i * 2 {
		td = append(td, tmpldata{
			Short:     bytesToShort(i),
			Bytes:     i,
			BytesHigh: i * 2,
			BytesLow:  i / 2,
		})
	}
	td[0].BytesLow = 0
	t.Execute(os.Stdout, td)
}

func bytesToShort(b int) string {
	switch {
	case b%(1024*1024*1024) == 0:
		return fmt.Sprintf("%dG", b/(1024*1024*1024))
	case b%(1024*1024) == 0:
		return fmt.Sprintf("%dM", b/(1024*1024))
	case b%1024 == 0:
		return fmt.Sprintf("%dK", b/1024)
	}
	return fmt.Sprintf("%d", b)
}
