// Author: LiuLin
// Date: 2022/5/10

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("close")
			}
			fmt.Println(v)
		}
	}
}
