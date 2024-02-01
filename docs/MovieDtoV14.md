# MovieDtoV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** | Id фильма с кинопоиска | 
**ExternalId** | [**ExternalId**](ExternalId.md) |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**AlternativeName** | Pointer to **NullableString** |  | [optional] 
**EnName** | Pointer to **NullableString** |  | [optional] 
**Names** | [**[]Name**](Name.md) |  | 
**Type** | **string** | Тип тайтла. Доступны: movie | tv-series | cartoon | anime | animated-series | tv-show | 
**TypeNumber** | **float32** | Тип тайтла в числовом обозначении. Доступны: 1 (movie) | 2 (tv-series) | 3 (cartoon) | 4 (anime) | 5 (animated-series) | 6 (tv-show) | 
**Year** | Pointer to **NullableFloat32** | Год премьеры. При поиске по этому полю, можно использовать интервалы 1860-2030 | [optional] 
**Description** | Pointer to **NullableString** | Описание тайтла | [optional] 
**ShortDescription** | Pointer to **NullableString** | Сокращенное описание | [optional] 
**Slogan** | Pointer to **NullableString** | Слоган | [optional] 
**Status** | Pointer to **NullableString** | Статус релиза тайтла. Доступные значения: filming | pre-production | completed | announced | post-production | [optional] 
**Rating** | Pointer to [**Rating**](Rating.md) |  | [optional] 
**Votes** | Pointer to [**Votes**](Votes.md) |  | [optional] 
**MovieLength** | Pointer to **NullableFloat32** | Продолжительность фильма | [optional] 
**RatingMpaa** | Pointer to **NullableString** | Возрастной рейтинг по MPAA | [optional] 
**AgeRating** | Pointer to **NullableFloat32** | Возрастной рейтинг | [optional] 
**Logo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**Poster** | Pointer to [**ShortImage**](ShortImage.md) |  | [optional] 
**Backdrop** | Pointer to [**ShortImage**](ShortImage.md) |  | [optional] 
**Videos** | Pointer to [**VideoTypes**](VideoTypes.md) |  | [optional] 
**Genres** | Pointer to [**[]ItemName**](ItemName.md) |  | [optional] 
**Countries** | Pointer to [**[]ItemName**](ItemName.md) |  | [optional] 
**Persons** | Pointer to [**[]PersonInMovie**](PersonInMovie.md) |  | [optional] 
**ReviewInfo** | Pointer to [**ReviewInfo**](ReviewInfo.md) |  | [optional] 
**SeasonsInfo** | Pointer to [**[]SeasonInfo**](SeasonInfo.md) |  | [optional] 
**Budget** | Pointer to [**CurrencyValue**](CurrencyValue.md) |  | [optional] 
**Fees** | Pointer to [**Fees**](Fees.md) |  | [optional] 
**Premiere** | Pointer to [**Premiere**](Premiere.md) |  | [optional] 
**SimilarMovies** | Pointer to [**[]LinkedMovieV14**](LinkedMovieV14.md) |  | [optional] 
**SequelsAndPrequels** | Pointer to [**[]LinkedMovieV14**](LinkedMovieV14.md) |  | [optional] 
**Watchability** | Pointer to [**Watchability**](Watchability.md) |  | [optional] 
**ReleaseYears** | Pointer to [**[]YearRange**](YearRange.md) |  | [optional] 
**Top10** | Pointer to **NullableFloat32** | Позиция тайтла в топ 10. Чтобы найти фильмы участвующие в рейтинге используйте: &#x60;!null&#x60; | [optional] 
**Top250** | Pointer to **NullableFloat32** | Позиция тайтла в топ 250. Чтобы найти фильмы участвующие в рейтинге используйте: &#x60;!null&#x60; | [optional] 
**TicketsOnSale** | Pointer to **NullableBool** | Признак того, что тайтл находится в прокате | [optional] 
**TotalSeriesLength** | Pointer to **NullableFloat32** | Продолжительность всех серий | [optional] 
**SeriesLength** | Pointer to **NullableFloat32** | Средняя продолжительность серии | [optional] 
**IsSeries** | **bool** | Признак сериала | 
**Audience** | Pointer to [**[]Audience**](Audience.md) |  | [optional] 
**Lists** | Pointer to **[]string** | Список коллекций, в которых находится тайтл. | [optional] 
**Networks** | [**[]NetworksV14**](NetworksV14.md) |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**Facts** | [**[]FactInMovie**](FactInMovie.md) |  | 
**ImagesInfo** | [**Images**](Images.md) |  | 

