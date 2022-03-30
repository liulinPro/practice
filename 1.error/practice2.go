// Author: LiuLin
// Date: 2022/3/28

package main

import "fmt"

type errorString1 struct {
	s string
}

func (e errorString1) Error() string {
	return e.s
}

func New1(text string) errorString1 {
	return errorString1{s: text}
}

var structError = New1("EOF")

func main() {
	errType := New1("EOF")
	// 比较的是里面的值
	if errType == structError {
		fmt.Println(errType.Error())
	}
}
