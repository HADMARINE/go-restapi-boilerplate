package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go-restapi-boilerplate/src/env"
	"go-restapi-boilerplate/src/routes"
	"go-restapi-boilerplate/src/world"
)

func main() {
	env.Init()
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

	routes.Route(router)
	// router.GET("/", func(con *gin.Context) {
	// 	con.JSON(http.StatusOK, gin.H{"result": true, "message": "ALIVE", "status": 200})
	// })
	// router.GET("/ping", func(con *gin.Context) {
	// 	con.JSON(http.StatusOK, gin.H{"message": "pong"})
	// })
	world.Hello("world")

	router.Run()
}