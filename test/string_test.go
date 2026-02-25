package goperformance

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// ========== 測試函數 ==========

// 方法 1: String Plus (+ 運算符)
func StringPlus(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "hello"
	}
	return s
}

// 方法 2: strings.Builder (推薦)
func StringsBuilder(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString("hello")
	}
	return sb.String()
}

// 方法 3: bytes.Buffer
func BytesBuffer(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString("hello")
	}
	return buf.String()
}

// 方法 4: fmt.Sprintf (額外測試)
func FmtSprintf(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, "hello")
	}
	return s
}

// ========== Benchmark 1: 小量串接 (10次) ==========

func BenchmarkStringPlus_10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringPlus(10)
	}
}

func BenchmarkStringsBuilder_10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringsBuilder(10)
	}
}

func BenchmarkBytesBuffer_10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BytesBuffer(10)
	}
}

func BenchmarkFmtSprintf_10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FmtSprintf(10)
	}
}

// ========== Benchmark 2: 中量串接 (100次) ==========

func BenchmarkStringPlus_100(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringPlus(100)
	}
}

func BenchmarkStringsBuilder_100(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringsBuilder(100)
	}
}

func BenchmarkBytesBuffer_100(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BytesBuffer(100)
	}
}

func BenchmarkFmtSprintf_100(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FmtSprintf(100)
	}
}

// ========== Benchmark 3: 大量串接 (1000次) ==========

func BenchmarkStringPlus_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringPlus(1000)
	}
}

func BenchmarkStringsBuilder_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = StringsBuilder(1000)
	}
}

func BenchmarkBytesBuffer_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = BytesBuffer(1000)
	}
}

func BenchmarkFmtSprintf_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = FmtSprintf(1000)
	}
}

// ========== Benchmark 4: 預分配容量 ==========

func BenchmarkStringsBuilderWithCap_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.Grow(5000) // 預分配容量
		for j := 0; j < 1000; j++ {
			sb.WriteString("hello")
		}
		_ = sb.String()
	}
}

func BenchmarkBytesBufferWithCap_1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf := bytes.NewBuffer(make([]byte, 0, 5000)) // 預分配容量
		for j := 0; j < 1000; j++ {
			buf.WriteString("hello")
		}
		_ = buf.String()
	}
}

// ========== Benchmark 5: 真實場景 - 構建 JSON ==========

func BenchmarkStringPlus_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := "{"
		result += `"name":"John",`
		result += `"age":30,`
		result += `"city":"New York",`
		result += `"email":"john@example.com"`
		result += "}"
		_ = result
	}
}

func BenchmarkStringsBuilder_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString("{")
		sb.WriteString(`"name":"John",`)
		sb.WriteString(`"age":30,`)
		sb.WriteString(`"city":"New York",`)
		sb.WriteString(`"email":"john@example.com"`)
		sb.WriteString("}")
		_ = sb.String()
	}
}

func BenchmarkFmtSprintf_JSON(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result := fmt.Sprintf(`{"name":"%s","age":%d,"city":"%s","email":"%s"}`,
			"John", 30, "New York", "john@example.com")
		_ = result
	}
}

// ========== Benchmark 6: 真實場景 - 構建 SQL ==========

func BenchmarkStringPlus_SQL(b *testing.B) {
	b.ReportAllocs()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		query := "SELECT * FROM users WHERE id IN ("
		for j, id := range ids {
			if j > 0 {
				query += ","
			}
			query += fmt.Sprintf("%d", id)
		}
		query += ")"
		_ = query
	}
}

func BenchmarkStringsBuilder_SQL(b *testing.B) {
	b.ReportAllocs()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		var sb strings.Builder
		sb.WriteString("SELECT * FROM users WHERE id IN (")
		for j, id := range ids {
			if j > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(fmt.Sprintf("%d", id))
		}
		sb.WriteString(")")
		_ = sb.String()
	}
}

func BenchmarkFmtSprintf_SQL(b *testing.B) {
	b.ReportAllocs()
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		parts := make([]string, len(ids))
		for j, id := range ids {
			parts[j] = fmt.Sprintf("%d", id)
		}
		query := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s)",
			strings.Join(parts, ","))
		_ = query
	}
}
