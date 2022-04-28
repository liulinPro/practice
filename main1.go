// Author: LiuLin
// Date: 2022/4/27

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		ch <- i
		go func(v int) {
			//ch <- i
			time.Sleep(1 * time.Second)
			fmt.Println(v, time.Now().Unix())
			<-ch
		}(i)
	}
	select {}
}
