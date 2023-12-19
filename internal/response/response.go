package response

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type ValidationError struct {
	Field  string `json:"field"`
	Errors string `json:"error"`
}
