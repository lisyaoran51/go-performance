package goperformance

import (
	"runtime"
	"testing"
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
