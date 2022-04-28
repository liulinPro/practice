// Author: LiuLin
// Date: 2022/4/27

package main

import (
	"fmt"
	"time"
)

var m = make(map[int]int)

func main() {
	for i := 0; i < 10; i++ {
		go func(v int) {
			routine(v)
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println(m)
}

func routine(i int) {
	m1 := make(map[int]int)
	m1[i] = i
	m = m1
}
