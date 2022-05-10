// Author: LiuLin
// Date: 2022/5/10

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	task := make(chan bool, 1)
	go func() {
		time.Sleep(3 * time.Second)
		task <- true
	}()
	routine3(timeout, task)
	time.Sleep(5 * time.Second) //不加sleep 可能造成假象goroutine泄露，其实没有泄露
	fmt.Println(runtime.NumGoroutine())
}

func routine3(timeout context.Context, task chan bool) {
	for {
		select {
		case <-timeout.Done():
			fmt.Println("timeout")
			return
		case <-task:
			fmt.Println("done")
			return
		}
	}
}
