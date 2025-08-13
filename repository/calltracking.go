/*
Репозиторий CallTracking описывает контракт, что нужно уметь для работы со звонками в Авито:

GetCallsByPeriod: Возвращает список звонков с фильтром по времени звонка
Doc: https://developers.avito.ru/api-catalog/calltracking/documentation#operation/get_calls
*/
package repository

import (
	"avito-sdk/entities"
	"context"
)

type CallTrackingRepository interface {
	GetCallsByPeriod(ctx context.Context, request *entities.CallsByPeriodRequest) (*entities.CallsByPeriodResponse, error)
}
