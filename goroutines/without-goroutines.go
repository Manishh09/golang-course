package main

import (
	"fmt"
	"net/http"
)

func withoutGoRoutines(links []string) {

	for _, link := range links {
		checkLinkStatus(link)

	}
}

/*
checkStatus checks the status of a given link by making an HTTP GET request.
It prints a message indicating whether the link is up or down.
it executes one by one and prints the status
one link status will be printed after the other synchronously
waits for another link status before proceeding to make a request for another link
*/
func checkLinkStatus(link string) {
	_, err := http.Get(link) // blocking call - code takes lil time to execute

	if err != nil {
		fmt.Println(link, "could be down!")
		return
	}

	fmt.Println(link, "is up !!")
}
