package services

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
)

type CallTrackingService struct {
	repo repository.CallTrackingRepository
}

func NewCallTrackingService(repo repository.CallTrackingRepository) *CallTrackingService {
	return &CallTrackingService{repo: repo}
}

// GetCallsByPeriod - получение списка звонков по периоду
func (s *CallTrackingService) GetCallsByPeriod(
	ctx context.Context,
	request *entities.CallsByPeriodRequest,
) (*entities.CallsByPeriodResponse, error) {
	return s.repo.GetCallsByPeriod(ctx, request)
}
