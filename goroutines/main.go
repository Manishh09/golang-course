package main

// main is the entry point of the program.
// It retrieves a list of links using the getLinks function,
// and then spawns a goroutine to check the status of each link concurrently.
func main() { // MAIN GOROUTINE get created by default!

	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.github.com",
	}

	//withoutGoRoutines(links)

	goRoutines(links) // child go routines get created
}
