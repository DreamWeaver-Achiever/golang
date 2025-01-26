package main

import (
        "encoding/json"
        "fmt"
)

type Person struct {
        Name   string `json:"full_name"`
        Age    int
        City   string
        Married bool
}

func main() {
        // Create a Person object
        person := Person{
                Name:   "John Doe",
                Age:    30,
                City:   "New York",
                Married: true,
        }

        // Encode to JSON
        jsonData, err := json.Marshal(person)
        CheckNilError(err, "Error during encoding")

        // Print JSON string
        fmt.Println("JSON:", string(jsonData))

        // Decode from JSON
        var decodedPerson Person
        err = json.Unmarshal(jsonData, &decodedPerson)
        CheckNilError(err, "Error during decoding")

        // Print decoded data
        fmt.Println("Decoded Person:", decodedPerson)
}

func CheckNilError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		return
	}
}