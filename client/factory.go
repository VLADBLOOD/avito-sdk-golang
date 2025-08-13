package client

import (
	"avito-sdk/entities"
	"avito-sdk/services"
	"avito-sdk/transport/http"
	"context"
	"fmt"
)

// Client - основной клиент SDK, агрегирующий все сервисы
type Client struct {
	Ads          *services.AdsService
	Autoloads    *services.AutoloadsService
	CallTracking *services.CallTrackingService
	Messenger    *services.MessengerService
}

// NewClient - создание нового клиента SDK с получением токена при инициализации
func NewClient(baseURL string, account *entities.Account) (*Client, error) {
	httpClient := http.NewHTTPClient(baseURL)
	ctx := context.Background()

	// Сервис авторизации
	authRepo := http.NewAuthorizationTransport(httpClient, baseURL)
	authService := services.NewAuthorizationService(authRepo)

	// Получаем токен
	token, err := authService.GetAccessToken(ctx, account.Credentials)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения токена доступа: %w", err)
	}

	fmt.Print(token)

	// Сохраняем токен в account
	account.Token = token

	// Создаем транспорты
	adsRepo := http.NewAdsTransport(httpClient, account)
	autoloadsRepo := http.NewAutoloadsTransport(httpClient, account)
	callTrackingRepo := http.NewCallTrackingTransport(httpClient, account)
	messengerRepo := http.NewMessengerTransport(httpClient, account)

	// Создаем сервисы
	adsService := services.NewAdsService(adsRepo)
	autoloadsService := services.NewAutoloadsService(autoloadsRepo)
	callTrackingService := services.NewCallTrackingService(callTrackingRepo)
	messengerService := services.NewMessengerService(messengerRepo)

	return &Client{
		Ads:          adsService,
		Autoloads:    autoloadsService,
		CallTracking: callTrackingService,
		Messenger:    messengerService,
	}, nil
}
