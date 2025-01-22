package main

import (
	"fmt"
	"net/url"
)

const urlPath = "https://www.example.com/path/to/resource?param1=value1"

func main() {
	fmt.Println("Handling Url In Golang")

	exampleUrl, err := url.Parse(urlPath)
	if err != nil {
		panic(err)
	}
	
	urlScheme := exampleUrl.Scheme
	fmt.Println("Url Scheme: ", urlScheme)
	
	host := exampleUrl.Hostname()
	fmt.Println("Host: ", host)

	port := exampleUrl.Port()
	fmt.Println("Port: ", port)

	path := exampleUrl.Path 
	fmt.Println("url Path: ", path)

	rawQuery := exampleUrl.RawQuery
	fmt.Println("rawQuery: ", rawQuery)

	queryParam := exampleUrl.Query()
	fmt.Println("Query parameter: ", queryParam)

	query := exampleUrl.Query()
	query.Set("param2", "value2")
	exampleUrl.RawQuery = query.Encode()
	fmt.Println("Modified url: ", exampleUrl.String())
	
	originalString := "Hello, World! & ? /"
	escapedString := url.QueryEscape(originalString)
	fmt.Println("Escaped String:", escapedString) // Output: Hello%2C+World%21+%26+%3F+%2F
	
	unescapedString, err := url.QueryUnescape(escapedString)
	if err != nil {
		fmt.Println("Error unescaping:", err) 
		} else { 
			fmt.Println("Unescaped String:", unescapedString) // Output: Hello, World! & ? /
			}

}