package model

import "time"

type Credentials struct {
	ClientID     string // Идентификатор клиента
	ClientSecret string // Секрет клиента для получения токена
}

// Token - структура токена доступа OAuth2 для Avito API.
type Token struct {
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int       `json:"expires_in"`
	ExpiresAt   time.Time `json:"-"` // Время истечения срока действия токена. Не сериализуется в JSON.
}
