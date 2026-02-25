package goperformance

import (
	"fmt"
	"unique"
)

func Unique() {
	s1 := "apple"
	s2 := string([]byte{'a', 'p', 'p', 'l', 'e'})

	// 使用 unique.Make 進行駐留
	h1 := unique.Make(s1)
	h2 := unique.Make(s2)

	// 直接比較 Handle 的指針地址
	if h1 == h2 {
		fmt.Println("h1 and h2 point to the same interned instance!")
	}

	// 拿回原始字串
	fmt.Println(h1.Value())
}
