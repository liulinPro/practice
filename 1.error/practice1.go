// Author: LiuLin
// Date: 2022/3/25

package main

import (
	"errors"
	"fmt"
)

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func New(text string) error {
	return errorString(text)
}

var ErrNamedTYpe = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	//本质上是字符串相匹配
	if ErrNamedTYpe == New("EOF") {
		fmt.Println("name error type", ErrNamedTYpe)
	}
	// 栈内存地址比较
	if ErrStructType == errors.New("EOF") {
		fmt.Println("struct error type", ErrStructType)
	}
}