## Methods

### NewMovieDtoV14

`func NewMovieDtoV14(id float32, externalId ExternalId, names []Name, type_ string, typeNumber float32, isSeries bool, networks []NetworksV14, updatedAt time.Time, createdAt time.Time, facts []FactInMovie, imagesInfo Images, ) *MovieDtoV14`

NewMovieDtoV14 instantiates a new MovieDtoV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieDtoV14WithDefaults

`func NewMovieDtoV14WithDefaults() *MovieDtoV14`

NewMovieDtoV14WithDefaults instantiates a new MovieDtoV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieDtoV14) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieDtoV14) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieDtoV14) SetId(v float32)`

SetId sets Id field to given value.


### GetExternalId

`func (o *MovieDtoV14) GetExternalId() ExternalId`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *MovieDtoV14) GetExternalIdOk() (*ExternalId, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *MovieDtoV14) SetExternalId(v ExternalId)`

SetExternalId sets ExternalId field to given value.


### GetName

`func (o *MovieDtoV14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MovieDtoV14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MovieDtoV14) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MovieDtoV14) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MovieDtoV14) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MovieDtoV14) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAlternativeName

`func (o *MovieDtoV14) GetAlternativeName() string`

GetAlternativeName returns the AlternativeName field if non-nil, zero value otherwise.

### GetAlternativeNameOk

`func (o *MovieDtoV14) GetAlternativeNameOk() (*string, bool)`

GetAlternativeNameOk returns a tuple with the AlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeName

`func (o *MovieDtoV14) SetAlternativeName(v string)`

SetAlternativeName sets AlternativeName field to given value.

### HasAlternativeName

`func (o *MovieDtoV14) HasAlternativeName() bool`

HasAlternativeName returns a boolean if a field has been set.

### SetAlternativeNameNil

`func (o *MovieDtoV14) SetAlternativeNameNil(b bool)`

 SetAlternativeNameNil sets the value for AlternativeName to be an explicit nil

### UnsetAlternativeName
`func (o *MovieDtoV14) UnsetAlternativeName()`

UnsetAlternativeName ensures that no value is present for AlternativeName, not even an explicit nil
### GetEnName

`func (o *MovieDtoV14) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *MovieDtoV14) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *MovieDtoV14) SetEnName(v string)`

SetEnName sets EnName field to given value.

### HasEnName

`func (o *MovieDtoV14) HasEnName() bool`

HasEnName returns a boolean if a field has been set.

### SetEnNameNil

`func (o *MovieDtoV14) SetEnNameNil(b bool)`

 SetEnNameNil sets the value for EnName to be an explicit nil

### UnsetEnName
`func (o *MovieDtoV14) UnsetEnName()`

UnsetEnName ensures that no value is present for EnName, not even an explicit nil
### GetNames

`func (o *MovieDtoV14) GetNames() []Name`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *MovieDtoV14) GetNamesOk() (*[]Name, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *MovieDtoV14) SetNames(v []Name)`

SetNames sets Names field to given value.


### GetType

`func (o *MovieDtoV14) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MovieDtoV14) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MovieDtoV14) SetType(v string)`

SetType sets Type field to given value.


### GetTypeNumber

`func (o *MovieDtoV14) GetTypeNumber() float32`

GetTypeNumber returns the TypeNumber field if non-nil, zero value otherwise.

### GetTypeNumberOk

`func (o *MovieDtoV14) GetTypeNumberOk() (*float32, bool)`

GetTypeNumberOk returns a tuple with the TypeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNumber

`func (o *MovieDtoV14) SetTypeNumber(v float32)`

SetTypeNumber sets TypeNumber field to given value.


### GetYear

`func (o *MovieDtoV14) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *MovieDtoV14) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *MovieDtoV14) SetYear(v float32)`

SetYear sets Year field to given value.

### HasYear

