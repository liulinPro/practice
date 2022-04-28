// Author: LiuLin
// Date: 2022/4/18

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a byte = '0'
	fmt.Printf("a 的类型 %T a占中的字节数是 %d\n", a, unsafe.Sizeof(a))
	var b = &a
	fmt.Printf("b 的类型 %T b占中的字节数是 %d", b, unsafe.Sizeof(b))
}
