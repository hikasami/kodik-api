package models

import "time"

// Translation описывает структуру озвучки (дублирования).
type Translation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"` // Возможные значения: "voice" или "subtitles"
}

// MaterialData описывает дополнительную информацию о материале из различных источников.
type MaterialData struct {
	// Основная информация
	Title         *string  `json:"title,omitempty"`
	AnimeTitle    *string  `json:"anime_title,omitempty"`
	TitleEn       *string  `json:"title_en,omitempty"`
	OtherTitles   []string `json:"other_titles,omitempty"`
	OtherTitlesEn []string `json:"other_titles_en,omitempty"`
	OtherTitlesJp []string `json:"other_titles_jp,omitempty"`

	// Аниме-специфичная информация
	AnimeLicenseName *string  `json:"anime_license_name,omitempty"`
	AnimeLicensedBy  []string `json:"anime_licensed_by,omitempty"`
	AnimeKind        *string  `json:"anime_kind,omitempty"`

	// MyDramaList теги
	MydramalistTags []string `json:"mydramalist_tags,omitempty"`

	// Статусы
	AllStatus   *string `json:"all_status,omitempty"`
	AnimeStatus *string `json:"anime_status,omitempty"`
	DramaStatus *string `json:"drama_status,omitempty"`

	// Основные данные
	Year             *int    `json:"year,omitempty"`
	Tagline          *string `json:"tagline,omitempty"`
	Description      *string `json:"description,omitempty"`
	AnimeDescription *string `json:"anime_description,omitempty"`

	// Постеры
	PosterUrl      *string  `json:"poster_url,omitempty"`
	AnimePosterUrl *string  `json:"anime_poster_url,omitempty"`
	DramaPosterUrl *string  `json:"drama_poster_url,omitempty"`
	Screenshots    []string `json:"screenshots,omitempty"`

	// Продолжительность и география
	Duration  *int     `json:"duration,omitempty"`
	Countries []string `json:"countries,omitempty"`

	// Жанры
	AllGenres    []string `json:"all_genres,omitempty"`
	Genres       []string `json:"genres,omitempty"`
	AnimeGenres  []string `json:"anime_genres,omitempty"`
	DramaGenres  []string `json:"drama_genres,omitempty"`
	AnimeStudios []string `json:"anime_studios,omitempty"`

	// Рейтинги
	KinopoiskRating   *float64 `json:"kinopoisk_rating,omitempty"`
	KinopoiskVotes    *int     `json:"kinopoisk_votes,omitempty"`
	ImdbRating        *float64 `json:"imdb_rating,omitempty"`
	ImdbVotes         *int     `json:"imdb_votes,omitempty"`
	ShikimoriRating   *float64 `json:"shikimori_rating,omitempty"`
	ShikimoriVotes    *int     `json:"shikimori_votes,omitempty"`
	MydramalistRating *float64 `json:"mydramalist_rating,omitempty"`
	MydramalistVotes  *int     `json:"mydramalist_votes,omitempty"`

	// Даты
	PremiereRu    *string `json:"premiere_ru,omitempty"`
	PremiereWorld *string `json:"premiere_world,omitempty"`
	AiredAt       *string `json:"aired_at,omitempty"`
	ReleasedAt    *string `json:"released_at,omitempty"`
	NextEpisodeAt *string `json:"next_episode_at,omitempty"`

	// Возрастные ограничения
	RatingMpaa *string `json:"rating_mpaa,omitempty"`
	MinimalAge *int    `json:"minimal_age,omitempty"`

	// Эпизоды
	EpisodesTotal *int `json:"episodes_total,omitempty"`
	EpisodesAired *int `json:"episodes_aired,omitempty"`

	// Участники производства
	Actors    []string `json:"actors,omitempty"`
	Directors []string `json:"directors,omitempty"`
	Producers []string `json:"producers,omitempty"`
	Writers   []string `json:"writers,omitempty"`
	Composers []string `json:"composers,omitempty"`
	Editors   []string `json:"editors,omitempty"`
	Designers []string `json:"designers,omitempty"`
	Operators []string `json:"operators,omitempty"`
}

// Material описывает основные данные материала (фильма, сериала и т.д.).
type Material struct {
	ID               string      `json:"id"`
	Title            string      `json:"title"`
	TitleOrig        string      `json:"title_orig"`
	OtherTitle       string      `json:"other_title"`
	Link             string      `json:"link"`
	Year             int         `json:"year"`
	KinopoiskID      string      `json:"kinopoisk_id"`
	ImdbID           string      `json:"imdb_id"`
	MdlID            string      `json:"mdl_id"`
	WorldartLink     string      `json:"worldart_link"`
	ShikimoriID      string      `json:"shikimori_id"`
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

	// Дополнительная информация из внешних источников (KinoPoisk, Shikimori, MyDramaList)
	MaterialData *MaterialData `json:"material_data,omitempty"`
}

// SearchResponse представляет структуру ответа для запроса /search.
type SearchResponse struct {
	Results []Material `json:"results"`
}
