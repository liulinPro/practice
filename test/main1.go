// Author: LiuLin
// Date: 2022/4/14

package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		go func(i int) {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}
