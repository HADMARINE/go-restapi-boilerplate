package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadJSON reads JSON file and returns the value
func ReadJSON(filename string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
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
