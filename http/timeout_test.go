// Author: LiuLin
// Date: 2022/3/28

package main

import "testing"

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("http://www.baidu.com/")
	}
}

func BenchmarkGet1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get1("http://www.baidu.com/")
	}
}
