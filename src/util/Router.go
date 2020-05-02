package util

import (
	"github.com/gin-gonic/gin"
)

type RouterContextType struct {
	Response customResponse
	Context  *gin.Context
}

type ResponseOptions struct {
	result         bool
	message        string
	code           string
	additionalData gin.H
}

type customResponse = func(status int, data gin.H)

// RouterContext provides customized context functions
func RouterContext(context *gin.Context) RouterContextType {
	customResponse := response(context)
	result := RouterContextType{Response: customResponse, Context: context}
	return result
}

func response(context *gin.Context) customResponse {
	return func(status int, response gin.H) {
		result := make(gin.H)

		if status == 0 {
			response["status"] = 200
		} else {
			response["status"] = status
		}

		if response["code"] == nil {
			response["code"] = DefaultCode(response["status"])
		}

		if response["message"] == nil {
			response["message"] = DefaultMessage(response["status"])
		}

		if response["result"] == nil {
			response["result"] = true
		}

		for key, value := range response {
			result[key] = value
		}

		context.JSON(status, result)
	}
}
