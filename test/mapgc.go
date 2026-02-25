package goperformance

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

func StringKeyMap() map[string]int {
	m := make(map[string]int)
	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("user:%d:session:%d", i, i*2)
		m[key] = i
	}
	return m
}

func XXHashKeyMap() map[uint64]int {
	m := make(map[uint64]int)
	for i := 0; i < 1000000; i++ {
		key := fmt.Sprintf("user:%d:session:%d", i, i*2)
		hash := xxhash.Sum64String(key)
		m[hash] = i
	}
	return m
}

// func main() {
// 	// String Key Map
// 	f1, _ := os.Create("string_heap.prof")
// 	defer f1.Close()

// 	stringMap := StringKeyMap()
// 	runtime.GC()
// 	pprof.WriteHeapProfile(f1)

// 	_ = stringMap
// 	stringMap = nil
// 	runtime.GC()

// 	// xxhash Key Map
// 	f2, _ := os.Create("xxhash_heap.prof")
// 	defer f2.Close()

// 	xxhashMap := XXHashKeyMap()
// 	runtime.GC()
// 	pprof.WriteHeapProfile(f2)

// 	_ = xxhashMap

// 	fmt.Println("✅ Profiles saved")
// 	fmt.Println("View with: go tool pprof string_heap.prof")
// 	fmt.Println("View with: go tool pprof xxhash_heap.prof")
// }
