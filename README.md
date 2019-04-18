# bytespool
A collection of pools for variable-sized objects: `[]byte`, `*bufio.Reader`, `*bufio.Writer` and `*bytes.Buffer`.

## High-level APIs

Once you import `github.com/CAFxX/bytespool` you have access to these high-level APIs:

- `GetBufferPool(size int) httputil.BufferPool`
- `GetBufioReader(pr io.Reader, size int) *bufio.Reader` and `PutBufioReader(r *bufio.Reader) bool`
- `GetBufioWriter(pw io.Writer, size int) *bufio.Writer` and `PutBufioWriter(w *bufio.Writer) bool`
- `GetBytesBuffer(size int) *bytes.Buffer` and `PutBytesBuffer(b *bytes.Buffer) bool`
- `GetBytesSlice(size int) []byte` and `PutBytesSlice(b []byte) bool`
- `GetBytesSlicePtr(size int) *[]byte` and `PutBytesSlicePtr(b *[]byte) bool`

`bytespool` relies on code generation to allow to control the granularity of the backing pools.

The [generated code](https://godoc.org/github.com/CAFxX/bytespool) contains specialized versions of all functions for each of the backing pools. These functions can be used to provide a small speed-up if you statically know the desired sizes at compile-time.

Note that it is allowed to pass to the `PutXxx` functions objects that have been obtained not from `GetXxx` or that have undergone resizing. If the `PutXxx` functions detect that the object can not be recycled they will return false without modifying the object.

## Examples

### Minimizing allocations in ReverseProxy

[httputil.ReverseProxy](https://golang.org/pkg/net/http/httputil/#ReverseProxy) can use an [httputil.BufferPool](https://golang.org/pkg/net/http/httputil/#BufferPool) to minimize allocations. You can provide one as follows:

```go
rp := httputil.NewSingleHostReverseProxy(target)
rp.BufferPool = bytespool.GetBufferPool(32*1024)
```
