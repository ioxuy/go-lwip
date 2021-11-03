package golwip

import "sync"

type LWIPPool interface {
	Get(size int) []byte
	Put(buf []byte) error
}

var bufferPool LWIPPool

// SetPoolAllocator set custom buffer pool allocator.
func SetPoolAllocator(p LWIPPool) {
	bufferPool = p
}

func NewBytes(size int) []byte {
	return bufferPool.Get(size)
}

func FreeBytes(b []byte) {
	_ = bufferPool.Put(b)
}

const bufSize = 2 * 1024

type localPool struct {
	pool *sync.Pool
}

func (p *localPool) Get(size int) []byte {
	if size <= bufSize {
		return p.pool.Get().([]byte)
	} else {
		return make([]byte, size)
	}
}

func (p *localPool) Put(b []byte) error {
	if len(b) >= bufSize {
		p.pool.Put(b)
	}
	return nil
}

// InitLocalPool init local buffer pool.
func InitLocalPool() {
	pool := &localPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, bufSize)
			},
		},
	}
	SetPoolAllocator(pool)
}
