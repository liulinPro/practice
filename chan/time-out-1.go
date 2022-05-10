// Author: LiuLin
// Date: 2022/5/10

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(4 * time.Second)
		timeout <- true
	}()

	task := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		task <- true
	}()
	routine(timeout, task)
	time.Sleep(5 * time.Second) //不加sleep 可能造成假象goroutine泄露，其实没有泄露
	fmt.Println(runtime.NumGoroutine())
}

func routine(timeout chan bool, task chan bool) {
	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			return
		case <-task:
			fmt.Println("done")
			return
		}
	}
}
