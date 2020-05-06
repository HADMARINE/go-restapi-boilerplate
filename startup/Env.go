package startup

import (
	"fmt"
	"go-restapi-boilerplate/util"
	"os"
	"strings"
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
	checkEnv()
}

func checkEnv() {
	envList := []string{"PORT", "GIN_MODE", "DB_HOST", "DB_NAME", "DB_USER", "DB_PASS", "REQUEST_URI"}
	var flag bool = false
	for _, env := range envList {
		if len(strings.TrimSpace(os.Getenv(env))) == 0 {
			fmt.Println(fmt.Sprintf("ENV setting \"%s\" is missing", env))
			flag = true
		}
	}
	if flag == true {
		panic("Stopping due to env settings...")
	}
}
