package error

type ResponseError struct {
	message string
	status  int
	code    string
}

func PageNotFound() ResponseError {
	return ResponseError{message: "Page not found", status: 404, code: "PAGE_NOT_FOUND"}
}
