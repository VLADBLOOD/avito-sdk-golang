package http

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
	"fmt"
)

type callTrackingTransport struct {
	client  HTTPClientI
	account *entities.Account
}

func NewCallTrackingTransport(client HTTPClientI, account *entities.Account) repository.CallTrackingRepository {
	return &callTrackingTransport{
		client:  client,
		account: account,
	}
}

func (t *callTrackingTransport) GetCallsByPeriod(
	ctx context.Context,
	request *entities.CallsByPeriodRequest,
) (*entities.CallsByPeriodResponse, error) {
	var response entities.CallsByPeriodResponse

	err := t.client.DoRequest(
		ctx,
		"POST",
		t.account,
		"/calltracking/v1/getCalls/",
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get calls by period: %w", err)
	}

	return &response, nil
}
