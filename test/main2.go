// Author: LiuLin
// Date: 2022/4/14

package main

import (
	"fmt"
	"math"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var wg sync.WaitGroup
		for i := 0; i < math.MaxInt32; i++ {
			wg.Add(1)
			go func(i int) {
				//wg.Add(1)
				defer wg.Done()
				fmt.Println(i)
				time.Sleep(time.Second)
			}(i)
		}
		wg.Wait()
	})
	http.ListenAndServe(":8080", nil)
}
