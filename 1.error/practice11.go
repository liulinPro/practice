// Author: LiuLin
// Date: 2022/4/14

package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("cxq")
	//error1 := fmt.Errorf("name not fount %v", err)
	error1 := fmt.Errorf("name not fount %w", err)
	if errors.Is(error1, err) {
		fmt.Println(error1.Error())
	}
}
