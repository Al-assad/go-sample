package main

import (
	"fmt"
	"testing"
)

// @desc 对 goroutine 进行基准测试

func main() {
	fmt.Println("sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
	/*          N             Time 1 op
	sync     6244950         188 ns/op
	buffered 19868823        59.6 ns/op
	*/
}

// 同步 channel 基准测试
func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

// 缓冲 channel 基准测试
func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
