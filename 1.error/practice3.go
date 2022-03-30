package main

import "fmt"

func main() {

	val1 := 234
	val2 := 567

	//创建和初始化指针
	var p1 *int
	p1 = &val1
	p2 := &val2
	p3 := &val1

	//使用==运算符比较指针
	res1 := &p1 == &p2
	fmt.Println("p1指针是否等于p2指针: ", res1)

	res2 := p1 == p2
	fmt.Println("p1指针是否等于p2指针: ", res2)

	res3 := p1 == p3
	fmt.Println("p1指针是否等于p3指针: ", res3)

	res4 := p2 == p3
	fmt.Println("p2指针是否等于p3指针: ", res4)

	res5 := &p3 == &p1
	fmt.Println("p3指针是否等于p1指针: ", res5)
}
