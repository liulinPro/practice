package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup 一个异步结构体
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

// NewPool 生成一个工作池, coreNum 限制
func NewPool(coreNum int) *WaitGroup {
	ch := make(chan int, coreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

// Add 添加
func (ap *WaitGroup) Add(num int) {
	for i := 0; i < num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

// Done 完结
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
	}
	ap.wg.Done()
}

// Wait 等待
func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}

func main() {
	work := NewPool(4)
	for i := 0; i < 50; i++ {
		work.Add(1)
		test(i, work)
	}
}

func test(i int, wg *WaitGroup) {
	defer wg.Done()
	fmt.Println(time.Now(), "", i)
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now(), "", i)
}
