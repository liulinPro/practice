// Author: LiuLin
// Date: 2022/4/20

package main

import (
	"encoding/json"
	"fmt"
)

type Restaurant struct {
	NumberOfCustomers *int `json:",omitempty"`
}

func main() {
	d1 := Restaurant{}
	b, _ := json.Marshal(d1)
	fmt.Println(string(b))
	//Prints: {}

	n := 0
	d2 := Restaurant{
		NumberOfCustomers: &n,
	}
	b, _ = json.Marshal(d2)
	fmt.Println(string(b))
	//Prints: {"NumberOfCustomers":0}
}
