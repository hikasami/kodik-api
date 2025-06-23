package models

import (
	"strconv"
	"strings"
)

// SearchParams описывает параметры запроса для эндпоинта /search API Kodik.
// Поля соответствуют описанным в таблице параметров.
// Если значение поля не установлено (нулевое), оно не будет включено в итоговый запрос.
type SearchParams struct {
	// Обязательные поля (хотя бы одно из них должно быть заполнено)
	Title     string `json:"title,omitempty"`
	TitleOrig string `json:"title_orig,omitempty"`

	// Параметры строгого поиска
	Strict    bool `json:"strict,omitempty"`
	FullMatch bool `json:"full_match,omitempty"`

	// Поиск по идентификаторам
	ID                  string `json:"id,omitempty"`
	PlayerLink          string `json:"player_link,omitempty"`
	KinopoiskID         int    `json:"kinopoisk_id,omitempty"`
	ImdbID              string `json:"imdb_id,omitempty"`
	MdlID               string `json:"mdl_id,omitempty"`
	WorldartAnimationID int    `json:"worldart_animation_id,omitempty"`
	WorldartCinemaID    int    `json:"worldart_cinema_id,omitempty"`
	WorldartLink        string `json:"worldart_link,omitempty"`
	ShikimoriID         int    `json:"shikimori_id,omitempty"`

	// Опциональные параметры
	Limit                     int    `json:"limit,omitempty"`
	Types                     string `json:"types,omitempty"` // список типов через запятую
	Year                      string `json:"year,omitempty"`
	TranslationID             int    `json:"translation_id,omitempty"`
	TranslationType           string `json:"translation_type,omitempty"`
	HasField                  string `json:"has_field,omitempty"`
	PrioritizeTranslations    string `json:"prioritize_translations,omitempty"`
	UnprioritizeTranslations  string `json:"unprioritize_translations,omitempty"`
	PrioritizeTranslationType string `json:"prioritize_translation_type,omitempty"`
	BlockTranslations         string `json:"block_translations,omitempty"`
	Camrip                    *bool  `json:"camrip,omitempty"` // указатель для отличия false и не заданного значения
	Lgbt                      *bool  `json:"lgbt,omitempty"`   // указатель для отличия false и не заданного значения
	WithSeasons               bool   `json:"with_seasons,omitempty"`
	Season                    int    `json:"season,omitempty"`
	WithEpisodes              bool   `json:"with_episodes,omitempty"`
	WithEpisodesData          bool   `json:"with_episodes_data,omitempty`
	Episode                   int    `json:"episode,omitempty"`
	WithPageLinks             bool   `json:"with_page_links,omitempty"`
	NotBlockedIn              string `json:"not_blocked_in,omitempty"` // список стран через запятую (без пробелов)
	NotBlockedForMe           *bool  `json:"not_blocked_for_me,omitempty"`
	WithMaterialData          bool   `json:"with_material_data,omitempty"`
	Countries                 string `json:"countries,omitempty"` // список стран через запятую
	Genres                    string `json:"genres,omitempty"`    // список жанров через запятую
	Duration                  string `json:"duration,omitempty"`
	KinopoiskRating           string `json:"kinopoisk_rating,omitempty"`
	ImdbRating                string `json:"imdb_rating,omitempty"`
	ShikimoriRating           string `json:"shikimori_rating,omitempty"`
	MydramalistRating         string `json:"mydramalist_rating,omitempty"`
	Actors                    string `json:"actors,omitempty"`    // список актёров через запятую
	Directors                 string `json:"directors,omitempty"` // список режиссёров через запятую
	Producers                 string `json:"producers,omitempty"` // список продюсеров через запятую
	Writers                   string `json:"writers,omitempty"`   // список сценаристов через запятую
	Composers                 string `json:"composers,omitempty"` // список композиторов через запятую
	Editors                   string `json:"editors,omitempty"`   // список монтажёров через запятую
	Designers                 string `json:"designers,omitempty"` // список дизайнеров через запятую
	Operators                 string `json:"operators,omitempty"` // список операторов через запятую
	RatingMPAA                string `json:"rating_mpaa,omitempty"`
	MinimalAge                string `json:"minimal_age,omitempty"`
	AnimeKind                 string `json:"anime_kind,omitempty"`
	MydramalistTags           string `json:"mydramalist_tags,omitempty"`
	AnimeStatus               string `json:"anime_status,omitempty"`
	DramaStatus               string `json:"drama_status,omitempty"`
	AllStatus                 string `json:"all_status,omitempty"`
	AnimeStudios              string `json:"anime_studios,omitempty"`
	AnimeLicensedBy           string `json:"anime_licensed_by,omitempty"`
}

