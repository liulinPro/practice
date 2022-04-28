// Author: LiuLin
// Date: 2022/4/25

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

func main() {
	m := make(map[string]interface{})
	m["test"] = "test"
	internalCallback(m)
	time.Sleep(3 * time.Second)
}

func internalCallback(data map[string]interface{}) {
	go func() {
		by, err := json.Marshal(data)
		if err != nil {
			return
		}
		_, _ = http.Post("https://winch.weipaitang.com/api/message/jd", "application/json; charset=utf-8", bytes.NewBuffer(by))
	}()
}
