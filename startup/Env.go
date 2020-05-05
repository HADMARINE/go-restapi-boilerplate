package startup

import (
	"fmt"
	"go-restapi-boilerplate/util"
	"os"
)

// Env function sets env vars.
func Env() {
	value, err := util.ReadJSON(".env.json")
	if err != nil {
		fmt.Println("Error while reading env settings : ", err)
	}
	for key, value := range value {
		os.Setenv(key, value.(string))
	}
}
