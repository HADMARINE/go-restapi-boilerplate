package user

import (
	"go-restapi-boilerplate/db"
	"go-restapi-boilerplate/errors"
	"go-restapi-boilerplate/models"
	"go-restapi-boilerplate/util"
	"net/http"

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
		data.ErrorHandler(errors.ParameterInvalid())
		return
	}

	userid := body["userid"]
	password := body["password"]
	authority := "normal"

	if util.IsNilArray(userid, password) == true {
		data.ErrorHandler(errors.ParameterInvalid())
		return
	}

	encryptedPassword, hashErr := util.HashPassword(password.(string))

	if hashErr != nil {
		data.ErrorHandler(errors.InternalLogicError())
		return
	}

	User := db.Database.Model("User")
	user := &models.User{
		Userid:    userid.(string),
		Password:  encryptedPassword,
		Authority: authority,
	}

	User.New(user)

	docSaveErr := user.Save()
	if docSaveErr != nil {
		data.ErrorHandler(errors.DatabaseSaveError())
		return
	}
	data.Response(http.StatusOK, gin.H{})
}
