package model

// UserInfoResponse - ответ API с информацией об авторизованном пользователе.
type UserInfoResponse struct {
	Email      string   `json:"email"`
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	Phone      string   `json:"phone"`
	Phones     []string `json:"phones"`
	ProfileURL string   `json:"profile_url"`
}
