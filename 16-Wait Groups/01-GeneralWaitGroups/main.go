package main

import (
	"fmt"
	"sync"
	"net/http"

)

var wg sync.WaitGroup

func main() {
	fmt.Println("Wait groups in golang")

	websiteList := []string {
		"https://go.dev/", 
		"https://google.com/", 
		"https://github.com/", 
	}
	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}
	wg.Wait()
}

func getStatusCode(website string) {
	defer wg.Done()
	
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("Error while fetching response from website: %s", website)
	} else {
		fmt.Printf("\nWebsite '%s' returned response: %d\n", website, resp.StatusCode)
		resp.Body.Close() // Close the response body to prevent resource leaks.
	}
}