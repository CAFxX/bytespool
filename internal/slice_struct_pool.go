package internal

import "sync"

var pools sync.Pool // *[]byte, containing zero slice
var emptySlice = []byte{}

// Puts returns a copy of the slice pointed to by p and recycles the heap-allocated slice structure.
// Only the slice is copied, not the backing array.
func Puts(p *[]byte) []byte {
	b := *p
	*p = emptySlice
	pools.Put(p)
	return b
}

// Gets stores the slice b into a slice struct allocated on the heap.
func Gets(b []byte) *[]byte {
	p := gets()
	*p = b
	return p
}

func gets() *[]byte {
	if p, _ := pools.Get().(*[]byte); p != nil {
		return p
	}
	return &[]byte{}
}
