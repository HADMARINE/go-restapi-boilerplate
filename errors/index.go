package errors

import "net/http"

type ResponseError struct {
	Message string
	Status  int
	Code    string
}

func PageNotFound() ResponseError {
	return ResponseError{Message: "Page not found", Status: http.StatusBadRequest, Code: "PAGE_NOT_FOUND"}
}

func ParameterInvalid() ResponseError {
	return ResponseError{Message: "Parameter is invalid", Status: http.StatusBadRequest, Code: "PARAMETER_INVALID"}
}

func DatabaseSaveError() ResponseError {
	return ResponseError{Message: "Failed to save document to database", Status: http.StatusInternalServerError, Code: "DATABASE_SAVE_FAIL"}
}

func InternalLogicError() ResponseError {
	return ResponseError{Message: "Internal Server Error", Status: http.StatusInternalServerError, Code: "INTERNAL_LOGIC_ERROR"}
}
