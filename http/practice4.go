// Author: LiuLin
// Date: 2022/4/22

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	get, _ := http.Get("https://cn.bing.com/images/search?q=Pic&FORM=IQFRBA&id=B06828ED54DAE76827E7FC54FB6D0F0D6397699C")
	all, _ := ioutil.ReadAll(get.Body)

	fmt.Println(string(all))
}
