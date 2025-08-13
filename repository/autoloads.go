/*
Репозиторий Autoloads описывает контракт, что нужно уметь для работы с автовыгрузками Авито:

GetAutoloadsListV2: Возвращает список отчётов автозагрузки
Doc: https://developers.avito.ru/api-catalog/autoload/documentation#operation/getReportsV2

GetAutoloadAdsListV2: С помощью этого метода можно получить результаты обработки каждого объявления в конкретной выгрузке.
Doc: https://developers.avito.ru/api-catalog/autoload/documentation#operation/getReportItemsFeesById

GetAutoloadStatsV3: Возвращает сводную статистику с результатами конкретной выгрузки
Doc: https://developers.avito.ru/api-catalog/autoload/documentation#operation/getReportByIdV3
*/
package repository

import (
	"avito-sdk/entities"
	"context"
)

type AutoloadsRepository interface {
	GetAutoloadsListV2(ctx context.Context, request *entities.AutoloadsListV2Request) (*entities.AutoloadsListV2Response, error)
	GetAutoloadAdsListV2(ctx context.Context, request *entities.AutoloadAdsListV2Request) (*entities.AutoloadAdsListV2Response, error)
	GetAutoloadStatsV3(ctx context.Context, request *entities.AutoloadStatsV3Request) (*entities.AutoloadStatsV3Response, error)
}
