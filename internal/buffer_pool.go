package internal

import (
	"net/http/httputil"
	"sync"
)

type BufferPool struct {
	httputil.BufferPool                     // implements
	P                   httputil.BufferPool // actual underlying buffer pool
	Sz                  int
}

func (b BufferPool) Get() []byte {
	p := b.P.Get()
	p = p[0:b.Sz]
	return p
}

func (b BufferPool) Put(p []byte) {
	b.P.Put(p)
}

type SingleSizeBufferPool struct {
	httputil.BufferPool           // implements
	p                   sync.Pool // underlying sync.Pool
	Sz                  int
}

func (b *SingleSizeBufferPool) Get() []byte {
	if p, _ := b.p.Get().([]byte); p != nil {
		return p
	}
	return make([]byte, b.Sz)
}

func (b *SingleSizeBufferPool) Put(p []byte) {
	if cap(p) < b.Sz || cap(p) > b.Sz*2 {
		return
	}
	b.p.Put(p[0:b.Sz])
}
