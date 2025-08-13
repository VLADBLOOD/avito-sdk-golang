package entities

import "time"

// AutoloadsListV2Request - фильтр запроса списка отчетов автозагрузки
type AutoloadsListV2Request struct {
	PerPage  *int       `json:"per_page,omitempty"`
	Page     *int       `json:"page,omitempty"`
	DateFrom *time.Time `json:"date_from,omitempty"`
	DateTo   *time.Time `json:"date_to,omitempty"`
}

// AutoloadsListV2Response - ответ API для списка отчетов автозагрузки
type AutoloadsListV2Response struct {
	Meta    AutoloadsListMeta `json:"meta"`
	Reports []*AutoloadReport `json:"reports"`
}

// AutoloadsListMeta - метаданные ответа для списка отчетов
type AutoloadsListMeta struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// AutoloadReport - отчет автозагрузки
type AutoloadReport struct {
	ID         int64          `json:"id"`
	Status     AutoloadStatus `json:"status"`
	StartedAt  *time.Time     `json:"started_at"`
	FinishedAt *NullableTime  `json:"finished_at,omitempty"`
}

// NullableTime - кастомный тип для времени, поддерживающий пустые строки
type NullableTime struct {
	time.Time
}

// Переопределяем метод для парсинга пустых строк
func (m *NullableTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	tt, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*m = NullableTime{tt}
	return err
}

// AutoloadStatus - статус отчета автозагрузки
type AutoloadStatus string

const (
	AutoloadStatusProcessing AutoloadStatus = "processing"
	AutoloadStatusSuccess    AutoloadStatus = "success"
	AutoloadStatusFailed     AutoloadStatus = "failed"
)

// AutoloadAdsListV2Request - фильтр запроса объявлений из выгрузки
type AutoloadAdsListV2Request struct {
	ReportID string `json:"-"`
	PerPage  *int   `json:"per_page,omitempty"`
	Page     *int   `json:"page,omitempty"`
	Query    string `json:"query,omitempty"`
	Sections string `json:"sections,omitempty"`
}

// AutoloadAdsListV2Response - ответ API для объявлений из выгрузки
type AutoloadAdsListV2Response struct {
	Items    []*AutoloadAd   `json:"items"`
	Meta     AutoloadAdsMeta `json:"meta"`
	ReportID int64           `json:"report_id"`
}

// AutoloadAdsMeta - метаданные ответа для объявлений
type AutoloadAdsMeta struct {
	Page    int `json:"page"`
	Pages   int `json:"pages"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// AutoloadAd - объявление из автозагрузки
type AutoloadAd struct {
	AdID         string            `json:"ad_id"`
	AppliedVas   []AppliedVAS      `json:"applied_vas,omitempty"`
	AvitoDateEnd string            `json:"avito_date_end,omitempty"`
	AvitoID      int64             `json:"avito_id,omitempty"`
	AvitoStatus  string            `json:"avito_status,omitempty"`
	FeedName     string            `json:"feed_name,omitempty"`
	Messages     []AutoloadMessage `json:"messages,omitempty"`
	Section      *AdSection        `json:"section,omitempty"`
	URL          string            `json:"url,omitempty"`
}

// AppliedVAS - примененные услуги
type AppliedVAS struct {
	Price float64 `json:"price"`
	Slug  string  `json:"slug"`
	Title string  `json:"title"`
}

// AutoloadMessage - сообщение об ошибке или предупреждении
type AutoloadMessage struct {
	Code        int       `json:"code"`
	Description string    `json:"description"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// AdSection - раздел объявления
type AdSection struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

// AutoloadStatsV3Request - фильтр запроса статистики по ID выгрузки
type AutoloadStatsV3Request struct {
	ReportID string `json:"-"`
}

// AutoloadStatsV3Response - ответ API со статистикой по ID выгрузки
type AutoloadStatsV3Response struct {
	ReportID     int64           `json:"report_id"`
	Status       AutoloadStatus  `json:"status"`
	Source       AutoloadSource  `json:"source"`
	StartedAt    time.Time       `json:"started_at"`
	FinishedAt   NullableTime    `json:"finished_at,omitempty"`
	Events       []AutoloadEvent `json:"events,omitempty"`
	FeedsUrls    []FeedURL       `json:"feeds_urls,omitempty"`
	ListingFees  *ListingFees    `json:"listing_fees,omitempty"`
	SectionStats *SectionStats   `json:"section_stats,omitempty"`
}

// AutoloadSource - источник автозагрузки
type AutoloadSource string

const (
	AutoloadSourceEmail    AutoloadSource = "email"
	AutoloadSourceFile     AutoloadSource = "file"
	AutoloadSourceAPI      AutoloadSource = "api"
	AutoloadSourceExternal AutoloadSource = "external"
)

// AutoloadEvent - событие в процессе автозагрузки
type AutoloadEvent struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// FeedURL - URL фида
type FeedURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// ListingFees - комиссии за размещение
type ListingFees struct {
	Packages []FeePackage `json:"packages,omitempty"`
	Single   *SingleFee   `json:"single,omitempty"`
}

// FeePackage - пакетная комиссия
type FeePackage struct {
	Count     int `json:"count"`
	PackageID int `json:"package_id"`
}

// SingleFee - одиночная комиссия
type SingleFee struct {
	Count       int     `json:"count"`
	TotalAmount float64 `json:"total_amount"`
}

// SectionStats - статистика по разделам
type SectionStats struct {
	Count    int           `json:"count"`
	Slug     string        `json:"slug"`
	Title    string        `json:"title"`
	Sections []SectionItem `json:"sections,omitempty"`
}

// SectionItem - элемент статистики раздела
type SectionItem struct {
	Count    int          `json:"count"`
	Slug     string       `json:"slug"`
	Title    string       `json:"title"`
	Sections []SubSection `json:"sections,omitempty"`
}

// SubSection - подраздел
type SubSection struct {
	Count int    `json:"count"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
