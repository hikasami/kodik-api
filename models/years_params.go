package models

import (
	"strconv"
	"strings"
)

// YearsParams описывает параметры запроса для эндпоинта /years API Kodik.
type YearsParams struct {
	// Фильтрация материалов
	Types             string `json:"types,omitempty"`              // Фильтрация по типу материала (например, foreign-movie, cartoon-serial и т.д.)
	TranslationID     int    `json:"translation_id,omitempty"`     // Фильтр по ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // IDs для исключения (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода: voice или subtitles
	HasField          string `json:"has_field,omitempty"`          // Наличие определённых полей (например, kinopoisk_id, imdb_id, mdl_id, worldart_link, shikimori_id)
	Lgbt              *bool  `json:"lgbt,omitempty"`               // Фильтрация по LGBT контенту (true/false)
	Sort              string `json:"sort,omitempty"`               // Сортировка: "year" или "count"

	// Внешняя фильтрация
	Countries         string `json:"countries,omitempty"`          // Фильтр по странам (через запятую)
	Genres            string `json:"genres,omitempty"`             // Фильтр по жанрам (через запятую)
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Фильтр по жанрам для аниме
	DramaGenres       string `json:"drama_genres,omitempty"`       // Фильтр по жанрам для драм
	AllGenres         string `json:"all_genres,omitempty"`         // Фильтр по всем жанрам
	Duration          string `json:"duration,omitempty"`           // Фильтр по длительности (в минутах, точное значение или диапазон, например, 30 или 40-80)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (точное значение или диапазон)
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг (точное значение или диапазон)
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori (точное значение или диапазон)
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList (точное значение или диапазон)
	Actors            string `json:"actors,omitempty"`             // Фильтр по актёрам (через запятую)
	Directors         string `json:"directors,omitempty"`          // Фильтр по режиссёрам (через запятую)
	Producers         string `json:"producers,omitempty"`          // Фильтр по продюсерам (через запятую)
	Writers           string `json:"writers,omitempty"`            // Фильтр по сценаристам (через запятую)
	Composers         string `json:"composers,omitempty"`          // Фильтр по композиторам (через запятую)
	Editors           string `json:"editors,omitempty"`            // Фильтр по монтажёрам (через запятую)
	Designers         string `json:"designers,omitempty"`          // Фильтр по дизайнерам (через запятую)
	Operators         string `json:"operators,omitempty"`          // Фильтр по операторам (через запятую)
	RatingMPAA        string `json:"rating_mpaa,omitempty"`        // Рейтинг MPAA (например, G, PG, PG-13, R и т.д.)
	MinimalAge        string `json:"minimal_age,omitempty"`        // Минимальный возраст (например, 16 или 12-16)
	AnimeKind         string `json:"anime_kind,omitempty"`         // Тип аниме (tv, movie, ova, ona и т.д.)
	MydramalistTags   string `json:"mydramalist_tags,omitempty"`   // Теги MyDramaList через запятую
	AnimeStatus       string `json:"anime_status,omitempty"`       // Статус аниме (anons, ongoing, released и т.д.)
	DramaStatus       string `json:"drama_status,omitempty"`       // Статус драмы (anons, ongoing, released и т.д.)
	AllStatus         string `json:"all_status,omitempty"`         // Универсальный статус (например, anons, ongoing, released)
	AnimeStudios      string `json:"anime_studios,omitempty"`      // Студии для аниме (через запятую, например, J.C.Staff, Studio Hibari)
	AnimeLicensedBy   string `json:"anime_licensed_by,omitempty"`  // Правообладатели для аниме (через запятую, например, Wakanim, Russian Reportage)
}

// ToMap преобразует структуру YearsParams в карту параметров для HTTP-запроса.
func (yp *YearsParams) ToMap() map[string]string {
	params := make(map[string]string)

	if yp.Types != "" {
		params["types"] = yp.Types
	}
	if yp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(yp.TranslationID)
	}
	if yp.BlockTranslations != "" {
		params["block_translations"] = strings.ReplaceAll(yp.BlockTranslations, " ", "")
	}
	if yp.TranslationType != "" {
		params["translation_type"] = yp.TranslationType
	}
	if yp.HasField != "" {
		params["has_field"] = yp.HasField
	}
	if yp.Lgbt != nil {
		if *yp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if yp.Sort != "" {
		params["sort"] = yp.Sort
	}
	if yp.Countries != "" {
		params["countries"] = yp.Countries
	}
	if yp.Genres != "" {
		params["genres"] = yp.Genres
	}
	if yp.AnimeGenres != "" {
		params["anime_genres"] = yp.AnimeGenres
	}
	if yp.DramaGenres != "" {
		params["drama_genres"] = yp.DramaGenres
	}
	if yp.AllGenres != "" {
		params["all_genres"] = yp.AllGenres
	}
	if yp.Duration != "" {
		params["duration"] = yp.Duration
	}
	if yp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = yp.KinopoiskRating
	}
	if yp.ImdbRating != "" {
		params["imdb_rating"] = yp.ImdbRating
	}
	if yp.ShikimoriRating != "" {
		params["shikimori_rating"] = yp.ShikimoriRating
	}
	if yp.MydramalistRating != "" {
		params["mydramalist_rating"] = yp.MydramalistRating
	}
	if yp.Actors != "" {
		params["actors"] = yp.Actors
	}
	if yp.Directors != "" {
		params["directors"] = yp.Directors
	}
	if yp.Producers != "" {
		params["producers"] = yp.Producers
	}
	if yp.Writers != "" {
		params["writers"] = yp.Writers
	}
	if yp.Composers != "" {
		params["composers"] = yp.Composers
	}
	if yp.Editors != "" {
		params["editors"] = yp.Editors
	}
	if yp.Designers != "" {
		params["designers"] = yp.Designers
	}
	if yp.Operators != "" {
		params["operators"] = yp.Operators
	}
	if yp.RatingMPAA != "" {
		params["rating_mpaa"] = yp.RatingMPAA
	}
	if yp.MinimalAge != "" {
		params["minimal_age"] = yp.MinimalAge
	}
	if yp.AnimeKind != "" {
		params["anime_kind"] = yp.AnimeKind
	}
	if yp.MydramalistTags != "" {
		params["mydramalist_tags"] = yp.MydramalistTags
	}
	if yp.AnimeStatus != "" {
		params["anime_status"] = yp.AnimeStatus
	}
	if yp.DramaStatus != "" {
		params["drama_status"] = yp.DramaStatus
	}
	if yp.AllStatus != "" {
		params["all_status"] = yp.AllStatus
	}
	if yp.AnimeStudios != "" {
		params["anime_studios"] = yp.AnimeStudios
	}
	if yp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = yp.AnimeLicensedBy
	}

	return params
}
