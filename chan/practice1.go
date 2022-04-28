// Author: LiuLin
// Date: 2022/4/14

package main

import (
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

//https://geektutu.com/post/hpg-concurrency-control.html
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var wg sync.WaitGroup
		c := make(chan struct{}, 3)
		for i := 0; i < math.MaxInt64; i++ {
			wg.Add(1)
			c <- struct{}{} //能控制
			go func(i int) {
				//c <- struct{}{}
				defer wg.Done()
				log.Println(i)
				time.Sleep(time.Second)
				<-c
			}(i)

		}
		wg.Wait()
	})
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
	select {}
}
