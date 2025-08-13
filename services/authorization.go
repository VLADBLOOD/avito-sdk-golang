package services

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
)

type AuthorizationService struct {
	repo repository.AuthorizationRepository
}

func NewAuthorizationService(repo repository.AuthorizationRepository) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) GetAccessToken(ctx context.Context, credentials entities.AccountCredentials) (*entities.Token, error) {
	return s.repo.GetAccessToken(ctx, credentials)
}
