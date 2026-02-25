package goperformance

import "fmt"

func main() {
	Func()
}

func PointerFunc() {

	var ptr *int

	ptr = new(int)

	*ptr = 1

	fmt.Println(*ptr)

	var num int
	num = 1

	fmt.Println(num)
}
