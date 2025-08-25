package model

// AdsInfoRequest - фильтр запроса информации по объявлениям
type AdsInfoRequest struct {
	PerPage       int32  `url:"per_page"`
	Page          int32  `url:"page"`
	Status        string `url:"status"`
	UpdatedAtFrom string `url:"updatedAtFrom"`
	Category      int32  `url:"category"`
}

// AdsStatsRequest - фильтр запроса статистики по объявлениям
type AdsStatsRequest struct {
	DateFrom       string   `json:"dateFrom"`
	DateTo         string   `json:"dateTo"`
	Fields         []string `json:"fields"`
	ItemIds        []int64  `json:"itemIds"`
	PeriodGrouping string   `json:"periodGrouping"`
}

// AdsInfoResponse - ответ API для списка объявлений
type AdsInfoResponse struct {
	Meta struct {
		Page    int32 `json:"page"`
		PerPage int32 `json:"per_page"`
		Pages   int32 `json:"pages"`
		Total   int32 `json:"total"`
	} `json:"meta"`
	Resources []struct {
		ID       int64   `json:"id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Status   string  `json:"status"`
		URL      string  `json:"url"`
		Address  string  `json:"address"`
		Category struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"category"`
	} `json:"resources"`
}

// AdsStatsResponse - ответ API со статистикой
type AdsStatsResponse struct {
	Result AdStatsResult `json:"result"`
}

// AdStatsResult - результат статистики
type AdStatsResult struct {
	Items []AdStatsItem `json:"items"`
}

// AdStatsItem - статистика по одному объявлению
type AdStatsItem struct {
	ItemID int64 `json:"itemId"`
	Stats  []struct {
		Date          string `json:"date,omitempty"`
		Views         int32  `json:"views,omitempty"`
		UniqViews     int32  `json:"uniqViews,omitempty"`
		Contacts      int32  `json:"contacts,omitempty"`
		UniqContacts  int32  `json:"uniqContacts,omitempty"`
		Favorites     int32  `json:"favorites,omitempty"`
		UniqFavorites int32  `json:"uniqFavorites,omitempty"`
	} `json:"stats"`
}
