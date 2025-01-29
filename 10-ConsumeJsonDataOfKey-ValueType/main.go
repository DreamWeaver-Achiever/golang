package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Consume json data of key-value type")

	jsonData := []byte(`{
	"name": "John Doe",
	"age": 42, 
	"City": "New York"
	}`)

	//Check if json is valid
	isJsonValid := json.Valid(jsonData)

	//Decode json data
	var decodedJson map[string]interface{}
	if isJsonValid {
		err := json.Unmarshal(jsonData, &decodedJson)
		checkNilError(err, "Error while unmarshalling data.")
	} else {
		fmt.Println("Json is not valid.")
		return
	}
	for k, v := range decodedJson {
		fmt.Printf("\nJson Key: %v, Json Value: %v, Json value Type: %T", k, v, v);
	}
}

func checkNilError(err error, errMsg string) {
	if err != nil {
		fmt.Println(errMsg)
		return
	}
}