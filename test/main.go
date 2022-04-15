// Author: LiuLin
// Date: 2022/4/14

package main

import "fmt"

func main() {
	var i = 0
Loop:
	for {
		i++
		if i == 3 {
			break Loop
		}
	}
	fmt.Println("liu")
}
