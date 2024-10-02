// models/response.go
package models

// ErrorResponse представляет структуру для ошибок API
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse представляет структуру для сообщений API
type MessageResponse struct {
	Message string `json:"message"`
}
