package startup

import (
	"go-restapi-boilerplate/util"
	"os"
)

// Env function sets env vars.
func Env() {
	for key, value := range util.ReadJSON(".env.json") {
		os.Setenv(key, value.(string))
	}
}
