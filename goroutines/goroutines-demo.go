package main

import (
	"fmt"
	"net/http"
)

// goRoutines is a function that demonstrates the usage of goroutines in Go.
// It retrieves a list of links using the getLinks function and checks the status of each link concurrently by creating a goroutines.
func goRoutines(links []string) {

	for _, link := range links {
		// a go routine will be created to process the requests concurrently.
		// GoRuntime will create this GoRoutine
		// checkStatus func will run inside the created goroutine.
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

// A GoRoutine will be created automatically when we start compiling any golang program

// What is a GoRoutine ?
// exists in running program / the process
// GoRoutine executes a program line by line
// if encounters a blocking function call - goroutines will stop
// a new one can br created
