package http

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
	"fmt"
)

type messengerTransport struct {
	client  HTTPClientI
	account *entities.Account
}

func NewMessengerTransport(client HTTPClientI, account *entities.Account) repository.MessengerRepository {
	return &messengerTransport{
		client:  client,
		account: account,
	}
}

func (t *messengerTransport) GetChatsListV2(
	ctx context.Context,
	request *entities.ChatsListV2Request,
) (*entities.ChatsListV2Response, error) {
	endpoint := fmt.Sprintf("/messenger/v2/accounts/%d/chats", t.account.ID)
	var response entities.ChatsListV2Response

	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		endpoint,
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get chats list: %w", err)
	}

	return &response, nil
}

func (t *messengerTransport) GetChatMessagesListV3(
	ctx context.Context,
	request *entities.ChatMessagesListV3Request,
) (*entities.ChatMessagesListV3Response, error) {
	endpoint := fmt.Sprintf("/messenger/v3/accounts/%d/chats/%s/messages/",
		t.account.ID,
		request.ChatID,
	)
	var response entities.ChatMessagesListV3Response

	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		endpoint,
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat messages: %w", err)
	}

	return &response, nil
}
