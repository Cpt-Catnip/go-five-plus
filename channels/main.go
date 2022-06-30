package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Printf("%s might be down!\n", l)
		c <- l
		return
	}

	fmt.Printf("%s is up!\n", l)
	c <- l
}