`func (o *MovieDtoV14) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *MovieDtoV14) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *MovieDtoV14) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetDescription

`func (o *MovieDtoV14) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MovieDtoV14) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MovieDtoV14) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MovieDtoV14) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MovieDtoV14) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MovieDtoV14) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetShortDescription

`func (o *MovieDtoV14) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *MovieDtoV14) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *MovieDtoV14) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *MovieDtoV14) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### SetShortDescriptionNil

`func (o *MovieDtoV14) SetShortDescriptionNil(b bool)`

 SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil

### UnsetShortDescription
`func (o *MovieDtoV14) UnsetShortDescription()`

UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
### GetSlogan

`func (o *MovieDtoV14) GetSlogan() string`

GetSlogan returns the Slogan field if non-nil, zero value otherwise.

### GetSloganOk

`func (o *MovieDtoV14) GetSloganOk() (*string, bool)`

GetSloganOk returns a tuple with the Slogan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlogan

`func (o *MovieDtoV14) SetSlogan(v string)`

SetSlogan sets Slogan field to given value.

### HasSlogan

`func (o *MovieDtoV14) HasSlogan() bool`

HasSlogan returns a boolean if a field has been set.

### SetSloganNil

`func (o *MovieDtoV14) SetSloganNil(b bool)`

 SetSloganNil sets the value for Slogan to be an explicit nil

### UnsetSlogan
`func (o *MovieDtoV14) UnsetSlogan()`

UnsetSlogan ensures that no value is present for Slogan, not even an explicit nil
### GetStatus

`func (o *MovieDtoV14) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MovieDtoV14) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MovieDtoV14) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MovieDtoV14) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MovieDtoV14) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MovieDtoV14) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRating

`func (o *MovieDtoV14) GetRating() Rating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MovieDtoV14) GetRatingOk() (*Rating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MovieDtoV14) SetRating(v Rating)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MovieDtoV14) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetVotes

`func (o *MovieDtoV14) GetVotes() Votes`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *MovieDtoV14) GetVotesOk() (*Votes, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *MovieDtoV14) SetVotes(v Votes)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *MovieDtoV14) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetMovieLength

`func (o *MovieDtoV14) GetMovieLength() float32`

GetMovieLength returns the MovieLength field if non-nil, zero value otherwise.

### GetMovieLengthOk

`func (o *MovieDtoV14) GetMovieLengthOk() (*float32, bool)`

GetMovieLengthOk returns a tuple with the MovieLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieLength

`func (o *MovieDtoV14) SetMovieLength(v float32)`

SetMovieLength sets MovieLength field to given value.

### HasMovieLength

`func (o *MovieDtoV14) HasMovieLength() bool`

HasMovieLength returns a boolean if a field has been set.

### SetMovieLengthNil

`func (o *MovieDtoV14) SetMovieLengthNil(b bool)`

 SetMovieLengthNil sets the value for MovieLength to be an explicit nil

### UnsetMovieLength
`func (o *MovieDtoV14) UnsetMovieLength()`

UnsetMovieLength ensures that no value is present for MovieLength, not even an explicit nil
### GetRatingMpaa

`func (o *MovieDtoV14) GetRatingMpaa() string`

GetRatingMpaa returns the RatingMpaa field if non-nil, zero value otherwise.

### GetRatingMpaaOk

`func (o *MovieDtoV14) GetRatingMpaaOk() (*string, bool)`

GetRatingMpaaOk returns a tuple with the RatingMpaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingMpaa

`func (o *MovieDtoV14) SetRatingMpaa(v string)`

SetRatingMpaa sets RatingMpaa field to given value.

### HasRatingMpaa

`func (o *MovieDtoV14) HasRatingMpaa() bool`

HasRatingMpaa returns a boolean if a field has been set.

### SetRatingMpaaNil

`func (o *MovieDtoV14) SetRatingMpaaNil(b bool)`

 SetRatingMpaaNil sets the value for RatingMpaa to be an explicit nil

### UnsetRatingMpaa
`func (o *MovieDtoV14) UnsetRatingMpaa()`

UnsetRatingMpaa ensures that no value is present for RatingMpaa, not even an explicit nil
### GetAgeRating

`func (o *MovieDtoV14) GetAgeRating() float32`

GetAgeRating returns the AgeRating field if non-nil, zero value otherwise.

### GetAgeRatingOk

`func (o *MovieDtoV14) GetAgeRatingOk() (*float32, bool)`

GetAgeRatingOk returns a tuple with the AgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRating

`func (o *MovieDtoV14) SetAgeRating(v float32)`

SetAgeRating sets AgeRating field to given value.

### HasAgeRating

`func (o *MovieDtoV14) HasAgeRating() bool`

HasAgeRating returns a boolean if a field has been set.

### SetAgeRatingNil

`func (o *MovieDtoV14) SetAgeRatingNil(b bool)`

 SetAgeRatingNil sets the value for AgeRating to be an explicit nil

### UnsetAgeRating
`func (o *MovieDtoV14) UnsetAgeRating()`

UnsetAgeRating ensures that no value is present for AgeRating, not even an explicit nil
### GetLogo

`func (o *MovieDtoV14) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *MovieDtoV14) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *MovieDtoV14) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *MovieDtoV14) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPoster

