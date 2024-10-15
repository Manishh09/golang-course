package main

import (
	"fmt"
	"net/http"
	"time"
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
		go checkRepeatGoRoutines(link, ch) // pass channel
	}
	// Main GORoutine will wait till the channel receives the value, so its a blocking line of code // waits for the resp from child goroutines

	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)

	//we need wait for all the goroutines,  since 5 go routines were created

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-ch) // wait for channel to receive data and resumes the loop execution
	// }

	// check repeated go routines to check the status of each
	// this for-loop runs continuously
	// for {
	// 	go checkRepeatGoRoutines(<-ch, ch) // wait for channel to receive data and resumes the loop execution
	// }

	// wait for ch till has its value
	// run body of the for loop
	// alternative to above for-loop
	// more concise

	/*
		for l := range ch {

			go checkRepeatGoRoutines(l, ch) // wait for channel to receive data and resumes the loop execution

		}
	*/

	// add sleep to each go routines

	for l := range ch {

		go func(link string) { // function literal syntax
			time.Sleep(5 * time.Second)
			checkRepeatGoRoutines(link, ch) // wait for channel to receive data and resumes the loop execution
		}(l)

	}

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

func checkRepeatGoRoutines(link string, ch chan string) {
	// time.Sleep(5 * time.Second) ;// not recommended

	_, err := http.Get(link) // blocking call - code takes lil time to execute --> Main GoRoutine stops here

	if err != nil {
		fmt.Println(link, "could be down!")
		ch <- link // keep message in channel
		return
	}

	fmt.Println(link, "is up !!")

	ch <- link //Child GoRoutines gets the value from server and  keep message in channel // once this gets executed, then the  MainGORoutine will be notified that something is inside the channel that is in the MainGORoutine , so in this way, MaiGoRoutine will not exit the program
}
