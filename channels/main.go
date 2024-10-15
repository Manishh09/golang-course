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

	ch := make(chan string) // creating a channel of type string

	for _, link := range links {
		go checkStatus(link, ch) // pass channel
	}

	fmt.Println(<-ch) // Main GORoutine will wait till the channel receives the value, so its a blocking line of code // waits for the resp from child goroutines
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//we need wait for all the goroutines,  since 5 go routines were created

}

func checkStatus(link string, ch chan string) {
	_, err := http.Get(link) // blocking call - code takes lil time to execute --> Main GoRoutine stops here

	if err != nil {
		fmt.Println(link, "could be down!")
		ch <- "Link might be down !" // keep message in channel
		return
	}

	fmt.Println(link, "is up !!")

	ch <- "Link is up !!" //Child GoRoutines gets the value from server and  keep message in channel // once this gets executed, then the  MainGORoutine will be notified that something is inside the channel that is in the MainGORoutine , so in this way, MaiGoRoutine will not exit the program
}

// Note: Receiving the data into a channel referred as a Blocking Call
