package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// StaticFile returns static file middleware
func StaticFile() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./public",false))
}