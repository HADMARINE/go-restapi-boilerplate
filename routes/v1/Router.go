package v1

import (
	"go-restapi-boilerplate/routes/v1/user"
	"go-restapi-boilerplate/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Route is a default router constructor
func Route(router *gin.RouterGroup) {
	router.Any("/", func(con *gin.Context) {
		anyIndex(util.RouterContext(con))
	})

	userRouter := router.Group("/user")
	{
		user.Route(userRouter)
	}
}

func anyIndex(data util.RouterContextType) {
	data.Response(http.StatusOK, gin.H{
		"data": gin.H{
			"status": "alive",
			"mode":   os.Getenv("GIN_MODE"),
		},
		"message": "Welcome to GOLANG V1 API",
	})
}
