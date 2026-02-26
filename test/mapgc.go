package goperformance

import (
	"fmt"

	"github.com/cespare/xxhash/v2"
)

func PrimitiveValueMap() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 1000000; i++ {
		m[i] = i
	}
	return m
}

func PointerValueMap() map[int]*int {
	m := make(map[int]*int)
	for i := 0; i < 1000000; i++ {
		m[i] = &i
	}
	return m
}

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

type BigStruct struct {
	A int64
	B int64
	C int64
	D int64
	E int64
	F int64
	G int64
	H int64
	I int64
	J int64
	K int64
	L int64
	M int64
	N int64
	O int64
	P int64
}

func BigStructValueMap() map[int]BigStruct {
	m := make(map[int]BigStruct)
	for i := 0; i < 1000000; i++ {
		m[i] = BigStruct{}
	}
	return m
}

type SmallStruct struct {
	A int64
}

func SmallStructValueMap() map[int]SmallStruct {
	m := make(map[int]SmallStruct)
	for i := 0; i < 1000000; i++ {
		m[i] = SmallStruct{}
	}
	return m
}
