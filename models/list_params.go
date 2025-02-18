package models

import (
	"strconv"
	"strings"
)

// ListParams описывает доступные параметры для запроса /list API Kodik.
// Если поле имеет нулевое значение, оно не включается в запрос.
type ListParams struct {
	// Основные параметры
	Limit             int    `json:"limit,omitempty"`              // Количество материалов на запрос (1-100)
	Sort              string `json:"sort,omitempty"`               // Поле сортировки (например, updated_at)
	Order             string `json:"order,omitempty"`              // Направление сортировки (asc или desc)
	Types             string `json:"types,omitempty"`              // Список типов материалов через запятую
	Year              string `json:"year,omitempty"`               // Год или диапазон годов выпуска
	TranslationID     int    `json:"translation_id,omitempty"`     // ID озвучки
	BlockTranslations string `json:"block_translations,omitempty"` // IDs озвучек, которые следует исключить (через запятую)
	TranslationType   string `json:"translation_type,omitempty"`   // Тип перевода: voice или subtitles
	HasField          string `json:"has_field,omitempty"`          // Фильтрация по наличию поля

	// Флаги маркеров
	Camrip *bool `json:"camrip,omitempty"` // Показывать только камрипы (true/false)
	Lgbt   *bool `json:"lgbt,omitempty"`   // Фильтрация материалов по содержанию LGBT-сцен

	// Параметры для сериалов
	WithSeasons      bool   `json:"with_seasons,omitempty"`       // Включать данные о сезонах
	WithEpisodes     bool   `json:"with_episodes,omitempty"`      // Включать данные об эпизодах
	WithEpisodesData bool   `json:"with_episodes_data,omitempty"` // Включать подробные данные об эпизодах
	WithPageLinks    bool   `json:"with_page_links,omitempty"`    // Заменять ссылки на плеер на ссылки на страницу
	NotBlockedIn     string `json:"not_blocked_in,omitempty"`     // Список стран (через запятую), в которых материал не должен быть заблокирован
	NotBlockedForMe  *bool  `json:"not_blocked_for_me,omitempty"` // Флаг исключения заблокированных материалов (автоматическое определение страны)

	// Внешняя фильтрация данных
	WithMaterialData bool   `json:"with_material_data,omitempty"` // Включать расширенные данные материала
	Countries        string `json:"countries,omitempty"`          // Страны через запятую
	Genres           string `json:"genres,omitempty"`             // Жанры через запятую
	AnimeGenres      string `json:"anime_genres,omitempty"`       // Жанры для аниме через запятую
	DramaGenres      string `json:"drama_genres,omitempty"`       // Жанры для драм через запятую
	AllGenres        string `json:"all_genres,omitempty"`         // Все жанры через запятую

	// Дополнительные параметры
	Duration          string `json:"duration,omitempty"`           // Длительность (точное значение или диапазон)
	KinopoiskRating   string `json:"kinopoisk_rating,omitempty"`   // Рейтинг по Кинопоиску или диапазон
	ImdbRating        string `json:"imdb_rating,omitempty"`        // IMDb рейтинг или диапазон
	ShikimoriRating   string `json:"shikimori_rating,omitempty"`   // Рейтинг Шикимори или диапазон
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
	RatingMPAA      string `json:"rating_mpaa,omitempty"`       // Рейтинг MPAA
	MinimalAge      string `json:"minimal_age,omitempty"`       // Минимальный возраст
	AnimeKind       string `json:"anime_kind,omitempty"`        // Тип аниме (tv, movie, ova, ona и т.д.)
	MydramalistTags string `json:"mydramalist_tags,omitempty"`  // Теги MyDramaList через запятую
	AnimeStatus     string `json:"anime_status,omitempty"`      // Статус аниме (anons, ongoing, released и т.д.)
	DramaStatus     string `json:"drama_status,omitempty"`      // Статус драмы
	AllStatus       string `json:"all_status,omitempty"`        // Универсальный статус
	AnimeStudios    string `json:"anime_studios,omitempty"`     // Студии для аниме через запятую
	AnimeLicensedBy string `json:"anime_licensed_by,omitempty"` // Владелец лицензионных прав через запятую
}

