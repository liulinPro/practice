// Author: LiuLin
// Date: 2022/4/27

package main

import (
	"fmt"
	"sync"
)

type Config struct {
	a []int
}

func main() {
	cfg := &Config{}
	go func() {
		i := 0
		cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4}
	}()
	wg := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				fmt.Printf("%v\n", cfg)
			}
			wg.Done()
		}()
	}
}
