// Author: LiuLin
// Date: 2022/5/11

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	for x := range ch {
		fmt.Printf("Process %d\n", x)
	}
}
