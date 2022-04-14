// Author: LiuLin
// Date: 2022/3/31

package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func ReadFile(path string) error {
	f, err := os.Open(path)
	// fmt.Printf("%T\n", err)
	defer f.Close()
	return errors.Wrap(err, "open file")
}

func warpError() {
	err := ReadFile("/saf")
	if err != nil {
		//错误类型&根
		fmt.Printf("original error: %T\n  %v\n", errors.Cause(err), errors.Cause(err))
		// 堆栈信息
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

func newErrors() error {
	return errors.New("message")
}

func main() {
	//err := lerrors.New("liulin")
	//fmt.Println(errors.Cause(err))
	fmt.Printf("%+v", newErrors())
}
