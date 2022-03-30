package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 3,
	}
}

func Get1(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, _ := http.NewRequest("Get", url, body)
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}

func main() {
	for {
		_, err := Get1("http://www.baidu.com/")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
