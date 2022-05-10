package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
	}()
	ch <- 1
	ch <- 1
	//for i := 0; i < 10; i++ {
	//	ch <- i

	//}
}
