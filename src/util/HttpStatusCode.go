package util

import (
	"strings"
)

// HTTPStatusCode contains the HTTP status code data
func HTTPStatusCode() map[interface{}]string {
	return map[interface{}]string{
		100: "Continue",
		101: "Switching Protocol",
		200: "OK",
		201: "Created",
		202: "Accepted",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		301: "Moved Permanently",
		302: "Found",
		304: "Not Modified",
		400: "Bad Request",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		409: "Conflict",
		418: "I'm a teapot",
		422: "Unprocessable Entity",
		429: "Too Many Requests",
		500: "Internal Server Error",
		503: "Service Unavailable",
	}
}

// DefaultCode returns default code
func DefaultCode(status interface{}) string {
	return strings.ToUpper(strings.ReplaceAll(HTTPStatusCode()[status], " ", "_"))
}

// DefaultMessage return default message
func DefaultMessage(status interface{}) string {
	return HTTPStatusCode()[status]
}
