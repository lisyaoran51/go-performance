package goperformance

import (
	"runtime/debug"
	"testing"
)

func BenchmarkGCFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCFunc()
	}
}

func BenchmarkGCFunc2(b *testing.B) {

	// 重置測試計時
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		GCFunc()
	}

	// 測試結束後獲取 GC 統計
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	b.Logf("GC Runs: %d, Total Pause: %v", stats.NumGC, stats.PauseTotal)
}
