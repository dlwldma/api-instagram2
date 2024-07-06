package utils

type HttpResponse struct {
	Code    int
	Message string
	Payload any
}

func CreateNewResponse(code int, message string, payload any) *HttpResponse {
	return &HttpResponse{code, message, payload}
}
