// Author: LiuLin
// Date: 2022/4/25

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	conns := make([]Conn, 0)
	conns = append(conns, Conn{ID: 1}, Conn{ID: 2}, Conn{ID: 3})
	query := Query(conns)
	fmt.Println(query.ID)
	fmt.Println(runtime.NumGoroutine())
}

type Conn struct {
	ID int
}

func (c Conn) DoQuery(k int) Result {
	time.Sleep(time.Duration(k) * time.Second)
	return Result{ID: k}
}

type Result struct {
	ID int
}

func Query(conns []Conn) Result {
	ch := make(chan Result, len(conns)) //
	//ch := make(chan Result) //下面default注释就会内存泄露
	for k, conn := range conns {
		go func(c Conn) {
			select {
			case ch <- c.DoQuery(k):
				//default: //不加default会导致goroutine泄露
			}
		}(conn)
	}
	return <-ch
}
