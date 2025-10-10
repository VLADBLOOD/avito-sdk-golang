package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/VLADBLOOD/avito-sdk-golang/model"
)

func (a *Authorization) isTokenFresh() bool {
	if a.token == nil {
		return false
	}
	return time.Now().Before(a.token.ExpiresAt)
}

func (a *Authorization) isTokenExpired() bool {
	if a.token == nil {
		return false
	}
	return !a.isTokenFresh()
}

type IAuthorization interface {
	// Возвращает токен для конкретных Credentials
	GetToken() (*model.Token, error)
	SetCredentials(creds *model.Credentials) error
	fetchToken() error
}

// Authorization управляет авторизацией в Avito API через OAuth2 client_credentials.
type Authorization struct {
	creds *model.Credentials
	token *model.Token
}

// NewHttpClient - конструктор HTTP-клиента с установленным таймаутом и токеном авторизации.
func NewAuthorization() *Authorization {
	return &Authorization{
		creds: nil,
		token: nil,
	}
}

func (a *Authorization) SetCredentials(creds *model.Credentials) error {
	a.creds = creds

	// Обновить токен после установки кредсов
	err := a.fetchToken()
	if err != nil {
		return fmt.Errorf("error at fetching token by creds: %w", err)
	}

	return nil
}

// Возвращает значение своего атрибута token
func (a *Authorization) GetToken() (*model.Token, error) {
	if a.creds == nil {
		return nil, fmt.Errorf("CREDENTIALS IS NOT INITIALIZED")
	}

	// Проверка актуальности токена перед передачей
	if a.isTokenExpired() {
		err := a.fetchToken()
		if err != nil {
			return nil, fmt.Errorf("failed to refresh token: %w", err)
		}
	}

	return a.token, nil
}

// GetToken - выполняет OAuth2 и возвращает токен доступа.
func (a *Authorization) fetchToken() error {
	ctx := context.Background()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", a.creds.ClientID)
	data.Set("client_secret", a.creds.ClientSecret)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/token", baseURL),
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		err = fmt.Errorf("could not create request: %w", err)
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: _defaultTimeoutToken,
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("could not do request: %w", err)
		return err
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
		return err
	}

	var t model.Token
	if decErr := json.NewDecoder(resp.Body).Decode(&t); decErr != nil {
		err = fmt.Errorf("could not decode response: %w", decErr)
		return err
	}

	// Устанавливаем время истечения при создании токена
	t.ExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)

	// Устанавливаем для атрибута структуры - значение токена
	a.token = &t

	return nil
}
