package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSON reads JSON file and returns the value
func ReadJSON(filename string) map[string]interface{} {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	return result
}
