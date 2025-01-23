package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Generating url in golang")

	generatedUrl := &url.URL {
		Scheme: "https",
		Host:"www.example.com",
		Path:"/path/to/resource",
		RawQuery:"param2=value2",
	}
	generatedUrls := generatedUrl.String()
	fmt.Println("Url generated: ", generatedUrls)
}