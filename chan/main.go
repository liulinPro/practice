// Author: LiuLin
// Date: 2022/3/28

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
	}()

	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println(<-ch)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(3 * time.Second)
}
