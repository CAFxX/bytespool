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

In addition, the following utility APIs are available:

- `Append` is a replacement for `append` that uses the backing pools to avoid allocations when appending to byte slices.

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

## Benchmarks

There is a small set of benchmarks you can run with `make bench` (you need to have [`benchstat`](https://golang.org/x/perf/cmd/benchstat) on your `PATH`), from which you can validate that in steady state no allocations are performed even for `[]byte`:

```
name                           time/op
GetBytesSlice-12               8.56ns ±14%
GetBytesSliceSemiStatic-12     8.26ns ±14%
GetBytesSliceStatic-12         8.19ns ±16%
GetBytesSlicePtr-12            5.18ns ±23%
GetBytesSlicePtrSemiStatic-12  4.74ns ±22%
GetBytesSlicePtrStatic-12      4.74ns ±22%
GetBytesBuffer-12              3.64ns ± 8%
GetBytesBufferSemiStatic-12    3.18ns ± 5%
GetBytesBufferStatic-12        3.12ns ± 5%
GetBufioReader-12              5.82ns ± 7%
GetBufioReaderSemiStatic-12    5.45ns ±10%
GetBufioReaderStatic-12        5.39ns ± 6%
GetBufioWriter-12              4.02ns ± 2%
GetBufioWriterSemiStatic-12    3.73ns ± 4%
GetBufioWriterStatic-12        3.56ns ± 3%

name                           alloc/op
GetBytesSlice-12                0.00B
GetBytesSliceSemiStatic-12      0.00B
GetBytesSliceStatic-12          0.00B
GetBytesSlicePtr-12             0.00B
GetBytesSlicePtrSemiStatic-12   0.00B
GetBytesSlicePtrStatic-12       0.00B
GetBytesBuffer-12               0.00B
GetBytesBufferSemiStatic-12     0.00B
GetBytesBufferStatic-12         0.00B
GetBufioReader-12               0.00B
GetBufioReaderSemiStatic-12     0.00B
GetBufioReaderStatic-12         0.00B
GetBufioWriter-12               0.00B
GetBufioWriterSemiStatic-12     0.00B
GetBufioWriterStatic-12         0.00B

name                           allocs/op
GetBytesSlice-12                 0.00
GetBytesSliceSemiStatic-12       0.00
GetBytesSliceStatic-12           0.00
GetBytesSlicePtr-12              0.00
GetBytesSlicePtrSemiStatic-12    0.00
GetBytesSlicePtrStatic-12        0.00
GetBytesBuffer-12                0.00
GetBytesBufferSemiStatic-12      0.00
GetBytesBufferStatic-12          0.00
GetBufioReader-12                0.00
GetBufioReaderSemiStatic-12      0.00
GetBufioReaderStatic-12          0.00
GetBufioWriter-12                0.00
GetBufioWriterSemiStatic-12      0.00
GetBufioWriterStatic-12          0.00
```

## Developing

Most of the code is in `gen/main.go`. You can run `make generate` to generate `bytespool.go`.

## TODO

- Split low-level APIs in a separate package
- Ensure that users of the low-level API are not forced to compile also the high-level APIs
- Define BufferPool-like interfaces for all objects (e.g. BytesBufferPool)
- Provide non-singleton pools
