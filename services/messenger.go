package services

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
)

type MessengerService struct {
	repo repository.MessengerRepository
}

func NewMessengerService(repo repository.MessengerRepository) *MessengerService {
	return &MessengerService{repo: repo}
}

// GetChatsListV2 - получение списка чатов
func (s *MessengerService) GetChatsListV2(
	ctx context.Context,
	request *entities.ChatsListV2Request,
) (*entities.ChatsListV2Response, error) {
	return s.repo.GetChatsListV2(ctx, request)
}

// GetChatMessagesListV3 - получение сообщений чата
func (s *MessengerService) GetChatMessagesListV3(
	ctx context.Context,
	request *entities.ChatMessagesListV3Request,
) (*entities.ChatMessagesListV3Response, error) {
	return s.repo.GetChatMessagesListV3(ctx, request)
}
