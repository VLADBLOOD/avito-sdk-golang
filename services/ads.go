package services

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
)

type AdsService struct {
	repo repository.AdsRepository
}

func NewAdsService(repo repository.AdsRepository) *AdsService {
	return &AdsService{repo: repo}
}

// GetAdsInfo - получение информации по объявлениям
func (s *AdsService) GetAdsInfo(ctx context.Context, request *entities.AdsInfoRequest) (*entities.AdsInfoResponse, error) {
	return s.repo.GetAdsInfo(ctx, request)
}

// GetAdsStats - получение статистики по объявлениям
func (s *AdsService) GetAdsStats(ctx context.Context, request *entities.AdsStatsRequest) (*entities.AdsStatsResponse, error) {
	return s.repo.GetAdsStats(ctx, request)
}
