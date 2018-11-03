package main

import (
	"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const (
	url = "https://v1b.es"
)

type ScanEntry struct {
	id   int    `json:"id"`
	url string `json:"url"`
	date string `json:"date`
	code int `json:"code`
	rptime int `json:"rptime`
}

func main(){

	t0 := time.Now()
	rc := httpGetResponseCode(url)
	t1 := time.Now()
	rptime := t1.Sub(t0)

	fmt.Println( url, " took ", rptime)	
	fmt.Println(rc)	

	b, err := MysqlAddScanEntry(url, strconv.Itoa(rc), strconv.FormatFloat(rptime.Seconds(), 'f', 2, 64))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(b)
}

func httpGetResponseCode(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return 500	
	}
	defer resp.Body.Close()
	return resp.StatusCode
}

func getNewRequestEntryId() int{

	id := 3
	return id
}