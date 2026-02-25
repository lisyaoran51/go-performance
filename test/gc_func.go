package goperformance

import (
	"fmt"
	"runtime"
	"time"
)

// 用於列印當前的記憶體統計資訊
func printMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// Alloc: 目前堆上對象佔用的位元組
	// TotalAlloc: 累計分配的位元組（不隨 GC 減少）
	// Sys: 從系統申請的總記憶體
	// NumGC: 已完成的 GC 次數
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func GCFunc() {
	// fmt.Println("--- 初始狀態 ---")
	// printMemUsage()

	// 1. 模擬大量分配記憶體
	fmt.Println("\n--- 分配大量記憶體中... ---")
	container := make([][]byte, 0)
	for i := 0; i < 10; i++ {
		// 每次分配約 50MiB 的空間
		data := make([]byte, 50*1024*1024)
		container = append(container, data)
		// printMemUsage()
		time.Sleep(100 * time.Millisecond)
	}
}

func GCFuncGC() {
	// 2. 釋放引用，讓對象變為「垃圾」
	fmt.Println("\n--- 釋放引用並等待 GC ---")
	// container = nil

	// 3. 手動觸發 GC (在實際生產環境通常讓它自動運行)
	// 你也可以嘗試註解掉下面這行，觀察 Go 如何自動觸發 GC
	runtime.GC()

	fmt.Println("--- 手動執行 runtime.GC() 後 ---")
	printMemUsage()

	// 保持程式運行一下，觀察系統是否將記憶體歸還給作業系統
	time.Sleep(2 * time.Second)
	fmt.Println("\n--- 結束前最終檢查 ---")
	printMemUsage()

}
