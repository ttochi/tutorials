package main

import (
	"fmt"
	"net/http"
)

type resps struct {
	url string
	status string
}

func routineMain() {
	results := map[string]string{}

	urls := []string{
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.airbnb.com",
		"https://www.facebook.com",
		"https://www.reddit.com",
		"https://www.reddit.com/hahaha",
	}

	// Check url synchronous
	for _, url := range urls {
		result := hitURLSync(url)
		results[result.url] = result.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

	// Check url asynchronous with go routine
	results = map[string]string{}
	c := make(chan resps)
	
	for _, url := range urls {
		go hitURLAsync(url, c)
	}

	for i:=0; i<len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURLSync(url string) resps {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode > 400 {
		status = "FAILED"
	}
	return resps{url: url, status: status}
}

func hitURLAsync(url string, c chan<- resps) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode > 400 {
		status = "FAILED"
	}
	c <- resps{url: url, status: status}
}
