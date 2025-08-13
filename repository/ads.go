/*
Репозиторий Ads описывает контракт, что нужно уметь для работы с объявлениями Авито:

GetAdsInfo: Возвращает список объявлений авторизованного пользователя - статус, категорию и ссылку на сайте.
Doc: https://developers.avito.ru/api-catalog/item/documentation#operation/getItemsInfo

GetAdsStats: Получение счетчиков по переданному списку объявлений
Doc: https://developers.avito.ru/api-catalog/item/documentation#operation/itemStatsShallow
*/
package repository

import (
	"avito-sdk/entities"
	"context"
)

type AdsRepository interface {
	GetAdsInfo(ctx context.Context, request *entities.AdsInfoRequest) (*entities.AdsInfoResponse, error)
	GetAdsStats(ctx context.Context, request *entities.AdsStatsRequest) (*entities.AdsStatsResponse, error)
}