`func (o *MovieDtoV14) GetPoster() ShortImage`

GetPoster returns the Poster field if non-nil, zero value otherwise.

### GetPosterOk

`func (o *MovieDtoV14) GetPosterOk() (*ShortImage, bool)`

GetPosterOk returns a tuple with the Poster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoster

`func (o *MovieDtoV14) SetPoster(v ShortImage)`

SetPoster sets Poster field to given value.

### HasPoster

`func (o *MovieDtoV14) HasPoster() bool`

HasPoster returns a boolean if a field has been set.

### GetBackdrop

`func (o *MovieDtoV14) GetBackdrop() ShortImage`

GetBackdrop returns the Backdrop field if non-nil, zero value otherwise.

### GetBackdropOk

`func (o *MovieDtoV14) GetBackdropOk() (*ShortImage, bool)`

GetBackdropOk returns a tuple with the Backdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdrop

`func (o *MovieDtoV14) SetBackdrop(v ShortImage)`

SetBackdrop sets Backdrop field to given value.

### HasBackdrop

`func (o *MovieDtoV14) HasBackdrop() bool`

HasBackdrop returns a boolean if a field has been set.

### GetVideos

`func (o *MovieDtoV14) GetVideos() VideoTypes`

GetVideos returns the Videos field if non-nil, zero value otherwise.

### GetVideosOk

`func (o *MovieDtoV14) GetVideosOk() (*VideoTypes, bool)`

GetVideosOk returns a tuple with the Videos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideos

`func (o *MovieDtoV14) SetVideos(v VideoTypes)`

SetVideos sets Videos field to given value.

### HasVideos

`func (o *MovieDtoV14) HasVideos() bool`

HasVideos returns a boolean if a field has been set.

### GetGenres

`func (o *MovieDtoV14) GetGenres() []ItemName`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *MovieDtoV14) GetGenresOk() (*[]ItemName, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *MovieDtoV14) SetGenres(v []ItemName)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *MovieDtoV14) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCountries

`func (o *MovieDtoV14) GetCountries() []ItemName`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *MovieDtoV14) GetCountriesOk() (*[]ItemName, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *MovieDtoV14) SetCountries(v []ItemName)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *MovieDtoV14) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetPersons

`func (o *MovieDtoV14) GetPersons() []PersonInMovie`

GetPersons returns the Persons field if non-nil, zero value otherwise.

### GetPersonsOk

`func (o *MovieDtoV14) GetPersonsOk() (*[]PersonInMovie, bool)`

GetPersonsOk returns a tuple with the Persons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersons

`func (o *MovieDtoV14) SetPersons(v []PersonInMovie)`

SetPersons sets Persons field to given value.

### HasPersons

`func (o *MovieDtoV14) HasPersons() bool`

HasPersons returns a boolean if a field has been set.

### GetReviewInfo

`func (o *MovieDtoV14) GetReviewInfo() ReviewInfo`

GetReviewInfo returns the ReviewInfo field if non-nil, zero value otherwise.

### GetReviewInfoOk

`func (o *MovieDtoV14) GetReviewInfoOk() (*ReviewInfo, bool)`

