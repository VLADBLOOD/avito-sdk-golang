package main

import (
	"avito-sdk/client"
	"avito-sdk/entities"
	"context"
	"fmt"
	"log"
)

func main() {
	account := setupAccount()
	sdkClient := setupSDKClient(account)
	ctx := context.Background()

	demoAdsInfo(ctx, sdkClient)
	demoAdsStats(ctx, sdkClient)
	demoCallTracking(ctx, sdkClient)
	demoMessenger(ctx, sdkClient)
	demoAutoloads(ctx, sdkClient)

	fmt.Println("\n=== Демонстрация завершена ===")
}

func setupAccount() *entities.Account {
	return &entities.Account{
		ID:   1,
		Name: "Авито Аккаунт",
		Credentials: entities.AccountCredentials{
			ClientID:     "client_id",
			ClientSecret: "client_secret",
		},
	}
}

func setupSDKClient(account *entities.Account) *client.Client {
	sdkClient, err := client.NewClient("https://api.avito.ru", account)
	if err != nil {
		log.Fatalf("Ошибка создания клиента SDK: %v", err)
	}
	return sdkClient
}

func demoAdsInfo(ctx context.Context, sdkClient *client.Client) {
	fmt.Println("=== Получение информации по объявлениям ===")
	adsRequest := &entities.AdsInfoRequest{
		PerPage: 10,
		Page:    1,
		Status:  "active",
	}

	adsResponse, err := sdkClient.Ads.GetAdsInfo(ctx, adsRequest)
	if err != nil {
		log.Printf("Error getting ads info: %v", err)
		return
	}

	fmt.Printf("Найдено %d объявлений\n", len(adsResponse.Resources))
	for _, ad := range adsResponse.Resources {
		fmt.Printf("- %s (ID: %d, Цена: %v)\n", ad.Title, ad.ID, ad.Price)
	}
}

func demoAdsStats(ctx context.Context, sdkClient *client.Client) {
	fmt.Println("\n=== Получение статистики по объявлениям ===")
	adsResponse, _ := sdkClient.Ads.GetAdsInfo(ctx, &entities.AdsInfoRequest{PerPage: 3})

	var adIDs []int64
	if adsResponse != nil && len(adsResponse.Resources) > 0 {
		for _, ad := range adsResponse.Resources {
			adIDs = append(adIDs, ad.ID)
		}
	}

	if len(adIDs) == 0 {
		fmt.Println("Нет объявлений для статистики")
		return
	}

	statsRequest := &entities.AdsStatsRequest{
		DateFrom: "2025-08-01",
		DateTo:   "2025-08-13",
		ItemIds:  adIDs,
	}

	statsResponse, err := sdkClient.Ads.GetAdsStats(ctx, statsRequest)
	if err != nil {
		log.Printf("Error getting ads stats: %v", err)
		return
	}

	fmt.Printf("Получена статистика для %d объявлений\n", len(statsResponse.Result.Items))
	for _, item := range statsResponse.Result.Items {
		fmt.Printf("Объявление ID %d: %d записей статистики\n", item.ItemID, len(item.Stats))
	}
}

func demoCallTracking(ctx context.Context, sdkClient *client.Client) {
	fmt.Println("\n=== Получение списка звонков ===")
	callsRequest := &entities.CallsByPeriodRequest{
		DateTimeFrom: "2025-08-01T00:00:00Z",
		Limit:        10,
		Offset:       0,
	}

	callsResponse, err := sdkClient.CallTracking.GetCallsByPeriod(ctx, callsRequest)
	if err != nil {
		log.Printf("Error getting calls list: %v", err)
		return
	}

	if callsResponse.Error.Code != 0 {
		fmt.Printf("API Error: %s (code: %d)\n", callsResponse.Error.Message, callsResponse.Error.Code)
		return
	}

	fmt.Printf("Найдено %d звонков\n", len(callsResponse.Calls))
	for _, call := range callsResponse.Calls {
		fmt.Printf("- Звонок ID %d, время: %s\n", call.CallID, call.CallTime)
	}
}

func demoMessenger(ctx context.Context, sdkClient *client.Client) {
	fmt.Println("\n=== Получение списка чатов ===")
	chatsRequest := &entities.ChatsListV2Request{
		Limit:  10,
		Offset: 0,
	}

	chatsResponse, err := sdkClient.Messenger.GetChatsListV2(ctx, chatsRequest)
	if err != nil {
		log.Printf("Error getting chats list: %v", err)
		return
	}

	fmt.Printf("Найдено %d чатов\n", len(chatsResponse.Chats))
	for _, chat := range chatsResponse.Chats {
		fmt.Printf("- Чат ID %s, обновлен: %d\n", chat.ID, chat.Updated)
	}

	demoChatMessages(ctx, sdkClient, chatsResponse)
}

