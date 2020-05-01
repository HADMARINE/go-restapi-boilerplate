package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-restapi-boilerplate/src/env"

	"go-restapi-boilerplate/src/world"
)

func main() {
	env.Init()
	router := gin.Default()
	router.GET("/ping", func(res *gin.Context) {
		res.JSON(http.StatusOK, gin.H{"message": "ponga"})
	})
	world.Hello("world")

	router.Run()
}
