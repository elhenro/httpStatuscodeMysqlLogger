package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"io/ioutil"
	"os"
)

type WebsiteList struct {
	Websites struct {
		URL1 string `json:"URL1"`
		URL2 string `json:"URL2"`
		URL3 string `json:"URL3"`
	} `json:"websites"`
}	/*
type WebsiteList struct {
	Websites []struct {
		URL1 string `json:"URL1,omitempty"`
		URL2 string `json:"URL2,omitempty"`
		URL3 string `json:"URL3,omitempty"`
	} `json:"websites"`
}*/

type ScanEntry struct {
	id   int    `json:"id"`
	url string `json:"url"`
	date string `json:"date`
	code int `json:"code`
	rptime int `json:"rptime`
}

func main(){

	l := getWeblist()
	
	//fmt.Println(l.Websites.URL2)

	//os.Exit(0)
	fmt.Println(scanSite(l.Websites.URL1))
	fmt.Println(scanSite(l.Websites.URL2))
	fmt.Println(scanSite(l.Websites.URL3))


	
}

func scanSite(url string) bool{
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
	return b
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

func getWeblist() WebsiteList {
	jsonFile, err := os.Open("sites.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result WebsiteList
	json.Unmarshal([]byte(byteValue), &result)

	return result
}