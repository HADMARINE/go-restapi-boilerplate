package startup

import (
	"fmt"
	"os"
	"strings"

	"github.com/zebresel-com/mongodm"
)

type dbEnvType struct {
	dbhost string
	dbname string
	dbuser string
	dbpass string
}

// ConnectDB connects database
func ConnectDB() {
	dbEnv := checkEnv()
	dbConfig := &mongodm.Config{
		DatabaseHosts:    []string{dbEnv.dbhost},
		DatabaseName:     dbEnv.dbname,
		DatabaseUser:     dbEnv.dbuser,
		DatabasePassword: dbEnv.dbpass,
		DatabaseSource:   "admin",
	}

	_, connerr := mongodm.Connect(dbConfig)
	if connerr != nil {
		fmt.Println("Database connection error :", connerr)
		os.Exit(1)
	}
}

func checkEnv() dbEnvType {
	dbhost := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")

	if len(strings.TrimSpace(dbhost)) == 0 || len(strings.TrimSpace(dbname)) == 0 || len(strings.TrimSpace(dbuser)) == 0 || len(strings.TrimSpace(dbpass)) == 0 {
		fmt.Println("ENV Settings for connecting database is not provided, stopping...")
		os.Exit(1)
	}

	return dbEnvType{
		dbhost: dbhost,
		dbname: dbname,
		dbuser: dbuser,
		dbpass: dbpass,
	}
}
