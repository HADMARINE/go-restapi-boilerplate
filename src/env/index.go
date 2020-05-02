package env

import (
	"os"
)

// Init function sets env vars.
func Init() {
	for key, value := range envData() {
		os.Setenv(key, value)
	}
}

func envData() map[string]string {
	return map[string]string{"PORT": "60020", "GIN_MODE": "debug"}
}
