package bytepool

import (
	"testing"
)

func BenchmarkAlloc256R(b *testing.B) {
	for i := 0; i < b.N; i++ {
		consumer(make([]byte, 256))
	}
}

func BenchmarkAlloc256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Free(Alloc(256))
	}
}

func BenchmarkAlloc1024R(b *testing.B) {
	for i := 0; i < b.N; i++ {
		consumer(make([]byte, 1024))
	}
}

func BenchmarkAlloc1024(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Free(Alloc(1024))
	}
}

func BenchmarkAlloc64kR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		consumer(make([]byte, 64*1024))
	}
}

func BenchmarkAlloc64k(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Free(Alloc(64 * 1024))
	}
}

var consumer = func([]byte) {}
