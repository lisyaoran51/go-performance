package goperformance

func Func() {
	a := 1
	go func() {
		a = 2
	}()
	_ = a
}
