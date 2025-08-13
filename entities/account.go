package entities

// Account - учетная запись Avito
type Account struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name,omitempty"`
	Credentials AccountCredentials `json:"credentials"`
	Token       *Token             `json:"-"`
}

type AccountCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// Token - OAuth2 токен доступа
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
