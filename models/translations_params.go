package models

import (
	"strings"
)

// TranslationsParams описывает параметры запроса для эндпоинта /translations/v2 API Kodik.
// Если значение поля является пустым, оно не включается в итоговый запрос.
type TranslationsParams struct {
	// Основные параметры
	Types           string `json:"types,omitempty"`            // Типы материалов, например: anime-serial, foreign-movie, etc.
	Year            string `json:"year,omitempty"`             // Фильтрация по году
	TranslationType string `json:"translation_type,omitempty"` // Тип перевода: voice или subtitles
	HasField        string `json:"has_field,omitempty"`        // Наличие определённых полей (например, kinopoisk_id, imdb_id и т.д.)
	Lgbt            *bool  `json:"lgbt,omitempty"`             // Фильтр по наличию LGBT контента: true/false
	Sort            string `json:"sort,omitempty"`             // Сортировка результатов: title или count

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Страны через запятую (например: USA,Russia)
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую (например: action,drama)
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность или диапазон длительности (в минутах)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска или диапазон
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг или диапазон
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori или диапазон
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList или диапазон

	// Персоналии
	Actors    string `json:"actors,omitempty"`    // Список актёров через запятую
	Directors string `json:"directors,omitempty"` // Список режиссёров через запятую
	Producers string `json:"producers,omitempty"` // Список продюсеров через запятую
	Writers   string `json:"writers,omitempty"`   // Список сценаристов через запятую
	Composers string `json:"composers,omitempty"` // Список композиторов через запятую
	Editors   string `json:"editors,omitempty"`   // Список монтажёров через запятую
	Designers string `json:"designers,omitempty"` // Список дизайнеров через запятую
	Operators string `json:"operators,omitempty"` // Список операторов через запятую

	// Другие параметры
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA (например, PG, PG-13, R)
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст (например, 16 или 12-16)
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме: tv, movie, ova, ona и т.д.
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме: anons, ongoing, released и т.д.
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Лицензионные правообладатели через запятую
}

// ToMap преобразует структуру TranslationsParams в карту параметров для HTTP-запроса.
func (tp *TranslationsParams) ToMap() map[string]string {
	params := make(map[string]string)

	if tp.Types != "" {
		params["types"] = tp.Types
	}
	if tp.Year != "" {
		params["year"] = tp.Year
	}
	if tp.TranslationType != "" {
		params["translation_type"] = tp.TranslationType
	}
	if tp.HasField != "" {
		params["has_field"] = tp.HasField
	}
	if tp.Lgbt != nil {
		if *tp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if tp.Sort != "" {
		params["sort"] = tp.Sort
	}

	// Внешняя фильтрация
	if tp.Countries != "" {
		// Убираем пробелы при необходимости
		params["countries"] = strings.ReplaceAll(tp.Countries, " ", "")
	}
	if tp.Genres != "" {
		params["genres"] = tp.Genres
	}
	if tp.AnimeGenres != "" {
		params["anime_genres"] = tp.AnimeGenres
	}
	if tp.DramaGenres != "" {
		params["drama_genres"] = tp.DramaGenres
	}
	if tp.AllGenres != "" {
		params["all_genres"] = tp.AllGenres
	}
	if tp.Duration != "" {
		params["duration"] = tp.Duration
	}
	if tp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = tp.KinopoiskRating
	}
	if tp.ImdbRating != "" {
		params["imdb_rating"] = tp.ImdbRating
	}
	if tp.ShikimoriRating != "" {
		params["shikimori_rating"] = tp.ShikimoriRating
	}
	if tp.MydramalistRating != "" {
		params["mydramalist_rating"] = tp.MydramalistRating
	}

	// Персоналии
	if tp.Actors != "" {
		params["actors"] = tp.Actors
	}
	if tp.Directors != "" {
		params["directors"] = tp.Directors
	}
	if tp.Producers != "" {
		params["producers"] = tp.Producers
	}
	if tp.Writers != "" {
		params["writers"] = tp.Writers
	}
	if tp.Composers != "" {
		params["composers"] = tp.Composers
	}
	if tp.Editors != "" {
		params["editors"] = tp.Editors
	}
	if tp.Designers != "" {
		params["designers"] = tp.Designers
	}
	if tp.Operators != "" {
		params["operators"] = tp.Operators
	}

	// Другие параметры
	if tp.RatingMPAA != "" {
		params["rating_mpaa"] = tp.RatingMPAA
	}
	if tp.MinimalAge != "" {
		params["minimal_age"] = tp.MinimalAge
	}
	if tp.AnimeKind != "" {
		params["anime_kind"] = tp.AnimeKind
	}
	if tp.MydramalistTags != "" {
		params["mydramalist_tags"] = tp.MydramalistTags
	}
	if tp.AnimeStatus != "" {
		params["anime_status"] = tp.AnimeStatus
	}
	if tp.DramaStatus != "" {
		params["drama_status"] = tp.DramaStatus
	}
	if tp.AllStatus != "" {
		params["all_status"] = tp.AllStatus
	}
	if tp.AnimeStudios != "" {
		params["anime_studios"] = tp.AnimeStudios
	}
	if tp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = tp.AnimeLicensedBy
	}

	return params
}
