package test

import (
	"testing"

	"github.com/CAFxX/bytespool"
)

func BenchmarkGetBytesSlice(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlice(1 << 10)
			bytespool.PutBytesSlice(b)
		}
	})
}

func BenchmarkGetBytesSliceSemiStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlice1K()
			bytespool.PutBytesSlice(b)
		}
	})
}

func BenchmarkGetBytesSliceStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlice1K()
			bytespool.PutBytesSlice1K(b)
		}
	})
}

func BenchmarkGetBytesSlicePtr(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlicePtr(1 << 10)
			bytespool.PutBytesSlicePtr(b)
		}
	})
}

func BenchmarkGetBytesSlicePtrSemiStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlicePtr1K()
			bytespool.PutBytesSlicePtr(b)
		}
	})
}

func BenchmarkGetBytesSlicePtrStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesSlicePtr1K()
			bytespool.PutBytesSlicePtr1K(b)
		}
	})
}

func BenchmarkGetBytesBuffer(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesBuffer(1 << 10)
			bytespool.PutBytesBuffer(b)
		}
	})
}

func BenchmarkGetBytesBufferSemiStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesBuffer1K()
			bytespool.PutBytesBuffer(b)
		}
	})
}

func BenchmarkGetBytesBufferStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBytesBuffer1K()
			bytespool.PutBytesBuffer1K(b)
		}
	})
}

func BenchmarkGetBufioReader(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioReader(nil, 1<<10)
			bytespool.PutBufioReader(b)
		}
	})
}

func BenchmarkGetBufioReaderSemiStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioReader1K(nil)
			bytespool.PutBufioReader(b)
		}
	})
}

func BenchmarkGetBufioReaderStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioReader1K(nil)
			bytespool.PutBufioReader1K(b)
		}
	})
}

func BenchmarkGetBufioWriter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioWriter(nil, 1<<10)
			bytespool.PutBufioWriter(b)
		}
	})
}

func BenchmarkGetBufioWriterSemiStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioWriter1K(nil)
			bytespool.PutBufioWriter(b)
		}
	})
}

func BenchmarkGetBufioWriterStatic(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			b := bytespool.GetBufioWriter1K(nil)
			bytespool.PutBufioWriter1K(b)
		}
	})
}
