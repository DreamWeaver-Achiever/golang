package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const (
	exampleUrl = "https://www.example.com/"
	statusOk = 200
)

func main() {
	fmt.Println("Get Request in Golang")

	response, err := http.Get(exampleUrl)
	checkNilError(err, "Error while calling url")
	defer response.Body.Close()

	responseStatusCode := response.StatusCode
	fmt.Println("Response status code: ", responseStatusCode)

	if responseStatusCode == statusOk {
		fmt.Println("Response of the get request: ", response)
		fmt.Println("Length of the response: ", response.ContentLength)
		responseBody, err := ioutil.ReadAll(response.Body)
		checkNilError(err, "Error while reading response body.")
		fmt.Println("Response Body in byte format: ", responseBody)
		fmt.Println("Response Body in string format: ", string(responseBody))
	} else {
		fmt.Println("Response Status Code is not ok. StatusCode received: ", responseStatusCode)
	}



}
func checkNilError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		return 
	}
}