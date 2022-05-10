// Author: LiuLin
// Date: 2022/5/10

package main

import "fmt"

func main() {
	in := product(10)
	ch := square(in)
	for v := range ch {
		fmt.Println(v)
	}
}

func product(nums int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < nums; i++ {
			out <- i
		}
	}()

	return out
}

func square(inch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range inch {
			out <- v * v
		}
	}()
	return out
}
