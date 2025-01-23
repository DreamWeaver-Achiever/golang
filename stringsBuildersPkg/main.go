package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
    "strings"
)

const (
	exampleUrl = "https://www.example.com/"
	statusOk = 200
)

func main() {
	fmt.Println("strings.Builder package in Golang")

	response, err := http.Get(exampleUrl)
	checkNilError(err, "Error while calling url")
	defer response.Body.Close()

	responseStatusCode := response.StatusCode
	fmt.Println("Response status code: ", responseStatusCode)

	if responseStatusCode == statusOk {
		var responseString strings.Builder
		responseBody, err := ioutil.ReadAll(response.Body)
		checkNilError(err, "Error while reading response body.")
        byteCount, err := responseString.Write(responseBody)
        checkNilError(err, "Error while writing response string builder.")
		fmt.Println("Byte Count of the respone: ", byteCount)
		fmt.Println("Response string: ", responseString.String())
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