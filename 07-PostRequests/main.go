package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
)

type MyData struct {
        Name  string `json:"name"`
        Value int    `json:"value"`
}

func main() {
        url := "http://httpbin.org/post" // Example endpoint (replace with your target)
        data := MyData{
                Name:  "John Doe",
                Value: 42,
        }

        // Encode data to JSON
        jsonData, err := json.Marshal(data)
        checkNilError(err, "Error encoding JSON response")

        // Create a request
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
        checkNilError(err, "Error creating request")

        // Set request headers
        req.Header.Set("Content-Type", "application/json")

        // Make the request
        client := &http.Client{}
        resp, err := client.Do(req)
        checkNilError(err, "Error making request")

        defer resp.Body.Close()

        // Read the response body
        body, err := io.ReadAll(resp.Body)
        checkNilError(err, "Error reading response body")

        fmt.Println("Response:", string(body))
}

func checkNilError(err error, errMsg string) {
        if err != nil {
                fmt.Println(errMsg)
                return
        }
}