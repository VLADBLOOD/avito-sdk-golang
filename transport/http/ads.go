package http

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
	"fmt"
)

type adsTransport struct {
	client  HTTPClientI
	account *entities.Account
}

func NewAdsTransport(client HTTPClientI, account *entities.Account) repository.AdsRepository {
	return &adsTransport{
		client:  client,
		account: account,
	}
}

func (t *adsTransport) GetAdsInfo(ctx context.Context, request *entities.AdsInfoRequest) (*entities.AdsInfoResponse, error) {
	var response entities.AdsInfoResponse
	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		"/core/v1/items",
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка получения ads info: %w", err)
	}

	return &response, nil
}

func (t *adsTransport) GetAdsStats(ctx context.Context, request *entities.AdsStatsRequest) (*entities.AdsStatsResponse, error) {
	endpoint := fmt.Sprintf("/stats/v1/accounts/%d/items", t.account.ID)
	var response entities.AdsStatsResponse
	err := t.client.DoRequest(
		ctx,
		"POST",
		t.account,
		endpoint,
		request,
		&response,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка получения ads stats: %w", err)
	}

	return &response, nil
}
