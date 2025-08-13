/*
Репозиторий Authorization описывает контракт получения AccessToken Авито:

GetAccessToken: Возвращает Bearer токен по client_id и client_secret.
Doc: https://developers.avito.ru/api-catalog/item/documentation#operation/getItemsInfo
*/
package repository

import (
	"avito-sdk/entities"
	"context"
)

type AuthorizationRepository interface {
	GetAccessToken(ctx context.Context, credentials entities.AccountCredentials) (*entities.Token, error)
}
