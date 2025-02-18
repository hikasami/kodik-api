package models

import "time"

// Translation описывает структуру озвучки (дублирования).
type Translation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"` // Возможные значения: "voice" или "subtitles"
}

// Material описывает основные данные материала (фильма, сериала и т.д.).
type Material struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	TitleOrig        string      `json:"title_orig"`
	OtherTitle       string      `json:"other_title"`
	Link             string      `json:"link"`
	Year             int         `json:"year"`
	KinopoiskID      int         `json:"kinopoisk_id"`
	ImdbID           string      `json:"imdb_id"`
	MdlID            string      `json:"mdl_id"`
	WorldartLink     string      `json:"worldart_link"`
	ShikimoriID      int         `json:"shikimori_id"`
	Type             string      `json:"type"`
	Quality          string      `json:"quality"`
	Camrip           bool        `json:"camrip"`
	Lgbt             bool        `json:"lgbt"`
	Translation      Translation `json:"translation"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	BlockedCountries []string    `json:"blocked_countries"`

	// Данные, специфичные для сериалов (для фильмов могут отсутствовать)
	Seasons        interface{} `json:"seasons,omitempty"`
	LastSeason     int         `json:"last_season,omitempty"`
	LastEpisode    int         `json:"last_episode,omitempty"`
	EpisodesCount  int         `json:"episodes_count,omitempty"`
	BlockedSeasons interface{} `json:"blocked_seasons,omitempty"`
	Screenshots    []string    `json:"screenshots,omitempty"`
}

// SearchResponse представляет структуру ответа для запроса /search.
type SearchResponse struct {
	Results []Material `json:"results"`
}
