package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// IHttpClient - интерфейс HTTP-клиента для выполнения запросов к Avito API.
// Метод request принимает относительный путь (path), тело запроса и указатель на структуру,
// в которую будет декодирован ответ (out). Возвращает HTTP статус код и ошибку.
// baseURL подставляется автоматически внутри реализации.
// ВНИМАНИЕ: out должен быть указателем на тип ответа (например, *model.AdsInfoResponse).
// Если out == nil, тело ответа не декодируется.
// Такой подход позволяет типобезопасно работать с ответами без лишних приведений типов.
//
// Пример:
//
//	var resp model.AdsInfoResponse
//	status, err := client.request(ctx, http.MethodGet, "/core/v1/items", nil, &resp)
//	if err != nil { ... }
//	_ = status
//
//goland:noinspection GoNameStartsWithPackageName
type IHttpClient interface {
	request(ctx context.Context, method string, path string, body io.Reader, out any) (int, error)
}

// HttpClient - реализация IHttpClient на базе стандартного http.Client.
type HTTPClient struct {
	client *http.Client
	auth   IAuthorization
}

// NewHttpClient - конструктор HTTP-клиента с установленным таймаутом и токеном авторизации.
func NewHTTPClient(auth IAuthorization) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: _defaultTimeoutHTTP,
		},
		auth: auth,
	}
}

// request - выполняет HTTP-запрос к Avito API по относительному пути.
// Добавляет заголовок авторизации и, при необходимости, декодирует JSON-ответ в out.
func (h *HTTPClient) request(ctx context.Context, method, path string, body io.Reader, out any) (int, error) {
	token, err := h.auth.GetToken()
	if err != nil {
		return 0, fmt.Errorf("error at getting token: %w", err)
	}

	fullURL := fmt.Sprintf("%s%s", baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return 0, fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := h.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil
	}

	if out != nil {
		if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
			return resp.StatusCode, fmt.Errorf("decode response: %w", err)
		}
	}

	return resp.StatusCode, nil
}
