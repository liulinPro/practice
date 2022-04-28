// Author: LiuLin
// Date: 2022/4/18

package main

import (
	"log"
	"net/http"
)

func Hee(s chan<- string) {
	s <- "liu"
}

func main() {
	ints := make(chan<- string, 1)
	Hee(ints)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
