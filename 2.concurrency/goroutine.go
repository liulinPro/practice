// Author: LiuLin
// Date: 2022/3/7

package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("sleep 10 second")
	}()
	fmt.Println("no sleep")
}
