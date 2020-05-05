package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSON reads JSON file and returns the value
func ReadJSON(filename string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return ReadJSONByte(byteValue)
}

// ReadJSONByte reads json value by byte
func ReadJSONByte(bytes []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(bytes), &result)
	return result, err
}
