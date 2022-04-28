// Author: LiuLin
// Date: 2022/4/20

package main

import "sync"

var wait sync.WaitGroup
var Count = 0

func main() {

	for i := 0; i < 2; i++ {

	}
}

func Routine(id int) {
	for i := 0; i < 2; i++ {
		value := Count
		value++
		Count = value
	}
	wait.Done()
}
