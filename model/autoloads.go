package model

import "time"

// AutoloadStatus - статус отчета автозагрузки
type AutoloadStatus string

const (
	AutoloadStatusProcessing AutoloadStatus = "processing"
	AutoloadStatusSuccess    AutoloadStatus = "success"
	AutoloadStatusFailed     AutoloadStatus = "failed"
)

// AutoloadSource - источник автозагрузки
type AutoloadSource string

const (
	AutoloadSourceEmail    AutoloadSource = "email"
	AutoloadSourceFile     AutoloadSource = "file"
	AutoloadSourceAPI      AutoloadSource = "api"
	AutoloadSourceExternal AutoloadSource = "external"
)

// AutoloadsListV2Request - фильтр запроса списка отчетов автозагрузки
type AutoloadsListV2Request struct {
	PerPage  int       `json:"per_page"`
	Page     int       `json:"page"`
	DateFrom time.Time `json:"date_from"`
	DateTo   time.Time `json:"date_to"`
}

// AutoloadsListV2Response - ответ API для списка отчетов автозагрузки
type AutoloadsListV2Response struct {
	Meta struct {
		Page    int `json:"page"`
		Pages   int `json:"pages"`
		PerPage int `json:"per_page"`
		Total   int `json:"total"`
	} `json:"meta"`
	Reports []struct {
		ID         int64          `json:"id,omitempty"`
		Status     AutoloadStatus `json:"status,omitempty"`
		StartedAt  string         `json:"started_at,omitempty"`
		FinishedAt string         `json:"finished_at,omitempty"`
	} `json:"reports"`
}

// AutoloadAdsListV2Request - фильтр запроса объявлений из выгрузки
type AutoloadAdsListV2Request struct {
	PerPage  int    `json:"per_page,omitempty"`
	Page     int    `json:"page,omitempty"`
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
	AdID       string `json:"ad_id"`
	AppliedVas []struct {
		Price float64 `json:"price"`
		Slug  string  `json:"slug"`
		Title string  `json:"title"`
	} `json:"applied_vas,omitempty"`
	AvitoDateEnd string `json:"avito_date_end,omitempty"`
	AvitoID      int64  `json:"avito_id,omitempty"`
	AvitoStatus  string `json:"avito_status,omitempty"`
	FeedName     string `json:"feed_name,omitempty"`
	Messages     []struct {
		Code        int       `json:"code"`
		Description string    `json:"description"`
		Title       string    `json:"title"`
		Type        string    `json:"type"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"messages,omitempty"`
	Section struct {
		Slug  string `json:"slug"`
		Title string `json:"title"`
	} `json:"section,omitempty"`
	URL string `json:"url,omitempty"`
}

// AutoloadStatsV3Response - ответ API со статистикой по ID выгрузки
type AutoloadStatsV3Response struct {
	ReportID   int64          `json:"report_id"`
	Status     AutoloadStatus `json:"status"`
	Source     AutoloadSource `json:"source"`
	StartedAt  string         `json:"started_at"`
	FinishedAt string         `json:"finished_at,omitempty"`
	Events     []struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
		Type        string `json:"type"`
	} `json:"events,omitempty"`
	FeedsUrls []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"feeds_urls,omitempty"`
	ListingFees struct {
		Packages []struct {
			Count     int `json:"count"`
			PackageID int `json:"package_id"`
		} `json:"packages,omitempty"`
		Single struct {
			Count       int     `json:"count"`
			TotalAmount float64 `json:"total_amount"`
		} `json:"single,omitempty"`
	} `json:"listing_fees,omitempty"`
	SectionStats struct {
		Count    int           `json:"count"`
		Slug     string        `json:"slug"`
		Title    string        `json:"title"`
		Sections []SectionItem `json:"sections,omitempty"`
	} `json:"section_stats,omitempty"`
}

// SectionItem - элемент статистики раздела
type SectionItem struct {
	Count    int    `json:"count"`
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	Sections []struct {
		Count int    `json:"count"`
		Slug  string `json:"slug"`
		Title string `json:"title"`
	} `json:"sections,omitempty"`
}
