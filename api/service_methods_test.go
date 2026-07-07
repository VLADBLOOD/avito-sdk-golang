package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/VLADBLOOD/avito-sdk-golang/model"
)

type stubHTTPClient struct {
	method   string
	path     string
	body     []byte
	status   int
	err      error
	response []byte
}

func (s *stubHTTPClient) request(_ context.Context, method, path string, body io.Reader, out any) (int, error) {
	s.method = method
	s.path = path
	if body != nil {
		data, err := io.ReadAll(body)
		if err != nil {
			return 0, err
		}
		s.body = data
	}
	if s.err != nil {
		return s.status, s.err
	}
	if out != nil && len(s.response) > 0 {
		if err := json.Unmarshal(s.response, out); err != nil {
			return 0, err
		}
	}
	return s.status, nil
}

func requireMethod(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("expected %s method, got %s", want, got)
	}
}

func requirePathContains(t *testing.T, got string, want ...string) {
	t.Helper()
	for _, part := range want {
		if !strings.Contains(got, part) {
			t.Fatalf("expected path to contain %q, got %q", part, got)
		}
	}
}

func TestADSService(t *testing.T) {
	t.Run("GetAccountSpendings", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"result":{"groupings":[{"date":"2024-01-01"}],"timestamp":1704067200}}`),
		}
		ads := NewADS(client)

		req := &model.AccountSpendingsRequest{
			DateFrom:      "2024-01-01",
			DateTo:        "2024-01-31",
			Grouping:      "day",
			SpendingTypes: []string{"promotion", "presence"},
			Filter: &model.AccountSpendingsFilter{
				CategoryIDs: []int64{111},
				ItemIDs:     []int64{999},
			},
		}

		resp, err := ads.GetAccountSpendings(context.Background(), 12345, req)
		if err != nil {
			t.Fatalf("GetAccountSpendings returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodPost)
		if client.path != "/stats/v2/accounts/12345/spendings" {
			t.Fatalf("unexpected path: %s", client.path)
		}
		if !bytes.Contains(client.body, []byte(`"grouping":"day"`)) {
			t.Fatalf("expected request body to contain grouping, got %s", string(client.body))
		}
		if !strings.Contains(string(client.body), `"promotion"`) {
			t.Fatalf("expected request body to contain spending types, got %s", string(client.body))
		}
		if !strings.Contains(string(client.body), `"itemIDs"`) || !strings.Contains(string(client.body), `"categoryIDs"`) {
			t.Fatalf("expected request body to contain filter fields, got %s", string(client.body))
		}
		if resp.Result.Timestamp != 1704067200 {
			t.Fatalf("unexpected timestamp: %d", resp.Result.Timestamp)
		}
	})

	t.Run("GetAdsInfo", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"meta":{"page":2,"per_page":10,"pages":1,"total":1},"resources":[{"id":123,"title":"Test","status":"active","url":"https://avito.ru/item","address":"Moscow","category":{"id":1,"name":"Товары"}}]}`),
		}
		ads := NewADS(client)

		req := &model.AdsInfoRequest{
			PerPage:       10,
			Page:          2,
			Status:        "active",
			UpdatedAtFrom: "2024-01-01",
			Category:      111,
		}

		resp, err := ads.GetAdsInfo(context.Background(), req)
		if err != nil {
			t.Fatalf("GetAdsInfo returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		requirePathContains(t, client.path, "per_page=10", "page=2", "status=active")
		if len(resp.Resources) != 1 || resp.Resources[0].ID != 123 {
			t.Fatalf("unexpected ads response: %+v", resp.Resources)
		}
	})

	t.Run("GetAdsStats", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"result":{"items":[{"itemId":77,"stats":[{"date":"2024-01-01","views":2}]}]}}`),
		}
		ads := NewADS(client)

		req := &model.AdsStatsRequest{
			DateFrom:       "2024-01-01",
			DateTo:         "2024-01-31",
			Fields:         []string{"views"},
			ItemIds:        []int64{77},
			PeriodGrouping: "day",
		}

		resp, err := ads.GetAdsStats(context.Background(), 42, req)
		if err != nil {
			t.Fatalf("GetAdsStats returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodPost)
		if client.path != "/stats/v1/accounts/42/items" {
			t.Fatalf("unexpected path: %s", client.path)
		}
		if !bytes.Contains(client.body, []byte(`"views"`)) || !bytes.Contains(client.body, []byte(`"periodGrouping":"day"`)) {
			t.Fatalf("unexpected request body: %s", string(client.body))
		}
		if len(resp.Result.Items) != 1 || resp.Result.Items[0].ItemID != 77 {
			t.Fatalf("unexpected stats response: %+v", resp.Result.Items)
		}
	})
}

func TestUserService(t *testing.T) {
	t.Run("GetUserInfoSelf", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"email":"test@example.com","id":42,"name":"Ivan","phone":"79998887766","phones":["79998887766"],"profile_url":"https://avito.ru/user/42/profile"}`),
		}
		user := NewUser(client)

		resp, err := user.GetUserInfoSelf(context.Background())
		if err != nil {
			t.Fatalf("GetUserInfoSelf returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		if client.path != "/core/v1/accounts/self" {
			t.Fatalf("unexpected path: %s", client.path)
		}
		if resp.ID != 42 {
			t.Fatalf("unexpected user id: %d", resp.ID)
		}
		if resp.Email != "test@example.com" {
			t.Fatalf("unexpected email: %s", resp.Email)
		}
	})
}

