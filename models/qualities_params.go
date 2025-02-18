package models

import (
	"strconv"
	"strings"
)

// QualitiesParams описывает параметры запроса для эндпоинта /qualities/v2 API Kodik.
// Если значение поля является пустым или нулевым, оно не включается в итоговый запрос.
type QualitiesParams struct {
	// Основные параметры
	Types           string `json:"types,omitempty"`            // Типы материалов (например, "anime-serial,foreign-movie,..." )
	Year            string `json:"year,omitempty"`             // Год или диапазон (например, "2020", "2010-2020")
	TranslationID   int    `json:"translation_id,omitempty"`   // ID озвучки
	TranslationType string `json:"translation_type,omitempty"` // Тип перевода: voice или subtitles
	HasField        string `json:"has_field,omitempty"`        // Наличие определённых полей (например, kinopoisk_id)
	Lgbt            *bool  `json:"lgbt,omitempty"`             // Фильтр по LGBT контенту (true/false)
	Sort            string `json:"sort,omitempty"`             // Сортировка результатов: title или count

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Страны через запятую (например, "USA,Russia")
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность или диапазон (например, "90", "80-120")
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (например, "7.0", "6.5-8.2")
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList

	// Персоналии (список имён через запятую)
	Actors    string `json:"actors,omitempty"`
	Directors string `json:"directors,omitempty"`
	Producers string `json:"producers,omitempty"`
	Writers   string `json:"writers,omitempty"`
	Composers string `json:"composers,omitempty"`
	Editors   string `json:"editors,omitempty"`
	Designers string `json:"designers,omitempty"`
	Operators string `json:"operators,omitempty"`

	// Другие параметры
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA (например, "PG-13")
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст (например, "16", "12-16")
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме (tv, movie, ova, ona и т.д.)
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме (anons, ongoing, released, etc.)
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Лицензионные правообладатели через запятую
}

// ToMap преобразует структуру QualitiesParams в карту параметров для HTTP-запроса.
func (qp *QualitiesParams) ToMap() map[string]string {
	params := make(map[string]string)

	if qp.Types != "" {
		params["types"] = qp.Types
	}
	if qp.Year != "" {
		params["year"] = qp.Year
	}
	if qp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(qp.TranslationID)
	}
	if qp.TranslationType != "" {
		params["translation_type"] = qp.TranslationType
	}
	if qp.HasField != "" {
		params["has_field"] = qp.HasField
	}
	if qp.Lgbt != nil {
		if *qp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if qp.Sort != "" {
		params["sort"] = qp.Sort
	}

	// Фильтрация по внешним данным
	if qp.Countries != "" {
		params["countries"] = strings.ReplaceAll(qp.Countries, " ", "")
	}
	if qp.Genres != "" {
		params["genres"] = qp.Genres
	}
	if qp.AnimeGenres != "" {
		params["anime_genres"] = qp.AnimeGenres
	}
	if qp.DramaGenres != "" {
		params["drama_genres"] = qp.DramaGenres
	}
	if qp.AllGenres != "" {
		params["all_genres"] = qp.AllGenres
	}
	if qp.Duration != "" {
		params["duration"] = qp.Duration
	}
	if qp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = qp.KinopoiskRating
	}
	if qp.ImdbRating != "" {
		params["imdb_rating"] = qp.ImdbRating
	}
	if qp.ShikimoriRating != "" {
		params["shikimori_rating"] = qp.ShikimoriRating
	}
	if qp.MydramalistRating != "" {
		params["mydramalist_rating"] = qp.MydramalistRating
	}

	// Персоналии
	if qp.Actors != "" {
		params["actors"] = qp.Actors
	}
	if qp.Directors != "" {
		params["directors"] = qp.Directors
	}
	if qp.Producers != "" {
		params["producers"] = qp.Producers
	}
	if qp.Writers != "" {
		params["writers"] = qp.Writers
	}
	if qp.Composers != "" {
		params["composers"] = qp.Composers
	}
	if qp.Editors != "" {
		params["editors"] = qp.Editors
	}
	if qp.Designers != "" {
		params["designers"] = qp.Designers
	}
	if qp.Operators != "" {
		params["operators"] = qp.Operators
	}

	// Другие параметры
	if qp.RatingMPAA != "" {
		params["rating_mpaa"] = qp.RatingMPAA
	}
	if qp.MinimalAge != "" {
		params["minimal_age"] = qp.MinimalAge
	}
	if qp.AnimeKind != "" {
		params["anime_kind"] = qp.AnimeKind
	}
	if qp.MydramalistTags != "" {
		params["mydramalist_tags"] = qp.MydramalistTags
	}
	if qp.AnimeStatus != "" {
		params["anime_status"] = qp.AnimeStatus
	}
	if qp.DramaStatus != "" {
		params["drama_status"] = qp.DramaStatus
	}
	if qp.AllStatus != "" {
		params["all_status"] = qp.AllStatus
	}
	if qp.AnimeStudios != "" {
		params["anime_studios"] = qp.AnimeStudios
	}
	if qp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = qp.AnimeLicensedBy
	}

	return params
}
