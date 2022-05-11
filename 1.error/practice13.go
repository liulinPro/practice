// Author: LiuLin
// Date: 2022/5/11

package main

import (
	"fmt"
	"golang.org/x/xerrors"
)

func main() {
	err := xerrors.New("liulin")
	fmt.Printf("%+v", err)
}
