package goperformance

import "testing"

const size = 1000000

func BenchmarkCopyForLoop(b *testing.B) {
	src := make([]int, size)
	dest := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			dest[j] = src[j]
		}
	}
}

func BenchmarkCopyAppend(b *testing.B) {
	src := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dest := make([]int, 0, size)
		dest = append(dest, src...)
	}
}

func BenchmarkCopyBuiltIn(b *testing.B) {
	src := make([]int, size)
	dest := make([]int, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(dest, src)
	}
}
