package util

import (
	"encoding/json"
	"ghepa/internal/core/domain"
	"net/http"
)

type Response struct {
}

func (r *Response) Success(w http.ResponseWriter, status int, message string, data interface{}) {
	res := domain.SuccessResponse{
		Message: message,
		Data:    data,
	}

	response, err := json.Marshal(res)
	if err != nil {
		r.Error(w, http.StatusInternalServerError, "Failed to marshal JSON response", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func (r *Response) Error(w http.ResponseWriter, status int, message string, err interface{}) {
	res := domain.ErrorResponse{
		Message: message,
		Errors:  err,
	}

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
