package entities

// AdsInfoRequest - фильтр запроса информации по объявлениям
type AdsInfoRequest struct {
	PerPage       int32    `url:"per_page,omitempty"`
	Page          int32    `url:"page,omitempty"`
	Status        AdStatus `url:"status,omitempty"`
	UpdatedAtFrom string   `url:"updatedAtFrom,omitempty"`
	Category      int32    `url:"category,omitempty"`
}

// AdStatus - статусы объявления
type AdStatus string

const (
	AdStatusActive   AdStatus = "active"
	AdStatusRemoved  AdStatus = "removed"
	AdStatusOld      AdStatus = "old"
	AdStatusBlocked  AdStatus = "blocked"
	AdStatusRejected AdStatus = "rejected"
)

// AdsInfoResponse - ответ API для списка объявлений
type AdsInfoResponse struct {
	Meta      AdsInfoMeta `json:"meta"`
	Resources []*AdInfo   `json:"resources"`
}

// AdsInfoMeta - метаданные ответа для списка объявлений
type AdsInfoMeta struct {
	Page    int32  `json:"page"`
	PerPage int32  `json:"per_page"`
	Pages   *int32 `json:"pages,omitempty"`
	Total   *int32 `json:"total,omitempty"`
}

// AdInfo - информация об одном объявлении
type AdInfo struct {
	ID       int64       `json:"id"`
	Title    string      `json:"title"`
	Price    *float64    `json:"price,omitempty"`
	Status   AdStatus    `json:"status"`
	URL      string      `json:"url"`
	Address  string      `json:"address,omitempty"`
	Category *AdCategory `json:"category,omitempty"`
}

// AdCategory - категория объявления
type AdCategory struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// AdsStatsRequest - фильтр запроса статистики по объявлениям
type AdsStatsRequest struct {
	DateFrom       string         `json:"dateFrom"`
	DateTo         string         `json:"dateTo"`
	Fields         []AdStatsField `json:"fields,omitempty"`
	ItemIds        []int64        `json:"itemIds"`
	PeriodGrouping *AdStatsPeriod `json:"periodGrouping,omitempty"`
}

// AdStatsField - поля статистики
type AdStatsField string

const (
	AdStatsFieldViews         AdStatsField = "views"
	AdStatsFieldUniqViews     AdStatsField = "uniqViews"
	AdStatsFieldContacts      AdStatsField = "contacts"
	AdStatsFieldUniqContacts  AdStatsField = "uniqContacts"
	AdStatsFieldFavorites     AdStatsField = "favorites"
	AdStatsFieldUniqFavorites AdStatsField = "uniqFavorites"
)

// AdStatsPeriod - период группировки
type AdStatsPeriod string

const (
	AdStatsPeriodDay   AdStatsPeriod = "day"
	AdStatsPeriodWeek  AdStatsPeriod = "week"
	AdStatsPeriodMonth AdStatsPeriod = "month"
)

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
	ItemID int64        `json:"itemId"`
	Stats  []AdStatData `json:"stats"`
}

// AdStatData - данные статистики за период
type AdStatData struct {
	Date          string `json:"date,omitempty"`
	Views         *int32 `json:"views,omitempty"`
	UniqViews     *int32 `json:"uniqViews,omitempty"`
	Contacts      *int32 `json:"contacts,omitempty"`
	UniqContacts  *int32 `json:"uniqContacts,omitempty"`
	Favorites     *int32 `json:"favorites,omitempty"`
	UniqFavorites *int32 `json:"uniqFavorites,omitempty"`
}
