package models

import (
	"strconv"
	"strings"
)

// CountriesParams описывает параметры запроса для эндпоинта /countries API Kodik.
type CountriesParams struct {
	// Фильтрация материалов
	Types             string `json:"types,omitempty"`              // Тип материала (например, "foreign-movie,cartoon-serial")
	Year              string `json:"year,omitempty"`               // Год или диапазон (например, "2020" или "2010-2020")
	TranslationID     int    `json:"translation_id,omitempty"`     // ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // IDs озвучек, которые нужно исключить (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода (voice или subtitles)
	HasField          string `json:"has_field,omitempty"`          // Наличие определённых полей (например: kinopoisk_id, imdb_id, mdl_id, worldart_link, shikimori_id)
	Lgbt              *bool  `json:"lgbt,omitempty"`               // Фильтр по контенту LGBT (true/false)
	Sort              string `json:"sort,omitempty"`               // Сортировка по названию страны (title) или количеству материалов (count)

	// Фильтрация по внешним данным
	Genres            string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres       string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres         string `json:"all_genres,omitempty"`         // Все жанры через запятую
	Duration          string `json:"duration,omitempty"`           // Длительность (в минутах, точное значение или диапазон, например, 30 или 40-80)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (например, 7.0 или 6.5-8.2)
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг (например, 7.0 или 6.5-8.2)
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori (например, 7.0 или 6.5-8.2)
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList (например, 7.0 или 6.5-8.2)
	Actors            string `json:"actors,omitempty"`             // Список актёров через запятую
	Directors         string `json:"directors,omitempty"`          // Список режиссёров через запятую
	Producers         string `json:"producers,omitempty"`          // Список продюсеров через запятую
	Writers           string `json:"writers,omitempty"`            // Список сценаристов через запятую
	Composers         string `json:"composers,omitempty"`          // Список композиторов через запятую
	Editors           string `json:"editors,omitempty"`            // Список монтажёров через запятую
	Designers         string `json:"designers,omitempty"`          // Список дизайнеров через запятую
	Operators         string `json:"operators,omitempty"`          // Список операторов через запятую
	RatingMPAA        string `json:"rating_mpaa,omitempty"`        // Рейтинг MPAA (например, G, PG, PG-13, R и т.д.)
	MinimalAge        string `json:"minimal_age,omitempty"`        // Минимальный возраст (например, 16 или 12-16)
	AnimeKind         string `json:"anime_kind,omitempty"`         // Тип аниме (tv, movie, ova, ona, и т.д.)
	MydramalistTags   string `json:"mydramalist_tags,omitempty"`   // Теги MyDramaList через запятую (например, Friendship, Violence, etc.)
	AnimeStatus       string `json:"anime_status,omitempty"`       // Статус аниме (anons, ongoing, released, и т.д.)
	DramaStatus       string `json:"drama_status,omitempty"`       // Статус драмы (anons, ongoing, released, и т.д.)
	AllStatus         string `json:"all_status,omitempty"`         // Универсальный статус (анонс, ongoing, released и т.д.)
	AnimeStudios      string `json:"anime_studios,omitempty"`      // Студии для аниме через запятую (например, J.C.Staff, Studio Hibari)
	AnimeLicensedBy   string `json:"anime_licensed_by,omitempty"`  // Лицензионные правообладатели через запятую (например, Wakanim, Russian Reportage)
}

// ToMap преобразует структуру CountriesParams в карту параметров для HTTP-запроса.
func (cp *CountriesParams) ToMap() map[string]string {
	params := make(map[string]string)

	if cp.Types != "" {
		params["types"] = cp.Types
	}
	if cp.Year != "" {
		params["year"] = cp.Year
	}
	if cp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(cp.TranslationID)
	}
	if cp.BlockTranslations != "" {
		params["block_translations"] = strings.ReplaceAll(cp.BlockTranslations, " ", "")
	}
	if cp.TranslationType != "" {
		params["translation_type"] = cp.TranslationType
	}
	if cp.HasField != "" {
		params["has_field"] = cp.HasField
	}
	if cp.Lgbt != nil {
		if *cp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if cp.Sort != "" {
		params["sort"] = cp.Sort
	}

	// Внешняя фильтрация
	if cp.Genres != "" {
		params["genres"] = cp.Genres
	}
	if cp.AnimeGenres != "" {
		params["anime_genres"] = cp.AnimeGenres
	}
	if cp.DramaGenres != "" {
		params["drama_genres"] = cp.DramaGenres
	}
	if cp.AllGenres != "" {
		params["all_genres"] = cp.AllGenres
	}
	if cp.Duration != "" {
		params["duration"] = cp.Duration
	}
	if cp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = cp.KinopoiskRating
	}
	if cp.ImdbRating != "" {
		params["imdb_rating"] = cp.ImdbRating
	}
	if cp.ShikimoriRating != "" {
		params["shikimori_rating"] = cp.ShikimoriRating
	}
	if cp.MydramalistRating != "" {
		params["mydramalist_rating"] = cp.MydramalistRating
	}
	if cp.Actors != "" {
		params["actors"] = cp.Actors
	}
	if cp.Directors != "" {
		params["directors"] = cp.Directors
	}
	if cp.Producers != "" {
		params["producers"] = cp.Producers
	}
	if cp.Writers != "" {
		params["writers"] = cp.Writers
	}
	if cp.Composers != "" {
		params["composers"] = cp.Composers
	}
	if cp.Editors != "" {
		params["editors"] = cp.Editors
	}
	if cp.Designers != "" {
		params["designers"] = cp.Designers
	}
	if cp.Operators != "" {
		params["operators"] = cp.Operators
	}
	if cp.RatingMPAA != "" {
		params["rating_mpaa"] = cp.RatingMPAA
	}
	if cp.MinimalAge != "" {
		params["minimal_age"] = cp.MinimalAge
	}
	if cp.AnimeKind != "" {
		params["anime_kind"] = cp.AnimeKind
	}
	if cp.MydramalistTags != "" {
		params["mydramalist_tags"] = cp.MydramalistTags
	}
	if cp.AnimeStatus != "" {
		params["anime_status"] = cp.AnimeStatus
	}
	if cp.DramaStatus != "" {
		params["drama_status"] = cp.DramaStatus
	}
	if cp.AllStatus != "" {
		params["all_status"] = cp.AllStatus
	}
	if cp.AnimeStudios != "" {
		params["anime_studios"] = cp.AnimeStudios
	}
	if cp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = cp.AnimeLicensedBy
	}

	return params
}
