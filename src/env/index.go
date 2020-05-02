package env

import (
	"os"
)

// Init function sets env vars.
func Init() {
	for key, value := range Data() {
		os.Setenv(key, value)
	}
}
