/*
Документация для неофициального API кинопоиска (kinopoisk.dev).

 <!-- Yandex.Metrika counter --> <div><img src=\"https://mc.yandex.ru/watch/62307766\" style=\"position:absolute; left:-9999px;\" alt=\"\" /></div> <!-- /Yandex.Metrika counter --> <p>Через этот API вы можете получить практически все данные из кинопоиска. Больше информации вы можете получить изучив эту документацию.</p> <h2>Как работать с документацией?</h2> <p> Для начала работы с API вам необходимо получить токен, который вы можете получить в боте <a href=\"https://t.me/kinopoiskdev_bot\">@kinopoiskdev_bot</a>. <br /> После получения токена, вам необходимо авторизоваться в документации, для этого нажмите на кнопку <strong>Authorize</strong> и введите токен в поле <strong>Value</strong>.<br /> После авторизации вы можете отправлять запросы к API, для этого нажмите на кнопку <strong>Try it out</strong> и заполните необходимые поля для составления нужного фильтра.<br /> После заполнения полей нажмите на кнопку <strong>Execute</strong> и получите ответ от API и пример запроса. </p> <h3>Как работать с API?</h3> <p> API работает по принципу REST, все запросы отправляются на адрес <code>https://api.kinopoisk.dev/</code> с указанием версии API и необходимого ресурса.<br /> Все запросы к API кинопоиска должны содержать заголовок <code>X-API-KEY</code> с вашим токеном. В противном случае вы получите ошибку <code>401</code>.<br /> При составлении запроса учитывайте, что все параметры должны быть в <code>query</code> и <code>path</code>. В зависимости от метода который вы используете. Например, вы хотите получить список фильмов за 2023 год в жанре <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?year=2023&genres.name=криминал</code>. Или вы хотите получить список фильмов с рейтингом выше 8, тогда ваш запрос будет выглядеть так: <code>https://api.kinopoisk.dev/v1.4/movie?rating.imdb=8-10</code>. Документация kinopoisk api может помочь вам составить нужный запрос, для этого воспользуйтесь ее конструктором. </p> <h3>Особенности синтекса query параметров</h3> <p> Ключи в query параметрах имеют разные типы значений. В зависимости от типа значения, вы можете использовать разные операторы для фильтрации для поиска максимально релевантного фильма, сериала и т.д. в базе. <br /> Поля с типом <code>Number</code> могут принимать значения в форматах: <code>rating.kp=1-10</code>, <code>rating.kp=1</code>, <code>year=2022&year=2023</code>. <br /> Поля с типом <code>Date</code> могут принимать значения в форматах: <code>premiere.russia=dd.mm.yyyy-dd.mm.yyyy</code>, <code>premiere.russia=dd.mm.yyyy</code>. <br /> Поля с типом <code>String</code> могут принимать значения в форматах: <code>genres.name=драма</code>, <code>genres.name=криминал</code>, <code>genres.name=криминал&genres.name=драма</code> <br/> Поля с типом <code>Boolean</code> могут принимать значения в форматах: <code>isSeries=true</code>, <code>isSeries=false</code>. <br /> Параметры жанров и стран могут принимать операторы <code>+</code> и <code>!</code>, для указания включаемых и исключаемых значений. Например, вы хотите получить список фильмов в жанрах <code>драма</code> и <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=+криминал</code>. Или вы хотите получить список фильмов с жанром <code>драма</code> и без жанра <code>криминал</code>, тогда ваш запрос будет выглядеть так: <code>genres.name=+драма&genres.name=!криминал</code>. <br /> </p> <p> Расшифровка операторов: <ul>   <li><code>!</code> - исключить. Этот символ нужно отправлять в кодированной форме <code>%21</code></li>   <li><code>+</code> - включить. Этот символ нужно отправлять в кодированной форме <code>%2B</code></li>   <li><code>-</code> - диапазон значений, используется в качестве разделителя.</li> </ul> </p>  <p>По вопросам работы с API обращайтесь в чат <a href=\"https://t.me/+jeHPZVXiLPFhODJi\">Developer Community KinopoiskDev</a>.</p>  <p>Если вы обнаружили ошибку или у вас есть предложения по улучшению, создавайте issue на <a href=\"https://github.com/mdwitr0/kinopoiskdev\">GitHub</a>.</p>  <h3>Полезные ссылки:</h3> <ul>   <li><a href=\"https://kinopoiskdev.readme.io\">Более удобная документация</a></li>   <li><a href=\"https://github.com/OpenMovieDB/kinopoiskdev_client\">JavaScript и TypeScript клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"https://github.com/odi1n/kinopoisk_dev\">Python клиент (Устарел, ждет обновления)</a></li>   <li><a href=\"/documentation-json\">OpenAPI Specification (JSON)</a></li>   <li><a href=\"/documentation-yaml\">OpenAPI Specification (YAML)</a></li> </ul> 

API version: 1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kinopoisk_api

import (
	"encoding/json"
)

// SearchMovieDtoV14 struct for SearchMovieDtoV14
type SearchMovieDtoV14 struct {
	Id float32 `json:"id"`
	Name string `json:"name"`
	AlternativeName string `json:"alternativeName"`
	EnName string `json:"enName"`
	Type string `json:"type"`
	Year float32 `json:"year"`
	Description string `json:"description"`
	ShortDescription string `json:"shortDescription"`
	MovieLength float32 `json:"movieLength"`
	Names []Name `json:"names"`
	ExternalId NullableSearchMovieDtoV14ExternalId `json:"externalId,omitempty"`
	Logo *Logo `json:"logo,omitempty"`
	Poster *ShortImage `json:"poster,omitempty"`
	Backdrop *ShortImage `json:"backdrop,omitempty"`
	Rating *Rating `json:"rating,omitempty"`
	Votes *Votes `json:"votes,omitempty"`
	Genres []ItemName `json:"genres,omitempty"`
	Countries []ItemName `json:"countries,omitempty"`
	ReleaseYears []YearRange `json:"releaseYears,omitempty"`
	IsSeries bool `json:"isSeries"`
	TicketsOnSale bool `json:"ticketsOnSale"`
	TotalSeriesLength float32 `json:"totalSeriesLength"`
	SeriesLength float32 `json:"seriesLength"`
	RatingMpaa string `json:"ratingMpaa"`
	AgeRating float32 `json:"ageRating"`
	Top10 NullableFloat32 `json:"top10,omitempty"`
	Top250 NullableFloat32 `json:"top250,omitempty"`
	TypeNumber float32 `json:"typeNumber"`
	Status string `json:"status"`
	InternalNames []string `json:"internalNames"`
	InternalRating float32 `json:"internalRating"`
	InternalVotes float32 `json:"internalVotes"`
}

// NewSearchMovieDtoV14 instantiates a new SearchMovieDtoV14 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMovieDtoV14(id float32, name string, alternativeName string, enName string, type_ string, year float32, description string, shortDescription string, movieLength float32, names []Name, isSeries bool, ticketsOnSale bool, totalSeriesLength float32, seriesLength float32, ratingMpaa string, ageRating float32, typeNumber float32, status string, internalNames []string, internalRating float32, internalVotes float32) *SearchMovieDtoV14 {
	this := SearchMovieDtoV14{}
	this.Id = id
	this.Name = name
	this.AlternativeName = alternativeName
	this.EnName = enName
	this.Type = type_
	this.Year = year
	this.Description = description
	this.ShortDescription = shortDescription
	this.MovieLength = movieLength
	this.Names = names
	this.IsSeries = isSeries
	this.TicketsOnSale = ticketsOnSale
	this.TotalSeriesLength = totalSeriesLength
	this.SeriesLength = seriesLength
	this.RatingMpaa = ratingMpaa
	this.AgeRating = ageRating
	this.TypeNumber = typeNumber
	this.Status = status
	this.InternalNames = internalNames
	this.InternalRating = internalRating
	this.InternalVotes = internalVotes
	return &this
}

// NewSearchMovieDtoV14WithDefaults instantiates a new SearchMovieDtoV14 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMovieDtoV14WithDefaults() *SearchMovieDtoV14 {
	this := SearchMovieDtoV14{}
	return &this
}

// GetId returns the Id field value
func (o *SearchMovieDtoV14) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SearchMovieDtoV14) SetId(v float32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *SearchMovieDtoV14) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SearchMovieDtoV14) SetName(v string) {
	o.Name = v
}

// GetAlternativeName returns the AlternativeName field value
func (o *SearchMovieDtoV14) GetAlternativeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AlternativeName
}

// GetAlternativeNameOk returns a tuple with the AlternativeName field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetAlternativeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlternativeName, true
}

// SetAlternativeName sets field value
func (o *SearchMovieDtoV14) SetAlternativeName(v string) {
	o.AlternativeName = v
}

// GetEnName returns the EnName field value
func (o *SearchMovieDtoV14) GetEnName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnName
}

// GetEnNameOk returns a tuple with the EnName field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetEnNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnName, true
}

// SetEnName sets field value
func (o *SearchMovieDtoV14) SetEnName(v string) {
	o.EnName = v
}

// GetType returns the Type field value
func (o *SearchMovieDtoV14) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SearchMovieDtoV14) SetType(v string) {
	o.Type = v
}

// GetYear returns the Year field value
func (o *SearchMovieDtoV14) GetYear() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetYearOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *SearchMovieDtoV14) SetYear(v float32) {
	o.Year = v
}

// GetDescription returns the Description field value
func (o *SearchMovieDtoV14) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SearchMovieDtoV14) SetDescription(v string) {
	o.Description = v
}

// GetShortDescription returns the ShortDescription field value
func (o *SearchMovieDtoV14) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value
func (o *SearchMovieDtoV14) SetShortDescription(v string) {
	o.ShortDescription = v
}

// GetMovieLength returns the MovieLength field value
func (o *SearchMovieDtoV14) GetMovieLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MovieLength
}

// GetMovieLengthOk returns a tuple with the MovieLength field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetMovieLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MovieLength, true
}

// SetMovieLength sets field value
func (o *SearchMovieDtoV14) SetMovieLength(v float32) {
	o.MovieLength = v
}

// GetNames returns the Names field value
func (o *SearchMovieDtoV14) GetNames() []Name {
	if o == nil {
		var ret []Name
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetNamesOk() ([]Name, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *SearchMovieDtoV14) SetNames(v []Name) {
	o.Names = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchMovieDtoV14) GetExternalId() SearchMovieDtoV14ExternalId {
	if o == nil || o.ExternalId.Get() == nil {
		var ret SearchMovieDtoV14ExternalId
		return ret
	}
	return *o.ExternalId.Get()
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchMovieDtoV14) GetExternalIdOk() (*SearchMovieDtoV14ExternalId, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalId.Get(), o.ExternalId.IsSet()
}

// HasExternalId returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasExternalId() bool {
	if o != nil && o.ExternalId.IsSet() {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given NullableSearchMovieDtoV14ExternalId and assigns it to the ExternalId field.
func (o *SearchMovieDtoV14) SetExternalId(v SearchMovieDtoV14ExternalId) {
	o.ExternalId.Set(&v)
}
// SetExternalIdNil sets the value for ExternalId to be an explicit nil
func (o *SearchMovieDtoV14) SetExternalIdNil() {
	o.ExternalId.Set(nil)
}

// UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
func (o *SearchMovieDtoV14) UnsetExternalId() {
	o.ExternalId.Unset()
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetLogo() Logo {
	if o == nil || o.Logo == nil {
		var ret Logo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetLogoOk() (*Logo, bool) {
	if o == nil || o.Logo == nil {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasLogo() bool {
	if o != nil && o.Logo != nil {
		return true
	}

	return false
}

// SetLogo gets a reference to the given Logo and assigns it to the Logo field.
func (o *SearchMovieDtoV14) SetLogo(v Logo) {
	o.Logo = &v
}

// GetPoster returns the Poster field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetPoster() ShortImage {
	if o == nil || o.Poster == nil {
		var ret ShortImage
		return ret
	}
	return *o.Poster
}

// GetPosterOk returns a tuple with the Poster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetPosterOk() (*ShortImage, bool) {
	if o == nil || o.Poster == nil {
		return nil, false
	}
	return o.Poster, true
}

// HasPoster returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasPoster() bool {
	if o != nil && o.Poster != nil {
		return true
	}

	return false
}

// SetPoster gets a reference to the given ShortImage and assigns it to the Poster field.
func (o *SearchMovieDtoV14) SetPoster(v ShortImage) {
	o.Poster = &v
}

// GetBackdrop returns the Backdrop field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetBackdrop() ShortImage {
	if o == nil || o.Backdrop == nil {
		var ret ShortImage
		return ret
	}
	return *o.Backdrop
}

// GetBackdropOk returns a tuple with the Backdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetBackdropOk() (*ShortImage, bool) {
	if o == nil || o.Backdrop == nil {
		return nil, false
	}
	return o.Backdrop, true
}

// HasBackdrop returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasBackdrop() bool {
	if o != nil && o.Backdrop != nil {
		return true
	}

	return false
}

// SetBackdrop gets a reference to the given ShortImage and assigns it to the Backdrop field.
func (o *SearchMovieDtoV14) SetBackdrop(v ShortImage) {
	o.Backdrop = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetRating() Rating {
	if o == nil || o.Rating == nil {
		var ret Rating
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetRatingOk() (*Rating, bool) {
	if o == nil || o.Rating == nil {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasRating() bool {
	if o != nil && o.Rating != nil {
		return true
	}

	return false
}

// SetRating gets a reference to the given Rating and assigns it to the Rating field.
func (o *SearchMovieDtoV14) SetRating(v Rating) {
	o.Rating = &v
}

// GetVotes returns the Votes field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetVotes() Votes {
	if o == nil || o.Votes == nil {
		var ret Votes
		return ret
	}
	return *o.Votes
}

// GetVotesOk returns a tuple with the Votes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetVotesOk() (*Votes, bool) {
	if o == nil || o.Votes == nil {
		return nil, false
	}
	return o.Votes, true
}

// HasVotes returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasVotes() bool {
	if o != nil && o.Votes != nil {
		return true
	}

	return false
}

// SetVotes gets a reference to the given Votes and assigns it to the Votes field.
func (o *SearchMovieDtoV14) SetVotes(v Votes) {
	o.Votes = &v
}

// GetGenres returns the Genres field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetGenres() []ItemName {
	if o == nil || o.Genres == nil {
		var ret []ItemName
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetGenresOk() ([]ItemName, bool) {
	if o == nil || o.Genres == nil {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasGenres() bool {
	if o != nil && o.Genres != nil {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []ItemName and assigns it to the Genres field.
func (o *SearchMovieDtoV14) SetGenres(v []ItemName) {
	o.Genres = v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetCountries() []ItemName {
	if o == nil || o.Countries == nil {
		var ret []ItemName
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetCountriesOk() ([]ItemName, bool) {
	if o == nil || o.Countries == nil {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasCountries() bool {
	if o != nil && o.Countries != nil {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []ItemName and assigns it to the Countries field.
func (o *SearchMovieDtoV14) SetCountries(v []ItemName) {
	o.Countries = v
}

// GetReleaseYears returns the ReleaseYears field value if set, zero value otherwise.
func (o *SearchMovieDtoV14) GetReleaseYears() []YearRange {
	if o == nil || o.ReleaseYears == nil {
		var ret []YearRange
		return ret
	}
	return o.ReleaseYears
}

// GetReleaseYearsOk returns a tuple with the ReleaseYears field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetReleaseYearsOk() ([]YearRange, bool) {
	if o == nil || o.ReleaseYears == nil {
		return nil, false
	}
	return o.ReleaseYears, true
}

// HasReleaseYears returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasReleaseYears() bool {
	if o != nil && o.ReleaseYears != nil {
		return true
	}

	return false
}

// SetReleaseYears gets a reference to the given []YearRange and assigns it to the ReleaseYears field.
func (o *SearchMovieDtoV14) SetReleaseYears(v []YearRange) {
	o.ReleaseYears = v
}

// GetIsSeries returns the IsSeries field value
func (o *SearchMovieDtoV14) GetIsSeries() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSeries
}

// GetIsSeriesOk returns a tuple with the IsSeries field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetIsSeriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSeries, true
}

// SetIsSeries sets field value
func (o *SearchMovieDtoV14) SetIsSeries(v bool) {
	o.IsSeries = v
}

// GetTicketsOnSale returns the TicketsOnSale field value
func (o *SearchMovieDtoV14) GetTicketsOnSale() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TicketsOnSale
}

// GetTicketsOnSaleOk returns a tuple with the TicketsOnSale field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetTicketsOnSaleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TicketsOnSale, true
}

// SetTicketsOnSale sets field value
func (o *SearchMovieDtoV14) SetTicketsOnSale(v bool) {
	o.TicketsOnSale = v
}

// GetTotalSeriesLength returns the TotalSeriesLength field value
func (o *SearchMovieDtoV14) GetTotalSeriesLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TotalSeriesLength
}

// GetTotalSeriesLengthOk returns a tuple with the TotalSeriesLength field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetTotalSeriesLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalSeriesLength, true
}

// SetTotalSeriesLength sets field value
func (o *SearchMovieDtoV14) SetTotalSeriesLength(v float32) {
	o.TotalSeriesLength = v
}

// GetSeriesLength returns the SeriesLength field value
func (o *SearchMovieDtoV14) GetSeriesLength() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SeriesLength
}

// GetSeriesLengthOk returns a tuple with the SeriesLength field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetSeriesLengthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SeriesLength, true
}

// SetSeriesLength sets field value
func (o *SearchMovieDtoV14) SetSeriesLength(v float32) {
	o.SeriesLength = v
}

// GetRatingMpaa returns the RatingMpaa field value
func (o *SearchMovieDtoV14) GetRatingMpaa() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RatingMpaa
}

// GetRatingMpaaOk returns a tuple with the RatingMpaa field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetRatingMpaaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingMpaa, true
}

// SetRatingMpaa sets field value
func (o *SearchMovieDtoV14) SetRatingMpaa(v string) {
	o.RatingMpaa = v
}

// GetAgeRating returns the AgeRating field value
func (o *SearchMovieDtoV14) GetAgeRating() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AgeRating
}

// GetAgeRatingOk returns a tuple with the AgeRating field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetAgeRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgeRating, true
}

// SetAgeRating sets field value
func (o *SearchMovieDtoV14) SetAgeRating(v float32) {
	o.AgeRating = v
}

// GetTop10 returns the Top10 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchMovieDtoV14) GetTop10() float32 {
	if o == nil || o.Top10.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Top10.Get()
}

// GetTop10Ok returns a tuple with the Top10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchMovieDtoV14) GetTop10Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Top10.Get(), o.Top10.IsSet()
}

// HasTop10 returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasTop10() bool {
	if o != nil && o.Top10.IsSet() {
		return true
	}

	return false
}

// SetTop10 gets a reference to the given NullableFloat32 and assigns it to the Top10 field.
func (o *SearchMovieDtoV14) SetTop10(v float32) {
	o.Top10.Set(&v)
}
// SetTop10Nil sets the value for Top10 to be an explicit nil
func (o *SearchMovieDtoV14) SetTop10Nil() {
	o.Top10.Set(nil)
}

// UnsetTop10 ensures that no value is present for Top10, not even an explicit nil
func (o *SearchMovieDtoV14) UnsetTop10() {
	o.Top10.Unset()
}

// GetTop250 returns the Top250 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SearchMovieDtoV14) GetTop250() float32 {
	if o == nil || o.Top250.Get() == nil {
		var ret float32
		return ret
	}
	return *o.Top250.Get()
}

// GetTop250Ok returns a tuple with the Top250 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchMovieDtoV14) GetTop250Ok() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Top250.Get(), o.Top250.IsSet()
}

// HasTop250 returns a boolean if a field has been set.
func (o *SearchMovieDtoV14) HasTop250() bool {
	if o != nil && o.Top250.IsSet() {
		return true
	}

	return false
}

// SetTop250 gets a reference to the given NullableFloat32 and assigns it to the Top250 field.
func (o *SearchMovieDtoV14) SetTop250(v float32) {
	o.Top250.Set(&v)
}
// SetTop250Nil sets the value for Top250 to be an explicit nil
func (o *SearchMovieDtoV14) SetTop250Nil() {
	o.Top250.Set(nil)
}

// UnsetTop250 ensures that no value is present for Top250, not even an explicit nil
func (o *SearchMovieDtoV14) UnsetTop250() {
	o.Top250.Unset()
}

// GetTypeNumber returns the TypeNumber field value
func (o *SearchMovieDtoV14) GetTypeNumber() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.TypeNumber
}

// GetTypeNumberOk returns a tuple with the TypeNumber field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetTypeNumberOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeNumber, true
}

// SetTypeNumber sets field value
func (o *SearchMovieDtoV14) SetTypeNumber(v float32) {
	o.TypeNumber = v
}

// GetStatus returns the Status field value
func (o *SearchMovieDtoV14) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SearchMovieDtoV14) SetStatus(v string) {
	o.Status = v
}

// GetInternalNames returns the InternalNames field value
func (o *SearchMovieDtoV14) GetInternalNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InternalNames
}

// GetInternalNamesOk returns a tuple with the InternalNames field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetInternalNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.InternalNames, true
}

// SetInternalNames sets field value
func (o *SearchMovieDtoV14) SetInternalNames(v []string) {
	o.InternalNames = v
}

// GetInternalRating returns the InternalRating field value
func (o *SearchMovieDtoV14) GetInternalRating() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InternalRating
}

// GetInternalRatingOk returns a tuple with the InternalRating field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetInternalRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalRating, true
}

// SetInternalRating sets field value
func (o *SearchMovieDtoV14) SetInternalRating(v float32) {
	o.InternalRating = v
}

// GetInternalVotes returns the InternalVotes field value
func (o *SearchMovieDtoV14) GetInternalVotes() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.InternalVotes
}

// GetInternalVotesOk returns a tuple with the InternalVotes field value
// and a boolean to check if the value has been set.
func (o *SearchMovieDtoV14) GetInternalVotesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InternalVotes, true
}

// SetInternalVotes sets field value
func (o *SearchMovieDtoV14) SetInternalVotes(v float32) {
	o.InternalVotes = v
}

func (o SearchMovieDtoV14) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["alternativeName"] = o.AlternativeName
	}
	if true {
		toSerialize["enName"] = o.EnName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["shortDescription"] = o.ShortDescription
	}
	if true {
		toSerialize["movieLength"] = o.MovieLength
	}
	if true {
		toSerialize["names"] = o.Names
	}
	if o.ExternalId.IsSet() {
		toSerialize["externalId"] = o.ExternalId.Get()
	}
	if o.Logo != nil {
		toSerialize["logo"] = o.Logo
	}
	if o.Poster != nil {
		toSerialize["poster"] = o.Poster
	}
	if o.Backdrop != nil {
		toSerialize["backdrop"] = o.Backdrop
	}
	if o.Rating != nil {
		toSerialize["rating"] = o.Rating
	}
	if o.Votes != nil {
		toSerialize["votes"] = o.Votes
	}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.Countries != nil {
		toSerialize["countries"] = o.Countries
	}
	if o.ReleaseYears != nil {
		toSerialize["releaseYears"] = o.ReleaseYears
	}
	if true {
		toSerialize["isSeries"] = o.IsSeries
	}
	if true {
		toSerialize["ticketsOnSale"] = o.TicketsOnSale
	}
	if true {
		toSerialize["totalSeriesLength"] = o.TotalSeriesLength
	}
	if true {
		toSerialize["seriesLength"] = o.SeriesLength
	}
	if true {
		toSerialize["ratingMpaa"] = o.RatingMpaa
	}
	if true {
		toSerialize["ageRating"] = o.AgeRating
	}
	if o.Top10.IsSet() {
		toSerialize["top10"] = o.Top10.Get()
	}
	if o.Top250.IsSet() {
		toSerialize["top250"] = o.Top250.Get()
	}
	if true {
		toSerialize["typeNumber"] = o.TypeNumber
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["internalNames"] = o.InternalNames
	}
	if true {
		toSerialize["internalRating"] = o.InternalRating
	}
	if true {
		toSerialize["internalVotes"] = o.InternalVotes
	}
	return json.Marshal(toSerialize)
}

type NullableSearchMovieDtoV14 struct {
	value *SearchMovieDtoV14
	isSet bool
}

func (v NullableSearchMovieDtoV14) Get() *SearchMovieDtoV14 {
	return v.value
}

func (v *NullableSearchMovieDtoV14) Set(val *SearchMovieDtoV14) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMovieDtoV14) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMovieDtoV14) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMovieDtoV14(val *SearchMovieDtoV14) *NullableSearchMovieDtoV14 {
	return &NullableSearchMovieDtoV14{value: val, isSet: true}
}

func (v NullableSearchMovieDtoV14) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMovieDtoV14) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


