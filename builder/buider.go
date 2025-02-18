package builder

import (
	"github.com/hikasami/kodik-api/api"
	"github.com/hikasami/kodik-api/models"
)

// Query представляет объект для построения запроса /search API Kodik.
type Query struct {
	params models.SearchParams
}

// Search — глобальный билдер для запросов /search, позволяющий формировать цепочку методов.
var Search = Query{
	params: models.SearchParams{},
}

// Anime устанавливает тип материала на "аниме".
func (q Query) Anime() Query {
	q.params.Types = "anime-serial,anime-movie"
	return q
}

// AnimeSerial устанавливает тип материала на "аниме сериал".
func (q Query) AnimeSerial() Query {
	q.params.Types = "anime-serial"
	return q
}

// AnimeMovie устанавливает тип материала на "аниме фильм".
func (q Query) AnimeMovie() Query {
	q.params.Types = "anime-movie"
	return q
}

// WithTitle устанавливает название материала для поиска.
func (q Query) WithTitle(title string) Query {
	q.params.Title = title
	return q
}

// WithTitleOrig устанавливает оригинальное название материала.
func (q Query) WithTitleOrig(titleOrig string) Query {
	q.params.TitleOrig = titleOrig
	return q
}

// SetStrict устанавливает флаг строгого поиска.
func (q Query) SetStrict(strict bool) Query {
	q.params.Strict = strict
	return q
}

// SetFullMatch устанавливает флаг полного совпадения.
func (q Query) SetFullMatch(fullMatch bool) Query {
	q.params.FullMatch = fullMatch
	return q
}

// WithID устанавливает идентификатор материала для поиска.
func (q Query) WithID(id string) Query {
	q.params.ID = id
	return q
}

// WithPlayerLink устанавливает ссылку на плеер.
func (q Query) WithPlayerLink(link string) Query {
	q.params.PlayerLink = link
	return q
}

// WithKinopoiskID устанавливает ID материала по Кинопоиску.
func (q Query) WithKinopoiskID(id int) Query {
	q.params.KinopoiskID = id
	return q
}

// WithImdbID устанавливает идентификатор IMDb.
func (q Query) WithImdbID(imdbID string) Query {
	q.params.ImdbID = imdbID
	return q
}

// WithMdlID устанавливает MDL идентификатор.
func (q Query) WithMdlID(mdlID string) Query {
	q.params.MdlID = mdlID
	return q
}

// WithWorldartAnimationID устанавливает Worldart Animation ID.
func (q Query) WithWorldartAnimationID(id int) Query {
	q.params.WorldartAnimationID = id
	return q
}

// WithWorldartCinemaID устанавливает Worldart Cinema ID.
func (q Query) WithWorldartCinemaID(id int) Query {
	q.params.WorldartCinemaID = id
	return q
}

// WithWorldartLink устанавливает ссылку Worldart.
func (q Query) WithWorldartLink(link string) Query {
	q.params.WorldartLink = link
	return q
}

// WithShikimoriID устанавливает идентификатор Shikimori.
func (q Query) WithShikimoriID(id int) Query {
	q.params.ShikimoriID = id
	return q
}

// Limit задает ограничение на количество возвращаемых результатов.
func (q Query) Limit(limit int) Query {
	q.params.Limit = limit
	return q
}

// WithTypes устанавливает типы материалов для поиска.
func (q Query) WithTypes(types string) Query {
	q.params.Types = types
	return q
}

// WithYear устанавливает год или диапазон лет для поиска.
func (q Query) WithYear(year string) Query {
	q.params.Year = year
	return q
}

// WithTranslationID устанавливает ID озвучки.
func (q Query) WithTranslationID(id int) Query {
	q.params.TranslationID = id
	return q
}

// WithTranslationType устанавливает тип перевода.
func (q Query) WithTranslationType(transType string) Query {
	q.params.TranslationType = transType
	return q
}

// WithHasField фильтрует материалы по наличию указанного поля.
func (q Query) WithHasField(field string) Query {
	q.params.HasField = field
	return q
}

// WithPrioritizeTranslations задает приоритет озвучек.
func (q Query) WithPrioritizeTranslations(prioritize string) Query {
	q.params.PrioritizeTranslations = prioritize
	return q
}

// WithUnprioritizeTranslations задает озвучки для исключения из приоритета.
func (q Query) WithUnprioritizeTranslations(unprioritize string) Query {
	q.params.UnprioritizeTranslations = unprioritize
	return q
}

// WithPrioritizeTranslationType устанавливает приоритетный тип перевода.
func (q Query) WithPrioritizeTranslationType(transType string) Query {
	q.params.PrioritizeTranslationType = transType
	return q
}

// WithBlockTranslations задает список блокируемых озвучек.
func (q Query) WithBlockTranslations(block string) Query {
	q.params.BlockTranslations = block
	return q
}

// WithCamrip устанавливает флаг camrip.
func (q Query) WithCamrip(camrip bool) Query {
	q.params.Camrip = &camrip
	return q
}

// WithLgbt устанавливает флаг LGBT.
func (q Query) WithLgbt(lgbt bool) Query {
	q.params.Lgbt = &lgbt
	return q
}

// WithSeasons устанавливает флаг получения сезонов.
func (q Query) WithSeasons(with bool) Query {
	q.params.WithSeasons = with
	return q
}

// WithSeason задает номер сезона для поиска.
func (q Query) WithSeason(season int) Query {
	q.params.Season = season
	return q
}

