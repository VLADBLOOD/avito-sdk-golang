package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/VLADBLOOD/avito-sdk-golang/model"

	"github.com/google/go-querystring/query"
)

// ADS - сервис работы с объявлениями Avito.
// Выполняет запросы для получения информации об объявлениях и их статистики.
type Messenger struct {
	client IHttpClient
}

// NewADS - конструктор сервиса объявлений.
func NewMessenger(client IHttpClient) *Messenger {
	return &Messenger{
		client: client,
	}
}

func (m *Messenger) GetChatsListV2(
	ctx context.Context,
	accountID int64,
	request *model.ChatsListV2Request,
) (*model.ChatsListV2Response, error) {
	path := fmt.Sprintf("/messenger/v2/accounts/%d/chats", accountID)
	response := new(model.ChatsListV2Response)

	if request != nil {
		values, err := query.Values(request)
		if err != nil {
			return nil, fmt.Errorf("get chats list: не удалось собрать query params: %w", err)
		}
		queryString := values.Encode()
		if queryString != "" {
			path += "?" + queryString
		}
	}

	status, err := m.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get chats list: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get chats list: сервер вернул код %d", status)
	}

	return response, nil
}

func (m *Messenger) GetChatMessagesListV3(
	ctx context.Context,
	accountID int64,
	chatID string,
	request *model.ChatMessagesListV3Request,
) (*model.ChatMessagesListV3Response, error) {
	path := fmt.Sprintf("/messenger/v3/accounts/%d/chats/%s/messages/",
		accountID,
		chatID,
	)
	response := new(model.ChatMessagesListV3Response)

	if request != nil {
		values, err := query.Values(request)
		if err != nil {
			return nil, fmt.Errorf("get chat messages: не удалось собрать query params: %w", err)
		}
		queryString := values.Encode()
		if queryString != "" {
			path += "?" + queryString
		}
	}

	status, err := m.client.request(ctx, http.MethodGet, path, io.Reader(nil), response)
	if err != nil {
		return response, fmt.Errorf("get chat messages: запрос завершился ошибкой: %w", err)
	}

	// TODO: централизованный обработчик кодов ответа
	if status != http.StatusOK {
		return response, fmt.Errorf("get chat messages: сервер вернул код %d", status)
	}

	return response, nil
}
