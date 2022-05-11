// Author: LiuLin
// Date: 2022/5/11

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	time.After(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
}
