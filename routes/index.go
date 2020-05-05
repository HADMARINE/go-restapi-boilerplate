package routes

import (
	v1 "go-restapi-boilerplate/routes/v1"
	"go-restapi-boilerplate/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is a default router constructor
func Route(router *gin.Engine) {
	router.GET("/", func(con *gin.Context) {
		getIndex(util.RouterContext(con))
	})
	v1Router := router.Group("/v1")
	{
		v1.Route(v1Router)
	}

}

func getIndex(data util.RouterContextType) {
	data.Response(http.StatusOK, gin.H{
		"data":    gin.H{"Hello": "world"},
		"message": "Golang API server is alive!",
	})
}
