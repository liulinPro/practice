// Author: LiuLin
// Date: 2022/4/27

package main

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	<-ch
}
