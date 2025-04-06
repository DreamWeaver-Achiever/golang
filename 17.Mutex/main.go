package main 

import (
	"fmt"
	"sync"
	"net/http"
)

var mut sync.Mutex
var wg sync.WaitGroup
var signals = []string{"test"}

func main() {
	fmt.Println("Mutex in Golang")
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
	fmt.Printf("Signals after mutex: %s\n", signals)

}

func getStatusCode(website string) {
	defer wg.Done()
	
	resp, err := http.Get(website)
	if err != nil {
		fmt.Printf("Error while fetching response from website: %s", website)
	} else {
		mut.Lock()
		signals = append(signals, website)
		defer mut.Unlock()
		fmt.Printf("\nWebsite '%s' returned response: %d\n", website, resp.StatusCode)
		resp.Body.Close() // Close the response body to prevent resource leaks.
	}
}