package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"http://google.com", "http://facebook.com", "http://golang.org", "http://amazon.com"}

	c := make(chan string)

	for _, url := range urls {
		go checkURL(url, c)
	}

	for u := range c {
		go func(url string) {
			time.Sleep(time.Second)
			go checkURL(url)
		}(u)
	}

}

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "error: ", err)
		c <- url
		return
	}
	fmt.Println(url, ":", resp.Status)
	c <- url
}
