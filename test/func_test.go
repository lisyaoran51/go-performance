package goperformance

import "testing"

func TestFunc(t *testing.T) {
	a := 1
	go func() {
		a = 2
	}()
	t.Log(a)
}
