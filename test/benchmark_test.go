package goperformance

import (
	"runtime/debug"
	"strings"
	"testing"
)

func BenchmarkMyFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyFunc()
	}
}

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Fatal("wrong result")
	}
}

type BStruct struct {
	A int
	B int
}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s BStruct
		s.A = 1
		s.B = 2
		// s := BStruct{A: 1, B: 2}
		_ = s
	}
}

func BenchmarkCopyStruct1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s BStruct

		s2 := s
		_ = s2
	}
}

func BenchmarkCopyStruct2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s BStruct

		s2 := &s
		_ = s2
	}
}

// 範例：使用 + 號進行字串拼接（效能較差，因為字串是不可變的，每次拼接都會重新分配記憶體）
func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "hello"
		}
	}
}

// 範例：使用 strings.Builder 進行字串拼接（效能較好，預先分配或動態增長緩衝區）
func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 100; j++ {
			builder.WriteString("hello")
		}
		_ = builder.String()
	}

	// 測試結束後獲取 GC 統計
	var stats debug.GCStats
	debug.ReadGCStats(&stats)
	b.Logf("GC Runs: %d, Total Pause: %v", stats.NumGC, stats.PauseTotal)
}