// ToMap преобразует структуру ListParams в карту параметров для HTTP-запроса.
func (lp *ListParams) ToMap() map[string]string {
	params := make(map[string]string)

	if lp.Limit != 0 {
		params["limit"] = strconv.Itoa(lp.Limit)
	}
	if lp.Sort != "" {
		params["sort"] = lp.Sort
	}
	if lp.Order != "" {
		params["order"] = lp.Order
	}
	if lp.Types != "" {
		params["types"] = lp.Types
	}
	if lp.Year != "" {
		params["year"] = lp.Year
	}
	if lp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(lp.TranslationID)
	}
	if lp.BlockTranslations != "" {
		params["block_translations"] = lp.BlockTranslations
	}
	if lp.TranslationType != "" {
		params["translation_type"] = lp.TranslationType
	}
	if lp.HasField != "" {
		params["has_field"] = lp.HasField
	}
	if lp.Camrip != nil {
		if *lp.Camrip {
			params["camrip"] = "true"
		} else {
			params["camrip"] = "false"
		}
	}
	if lp.Lgbt != nil {
		if *lp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	// Параметры сериалов (флаги всегда включаются, даже если false)
	if lp.WithSeasons {
		params["with_seasons"] = "true"
	} else {
		params["with_seasons"] = "false"
	}
	if lp.WithEpisodes {
		params["with_episodes"] = "true"
	} else {
		params["with_episodes"] = "false"
	}
	if lp.WithEpisodesData {
		params["with_episodes_data"] = "true"
	} else {
		params["with_episodes_data"] = "false"
	}
	if lp.WithPageLinks {
		params["with_page_links"] = "true"
	} else {
		params["with_page_links"] = "false"
	}
	if lp.NotBlockedIn != "" {
		params["not_blocked_in"] = strings.ReplaceAll(lp.NotBlockedIn, " ", "")
	}
	if lp.NotBlockedForMe != nil {
		if *lp.NotBlockedForMe {
			params["not_blocked_for_me"] = "true"
		} else {
			params["not_blocked_for_me"] = "false"
		}
	}
	// Внешняя фильтрация
	if lp.WithMaterialData {
		params["with_material_data"] = "true"
	} else {
		params["with_material_data"] = "false"
	}
	if lp.Countries != "" {
		params["countries"] = lp.Countries
	}
	if lp.Genres != "" {
		params["genres"] = lp.Genres
	}
	if lp.AnimeGenres != "" {
		params["anime_genres"] = lp.AnimeGenres
	}
	if lp.DramaGenres != "" {
		params["drama_genres"] = lp.DramaGenres
	}
	if lp.AllGenres != "" {
		params["all_genres"] = lp.AllGenres
	}
	// Дополнительные параметры
	if lp.Duration != "" {
		params["duration"] = lp.Duration
	}
	if lp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = lp.KinopoiskRating
	}
	if lp.ImdbRating != "" {
		params["imdb_rating"] = lp.ImdbRating
	}
	if lp.ShikimoriRating != "" {
		params["shikimori_rating"] = lp.ShikimoriRating
	}
	if lp.MydramalistRating != "" {
		params["mydramalist_rating"] = lp.MydramalistRating
	}
	// Персоналии
	if lp.Actors != "" {
		params["actors"] = lp.Actors
	}
	if lp.Directors != "" {
		params["directors"] = lp.Directors
	}
	if lp.Producers != "" {
		params["producers"] = lp.Producers
	}
	if lp.Writers != "" {
		params["writers"] = lp.Writers
	}
	if lp.Composers != "" {
		params["composers"] = lp.Composers
	}
	if lp.Editors != "" {
		params["editors"] = lp.Editors
	}
	if lp.Designers != "" {
		params["designers"] = lp.Designers
	}
	if lp.Operators != "" {
		params["operators"] = lp.Operators
	}
	// Прочее
	if lp.RatingMPAA != "" {
		params["rating_mpaa"] = lp.RatingMPAA
	}
	if lp.MinimalAge != "" {
		params["minimal_age"] = lp.MinimalAge
	}
	if lp.AnimeKind != "" {
		params["anime_kind"] = lp.AnimeKind
	}
	if lp.MydramalistTags != "" {
		params["mydramalist_tags"] = lp.MydramalistTags
	}
	if lp.AnimeStatus != "" {
		params["anime_status"] = lp.AnimeStatus
	}
	if lp.DramaStatus != "" {
		params["drama_status"] = lp.DramaStatus
	}
	if lp.AllStatus != "" {
		params["all_status"] = lp.AllStatus
	}
	if lp.AnimeStudios != "" {
		params["anime_studios"] = lp.AnimeStudios
	}
	if lp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = lp.AnimeLicensedBy
	}

	return params
}
