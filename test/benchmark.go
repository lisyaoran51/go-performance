package goperformance

func MyFunc() {
	var ptr *int

	ptr = new(int)

	*ptr = 1
}

func Add(a, b int) int {
	return a + b
}
