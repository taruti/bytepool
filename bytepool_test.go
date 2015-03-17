package bytepool

import (
	"testing"
)

func TestBytepool(t *testing.T) {
	for i := 0; i < 99; i++ {
		bs := Alloc(i)
		if len(bs) != i {
			t.Fatal("Allocated slice has wrong length")
		}
		Free(bs[:0])
	}
	for i := 0; i < MaxSize; i += 3 + i/2 {
		t.Logf("Alloc %d %x", i, i)
		bs := Alloc(i)
		if len(bs) != i {
			t.Fatal("Allocated slice has wrong length")
		}
		Free(bs)
	}
	for i := 0; i < MaxSize; i += 3 + i/3 {
		bs := Alloc(i)
		if len(bs) != i {
			t.Fatal("Allocated slice has wrong length")
		}
		Free(bs)
	}
}

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