GetReviewInfoOk returns a tuple with the ReviewInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewInfo

`func (o *MovieDtoV14) SetReviewInfo(v ReviewInfo)`

SetReviewInfo sets ReviewInfo field to given value.

### HasReviewInfo

`func (o *MovieDtoV14) HasReviewInfo() bool`

HasReviewInfo returns a boolean if a field has been set.

### GetSeasonsInfo

`func (o *MovieDtoV14) GetSeasonsInfo() []SeasonInfo`

GetSeasonsInfo returns the SeasonsInfo field if non-nil, zero value otherwise.

### GetSeasonsInfoOk

`func (o *MovieDtoV14) GetSeasonsInfoOk() (*[]SeasonInfo, bool)`

GetSeasonsInfoOk returns a tuple with the SeasonsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeasonsInfo

`func (o *MovieDtoV14) SetSeasonsInfo(v []SeasonInfo)`

SetSeasonsInfo sets SeasonsInfo field to given value.

### HasSeasonsInfo

`func (o *MovieDtoV14) HasSeasonsInfo() bool`

HasSeasonsInfo returns a boolean if a field has been set.

### GetBudget

`func (o *MovieDtoV14) GetBudget() CurrencyValue`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *MovieDtoV14) GetBudgetOk() (*CurrencyValue, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *MovieDtoV14) SetBudget(v CurrencyValue)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *MovieDtoV14) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetFees

`func (o *MovieDtoV14) GetFees() Fees`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *MovieDtoV14) GetFeesOk() (*Fees, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *MovieDtoV14) SetFees(v Fees)`

SetFees sets Fees field to given value.

### HasFees

`func (o *MovieDtoV14) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetPremiere

`func (o *MovieDtoV14) GetPremiere() Premiere`

GetPremiere returns the Premiere field if non-nil, zero value otherwise.

### GetPremiereOk

`func (o *MovieDtoV14) GetPremiereOk() (*Premiere, bool)`

GetPremiereOk returns a tuple with the Premiere field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiere

`func (o *MovieDtoV14) SetPremiere(v Premiere)`

SetPremiere sets Premiere field to given value.

### HasPremiere

`func (o *MovieDtoV14) HasPremiere() bool`

HasPremiere returns a boolean if a field has been set.

### GetSimilarMovies

`func (o *MovieDtoV14) GetSimilarMovies() []LinkedMovieV14`

GetSimilarMovies returns the SimilarMovies field if non-nil, zero value otherwise.

### GetSimilarMoviesOk

`func (o *MovieDtoV14) GetSimilarMoviesOk() (*[]LinkedMovieV14, bool)`

GetSimilarMoviesOk returns a tuple with the SimilarMovies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarMovies

`func (o *MovieDtoV14) SetSimilarMovies(v []LinkedMovieV14)`

SetSimilarMovies sets SimilarMovies field to given value.

### HasSimilarMovies

`func (o *MovieDtoV14) HasSimilarMovies() bool`

HasSimilarMovies returns a boolean if a field has been set.

### GetSequelsAndPrequels

`func (o *MovieDtoV14) GetSequelsAndPrequels() []LinkedMovieV14`

GetSequelsAndPrequels returns the SequelsAndPrequels field if non-nil, zero value otherwise.

### GetSequelsAndPrequelsOk

`func (o *MovieDtoV14) GetSequelsAndPrequelsOk() (*[]LinkedMovieV14, bool)`

GetSequelsAndPrequelsOk returns a tuple with the SequelsAndPrequels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequelsAndPrequels

`func (o *MovieDtoV14) SetSequelsAndPrequels(v []LinkedMovieV14)`

SetSequelsAndPrequels sets SequelsAndPrequels field to given value.

### HasSequelsAndPrequels

`func (o *MovieDtoV14) HasSequelsAndPrequels() bool`

HasSequelsAndPrequels returns a boolean if a field has been set.

### GetWatchability

`func (o *MovieDtoV14) GetWatchability() Watchability`

GetWatchability returns the Watchability field if non-nil, zero value otherwise.

### GetWatchabilityOk

`func (o *MovieDtoV14) GetWatchabilityOk() (*Watchability, bool)`

