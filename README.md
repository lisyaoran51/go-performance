# go-performance


```sh
go run -gcflags="-m" main.go
go run -gcflags="-m" .
go test -gcflags="-m"
```

## cpu
```sh
# Benchmark開頭的函式
go test -bench=. -cpuprofile=cpu.out
go tool pprof cpu.out

    top
    list MyFunc
    web 
```

## memory
```sh
# Benchmark開頭的函式
go test -bench=. -benchmem
go test -bench=. -memprofile=mem.out
go tool pprof mem.out
    top
    top -alloc_space
    top -alloc_objects
```

## gc
```sh
GODEBUG=gctrace=1 go test -bench=.

go test -bench=. -trace=trace.out
go tool trace trace.out
```

## test
```sh
go test -run TestAdd
# 只匹配名稱完全等於 TestPull 的函式
go test -v -run "^TestPull$"
```

## benchmark
```sh
go test -bench BenchmarkMyFunc
go test -bench=. -benchmem
```

## 執行次數
```sh
go test -bench=BenchmarkXXHashKeyMapGC -benchtime=1000x
```

## 物間個數
```sh
go tool pprof -sample_index=alloc_objects mem.out
```

## 競爭
```sh
go test -race -v race_test.go
go test -race -v -run TestDataRaceMutex
```