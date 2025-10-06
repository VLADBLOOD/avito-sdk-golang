package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"avito-sdk-golang/model"
)

// ADS - сервис работы с объявлениями Avito.
// Выполняет запросы для получения информации об объявлениях и их статистики.
type CallTracking struct {
	client IHttpClient
}

// NewADS - конструктор сервиса объявлений.
func NewCallTracking(client IHttpClient) *CallTracking {
	return &CallTracking{
		client: client,
	}
}

func (c *CallTracking) GetCallsByPeriod(
	ctx context.Context,
	request *model.CallsByPeriodRequest,
) (*model.CallsByPeriodResponse, error) {
	path := "/calltracking/v1/getCalls/"
	response := new(model.CallsByPeriodResponse)

	var bodyReader io.Reader

	if request != nil {
		jsonBody, err := json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("get calls by period: не удалось сериализовать тело запроса: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	status, err := c.client.request(ctx, http.MethodPost, path, bodyReader, response)
	if err != nil {
		return response, fmt.Errorf("get calls by period: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get calls by period: сервер вернул код %d", status)
	}

	return response, nil
}
