package main

import (
	"fmt"
	"time"
)

func job1(count int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for i := 0; i < count; i++ {
			time.Sleep(time.Second)
			fmt.Println("job1 finish:", 1)
			outCh <- 1
		}
	}()

	return outCh
}

func job2(inCh <-chan int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for val := range inCh {
			// 耗时2秒
			time.Sleep(time.Second * 2)
			val++
			fmt.Println("job2 finish:", val)
			outCh <- val
		}
	}()

	return outCh
}

func job3(inCh <-chan int) <-chan int {
	outCh := make(chan int, 2)

	go func() {
		defer close(outCh)
		for val := range inCh {
			val++
			fmt.Println("job3 finish:", val)
			outCh <- val
		}
	}()

	return outCh
}

func main() {
	t := time.Now()

	firstResult := job1(10)
	secondResult := job2(firstResult)
	thirdResult := job3(secondResult)

	for v := range thirdResult {
		fmt.Println(v)
	}

	fmt.Println("all finish")
	fmt.Println("duration:", time.Since(t).String())
}
