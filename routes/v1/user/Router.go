package user

import (
	"fmt"
	"go-restapi-boilerplate/db"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/util"

	"github.com/gin-gonic/gin"
)

// Route is a default router constructor
func Route(router *gin.RouterGroup) {
	router.POST("/", func(con *gin.Context) {
		postUser(util.RouterContext(con))
	})
}

func postUser(data util.RouterContextType) {
	body, bodyErr := data.ParseBody(data.Context.Request.Body)
	if bodyErr != nil {
		// data.Response(http.StatusBadRequest, {})
	}
	fmt.Println(body["hellow"])

	User := db.Database.Model("User")
	user := &models.User{Userid: "test"}
	User.New(user)
}
