package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Token - структура токена доступа OAuth2 для Avito API.
 type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetToken - выполняет Client Credentials OAuth2-флоу и возвращает токен доступа.
 func GetToken(clientID, clientSecret string) (token *Token, err error) {
	ctx := context.Background()

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/token", baseURL),
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		err = fmt.Errorf("could not create request: %w", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{
		Timeout: _defaultTimeoutToken,
	}

	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("could not do request: %w", err)
		return
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("invalid status code: %d", resp.StatusCode)
		return
	}

	var t Token
	if decErr := json.NewDecoder(resp.Body).Decode(&t); decErr != nil {
		err = fmt.Errorf("could not decode response: %w", decErr)
		return
	}

	token = &t
	return
}

// TODO: метод проверки токена перед запросом (проверка срока действия и т.п.)
func (t *Token) validation() error {
	return nil
}
