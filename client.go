package client

import (
	"context"
	"fmt"

	"github.com/VLADBLOOD/avito-sdk-golang/api"
	"github.com/VLADBLOOD/avito-sdk-golang/model"
)

// IAdvertising - интерфейс сервиса объявлений.
// Позволяет получать информацию об объявлениях и их статистику.
type IAdvertising interface {
	GetAdsInfo(ctx context.Context, request *model.AdsInfoRequest) (*model.AdsInfoResponse, error)
	GetAdsStats(ctx context.Context, accountID int64, request *model.AdsStatsRequest) (*model.AdsStatsResponse, error)
}

// IAutoloads - интерфейс сервиса автозагрузок.
// Позволяет получать информацию по автозагрузкам, их статистику и содержащиеся в них объявления
type IAutoloads interface {
	GetAutoloadsListV2(ctx context.Context, request *model.AutoloadsListV2Request) (*model.AutoloadsListV2Response, error)
	GetAutoloadAdsListV2(
		ctx context.Context,
		reportID string,
		request *model.AutoloadAdsListV2Request,
	) (*model.AutoloadAdsListV2Response, error)
	GetAutoloadStatsV3(ctx context.Context, reportID string) (*model.AutoloadStatsV3Response, error)
}

// ICallTracking - интерфейс сервиса звонков.
// Позволяет получать количество звонков за период.
type ICallTracking interface {
	GetCallsByPeriod(ctx context.Context, request *model.CallsByPeriodRequest) (*model.CallsByPeriodResponse, error)
}

// IMessenger - интерфейс сервиса чатов.
// Позволяет получать информацию о чатах и сообщения в них
type IMessenger interface {
	GetChatsListV2(ctx context.Context, accountID int64, request *model.ChatsListV2Request) (*model.ChatsListV2Response, error)
	GetChatMessagesListV3(
		ctx context.Context,
		accountID int64,
		chatID string,
		request *model.ChatMessagesListV3Request,
	) (*model.ChatMessagesListV3Response, error)
}

// Client - основной клиент SDK, агрегирует сервисы Avito API.
type Client struct {
	ADS          IAdvertising  // Сервис объявлений
	Autoloads    IAutoloads    // Сервис автозагрузок
	CallTracking ICallTracking // Сервис звонков
	Messenger    IMessenger    // Сервис чатов
}

// NewClient - конструктор клиента SDK.
// Получает токен по Client Credentials и создает HTTP-клиент для дальнейших запросов.
func NewClient(creds *api.Credentials) (*Client, error) {
	if creds == nil || creds.ClientID == "" || creds.ClientSecret == "" {
		return nil, fmt.Errorf("invalid creds: clientID and clientSecret are required")
	}

	token, err := api.GetToken(creds)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	http := api.NewHTTPClient(token, creds)

	return &Client{
		ADS:          api.NewADS(http),
		Autoloads:    api.NewAutoloads(http),
		CallTracking: api.NewCallTracking(http),
		Messenger:    api.NewMessenger(http),
	}, nil
}
