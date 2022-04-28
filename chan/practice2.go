// Author: LiuLin
// Date: 2022/4/18

package main

import (
	"fmt"
	"time"
)

func test(dir string) chan string {
	ch := make(chan string)
	go func() {
		ch <- "dir"
		close(ch)
		time.Sleep(1 * time.Second)
		ch <- "dir"
	}()
	return ch
}

func main() {
	ch := test("/")
	s := <-ch
	fmt.Println(s)
}
