// Author: LiuLin
// Date: 2022/4/24

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var count int

func main() {
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go routine(i)
	}
	wg.Wait()
	fmt.Println(count) // data race 2||4
}

func routine(i int) {
	for i := 1; i <= 2; i++ {
		//value := count
		//value++
		//count = value
		count = count + 1 //并不安全 三条指令 使用互斥锁或者原子自增
	}
	wg.Done()
}
