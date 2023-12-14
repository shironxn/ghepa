package port

import (
	"net/http"
)

type Response interface {
	Success(w http.ResponseWriter, status int, message string, data interface{})
	Error(w http.ResponseWriter, status int, message string, err interface{})
}
