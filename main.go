package main

import (
	"github.com/gin-gonic/gin"

	"go-restapi-boilerplate/errors"
	"go-restapi-boilerplate/middleware"
	"go-restapi-boilerplate/routes"
	"go-restapi-boilerplate/startup"
	"go-restapi-boilerplate/util"
)

func main() {
	startup.Env()
	startup.ConnectDB()

	router := gin.New()

	// Middlewares
	router.Use(gin.Recovery())
	router.Use(middleware.GinCustomLogger())
	router.Use(middleware.GinCors())
	router.Use(middleware.StaticFile())

	// Middleware functions
	middleware.Cron()

	router.NoRoute(func(con *gin.Context) {
		util.RouterContext(con).ErrorHandler(errors.PageNotFound())
	})

	// Routes
	routes.Route(router)

	router.Run()
}
