// Author: LiuLin
// Date: 2022/5/11

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}
	ctx, cancel := context.WithCancel(context.Background())
	for v := range gen(ctx) {
		fmt.Println(v)
		//break
		if v == 5 {
			break
		}
	}
	fmt.Println(runtime.NumGoroutine())
	cancel() // 不然会造成goroutine泄露
	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
