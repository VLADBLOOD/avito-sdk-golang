package services

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
)

type AutoloadsService struct {
	repo repository.AutoloadsRepository
}

func NewAutoloadsService(repo repository.AutoloadsRepository) *AutoloadsService {
	return &AutoloadsService{repo: repo}
}

// GetAutoloadsListV2 - получение списка отчетов автозагрузки
func (s *AutoloadsService) GetAutoloadsListV2(
	ctx context.Context,
	request *entities.AutoloadsListV2Request,
) (*entities.AutoloadsListV2Response, error) {
	return s.repo.GetAutoloadsListV2(ctx, request)
}

// GetAutoloadAdsListV2 - получение объявлений из выгрузки
func (s *AutoloadsService) GetAutoloadAdsListV2(
	ctx context.Context,
	request *entities.AutoloadAdsListV2Request,
) (*entities.AutoloadAdsListV2Response, error) {
	return s.repo.GetAutoloadAdsListV2(ctx, request)
}

// GetAutoloadStatsV3 - получение статистики автозагрузки
func (s *AutoloadsService) GetAutoloadStatsV3(
	ctx context.Context,
	request *entities.AutoloadStatsV3Request,
) (*entities.AutoloadStatsV3Response, error) {
	return s.repo.GetAutoloadStatsV3(ctx, request)
}
