package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Token - структура токена доступа OAuth2 для Avito API.
type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"-"` // Время истечения срока действия токена. Не сериализуется в JSON.
}

type Credentials struct {
	ClientID     string // Идентификатор клиента
	ClientSecret string // Секрет клиента для получения токена
}

// GetToken - выполняет Client Credentials OAuth2-флоу и возвращает токен доступа.
func GetToken(creds *Credentials) (token *Token, err error) {
	ctx := context.Background()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", creds.ClientID)
	data.Set("client_secret", creds.ClientSecret)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/token", baseURL),
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		err = fmt.Errorf("could not create request: %w", err)
		return token, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: _defaultTimeoutToken,
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("could not do request: %w", err)
		return token, err
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
		return token, err
	}

	var t Token
	if decErr := json.NewDecoder(resp.Body).Decode(&t); decErr != nil {
		err = fmt.Errorf("could not decode response: %w", decErr)
		return token, err
	}

	// Устанавливаем время истечения при создании токена
	t.ExpiresAt = time.Now().Add(time.Duration(t.ExpiresIn) * time.Second)

	token = &t
	return token, err
}

func (t *Token) IsFresh() bool {
	if t == nil {
		return false
	}
	return time.Now().Before(t.ExpiresAt)
}

func (t *Token) IsExpired() bool {
	if t == nil {
		return false
	}
	return !t.IsFresh()
}
