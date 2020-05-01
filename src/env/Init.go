package env

import (
	"os"
)

// Init function sets env vars.
func Init() {
	envData := [][]string{{"PORT", "60020"}}
	for _, data := range envData {
		os.Setenv(data[0], data[1])
	}
}
