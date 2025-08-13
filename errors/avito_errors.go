/*
Пакет с ошибками от API Avito
*/
package errors

// AvitoAPIError represents the base structure of Avito API errors.
type AvitoAPIError struct {
	Error ErrorDetails `json:"error"`
}

// ErrorDetails contains the specific error information.
type ErrorDetails struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BadRequestError represents a 400 Bad Request error from Avito API.
type BadRequestError struct {
	AvitoAPIError
}

// UnauthorizedError represents a 401 Unauthorized error from Avito API.
type UnauthorizedError struct {
	AvitoAPIError
}

// ForbiddenError represents a 403 Forbidden error from Avito API.
type ForbiddenError struct {
	AvitoAPIError
}

// RateLimitError represents a 429 Too Many Requests error from Avito API.
type RateLimitError struct {
	AvitoAPIError
}

// InternalServerError represents a 500 Internal Server Error from Avito API.
type InternalServerError struct {
	AvitoAPIError
}
