package http

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
	"encoding/json"
)

type authorizationTransport struct {
	client  HTTPClientI
	baseURL string
}

func NewAuthorizationTransport(client HTTPClientI, baseURL string) repository.AuthorizationRepository {
	return &authorizationTransport{
		client:  client,
		baseURL: baseURL,
	}
}

func (t *authorizationTransport) GetAccessToken(ctx context.Context, credentials entities.AccountCredentials) (*entities.Token, error) {
	resp, err := t.client.GetAccessToken(ctx, credentials, "/token")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Парсинг ответа
	var response entities.Token
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
