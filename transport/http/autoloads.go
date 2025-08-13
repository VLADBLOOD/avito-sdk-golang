package http

import (
	"avito-sdk/entities"
	"avito-sdk/repository"
	"context"
	"fmt"
)

type autoloadsTransport struct {
	client  HTTPClientI
	account *entities.Account
}

func NewAutoloadsTransport(client HTTPClientI, account *entities.Account) repository.AutoloadsRepository {
	return &autoloadsTransport{
		client:  client,
		account: account,
	}
}

func (t *autoloadsTransport) GetAutoloadsListV2(
	ctx context.Context,
	request *entities.AutoloadsListV2Request,
) (*entities.AutoloadsListV2Response, error) {
	var response entities.AutoloadsListV2Response
	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		"/autoload/v2/reports",
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get autoloads list: %w", err)
	}
	return &response, nil
}

func (t *autoloadsTransport) GetAutoloadAdsListV2(
	ctx context.Context,
	request *entities.AutoloadAdsListV2Request,
) (*entities.AutoloadAdsListV2Response, error) {
	endpoint := fmt.Sprintf("/autoload/v2/reports/%s/items", request.ReportID)
	var response entities.AutoloadAdsListV2Response
	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		endpoint,
		request,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get autoload ads list: %w", err)
	}
	return &response, nil
}

func (t *autoloadsTransport) GetAutoloadStatsV3(
	ctx context.Context,
	request *entities.AutoloadStatsV3Request,
) (*entities.AutoloadStatsV3Response, error) {
	endpoint := fmt.Sprintf("/autoload/v3/reports/%s", request.ReportID)
	var response entities.AutoloadStatsV3Response
	err := t.client.DoRequest(
		ctx,
		"GET",
		t.account,
		endpoint,
		nil,
		&response,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get autoload stats: %w", err)
	}
	return &response, nil
}
