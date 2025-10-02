package api

import "time"

// Константы таймаутов и базового URL для Avito API.
const (
	// _defaultTimeoutHTTP - таймаут для обычных HTTP-запросов SDK
	_defaultTimeoutHTTP = 30 * time.Second
	// _defaultTimeoutToken - таймаут для запроса токена OAuth
	_defaultTimeoutToken = 10 * time.Second
	// baseURL - базовый URL Avito API
	baseURL = "https://api.avito.ru"
)
