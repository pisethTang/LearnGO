package main

import (
	"fmt"
	"net/http"
	// "sync"
)

// Reference: https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1

type urlStatus struct {
	url string
	status bool
}

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

	// var wg sync.WaitGroup
	// A `WaitGroup` waits for a collection of go routines to finish. 


	c := make(chan urlStatus)
	for _, u := range urls {
		go checkUrl(u, c)

		// increment the wait group counter
		// wg.Add(1)
		// go func(url string) {
		// 	// decrement the counter when the go routine completes
		// 	defer wg.Done()
		// 	// call the function check
		// 	checkUrl(url)
		// }(u)
		// go checkUrl(url)
	}
	// time.Sleep(5 * time.Second)
	// wait for the all the website calls to finish
	result := make([]urlStatus, len(urls))
	for i, _ := range result {
		result[i] = <- c
		if result[i].status {
			fmt.Println(result[i].url, "is up.")
		} else {
			fmt.Println(result[i].url, "is down !!")
		}
	}
	
}



// checks and prints a message if the given website is up or down
func checkUrl(url string, c chan urlStatus){
	_, err := http.Get(url)
	if err != nil{
		// fmt.Println(url, "is down!!!")
		// the website is down 
		c <- urlStatus{url, false}
		return 
	} else {
		// the website is up
		c <- urlStatus{url, true}
	}
	// fmt.Println(url, "is up and running.")
}
