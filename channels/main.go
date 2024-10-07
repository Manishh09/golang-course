package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	for _, link := range links {
		go checkStatus(link)
	}

}

func checkStatus(link string) {
	_, err := http.Get(link) // blocking call - code takes lil time to execute --> Main GoRoutine stops here

	if err != nil {
		fmt.Println(link, "could be down!")
		return
	}

	fmt.Println(link, "is up !!")
}
