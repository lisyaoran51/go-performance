package goperformance

import (
	"context"
	"testing"
	"time"

	"go.uber.org/goleak"
)

func TestWaitPush(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		ch <- 0
	}()
	time.Sleep(time.Second)
}

func TestWaitPull(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		ch <- 0
	}()
	time.Sleep(time.Second)
}

func TestPushNil(t *testing.T) {
	defer goleak.VerifyNone(t)
	var ch chan int
	go func() {
		ch <- 0
		t.Log("goroutine end")
	}()
	time.Sleep(time.Second)
}

func TestPullNil(t *testing.T) {
	defer goleak.VerifyNone(t)
	var ch chan int
	go func() {
		i := <-ch
		t.Log(i)
		t.Log("goroutine end")
	}()
	time.Sleep(time.Second)
}

func TestPull(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int, 1)
	go func() {
		i := <-ch
		t.Log(i)
		t.Log("goroutine end")
	}()
	ch <- 0
	time.Sleep(time.Second)
}

func TestPush(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int, 1)
	go func() {
		ch <- 0
		t.Log("goroutine end")
	}()
	time.Sleep(time.Second)
}

func TestPullClose(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		i := <-ch
		t.Log(i)
		t.Log("goroutine end")
	}()
	close(ch)
	time.Sleep(time.Second)
}

func TestPushClose(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		ch <- 0
		t.Log("goroutine end")
	}()
	close(ch)
	time.Sleep(time.Second)
}

func TestPullBufferedClose(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		time.Sleep(10 * time.Millisecond)
		i := <-ch
		t.Log(i)
		i = <-ch
		t.Log(i)
		t.Log("goroutine end")
	}()
	ch <- 1
	close(ch)
	time.Sleep(time.Second)
}

func TestPullCloseOK(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		i, ok := <-ch
		if !ok {
			t.Log("channel closed")
			return
		}
		t.Log(i)
		t.Log("goroutine end")
	}()
	time.Sleep(time.Second)
	close(ch)
	time.Sleep(time.Second)
}

func TestPullCloseFor(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		for i := range ch {
			t.Log(i)
		}
		t.Log("goroutine end")
	}()
	ch <- 1
	close(ch)
	time.Sleep(time.Second)
}

func TestPullSelectCtx(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				t.Log(i)
			case <-ctx.Done():
				t.Log("goroutine end")
				cancel()
				return
			}
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	cancel()
	time.Sleep(time.Second)
}

func TestPullSelectClose(t *testing.T) {
	defer goleak.VerifyNone(t)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				t.Log(i)
			case <-ctx.Done():
				t.Log("goroutine end")
				cancel()
				return
			}
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	time.Sleep(time.Second)
}

func TestPullSelectTimer(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				t.Log(i)
			case <-time.After(time.Second):
				t.Log("goroutine end")
				return
			}
		}
	}()
	time.Sleep(time.Second)
	time.Sleep(time.Second)
}

func TestPullSelectTicker(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case i, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				t.Log(i)
			case <-ticker.C:
				t.Log("goroutine end")
				return
			}
		}
	}()
	time.Sleep(time.Second)
	time.Sleep(time.Second)
}

func TestPullSelectTimerLeak(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				// t.Log(i)
			case <-time.After(time.Second * 300):
				t.Log("goroutine end")
				return
			}
		}
	}()
	for i := 0; i < 10000000; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(time.Second)
}

func TestPullSelectTimerLeak2(t *testing.T) {
	defer goleak.VerifyNone(t)
	ch := make(chan int)
	go func() {
		timer := time.NewTimer(time.Second * 1)
		for {
			select {
			case _, ok := <-ch:
				if !ok {
					t.Log("channel closed")
					return
				}
				// t.Log(i)
			case <-timer.C:
				t.Log("goroutine end")
				return
			}
			if !timer.Stop() {
				// 如果 Stop 回傳 false，代表 timer 已經過期或已停止
				// 必須嘗試從 channel 中抽走那筆過期資料，避免下次 select 誤觸
				select {
				case <-timer.C:
				default:
				}
			}
			timer.Reset(time.Second * 1)
		}
	}()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(time.Second * 2)
}
