package structs

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}