func demoChatMessages(ctx context.Context, sdkClient *client.Client, chatsResponse *entities.ChatsListV2Response) {
	if chatsResponse == nil || len(chatsResponse.Chats) == 0 {
		return
	}

	fmt.Println("\n=== Получение сообщений чата ===")
	chatID := chatsResponse.Chats[0].ID
	var limit = 5
	var offset = 0

	messagesRequest := &entities.ChatMessagesListV3Request{
		ChatID: chatID,
		Limit:  intPtr(limit),
		Offset: intPtr(offset),
	}

	messagesResponse, err := sdkClient.Messenger.GetChatMessagesListV3(ctx, messagesRequest)
	if err != nil {
		log.Printf("Error getting chat messages: %v", err)
		return
	}

	fmt.Printf("Найдено %d сообщений в чате %s\n", len(messagesResponse.Messages), chatID)
	for _, message := range messagesResponse.Messages {
		fmt.Printf("- Сообщение ID %s от пользователя %d\n", message.ID, message.AuthorID)
		if message.Content.Text != "" {
			fmt.Printf("  Текст: %s\n", message.Content.Text)
		}
	}
}

func demoAutoloads(ctx context.Context, sdkClient *client.Client) {
	fmt.Println("\n=== Получение списка отчетов автозагрузки ===")

	var perPage = 5
	var page = 0

	autoloadsRequest := &entities.AutoloadsListV2Request{
		PerPage: intPtr(perPage),
		Page:    intPtr(page),
	}

	autoloadsResponse, err := sdkClient.Autoloads.GetAutoloadsListV2(ctx, autoloadsRequest)
	if err != nil {
		log.Printf("Error getting autoloads list: %v", err)
		return
	}

	fmt.Printf("Найдено %d отчетов автозагрузки\n", len(autoloadsResponse.Reports))
	for _, report := range autoloadsResponse.Reports {
		fmt.Printf("- Отчет ID %d, статус: %s\n", report.ID, report.Status)
	}

	demoAutoloadAds(ctx, sdkClient, autoloadsResponse)
	demoAutoloadStats(ctx, sdkClient, autoloadsResponse)
}

func demoAutoloadAds(ctx context.Context, sdkClient *client.Client, autoloadsResponse *entities.AutoloadsListV2Response) {
	if autoloadsResponse == nil || len(autoloadsResponse.Reports) == 0 {
		return
	}

	fmt.Println("\n=== Получение объявлений из выгрузки ===")
	reportID := fmt.Sprintf("%d", autoloadsResponse.Reports[0].ID)

	var perPage = 5
	var page = 0

	autoloadAdsRequest := &entities.AutoloadAdsListV2Request{
		ReportID: reportID,
		PerPage:  intPtr(perPage),
		Page:     intPtr(page),
	}

	autoloadAdsResponse, err := sdkClient.Autoloads.GetAutoloadAdsListV2(ctx, autoloadAdsRequest)
	if err != nil {
		log.Printf("Error getting autoload ads list: %v", err)
		return
	}

	fmt.Printf("Найдено %d объявлений в отчете %s\n", len(autoloadAdsResponse.Items), reportID)
	for _, ad := range autoloadAdsResponse.Items {
		fmt.Printf("- Объявление ID: %s, статус: %s\n", ad.AdID, ad.AvitoStatus)
	}
}

func demoAutoloadStats(ctx context.Context, sdkClient *client.Client, autoloadsResponse *entities.AutoloadsListV2Response) {
	if autoloadsResponse == nil || len(autoloadsResponse.Reports) == 0 {
		return
	}

	fmt.Println("\n=== Получение статистики автозагрузки ===")
	reportID := fmt.Sprintf("%d", autoloadsResponse.Reports[0].ID)

	autoloadStatsRequest := &entities.AutoloadStatsV3Request{
		ReportID: reportID,
	}

	autoloadStatsResponse, err := sdkClient.Autoloads.GetAutoloadStatsV3(ctx, autoloadStatsRequest)
	if err != nil {
		log.Printf("Error getting autoload stats: %v", err)
		return
	}

	fmt.Printf("Статистика для отчета %s:\n", reportID)
	fmt.Printf("- Статус: %s\n", autoloadStatsResponse.Status)
	fmt.Printf("- Источник: %s\n", autoloadStatsResponse.Source)

	if autoloadStatsResponse.SectionStats != nil {
		fmt.Printf("- Всего разделов: %d\n", autoloadStatsResponse.SectionStats.Count)
	}

	if len(autoloadStatsResponse.Events) > 0 {
		fmt.Printf("- Событий: %d\n", len(autoloadStatsResponse.Events))
	}
}

func intPtr(v int) *int {
	return &v
}
