// Author: LiuLin
// Date: 2022/3/28

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//ch := make(chan int)
	ch := make(chan int, 4)
	for i := 0; i < 5; i++ {
		go func(v int) {
			ch <- v
		}(i)
	}
	go func() {
		<-ch
	}()
	time.Sleep(5 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
