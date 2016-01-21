package go_http_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(url string, b []byte) ([]byte,error) {	
	req, err3 := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err3!=nil {
		return nil,err3
	}
	
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err!=nil {
		return nil,err
	}
	
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2!=nil {
		return nil,err2
	}
	
	defer resp.Body.Close()
	return body,nil
}

func PostStruct(url string, i interface{}) ([]byte,error) {
	iStr, err1 := json.Marshal(i)
	if err1!=nil {
		return nil,err1
	}
	
	req, err2 := http.NewRequest("POST", url, bytes.NewBuffer(iStr))
	if err2!=nil {
		return nil,err2
	}
	
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err3 := client.Do(req)
	defer resp.Body.Close()
	if err3!=nil {
		return nil,err3
	}
	
	body, err4 := ioutil.ReadAll(resp.Body)
	if err4!=nil {
		return nil,err4
	}

	return body,nil
}

func PostInterface(url string, b []byte, v interface{}) ([]byte,error) {
	
	body,err:=Post(url,b)
	if err!=nil {
		return nil,err
	} else {
		return body,json.Unmarshal([]byte(body), &v)
	}
}

func PostStructInterface(url string ,i interface{}, o interface{}) ([]byte,error) {
	body,err:=PostStruct(url,i)
	if err!=nil {
		return nil,err
	} else {
		return body,json.Unmarshal([]byte(body), &o)
	}
}

func Get(url string, timeOutInSec int) ([]byte,error) {
	timeout := time.Duration(timeOutInSec * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, err1 := client.Get(url)
	if resp==nil || resp.Body==nil {
		return nil,err1
	}
	defer resp.Body.Close()
	
	if err1!=nil {
		return nil,errors.New("Get: Can't connect to " + url)
	}
	
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2!=nil {
		return nil,errors.New("Get: Error result 2 from " + url)
	}

	return body,nil
}

func GetInterface(url string, v interface{}) ([]byte,error) {
	body,err:=Get(url)
	if err!=nil {
		return nil,err
	} else {
		return body,json.Unmarshal([]byte(body), &v)
	}
}