// ToMap преобразует структуру SearchParams в карту параметров для запроса.
// Поля со значением по умолчанию пропускаются (за исключением булевых, для которых значение false явно передаётся).
func (sp *SearchParams) ToMap() map[string]string {
	params := make(map[string]string)

	if sp.Title != "" {
		params["title"] = sp.Title
	}
	if sp.TitleOrig != "" {
		params["title_orig"] = sp.TitleOrig
	}
	// Булевые значения передаются всегда явно
	if sp.Strict {
		params["strict"] = "true"
	} else {
		params["strict"] = "false"
	}
	if sp.FullMatch {
		params["full_match"] = "true"
	} else {
		params["full_match"] = "false"
	}
	if sp.ID != "" {
		params["id"] = sp.ID
	}
	if sp.PlayerLink != "" {
		params["player_link"] = sp.PlayerLink
	}
	if sp.KinopoiskID != 0 {
		params["kinopoisk_id"] = strconv.Itoa(sp.KinopoiskID)
	}
	if sp.ImdbID != "" {
		params["imdb_id"] = sp.ImdbID
	}
	if sp.MdlID != "" {
		params["mdl_id"] = sp.MdlID
	}
	if sp.WorldartAnimationID != 0 {
		params["worldart_animation_id"] = strconv.Itoa(sp.WorldartAnimationID)
	}
	if sp.WorldartCinemaID != 0 {
		params["worldart_cinema_id"] = strconv.Itoa(sp.WorldartCinemaID)
	}
	if sp.WorldartLink != "" {
		params["worldart_link"] = sp.WorldartLink
	}
	if sp.ShikimoriID != 0 {
		params["shikimori_id"] = strconv.Itoa(sp.ShikimoriID)
	}
	if sp.Limit != 0 {
		params["limit"] = strconv.Itoa(sp.Limit)
	}
	if sp.Types != "" {
		params["types"] = sp.Types
	}
	if sp.Year != "" {
		params["year"] = sp.Year
	}
	if sp.TranslationID != 0 {
		params["translation_id"] = strconv.Itoa(sp.TranslationID)
	}
	if sp.TranslationType != "" {
		params["translation_type"] = sp.TranslationType
	}
	if sp.HasField != "" {
		params["has_field"] = sp.HasField
	}
	if sp.PrioritizeTranslations != "" {
		params["prioritize_translations"] = sp.PrioritizeTranslations
	}
	if sp.UnprioritizeTranslations != "" {
		params["unprioritize_translations"] = sp.UnprioritizeTranslations
	}
	if sp.PrioritizeTranslationType != "" {
		params["prioritize_translation_type"] = sp.PrioritizeTranslationType
	}
	if sp.BlockTranslations != "" {
		params["block_translations"] = sp.BlockTranslations
	}
	if sp.Camrip != nil {
		if *sp.Camrip {
			params["camrip"] = "true"
		} else {
			params["camrip"] = "false"
		}
	}
	if sp.Lgbt != nil {
		if *sp.Lgbt {
			params["lgbt"] = "true"
		} else {
			params["lgbt"] = "false"
		}
	}
	if sp.WithSeasons {
		params["with_seasons"] = "true"
	} else {
		params["with_seasons"] = "false"
	}
	if sp.Season != 0 {
		params["season"] = strconv.Itoa(sp.Season)
	}
	if sp.WithEpisodes {
		params["with_episodes"] = "true"
	} else {
		params["with_episodes"] = "false"
	}
	if sp.WithEpisodesData {
		params["with_episodes_data"] = "true"
	} else {
		params["with_episodes_data"] = "false"
	}
	if sp.Episode != 0 {
		params["episode"] = strconv.Itoa(sp.Episode)
	}
	if sp.WithPageLinks {
		params["with_page_links"] = "true"
	} else {
		params["with_page_links"] = "false"
	}
	if sp.NotBlockedIn != "" {
		// Удаляем пробелы и оставляем список через запятую
		params["not_blocked_in"] = strings.ReplaceAll(sp.NotBlockedIn, " ", "")
	}
	if sp.NotBlockedForMe != nil {
		if *sp.NotBlockedForMe {
			params["not_blocked_for_me"] = "true"
		} else {
			params["not_blocked_for_me"] = "false"
		}
	}
	if sp.WithMaterialData {
		params["with_material_data"] = "true"
	} else {
		params["with_material_data"] = "false"
	}
	if sp.Countries != "" {
		params["countries"] = sp.Countries
	}
	if sp.Genres != "" {
		params["genres"] = sp.Genres
	}
	if sp.Duration != "" {
		params["duration"] = sp.Duration
	}
	if sp.KinopoiskRating != "" {
		params["kinopoisk_rating"] = sp.KinopoiskRating
	}
	if sp.ImdbRating != "" {
		params["imdb_rating"] = sp.ImdbRating
	}
	if sp.ShikimoriRating != "" {
		params["shikimori_rating"] = sp.ShikimoriRating
	}
	if sp.MydramalistRating != "" {
		params["mydramalist_rating"] = sp.MydramalistRating
	}
	if sp.Actors != "" {
		params["actors"] = sp.Actors
	}
	if sp.Directors != "" {
		params["directors"] = sp.Directors
	}
	if sp.Producers != "" {
		params["producers"] = sp.Producers
	}
	if sp.Writers != "" {
		params["writers"] = sp.Writers
	}
	if sp.Composers != "" {
		params["composers"] = sp.Composers
	}
	if sp.Editors != "" {
		params["editors"] = sp.Editors
	}
	if sp.Designers != "" {
		params["designers"] = sp.Designers
	}
	if sp.Operators != "" {
		params["operators"] = sp.Operators
	}
	if sp.RatingMPAA != "" {
		params["rating_mpaa"] = sp.RatingMPAA
	}
	if sp.MinimalAge != "" {
		params["minimal_age"] = sp.MinimalAge
	}
	if sp.AnimeKind != "" {
		params["anime_kind"] = sp.AnimeKind
	}
	if sp.MydramalistTags != "" {
		params["mydramalist_tags"] = sp.MydramalistTags
	}
	if sp.AnimeStatus != "" {
		params["anime_status"] = sp.AnimeStatus
	}
	if sp.DramaStatus != "" {
		params["drama_status"] = sp.DramaStatus
	}
	if sp.AllStatus != "" {
		params["all_status"] = sp.AllStatus
	}
	if sp.AnimeStudios != "" {
		params["anime_studios"] = sp.AnimeStudios
	}
	if sp.AnimeLicensedBy != "" {
		params["anime_licensed_by"] = sp.AnimeLicensedBy
	}

	return params
}
