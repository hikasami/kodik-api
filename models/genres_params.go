package models

import (
	"strconv"
	"strings"
)

// GenresParams описывает параметры запроса для эндпоинта /genres API Kodik.
// Токен не передаётся, так как он уже установлен в клиенте.
type GenresParams struct {
	// Основные параметры
	GenresType        string `json:"genres_type,omitempty"`        // Выбор источника жанров: kinopoisk (по умолчанию), shikimori, mydramalist или all
	Types             string `json:"types,omitempty"`              // Фильтрация по типу материала (например, foreign-movie, cartoon-serial)
	Year              string `json:"year,omitempty"`               // Фильтр по году или диапазону лет
	TranslationID     int    `json:"translation_id,omitempty"`     // Фильтр по ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // Исключение указанных озвучек (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода: voice или subtitles
	HasField          string `json:"has_field,omitempty"`          // Фильтрация по наличию определённых полей (например, kinopoisk_id, imdb_id, mdl_id, worldart_link, shikimori_id)
	Lgbt              *bool  `json:"lgbt,omitempty"`               // Фильтрация по содержанию LGBT (true/false)
	Sort              string `json:"sort,omitempty"`               // Сортировка результатов по названию жанра (title) или по числу материалов (count)

	// Фильтрация по внешним данным
	Countries         string `json:"countries,omitempty"`          // Фильтр по странам (через запятую)
	Genres            string `json:"genres,omitempty"`             // Фильтр по жанрам (через запятую)
	AnimeGenres       string `json:"anime_genres,omitempty"`       // Фильтр по жанрам для аниме
	DramaGenres       string `json:"drama_genres,omitempty"`       // Фильтр по жанрам для драм
	AllGenres         string `json:"all_genres,omitempty"`         // Фильтр по всем жанрам
	Duration          string `json:"duration,omitempty"`           // Длительность (точное значение или диапазон в минутах)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг Кинопоиска (точное значение или диапазон)
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг (точное значение или диапазон)
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Shikimori (точное значение или диапазон)
	MydramalistRating string `json:"mydramalist_rating,omitempty"` // Рейтинг MyDramaList (точное значение или диапазон)
	Actors            string `json:"actors,omitempty"`             // Список актёров через запятую
	Directors         string `json:"directors,omitempty"`          // Список режиссёров через запятую
	Producers         string `json:"producers,omitempty"`          // Список продюсеров через запятую
	Writers           string `json:"writers,omitempty"`            // Список сценаристов через запятую
	Composers         string `json:"composers,omitempty"`          // Список композиторов через запятую
	Editors           string `json:"editors,omitempty"`            // Список монтажёров через запятую
	Designers         string `json:"designers,omitempty"`          // Список дизайнеров через запятую
	Operators         string `json:"operators,omitempty"`          // Список операторов через запятую
	RatingMPAA        string `json:"rating_mpaa,omitempty"`        // Рейтинг MPAA (например, G, PG, PG-13, R)
	MinimalAge        string `json:"minimal_age,omitempty"`        // Минимальный возраст для просмотра (например, 16 или 12-16)
	AnimeKind         string `json:"anime_kind,omitempty"`         // Тип аниме: tv, movie, ova, ona и т.д.
	MydramalistTags   string `json:"mydramalist_tags,omitempty"`   // Теги MyDramaList через запятую
	AnimeStatus       string `json:"anime_status,omitempty"`       // Статус аниме (anons, ongoing, released и т.д.)
	DramaStatus       string `json:"drama_status,omitempty"`       // Статус драмы (anons, ongoing, released и т.д.)
	AllStatus         string `json:"all_status,omitempty"`         // Универсальный статус (например, anons, ongoing, released)
	AnimeStudios      string `json:"anime_studios,omitempty"`      // Студии для аниме через запятую
	AnimeLicensedBy   string `json:"anime_licensed_by,omitempty"`  // Правообладатели для аниме через запятую
}

// ToMap преобразует структуру GenresParams в карту параметров для формирования HTTP-запроса.
func (gp *GenresParams) ToMap() map[string]string {
	params := make(map[string]string)

	// Основные параметры
	if gp.GenresType != "" {
		params["genres_type"] = gp.GenresType
	}
	if gp.Types != "" {
		params["types"] = gp.Types
	}
	if gp.Year != "" {
		params["year"] = gp.Year
	}
	if gp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(gp.TranslationID)
	}
	if gp.BlockTranslations != "" {
		params["block_translations"] = strings.ReplaceAll(gp.BlockTranslations, " ", "")
	}
	if gp.TranslationType != "" {
		params["translation_type"] = gp.TranslationType
	}
	if gp.HasField != "" {
		params["has_field"] = gp.HasField
	}
	if gp.Lgbt != nil {
		if *gp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if gp.Sort != "" {
		params["sort"] = gp.Sort
	}

	// Фильтрация по внешним данным
	if gp.Countries != "" {
		params["countries"] = gp.Countries
	}
	if gp.Genres != "" {
		params["genres"] = gp.Genres
	}
	if gp.AnimeGenres != "" {
		params["anime_genres"] = gp.AnimeGenres
	}
	if gp.DramaGenres != "" {
		params["drama_genres"] = gp.DramaGenres
	}
	if gp.AllGenres != "" {
		params["all_genres"] = gp.AllGenres
	}
	if gp.Duration != "" {
		params["duration"] = gp.Duration
	}
	if gp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = gp.KinopoiskRating
	}
	if gp.ImdbRating != "" {
		params["imdb_rating"] = gp.ImdbRating
	}
	if gp.ShikimoriRating != "" {
		params["shikimori_rating"] = gp.ShikimoriRating
	}
	if gp.MydramalistRating != "" {
		params["mydramalist_rating"] = gp.MydramalistRating
	}
	if gp.Actors != "" {
		params["actors"] = gp.Actors
	}
	if gp.Directors != "" {
		params["directors"] = gp.Directors
	}
	if gp.Producers != "" {
		params["producers"] = gp.Producers
	}
	if gp.Writers != "" {
		params["writers"] = gp.Writers
	}
	if gp.Composers != "" {
		params["composers"] = gp.Composers
	}
	if gp.Editors != "" {
		params["editors"] = gp.Editors
	}
	if gp.Designers != "" {
		params["designers"] = gp.Designers
	}
	if gp.Operators != "" {
		params["operators"] = gp.Operators
	}
	if gp.RatingMPAA != "" {
		params["rating_mpaa"] = gp.RatingMPAA
	}
	if gp.MinimalAge != "" {
		params["minimal_age"] = gp.MinimalAge
	}
	if gp.AnimeKind != "" {
		params["anime_kind"] = gp.AnimeKind
	}
	if gp.MydramalistTags != "" {
		params["mydramalist_tags"] = gp.MydramalistTags
	}
	if gp.AnimeStatus != "" {
		params["anime_status"] = gp.AnimeStatus
	}
	if gp.DramaStatus != "" {
		params["drama_status"] = gp.DramaStatus
	}
	if gp.AllStatus != "" {
		params["all_status"] = gp.AllStatus
	}
	if gp.AnimeStudios != "" {
		params["anime_studios"] = gp.AnimeStudios
	}
	if gp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = gp.AnimeLicensedBy
	}

	return params
}
