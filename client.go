package client

import (
	"avito-sdk/api"
	"avito-sdk/model"
	"context"
	"fmt"
)

// IAdvertising - интерфейс сервиса объявлений.
// Позволяет получать информацию об объявлениях и их статистику.
 type IAdvertising interface {
	GetAdsInfo(ctx context.Context, request *model.AdsInfoRequest) (*model.AdsInfoResponse, error)
	GetAdsStats(ctx context.Context, request *model.AdsStatsRequest) (*model.AdsStatsResponse, error)
}

// Config - конфигурация клиента SDK.
 type Config struct {
	ClientID     string // Идентификатор клиента (account_id)
	ClientSecret string // Секрет клиента для получения токена
}

// Client - основной клиент SDK, агрегирует сервисы Avito API.
 type Client struct {
	ADS IAdvertising // Сервис объявлений
}

// NewClient - конструктор клиента SDK.
// Получает токен по Client Credentials и создает HTTP-клиент для дальнейших запросов.
 func NewClient(config *Config) (*Client, error) {
	if config == nil || config.ClientID == "" || config.ClientSecret == "" {
		return nil, fmt.Errorf("invalid config: clientID and clientSecret are required")
	}

	token, err := api.GetToken(config.ClientID, config.ClientSecret)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	http := api.NewHttpClient(token)

	return &Client{
		ADS: api.NewADS(config.ClientID, http),
	}, nil
}
