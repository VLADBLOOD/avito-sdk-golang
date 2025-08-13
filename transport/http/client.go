package http

import (
	"avito-sdk/entities"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

// HTTPClient - реализация HTTP клиента
type HTTPClient struct {
	client  *http.Client
	baseURL string
}

// NewHTTPClient создает новый HTTP клиент
func NewHTTPClient(baseURL string) HTTPClientI {
	return &HTTPClient{
		client:  &http.Client{Timeout: 30 * time.Second},
		baseURL: baseURL,
	}
}

func (c *HTTPClient) GetAccessToken(ctx context.Context, credentials entities.AccountCredentials, endpoint string) (*http.Response, error) {
	// URL для получения токена
	fullURL := c.baseURL + endpoint

	// Подготовка данных для запроса
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", credentials.ClientID)
	data.Set("client_secret", credentials.ClientSecret)

	// Создание HTTP запроса
	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}

	// Установка заголовков
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	// Выполнение запроса
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса на получение токена доступа: %w", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("получен статус ошибки в запросе при получении токена: %d", resp.StatusCode)
	}

	return resp, nil
}

// DoRequest реализация
func (c *HTTPClient) DoRequest(
	ctx context.Context,
	method string,
	account *entities.Account,
	endpoint string,
	requestData any,
	resultEntity any,
) error {
	var resp *http.Response
	var err error

	switch method {
	case "GET":
		resp, err = c.DoGET(ctx, account, endpoint, requestData)
	case "POST":
		resp, err = c.DoPOST(ctx, account, endpoint, requestData)
	default:
		return fmt.Errorf("неподдерживаемый HTTP метод: %s", method)
	}

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(resultEntity)
}

// DoGET реализует GET запрос
func (c *HTTPClient) DoGET(ctx context.Context, account *entities.Account, endpoint string, queryParams any) (*http.Response, error) {
	// Формируем полный URL с query параметрами
	fullURL := c.baseURL + endpoint

	// Добавляем query параметры
	if queryParams != nil {
		values, err := query.Values(queryParams)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования queryParams: %w", err)
		}

		queryString := values.Encode()
		if queryString != "" {
			fullURL += "?" + queryString
		}
	}

	// Создаем GET запрос
	req, err := http.NewRequestWithContext(ctx, "GET", fullURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания GET request: %w", err)
	}

	// Добавляем заголовки авторизации
	return c.executeRequestWithAuth(ctx, req, account)
}

// DoPOST реализует POST запрос
func (c *HTTPClient) DoPOST(ctx context.Context, account *entities.Account, endpoint string, body any) (*http.Response, error) {
	fullURL := c.baseURL + endpoint

	var bodyReader io.Reader

	// Создаем тело запроса если есть данные
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования структуры в json: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	// Создаем POST запрос
	req, err := http.NewRequestWithContext(ctx, "POST", fullURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания POST запроса: %w", err)
	}

	// Добавляем заголовки авторизации
	return c.executeRequestWithAuth(ctx, req, account)
}

// executeRequestWithAuth - выполнение запроса с авторизацией и обработкой ошибок
func (c *HTTPClient) executeRequestWithAuth(ctx context.Context, req *http.Request, account *entities.Account) (*http.Response, error) {
	// Добавляем заголовки авторизации из account.Token
	if account.Token != nil {
		req.Header.Set("Authorization", "Bearer "+account.Token.AccessToken)
	}
	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос с обработкой повторов при 429
	return c.executeWithRetry(ctx, req)
}

// executeWithRetry - выполнение запроса с повтором при 429
func (c *HTTPClient) executeWithRetry(ctx context.Context, req *http.Request) (*http.Response, error) {
	maxRetries := 3
	retryDelay := 20 * time.Second

	for i := range maxRetries {
		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}

		switch resp.StatusCode {
		// 200 - успешный ответ
		case http.StatusOK:
			return resp, nil

		// 400 - неверный запрос
		case http.StatusBadRequest:
			return nil, fmt.Errorf("bad request (400)")

		// 403 - доступ запрещен
		case http.StatusForbidden:
			return nil, fmt.Errorf("forbidden (403)")

		// 404 - данные не найдены
		case http.StatusNotFound:
			return nil, fmt.Errorf("not found (404)")

		// 500 - внутренняя ошибка API
		case http.StatusInternalServerError:
			return nil, fmt.Errorf("internal server error (500)")

		// 429 - слишком много запросов
		case http.StatusTooManyRequests:
			if i < maxRetries-1 {
				// Ждем и повторяем
				select {
				case <-time.After(retryDelay):
					// Продолжаем цикл для повтора
				case <-ctx.Done():
					return nil, ctx.Err()
				}
			}

		// Другие ошибки
		default:
			return nil, fmt.Errorf("непредвиденная ошибка, status code: %d", resp.StatusCode)
		}
	}

	return nil, fmt.Errorf("превышено допустимое кол-во попыток выполнить запрос")
}
