package main

import (
	"fmt"
	"net/url"
	"net/http"
)

const (
	exampleUrl = "https://httpbin.org/post"
)
func main() {
	fmt.Println("PostForm Request In Golang")

	data := url.Values{}
	data.Set("Name", "John Doe")
	data.Set("Age", "42")

	response, err := http.PostForm(exampleUrl, data)
	CheckNilError(err, "Error while making postForm request")
	defer response.Body.Close()
	fmt.Println("Response: ", response)
}

func CheckNilError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		return
	}
}