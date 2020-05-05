package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go-restapi-boilerplate/routes"
	"go-restapi-boilerplate/startup"
	"go-restapi-boilerplate/util"
)

func main() {
	startup.Env()
	startup.ConnectDB()

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s]: %s %s - %d(%s) [%s %s %s]",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Request.UserAgent(),
			param.Request.Proto)
	}))

	result, err := util.HashPassword("helloworld")
	if err == nil {
		fmt.Println(result)
	}
	fmt.Println(util.VerifyPassword("helloworld", result))
	fmt.Println(util.VerifyPassword("heLloworld", result))

	routes.Route(router)
	router.Run()
}
