package models

type Response struct {
	Success bool `json:"success"`
	Message    string `json:"message"`
	Error      error  `json:"error"`
	Score       any    `json:"score"`
}
