package goperformance

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestDataRace(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ❌ 這裡發生了競爭！
			// counter++ 實際上是：讀取 -> 加1 -> 寫回
			// 多個協程可能同時讀到舊的值
			counter++
		}()
	}

	wg.Wait()
	t.Logf("Final counter: %d", counter)
}

func TestDataRaceMutex(t *testing.T) {
	counter := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ❌ 這裡發生了競爭！
			// counter++ 實際上是：讀取 -> 加1 -> 寫回
			// 多個協程可能同時讀到舊的值
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}

	wg.Wait()
	t.Logf("Final counter: %d", counter)
}

func TestDataRaceAtomic(t *testing.T) {
	var counter int64 // 必須使用 int32 或 int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// ✅ 原子加法，不會有競爭
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	t.Logf("Final counter: %d", counter)
}
