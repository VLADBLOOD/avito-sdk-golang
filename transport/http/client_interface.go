package http

import (
	"avito-sdk/entities"
	"context"
	"net/http"
)

// HTTPClientI - интерфейс для HTTP клиента
type HTTPClientI interface {
	// GetAccessToken выполняет запрос с заголовком x-www-form-urlencoded для получения access_token
	GetAccessToken(
		ctx context.Context,
		credentials entities.AccountCredentials,
		endpoint string,
	) (*http.Response, error)

	// Общий метод для HTTP запросов
	DoRequest(
		ctx context.Context,
		method string,
		account *entities.Account,
		endpoint string,
		requestData any,
		result any,
	) error

	// DoGET выполняет GET HTTP запрос с авторизацией из account
	DoGET(
		ctx context.Context,
		account *entities.Account,
		endpoint string,
		queryParams any,
	) (*http.Response, error)

	// DoPOST выполняет POST HTTP запрос с авторизацией из account
	DoPOST(
		ctx context.Context,
		account *entities.Account,
		endpoint string,
		body any,
	) (*http.Response, error)
}
