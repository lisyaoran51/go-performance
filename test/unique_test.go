package goperformance

import "testing"

func BenchmarkUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unique()
	}
}
