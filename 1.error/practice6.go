// Author: LiuLin
// Date: 2022/3/30

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
	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		line++
	}
	fmt.Println(line, scanner.Err())
}
