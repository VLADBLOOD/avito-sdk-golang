package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/VLADBLOOD/avito-sdk-golang/model"

	"github.com/google/go-querystring/query"
)

type Autoloads struct {
	client IHttpClient
}

// NewADS - конструктор сервиса объявлений.
func NewAutoloads(client IHttpClient) *Autoloads {
	return &Autoloads{
		client: client,
	}
}

func (a *Autoloads) GetAutoloadsListV2(
	ctx context.Context,
	request *model.AutoloadsListV2Request,
) (*model.AutoloadsListV2Response, error) {
	path := "/autoload/v2/reports"
	response := new(model.AutoloadsListV2Response)

	if request != nil {
		values, err := query.Values(request)
		if err != nil {
			return nil, fmt.Errorf("get autoloas list: не удалось собрать query params: %w", err)
		}
		queryString := values.Encode()
		if queryString != "" {
			path += "?" + queryString
		}
	}

	fmt.Print(path)

	status, err := a.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get autoloads list: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get autoloads list: сервер вернул код %d", status)
	}

	return response, nil
}

func (a *Autoloads) GetAutoloadAdsListV2(
	ctx context.Context,
	reportID string,
	request *model.AutoloadAdsListV2Request,
) (*model.AutoloadAdsListV2Response, error) {
	path := fmt.Sprintf("/autoload/v2/reports/%s/items", reportID)
	response := new(model.AutoloadAdsListV2Response)

	if request != nil {
		values, err := query.Values(request)
		if err != nil {
			return nil, fmt.Errorf("get autoload ads: не удалось собрать query params: %w", err)
		}
		queryString := values.Encode()
		if queryString != "" {
			path += "?" + queryString
		}
	}

	status, err := a.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get autoload ads: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get autoload ads: сервер вернул код %d", status)
	}

	return response, nil
}

func (a *Autoloads) GetAutoloadStatsV3(
	ctx context.Context,
	reportID string,
) (*model.AutoloadStatsV3Response, error) {
	path := fmt.Sprintf("/autoload/v3/reports/%s", reportID)
	response := new(model.AutoloadStatsV3Response)
	status, err := a.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get autoload stats: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get autoload stats: сервер вернул код %d", status)
	}

	return response, nil
}
