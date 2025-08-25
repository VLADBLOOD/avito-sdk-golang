# Avito SDK для Go

Легкий и типобезопасный Go SDK для работы с Avito API (объявления, статистика и др.).

## Что внутри
- Упрощенная аутентификация (OAuth2 Client Credentials)
- Типизированные модели запросов/ответов
- Удобный клиент с сервисами (на текущий момент: Ads)
- Встроенные таймауты и обработка ошибок

> В проекте заложена структура для расширения (Autoloads, Messenger, CallTracking и пр.),
> но из коробки реализованы методы по объявлениям (Ads).

## Установка

```bash
go get github.com/eserg-key/avito-sdk-golang
```

Go-модуль проекта указан как `avito-sdk`, поэтому в импортах используется:

```go
import (
    "avito-sdk/client"
    "avito-sdk/model"
)
```

## Быстрый старт

```go
package main

import (
    "avito-sdk/client"
    "avito-sdk/model"
    "context"
    "fmt"
)

func main() {
    cfg := &client.Config{ClientID: "<account_id>", ClientSecret: "<client_secret>"}
    c, err := client.NewClient(cfg)
    if err != nil { panic(err) }

    // Получить список объявлений
    ctx := context.Background()
    ads, err := c.ADS.GetAdsInfo(ctx, &model.AdsInfoRequest{PerPage: 10, Page: 1})
    if err != nil { panic(err) }

    fmt.Println("Всего объявлений:", len(ads.Resources))
}
```

## Поддерживаемые методы

### Ads (объявления)
- GetAdsInfo — получение информации по объявлениям
- GetAdsStats — получение статистики по объявлениям

## Требования
- Go 1.21+ (рекомендуется последняя минорная версия)

## Дизайн и оговорки
- HttpClient.request теперь принимает указатель на структуру ответа (out), поэтому отпадают небезопасные приведения типов.
- Базовый URL Avito API задается в константах (`https://api.avito.ru`).
- Примеры в файле `examples.go` помечены build-тегом `//go:build ignore` и не компилируются вместе с библиотекой.

## Дорожная карта
- [ ] Централизованный обработчик кодов ошибок и тела ошибок API
- [ ] Автоматическое обновление токена перед запросами
- [ ] Поддержка дополнительных сервисов (Messenger, CallTracking, Autoloads)
- [ ] Юнит-тесты основных сервисов

## Лицензия
MIT