GetWatchabilityOk returns a tuple with the Watchability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchability

`func (o *MovieDtoV14) SetWatchability(v Watchability)`

SetWatchability sets Watchability field to given value.

### HasWatchability

`func (o *MovieDtoV14) HasWatchability() bool`

HasWatchability returns a boolean if a field has been set.

### GetReleaseYears

`func (o *MovieDtoV14) GetReleaseYears() []YearRange`

GetReleaseYears returns the ReleaseYears field if non-nil, zero value otherwise.

### GetReleaseYearsOk

`func (o *MovieDtoV14) GetReleaseYearsOk() (*[]YearRange, bool)`

GetReleaseYearsOk returns a tuple with the ReleaseYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseYears

`func (o *MovieDtoV14) SetReleaseYears(v []YearRange)`

SetReleaseYears sets ReleaseYears field to given value.

### HasReleaseYears

`func (o *MovieDtoV14) HasReleaseYears() bool`

HasReleaseYears returns a boolean if a field has been set.

### GetTop10

`func (o *MovieDtoV14) GetTop10() float32`

GetTop10 returns the Top10 field if non-nil, zero value otherwise.

### GetTop10Ok

`func (o *MovieDtoV14) GetTop10Ok() (*float32, bool)`

GetTop10Ok returns a tuple with the Top10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10

`func (o *MovieDtoV14) SetTop10(v float32)`

SetTop10 sets Top10 field to given value.

### HasTop10

`func (o *MovieDtoV14) HasTop10() bool`

HasTop10 returns a boolean if a field has been set.

### SetTop10Nil

`func (o *MovieDtoV14) SetTop10Nil(b bool)`

 SetTop10Nil sets the value for Top10 to be an explicit nil

### UnsetTop10
`func (o *MovieDtoV14) UnsetTop10()`

UnsetTop10 ensures that no value is present for Top10, not even an explicit nil
### GetTop250

`func (o *MovieDtoV14) GetTop250() float32`

GetTop250 returns the Top250 field if non-nil, zero value otherwise.

### GetTop250Ok

`func (o *MovieDtoV14) GetTop250Ok() (*float32, bool)`

GetTop250Ok returns a tuple with the Top250 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop250

`func (o *MovieDtoV14) SetTop250(v float32)`

SetTop250 sets Top250 field to given value.

### HasTop250

`func (o *MovieDtoV14) HasTop250() bool`

HasTop250 returns a boolean if a field has been set.

### SetTop250Nil

`func (o *MovieDtoV14) SetTop250Nil(b bool)`

 SetTop250Nil sets the value for Top250 to be an explicit nil

### UnsetTop250
`func (o *MovieDtoV14) UnsetTop250()`

UnsetTop250 ensures that no value is present for Top250, not even an explicit nil
### GetTicketsOnSale

`func (o *MovieDtoV14) GetTicketsOnSale() bool`

GetTicketsOnSale returns the TicketsOnSale field if non-nil, zero value otherwise.

### GetTicketsOnSaleOk

`func (o *MovieDtoV14) GetTicketsOnSaleOk() (*bool, bool)`

GetTicketsOnSaleOk returns a tuple with the TicketsOnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketsOnSale

`func (o *MovieDtoV14) SetTicketsOnSale(v bool)`

SetTicketsOnSale sets TicketsOnSale field to given value.

### HasTicketsOnSale

`func (o *MovieDtoV14) HasTicketsOnSale() bool`

HasTicketsOnSale returns a boolean if a field has been set.

### SetTicketsOnSaleNil

`func (o *MovieDtoV14) SetTicketsOnSaleNil(b bool)`

 SetTicketsOnSaleNil sets the value for TicketsOnSale to be an explicit nil

### UnsetTicketsOnSale
`func (o *MovieDtoV14) UnsetTicketsOnSale()`

UnsetTicketsOnSale ensures that no value is present for TicketsOnSale, not even an explicit nil
### GetTotalSeriesLength

`func (o *MovieDtoV14) GetTotalSeriesLength() float32`

GetTotalSeriesLength returns the TotalSeriesLength field if non-nil, zero value otherwise.

### GetTotalSeriesLengthOk

`func (o *MovieDtoV14) GetTotalSeriesLengthOk() (*float32, bool)`

