package main

import (
	"fmt"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for all the goroutines to finish

var wg sync.WaitGroup

var mtx sync.Mutex

var signals []string

func main() {

	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, web := range websites {
		go getStatusCode(web)

		// As it indicates there is one go routine in progress
		wg.Add(1)
	}

	// to wait for all the goroutines to finish. so we add it at the end of main function so it won't end before all the goroutines finish their execution
	wg.Wait()

	fmt.Println(signals)

}

func getStatusCode(url string) {

	// this will indicates all the goroutines are done
	defer wg.Done()

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	} else {

		// this will lock the mutex so that only one goroutine can access the mutex at a time and only one goroutine can unlock the mutex at a time. so it prevents the race condition

		mtx.Lock()

		signals = append(signals, url)

		mtx.Unlock()

		fmt.Printf("statusCode: %d from website %s\n", res.StatusCode, url)
	}
}