// WithEpisodes устанавливает флаг получения эпизодов.
func (q Query) WithEpisodes(with bool) Query {
	q.params.WithEpisodes = with
	return q
}

// WithEpisode задает номер эпизода.
func (q Query) WithEpisode(episode int) Query {
	q.params.Episode = episode
	return q
}

// WithPageLinks устанавливает флаг получения ссылок на страницы.
func (q Query) WithPageLinks(with bool) Query {
	q.params.WithPageLinks = with
	return q
}

// WithNotBlockedIn задает список стран, в которых материал не должен быть заблокирован.
func (q Query) WithNotBlockedIn(countries string) Query {
	q.params.NotBlockedIn = countries
	return q
}

// WithNotBlockedForMe устанавливает флаг, что материал не заблокирован для пользователя.
func (q Query) WithNotBlockedForMe(notBlocked bool) Query {
	q.params.NotBlockedForMe = &notBlocked
	return q
}

// WithMaterialData устанавливает флаг получения дополнительных данных о материале.
func (q Query) WithMaterialData(with bool) Query {
	q.params.WithMaterialData = with
	return q
}

// WithCountries задает фильтрацию по странам.
func (q Query) WithCountries(countries string) Query {
	q.params.Countries = countries
	return q
}

// WithGenres задает фильтрацию по жанрам.
func (q Query) WithGenres(genres string) Query {
	q.params.Genres = genres
	return q
}

// WithDuration задает фильтрацию по длительности.
func (q Query) WithDuration(duration string) Query {
	q.params.Duration = duration
	return q
}

// WithKinopoiskRating задает фильтрацию по рейтингу Кинопоиска.
func (q Query) WithKinopoiskRating(rating string) Query {
	q.params.KinopoiskRating = rating
	return q
}

// WithImdbRating задает фильтрацию по рейтингу IMDb.
func (q Query) WithImdbRating(rating string) Query {
	q.params.ImdbRating = rating
	return q
}

// WithShikimoriRating задает фильтрацию по рейтингу Shikimori.
func (q Query) WithShikimoriRating(rating string) Query {
	q.params.ShikimoriRating = rating
	return q
}

// WithMydramalistRating задает фильтрацию по рейтингу MyDramaList.
func (q Query) WithMydramalistRating(rating string) Query {
	q.params.MydramalistRating = rating
	return q
}

// WithActors задает фильтрацию по актерам.
func (q Query) WithActors(actors string) Query {
	q.params.Actors = actors
	return q
}

// WithDirectors задает фильтрацию по режиссерам.
func (q Query) WithDirectors(directors string) Query {
	q.params.Directors = directors
	return q
}

// WithProducers задает фильтрацию по продюсерам.
func (q Query) WithProducers(producers string) Query {
	q.params.Producers = producers
	return q
}

// WithWriters задает фильтрацию по сценаристам.
func (q Query) WithWriters(writers string) Query {
	q.params.Writers = writers
	return q
}

// WithComposers задает фильтрацию по композиторам.
func (q Query) WithComposers(composers string) Query {
	q.params.Composers = composers
	return q
}

// WithEditors задает фильтрацию по монтажерам.
func (q Query) WithEditors(editors string) Query {
	q.params.Editors = editors
	return q
}

// WithDesigners задает фильтрацию по дизайнерам.
func (q Query) WithDesigners(designers string) Query {
	q.params.Designers = designers
	return q
}

// WithOperators задает фильтрацию по операторам.
func (q Query) WithOperators(operators string) Query {
	q.params.Operators = operators
	return q
}

// WithRatingMPAA задает фильтрацию по рейтингу MPAA.
func (q Query) WithRatingMPAA(ratingMPAA string) Query {
	q.params.RatingMPAA = ratingMPAA
	return q
}

// WithMinimalAge задает фильтрацию по минимальному возрасту.
func (q Query) WithMinimalAge(age string) Query {
	q.params.MinimalAge = age
	return q
}

// WithAnimeKind задает фильтрацию по типу аниме.
func (q Query) WithAnimeKind(kind string) Query {
	q.params.AnimeKind = kind
	return q
}

// WithMydramalistTags задает фильтрацию по тегам MyDramaList.
func (q Query) WithMydramalistTags(tags string) Query {
	q.params.MydramalistTags = tags
	return q
}

// WithAnimeStatus задает фильтрацию по статусу аниме.
func (q Query) WithAnimeStatus(status string) Query {
	q.params.AnimeStatus = status
	return q
}

// WithDramaStatus задает фильтрацию по статусу драмы.
func (q Query) WithDramaStatus(status string) Query {
	q.params.DramaStatus = status
	return q
}

// WithAllStatus задает фильтрацию по универсальному статусу.
func (q Query) WithAllStatus(status string) Query {
	q.params.AllStatus = status
	return q
}

// WithAnimeStudios задает фильтрацию по студиям аниме.
func (q Query) WithAnimeStudios(studios string) Query {
	q.params.AnimeStudios = studios
	return q
}

// WithAnimeLicensedBy задает фильтрацию по правообладателям аниме.
func (q Query) WithAnimeLicensedBy(licensedBy string) Query {
	q.params.AnimeLicensedBy = licensedBy
	return q
}

// Execute выполняет запрос /search с текущими параметрами.
func (q Query) Execute() (*models.SearchResponse, error) {
	return api.Search(&q.params)
}