GetTotalSeriesLengthOk returns a tuple with the TotalSeriesLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeriesLength

`func (o *MovieDtoV14) SetTotalSeriesLength(v float32)`

SetTotalSeriesLength sets TotalSeriesLength field to given value.

### HasTotalSeriesLength

`func (o *MovieDtoV14) HasTotalSeriesLength() bool`

HasTotalSeriesLength returns a boolean if a field has been set.

### SetTotalSeriesLengthNil

`func (o *MovieDtoV14) SetTotalSeriesLengthNil(b bool)`

 SetTotalSeriesLengthNil sets the value for TotalSeriesLength to be an explicit nil

### UnsetTotalSeriesLength
`func (o *MovieDtoV14) UnsetTotalSeriesLength()`

UnsetTotalSeriesLength ensures that no value is present for TotalSeriesLength, not even an explicit nil
### GetSeriesLength

`func (o *MovieDtoV14) GetSeriesLength() float32`

GetSeriesLength returns the SeriesLength field if non-nil, zero value otherwise.

### GetSeriesLengthOk

`func (o *MovieDtoV14) GetSeriesLengthOk() (*float32, bool)`

GetSeriesLengthOk returns a tuple with the SeriesLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLength

`func (o *MovieDtoV14) SetSeriesLength(v float32)`

SetSeriesLength sets SeriesLength field to given value.

### HasSeriesLength

`func (o *MovieDtoV14) HasSeriesLength() bool`

HasSeriesLength returns a boolean if a field has been set.

### SetSeriesLengthNil

`func (o *MovieDtoV14) SetSeriesLengthNil(b bool)`

 SetSeriesLengthNil sets the value for SeriesLength to be an explicit nil

### UnsetSeriesLength
`func (o *MovieDtoV14) UnsetSeriesLength()`

UnsetSeriesLength ensures that no value is present for SeriesLength, not even an explicit nil
### GetIsSeries

`func (o *MovieDtoV14) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *MovieDtoV14) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *MovieDtoV14) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.


### GetAudience

`func (o *MovieDtoV14) GetAudience() []Audience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *MovieDtoV14) GetAudienceOk() (*[]Audience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *MovieDtoV14) SetAudience(v []Audience)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *MovieDtoV14) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *MovieDtoV14) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *MovieDtoV14) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetLists

`func (o *MovieDtoV14) GetLists() []string`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *MovieDtoV14) GetListsOk() (*[]string, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *MovieDtoV14) SetLists(v []string)`

SetLists sets Lists field to given value.

### HasLists

`func (o *MovieDtoV14) HasLists() bool`

HasLists returns a boolean if a field has been set.

### SetListsNil

`func (o *MovieDtoV14) SetListsNil(b bool)`

 SetListsNil sets the value for Lists to be an explicit nil

### UnsetLists
`func (o *MovieDtoV14) UnsetLists()`

UnsetLists ensures that no value is present for Lists, not even an explicit nil
### GetNetworks

`func (o *MovieDtoV14) GetNetworks() []NetworksV14`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *MovieDtoV14) GetNetworksOk() (*[]NetworksV14, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *MovieDtoV14) SetNetworks(v []NetworksV14)`

SetNetworks sets Networks field to given value.


### GetUpdatedAt

`func (o *MovieDtoV14) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MovieDtoV14) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MovieDtoV14) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *MovieDtoV14) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MovieDtoV14) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MovieDtoV14) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFacts

`func (o *MovieDtoV14) GetFacts() []FactInMovie`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *MovieDtoV14) GetFactsOk() (*[]FactInMovie, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *MovieDtoV14) SetFacts(v []FactInMovie)`

SetFacts sets Facts field to given value.


### GetImagesInfo

`func (o *MovieDtoV14) GetImagesInfo() Images`

GetImagesInfo returns the ImagesInfo field if non-nil, zero value otherwise.

### GetImagesInfoOk

`func (o *MovieDtoV14) GetImagesInfoOk() (*Images, bool)`

GetImagesInfoOk returns a tuple with the ImagesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagesInfo

`func (o *MovieDtoV14) SetImagesInfo(v Images)`

SetImagesInfo sets ImagesInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


