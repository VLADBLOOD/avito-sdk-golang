/*
Репозиторий Messenger описывает контракт, что нужно уметь для работы в Авито с чатами и их сообщениями:

GetChatsListV2: Возвращает список чатов
Doc: https://developers.avito.ru/api-catalog/messenger/documentation#operation/getChatsV2

GetAutoloadAdsListV2: Получение списка сообщений из чата
Doc: https://developers.avito.ru/api-catalog/messenger/documentation#operation/getMessagesV3
*/
package repository

import (
	"avito-sdk/entities"
	"context"
)

type MessengerRepository interface {
	GetChatsListV2(ctx context.Context, request *entities.ChatsListV2Request) (*entities.ChatsListV2Response, error)
	GetChatMessagesListV3(ctx context.Context, request *entities.ChatMessagesListV3Request) (*entities.ChatMessagesListV3Response, error)
}
