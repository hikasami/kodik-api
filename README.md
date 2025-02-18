# Kodik API Go Client Library

Библиотека для работы с [Kodik API](https://kodikapi.com) на языке Go. Данная библиотека позволяет выполнять запросы к API с использованием структурированных параметров.

## Структура проекта

```
├── api
│   └── search.go         # Функции для работы с эндпоинтами API (пример – /search)
├── client
│   └── client.go         # HTTP-клиент, инициализация и функция выполнения запросов
├── errors
│   └── errors.go         # Пользовательские ошибки библиотеки
├── models
│   ├── search.go         # Модели для ответов API
│   └── search_params.go  # Структура параметров запроса для /search
└── main.go               # Пример использования библиотеки
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

Перед использованием функций API необходимо инициализировать клиента с вашим API токеном. Это можно сделать с помощью функции `client.Init`.

```go
package main

import (
	"github.com/hikasami/kodik-api/client"
)

func main() {
    // Инициализация клиента с указанием токена и использованием HTTPS протокола
    client.Init("YOUR_API_TOKEN_HERE", true)
    // Далее можно использовать api.{method} для вызова API
}
```

## Пример запроса `/search`

Ниже приведён пример использования запроса `/search` с использованием структуры параметров `SearchParams`. Токен задается глобально при инициализации клиента, поэтому он не передается в структуру параметров.

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
	// Инициализация глобального клиента API с использованием HTTPS.
	client.Init("YOUR_API_TOKEN_HERE", true)

	// Заполнение параметров запроса /search.
	searchParams := &models.SearchParams{
		Title:     "Inception", // Поиск по названию фильма
		Strict:    false,
		FullMatch: false,
		Limit:     5,           // Ограничение на количество результатов
		// Можно установить дополнительные параметры:
		// TitleOrig, TranslationType, Genres, Year, и др.
	}

	// Выполнение запроса /search к API Kodik.
	searchResponse, err := api.Search(searchParams)
	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса /search: %v", err)
	}

	// Вывод результатов запроса.
	fmt.Printf("Найдено материалов: %d\n", len(searchResponse.Results))
	if len(searchResponse.Results) > 0 {
		first := searchResponse.Results[0]
		fmt.Println("Пример первого найденного материала:")
		fmt.Printf("ID: %s\nНазвание: %s\nОригинальное название: %s\nГод: %d\nКачество: %s\n",
			first.ID, first.Title, first.TitleOrig, first.Year, first.Quality)
	} else {
		fmt.Println("Материалы не найдены.")
	}
}
```

## Дополнительные возможности

Библиотека разработана по модульной архитектуре, что позволяет легко расширять функциональность. Помимо `/search`, будет добавлена поддержка других эндпоинтов, таких как:

- **`/list`**
- **`/qualities`**
- **`/translations/v2`**
- **`/countries`**, **`/genres`**, **`/years`**

## Обработка ошибок

Все ошибки, возникающие при выполнении запросов, возвращаются функциями API в виде стандартного типа `error`. Пользовательские ошибки, например, `ErrUnexpectedStatus` или `ErrUnsupportedMethod`, определены в пакете `errors`.