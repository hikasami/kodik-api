# Kodik API Go Client Library

Библиотека для работы с [Kodik API](https://kodikapi.com) на языке Go. Данная библиотека позволяет выполнять запросы к API с использованием структурированных параметров для различных эндпоинтов, таких как `/search`, `/list`, `/translations/v2`, `/qualities/v2`, `/countries`, `/genres` и `/years`.

## Структура проекта

```
├── api
│   ├── search.go                # Функции для работы с эндпоинтами (/search, /list, /translations/v2, /qualities/v2, /countries, /genres, /years)
├── client
│   └── client.go                # HTTP-клиент, инициализация и функция выполнения запросов
├── errors
│   └── errors.go                # Пользовательские ошибки библиотеки
├── models
│   ├── search.go                # Модели для ответов API (/search)
│   ├── search_params.go         # Структура параметров запроса для /search
│   ├── ...
```

## Установка

1. **Клонирование репозитория:**

   ```bash
   git clone https://github.com/hikasami/kodik-api
   cd kodik-api
   ```

2. **Сборка проекта:**

   Для сборки выполните в терминале:
   ```bash
   go build
   ```

## Быстрый старт

Перед использованием функций API необходимо инициализировать клиента с вашим API токеном. Для этого используйте функцию `client.Init`.

```go
package main

import (
	"github.com/hikasami/kodik-api/client"
)

func main() {
	// Инициализация клиента с указанием токена и использованием HTTPS
	client.Init("YOUR_API_TOKEN_HERE", true)
	// Теперь можно использовать различные методы API
}
```

## Пример запроса `/search`

Ниже приведён пример использования запроса `/search` с использованием структуры параметров `SearchParams`. Токен задается глобально при инициализации клиента.

```go
package main

import (
	"fmt"
	"log"

	"github.com/hikasami/kodik-api/api"
	"github.com/hikasami/kodik-api/client"
	"github.com/hikasami/kodik-api/models"
)

func main() {
	// Инициализация клиента API
	client.Init("YOUR_API_TOKEN_HERE", true)

	// Заполнение параметров запроса /search
	searchParams := &models.SearchParams{
		Title:     "Inception",
		Strict:    false,
		FullMatch: false,
		Limit:     5,
		// Дополнительные параметры можно устанавливать по необходимости
	}

	// Выполнение запроса /search к API Kodik
	searchResponse, err := api.Search(searchParams)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса /search: %v", err)
	}

	// Вывод результатов запроса
	fmt.Printf("Найдено материалов: %d\n", len(searchResponse.Results))
}
```

## Дополнительные методы API

Библиотека поддерживает следующие эндпоинты:

- **`/list`**: Получение списка материалов с заданными параметрами.
- **`/translations/v2`**: Получение данных по озвучкам.
- **`/qualities/v2`**: Получение информации по качествам видео.
- **`/countries`**: Получение списка стран с количеством материалов.
- **`/genres`**: Получение списка жанров с количеством материалов.
- **`/years`**: Получение списка годов с количеством материалов.

Каждый эндпоинт имеет отдельные структуры параметров и моделей ответа.

## Обработка ошибок

Все ошибки, возникающие при выполнении запросов, возвращаются функциями API в виде стандартного объекта `error`. Пользовательские ошибки, например, `ErrUnexpectedStatus` или `ErrUnsupportedMethod`, определены в пакете `errors`.