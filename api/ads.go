package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/VLADBLOOD/avito-sdk-golang/model"

	"github.com/google/go-querystring/query"
)

// ADS - сервис работы с объявлениями Avito.
// Выполняет запросы для получения информации об объявлениях и их статистики.
type ADS struct {
	client IHttpClient
}

// NewADS - конструктор сервиса объявлений.
func NewADS(client IHttpClient) *ADS {
	return &ADS{
		client: client,
	}
}

// GetAdsInfo - получение информации по объявлениям с возможностью фильтрации.
// Возвращает структуру AdsInfoResponse или ошибку.
func (a *ADS) GetAdsInfo(ctx context.Context, request *model.AdsInfoRequest) (*model.AdsInfoResponse, error) {
	path := "/core/v1/items"
	response := new(model.AdsInfoResponse)

	if request != nil {
		values, err := query.Values(request)
		if err != nil {
			return nil, fmt.Errorf("get ads info: не удалось собрать query params: %w", err)
		}
		queryString := values.Encode()
		if queryString != "" {
			path += "?" + queryString
		}
	}

	status, err := a.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get ads info: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get ads info: сервер вернул код %d", status)
	}

	return response, nil
}

// GetAdsStats - получение статистики по объявлениям за период.
// Возвращает структуру AdsStatsResponse или ошибку.
func (a *ADS) GetAdsStats(ctx context.Context, accountID int64, request *model.AdsStatsRequest) (*model.AdsStatsResponse, error) {
	path := fmt.Sprintf("/stats/v1/accounts/%d/items", accountID)
	response := new(model.AdsStatsResponse)

	var bodyReader io.Reader

	if request != nil {
		jsonBody, err := json.Marshal(request)
		if err != nil {
			return nil, fmt.Errorf("get ads stats: не удалось сериализовать тело запроса: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	status, err := a.client.request(ctx, http.MethodPost, path, bodyReader, response)
	if err != nil {
		return response, fmt.Errorf("get ads stats: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get ads stats: сервер вернул код %d", status)
	}

	return response, nil
}
