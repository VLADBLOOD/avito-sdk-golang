package entities

// CallsByPeriodRequest - фильтр запроса списка звонков по времени
type CallsByPeriodRequest struct {
	DateTimeFrom string `json:"dateTimeFrom"`
	DateTimeTo   string `json:"dateTimeTo,omitempty"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
}

// CallsByPeriodResponse - ответ API для списка звонков
type CallsByPeriodResponse struct {
	Calls []*CallRecord `json:"calls,omitempty"`
	Error *CallError    `json:"error,omitempty"`
}

// CallRecord - запись о звонке
type CallRecord struct {
	CallID          int64  `json:"callId"`
	BuyerPhone      string `json:"buyerPhone"`
	SellerPhone     string `json:"sellerPhone"`
	VirtualPhone    string `json:"virtualPhone"`
	ItemID          int64  `json:"itemId"`
	CallTime        string `json:"callTime"`
	TalkDuration    int    `json:"talkDuration"`
	WaitingDuration int    `json:"waitingDuration"`
}

// CallError - ошибка в ответе
type CallError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
