package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/VLADBLOOD/avito-sdk-golang/model"
)

// User - сервис работы с данными авторизованного пользователя Avito.
type User struct {
	client IHttpClient
}

// NewUser - конструктор сервиса пользователя.
func NewUser(client IHttpClient) *User {
	return &User{client: client}
}

// GetUserInfoSelf - получение информации об авторизованном пользователе.
func (u *User) GetUserInfoSelf(ctx context.Context) (*model.UserInfoResponse, error) {
	path := "/core/v1/accounts/self"
	response := new(model.UserInfoResponse)

	status, err := u.client.request(ctx, http.MethodGet, path, nil, response)
	if err != nil {
		return response, fmt.Errorf("get user info self: запрос завершился ошибкой: %w", err)
	}

	if status != http.StatusOK {
		return response, fmt.Errorf("get user info self: сервер вернул код %d", status)
	}

	return response, nil
}
