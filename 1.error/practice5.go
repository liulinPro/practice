// Author: LiuLin
// Date: 2022/3/29

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	open, err := os.Open("/Users/liulin/practice/1.error/practice3.go")
	if err != nil {

	}
	var line int
	reader := bufio.NewReader(open)
	for {
		readString, err := reader.ReadString('\n')
		line++
		fmt.Println(readString)
		if err != nil {
			break
		}
	}
	fmt.Println(line)
}
