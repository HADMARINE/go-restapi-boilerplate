package routes

import (
	"go-restapi-boilerplate/util"

	"github.com/gin-gonic/gin"
)

// Route is a default router
func Route(router *gin.Engine) {
	router.GET("/", func(con *gin.Context) {
		getIndex(util.RouterContext(con))
	})

}

func getIndex(data util.RouterContextType) {
	data.Response(200, gin.H{
		"data":    gin.H{"Hello": "world"},
		"message": "Golang API server is alive!",
	})
}
