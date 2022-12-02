package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string {
	"https://coupang.com/",
	"https://www.google.com/",
	"https://www.naver.com/",
	"https://www.youtube.com/",
	"https://www.sonarpod.com/",
	"https://www.daum.net/",
	"https://www.gmarket.co.kr/",
	"https://medium.com/",
}

type urlChecked struct {
	url string
	result string
}

func main () {
	var results = make(map[string]string)
	// var results = make(map[string]string)
	var c = make(chan urlChecked)
	// var people = []string{"taewon", "zomb", "mark", "felix"};
	for _, url := range urls {
		go isValid(url, c)
	}
	for i:=0; i<len(urls); i++{
		result := <- c
		results[result.url] = result.result
	}
	for url, result := range results{
		fmt.Println(url, result)
	}
}

func isSexy (name string, c chan string) {
	time.Sleep(time.Second)
	c <- name + " is sexy."
}

func isValid (url string, c chan<- urlChecked) {
	res, err := http.Get(url)
	result := "ok"
	if(err != nil || res.StatusCode >= 400){
		result = "failed"
	}
	c <- urlChecked{
		url: url,
		result: result,
	}
}