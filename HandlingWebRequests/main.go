package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.example.com"

func main() {
	fmt.Println("Handling Web Requests in Golang")
	resp, err := http.Get(url)
	checkNilErr(err)
	fmt.Printf("\nType of response: %T\n", resp)

	defer resp.Body.Close() //callers responsiblity to close the request everytime.

	dataBytes, err := ioutil.ReadAll(resp.Body)
	checkNilErr(err)
	content := string(dataBytes)
	fmt.Printf("Contents received from web url/web request: %s", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}