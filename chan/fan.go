// Author: LiuLin
// Date: 2022/5/10

package main

import "sync"

func main() {
	in := product(10)

	// FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	// consumer
	for _ = range merge(c1, c2, c3) {
	}
}
func mergeCh(cs ...<-chan int) <-chan int {
	out := make(chan int)
	//out := make(chan int,100)
	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	// FAN-IN
	for _, c := range cs {
		go collect(c)
	}

	// 错误方式：直接等待是bug，死锁，因为merge写了out，main却没有读
	// wg.Wait()
	// close(out)

	// 正确方式
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
