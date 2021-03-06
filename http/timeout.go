package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

var tr *http.Transport

//var client *http.Client

func init() {
	tr = &http.Transport{
		MaxIdleConns: 100,
		// 下面的代码被干掉了
		//Dial: func(netw, addr string) (net.Conn, error) {
		//    conn, err := net.DialTimeout(netw, addr, time.Second*2) //设置建立连接超时
		//    if err != nil {
		//        return nil, err
		//    }
		//    err = conn.SetDeadline(time.Now().Add(time.Second * 3)) //设置发送接受数据超时
		//    if err != nil {
		//        return nil, err
		//    }
		//    return conn, nil
		//},
	}
	//client = &http.Client{
	//	Transport: tr,
	//	Timeout:   time.Second * 3,
	//}
}

func Get(url string) ([]byte, error) {
	m := make(map[string]interface{})
	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	req, _ := http.NewRequest("Get", url, body)
	req.Header.Add("content-type", "application/json")

	client := &http.Client{Transport: tr, Timeout: time.Second * 3}
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
