// Author: LiuLin
// Date: 2022/4/19

package main

import (
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//var values url.Values
	values := make(url.Values)
	values.Add("sdafas", "")
	_, _ = http.Post("http://winch.wptqc.com/api/message/jd", "application/json; charset=utf-8", strings.NewReader(values.Encode()))

}