func TestCallTrackingService(t *testing.T) {
	t.Run("GetCallsByPeriod", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"calls":[{"callId":1,"buyerPhone":"+7999","callTime":"2024-01-01T00:00:00Z"}],"error":{"code":0,"message":""}}`),
		}
		tracking := NewCallTracking(client)

		req := &model.CallsByPeriodRequest{
			DateTimeFrom: "2024-01-01T00:00:00Z",
			DateTimeTo:   "2024-01-02T00:00:00Z",
			Limit:        5,
			Offset:       1,
		}

		resp, err := tracking.GetCallsByPeriod(context.Background(), req)
		if err != nil {
			t.Fatalf("GetCallsByPeriod returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodPost)
		if client.path != "/calltracking/v1/getCalls/" {
			t.Fatalf("unexpected path: %s", client.path)
		}
		if !bytes.Contains(client.body, []byte(`"limit":5`)) || !bytes.Contains(client.body, []byte(`"offset":1`)) {
			t.Fatalf("unexpected request body: %s", string(client.body))
		}
		if len(resp.Calls) != 1 || resp.Calls[0].CallID != 1 {
			t.Fatalf("unexpected call response: %+v", resp.Calls)
		}
	})
}

func TestAutoloadsService(t *testing.T) {
	t.Run("GetAutoloadsListV2", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"meta":{"page":1,"pages":1,"per_page":25,"total":1},"reports":[{"id":10,"status":"success"}]}`),
		}
		autos := NewAutoloads(client)

		req := &model.AutoloadsListV2Request{PerPage: 25, Page: 1}

		resp, err := autos.GetAutoloadsListV2(context.Background(), req)
		if err != nil {
			t.Fatalf("GetAutoloadsListV2 returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		requirePathContains(t, client.path, "/autoload/v2/reports", "PerPage=25", "Page=1")
		if len(resp.Reports) != 1 || resp.Reports[0].ID != 10 {
			t.Fatalf("unexpected reports response: %+v", resp.Reports)
		}
	})

	t.Run("GetAutoloadAdsListV2", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"items":[{"avito_id":77}],"meta":{"page":1,"pages":1,"per_page":10,"total":1},"report_id":100}`),
		}
		autos := NewAutoloads(client)

		req := &model.AutoloadAdsListV2Request{PerPage: 10, Page: 1, Query: "phone", Sections: "main"}

		resp, err := autos.GetAutoloadAdsListV2(context.Background(), "abc", req)
		if err != nil {
			t.Fatalf("GetAutoloadAdsListV2 returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		requirePathContains(t, client.path, "/autoload/v2/reports/abc/items", "Query=phone", "Sections=main", "PerPage=10")
		if resp.ReportID != 100 || len(resp.Items) != 1 {
			t.Fatalf("unexpected autoload ads response: %+v", resp)
		}
	})

	t.Run("GetAutoloadStatsV3", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"report_id":100,"status":"success","source":"file"}`),
		}
		autos := NewAutoloads(client)

		resp, err := autos.GetAutoloadStatsV3(context.Background(), "abc")
		if err != nil {
			t.Fatalf("GetAutoloadStatsV3 returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		if client.path != "/autoload/v3/reports/abc" {
			t.Fatalf("unexpected path: %s", client.path)
		}
		if resp.ReportID != 100 {
			t.Fatalf("unexpected stats response: %+v", resp)
		}
	})
}

func TestMessengerService(t *testing.T) {
	t.Run("GetChatsListV2", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"chats":[{"id":"1","created":1,"updated":1,"context":{"type":"item","value":{"id":10,"title":"Item","url":"https://avito.ru/item","user_id":1}},"users":[]}]}`),
		}
		messenger := NewMessenger(client)

		req := &model.ChatsListV2Request{ItemIDs: []int64{10}, UnreadOnly: true, ChatTypes: []string{"support"}, Limit: 5, Offset: 1}

		resp, err := messenger.GetChatsListV2(context.Background(), 42, req)
		if err != nil {
			t.Fatalf("GetChatsListV2 returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		requirePathContains(t, client.path, "/messenger/v2/accounts/42/chats", "Limit=5", "ItemIDs=10")
		if len(resp.Chats) != 1 || resp.Chats[0].ID != "1" {
			t.Fatalf("unexpected chat response: %+v", resp.Chats)
		}
	})

	t.Run("GetChatMessagesListV3", func(t *testing.T) {
		client := &stubHTTPClient{
			status:   http.StatusOK,
			response: []byte(`{"messages":[{"id":"msg1","author_id":1,"created":1,"direction":"in","type":"text","content":{"text":"hello"}}]}`),
		}
		messenger := NewMessenger(client)

		req := &model.ChatMessagesListV3Request{Limit: 10, Offset: 2}

		resp, err := messenger.GetChatMessagesListV3(context.Background(), 42, "chat-1", req)
		if err != nil {
			t.Fatalf("GetChatMessagesListV3 returned error: %v", err)
		}
		requireMethod(t, client.method, http.MethodGet)
		requirePathContains(t, client.path, "/messenger/v3/accounts/42/chats/chat-1/messages", "Limit=10", "Offset=2")
		if len(resp.Messages) != 1 || resp.Messages[0].ID != "msg1" {
			t.Fatalf("unexpected message response: %+v", resp.Messages)
		}
	})
}
