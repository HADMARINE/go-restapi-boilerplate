package env

import (
	"go-restapi-boilerplate/src/util"
	"os"
)

// Init function sets env vars.
func Init() {
	for key, value := range util.ReadJSON(".env.json") {
		os.Setenv(key, value.(string))
	}
}
