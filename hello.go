package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Checks whether a set of websites is up and running.
func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
		"https://www.swiss.com/",
		"https://www.swiss.com/",
		"https://www.swiss.com/",
		"https://www.swiss.com/",
		"https://www.swiss.com/",
	}

	var wg sync.WaitGroup
	// A `WaitGroup` waits for a collection of go routines to finish. 


	for _, u := range urls {

		// increment the wait group counter
		wg.Add(1)
		go func(url string) {
			// decrement the counter when the go routine completes
			defer wg.Done()
			// call the function check
			checkUrl(url)
		}(u)
		// go checkUrl(url)
	}
	// time.Sleep(5 * time.Second)
	// wait for the all the website calls to finish
	wg.Wait()
}



// checks and prints a message if the given website is up or down
func checkUrl(url string){
	_, err := http.Get(url)
	if err != nil{
		fmt.Println(url, "is down!!!")
		return 
	}
	fmt.Println(url, "is up and running.")
}
