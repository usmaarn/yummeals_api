package dto

type ApiResponse struct {
	Success    bool   `json:"success"`
	Data       any    `json:"data,omitempty"`
	Message    string `json:"message,omitempty"`
	StatusCode int    `json:"status_code"`
	Errors     any    `json:"errors,omitempty"`
}
