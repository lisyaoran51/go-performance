package goperformance

import (
	"runtime"
	"testing"
	"unsafe"
)

func BenchmarkStringKeyMapGC(b *testing.B) {
	StringKeyMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func BenchmarkXXHashKeyMapGC(b *testing.B) {
	XXHashKeyMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func BenchmarkPrimitiveValueMapGC(b *testing.B) {
	PrimitiveValueMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func BenchmarkPointerValueMapGC(b *testing.B) {
	PointerValueMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func BenchmarkBigStructValueMapGC(b *testing.B) {
	BigStructValueMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func BenchmarkSmallStructValueMapGC(b *testing.B) {
	SmallStructValueMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		runtime.GC() // 在測試中強制觸發 GC，觀察耗時
	}
}

func TestBigStructSize(t *testing.T) {
	var s BigStruct
	t.Logf("Size of BigStruct: %d", unsafe.Sizeof(s))
}
