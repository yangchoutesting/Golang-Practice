package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://goole.com",
		"http://facebook.com",
		"http://stackoverflow.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for l := range channel {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, channel)
		}(l)

	}
}

func checkLink(link string, channel chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		channel <- link
		return
	}
	fmt.Println(link, "is up.")
	channel <- link
}
