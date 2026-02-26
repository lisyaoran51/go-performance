package main

import (
	"fmt"
	"unsafe"
)

// 模擬 Go runtime 內部的 maptype 結構 (簡化版)
type maptype struct {
	size       uintptr
	ptrdata    uintptr
	hash       uint32
	tflag      uint8
	align      uint8
	fieldAlign uint8
	kind       uint8
	equal      func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata     *byte
	str        int32
	ptrToThis  int32
	key        unsafe.Pointer
	elem       unsafe.Pointer
	bucket     unsafe.Pointer
	hasher     func(unsafe.Pointer, uintptr) uintptr
	keysize    uint8  // key size
	valuesize  uint8  // value size
	bucketsize uint16 // bucket size
	flags      uint32
}

const (
	indirectKey   = 1 // 標記 Key 是否為指標存儲
	indirectValue = 2 // 標記 Value 是否為指標存儲
)

func isIndirect(m interface{}) (keyInd, valInd bool) {
	// 透過空介面取得底層 map 類型描述
	type emptyInterface struct {
		typ unsafe.Pointer
		val unsafe.Pointer
	}
	ei := (*emptyInterface)(unsafe.Pointer(&m))
	mt := (*maptype)(ei.typ)

	// 檢查 flags 位元
	return (mt.flags & indirectKey) != 0, (mt.flags & indirectValue) != 0
}

func main() {
	m1 := make(map[int]int)             // 小於 128 bytes
	m2 := make(map[[129]byte][120]byte) // 大於 128 bytes

	k1, v1 := isIndirect(m1)
	k2, v2 := isIndirect(m2)

	fmt.Printf("map[int]int -> Key 指標: %v, Val 指標: %v\n", k1, v1)
	fmt.Printf("map[129byte] -> Key 指標: %v, Val 指標: %v\n", k2, v2)
}
