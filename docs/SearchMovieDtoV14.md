# SearchMovieDtoV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**AlternativeName** | **string** |  | 
**EnName** | **string** |  | 
**Type** | **string** |  | 
**Year** | **float32** |  | 
**Description** | **string** |  | 
**ShortDescription** | **string** |  | 
**MovieLength** | **float32** |  | 
**Names** | [**[]Name**](Name.md) |  | 
**ExternalId** | Pointer to [**NullableSearchMovieDtoV14ExternalId**](SearchMovieDtoV14ExternalId.md) |  | [optional] 
**Logo** | Pointer to [**Logo**](Logo.md) |  | [optional] 
**Poster** | Pointer to [**ShortImage**](ShortImage.md) |  | [optional] 
**Backdrop** | Pointer to [**ShortImage**](ShortImage.md) |  | [optional] 
**Rating** | Pointer to [**Rating**](Rating.md) |  | [optional] 
**Votes** | Pointer to [**Votes**](Votes.md) |  | [optional] 
**Genres** | Pointer to [**[]ItemName**](ItemName.md) |  | [optional] 
**Countries** | Pointer to [**[]ItemName**](ItemName.md) |  | [optional] 
**ReleaseYears** | Pointer to [**[]YearRange**](YearRange.md) |  | [optional] 
**IsSeries** | **bool** |  | 
**TicketsOnSale** | **bool** |  | 
**TotalSeriesLength** | **float32** |  | 
**SeriesLength** | **float32** |  | 
**RatingMpaa** | **string** |  | 
**AgeRating** | **float32** |  | 
**Top10** | Pointer to **NullableFloat32** |  | [optional] 
**Top250** | Pointer to **NullableFloat32** |  | [optional] 
**TypeNumber** | **float32** |  | 
**Status** | **string** |  | 
**InternalNames** | **[]string** |  | 
**InternalRating** | **float32** |  | 
**InternalVotes** | **float32** |  | 

## Methods

### NewSearchMovieDtoV14

`func NewSearchMovieDtoV14(id float32, name string, alternativeName string, enName string, type_ string, year float32, description string, shortDescription string, movieLength float32, names []Name, isSeries bool, ticketsOnSale bool, totalSeriesLength float32, seriesLength float32, ratingMpaa string, ageRating float32, typeNumber float32, status string, internalNames []string, internalRating float32, internalVotes float32, ) *SearchMovieDtoV14`

NewSearchMovieDtoV14 instantiates a new SearchMovieDtoV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchMovieDtoV14WithDefaults

`func NewSearchMovieDtoV14WithDefaults() *SearchMovieDtoV14`

NewSearchMovieDtoV14WithDefaults instantiates a new SearchMovieDtoV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchMovieDtoV14) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchMovieDtoV14) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchMovieDtoV14) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *SearchMovieDtoV14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchMovieDtoV14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchMovieDtoV14) SetName(v string)`

SetName sets Name field to given value.


### GetAlternativeName

`func (o *SearchMovieDtoV14) GetAlternativeName() string`

GetAlternativeName returns the AlternativeName field if non-nil, zero value otherwise.

### GetAlternativeNameOk

`func (o *SearchMovieDtoV14) GetAlternativeNameOk() (*string, bool)`

GetAlternativeNameOk returns a tuple with the AlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeName

`func (o *SearchMovieDtoV14) SetAlternativeName(v string)`

SetAlternativeName sets AlternativeName field to given value.


### GetEnName

`func (o *SearchMovieDtoV14) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *SearchMovieDtoV14) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *SearchMovieDtoV14) SetEnName(v string)`

SetEnName sets EnName field to given value.


### GetType

`func (o *SearchMovieDtoV14) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SearchMovieDtoV14) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SearchMovieDtoV14) SetType(v string)`

SetType sets Type field to given value.


### GetYear

`func (o *SearchMovieDtoV14) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *SearchMovieDtoV14) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *SearchMovieDtoV14) SetYear(v float32)`

SetYear sets Year field to given value.


### GetDescription

`func (o *SearchMovieDtoV14) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SearchMovieDtoV14) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SearchMovieDtoV14) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetShortDescription

`func (o *SearchMovieDtoV14) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *SearchMovieDtoV14) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *SearchMovieDtoV14) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.


### GetMovieLength

`func (o *SearchMovieDtoV14) GetMovieLength() float32`

GetMovieLength returns the MovieLength field if non-nil, zero value otherwise.

### GetMovieLengthOk

`func (o *SearchMovieDtoV14) GetMovieLengthOk() (*float32, bool)`

GetMovieLengthOk returns a tuple with the MovieLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieLength

`func (o *SearchMovieDtoV14) SetMovieLength(v float32)`

SetMovieLength sets MovieLength field to given value.


### GetNames

`func (o *SearchMovieDtoV14) GetNames() []Name`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *SearchMovieDtoV14) GetNamesOk() (*[]Name, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *SearchMovieDtoV14) SetNames(v []Name)`

SetNames sets Names field to given value.


### GetExternalId

`func (o *SearchMovieDtoV14) GetExternalId() SearchMovieDtoV14ExternalId`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SearchMovieDtoV14) GetExternalIdOk() (*SearchMovieDtoV14ExternalId, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SearchMovieDtoV14) SetExternalId(v SearchMovieDtoV14ExternalId)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SearchMovieDtoV14) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *SearchMovieDtoV14) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *SearchMovieDtoV14) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetLogo

`func (o *SearchMovieDtoV14) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *SearchMovieDtoV14) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *SearchMovieDtoV14) SetLogo(v Logo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *SearchMovieDtoV14) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetPoster

`func (o *SearchMovieDtoV14) GetPoster() ShortImage`

GetPoster returns the Poster field if non-nil, zero value otherwise.

### GetPosterOk

`func (o *SearchMovieDtoV14) GetPosterOk() (*ShortImage, bool)`

GetPosterOk returns a tuple with the Poster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoster

`func (o *SearchMovieDtoV14) SetPoster(v ShortImage)`

SetPoster sets Poster field to given value.

### HasPoster

`func (o *SearchMovieDtoV14) HasPoster() bool`

HasPoster returns a boolean if a field has been set.

### GetBackdrop

`func (o *SearchMovieDtoV14) GetBackdrop() ShortImage`

GetBackdrop returns the Backdrop field if non-nil, zero value otherwise.

### GetBackdropOk

`func (o *SearchMovieDtoV14) GetBackdropOk() (*ShortImage, bool)`

GetBackdropOk returns a tuple with the Backdrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackdrop

`func (o *SearchMovieDtoV14) SetBackdrop(v ShortImage)`

SetBackdrop sets Backdrop field to given value.

### HasBackdrop

`func (o *SearchMovieDtoV14) HasBackdrop() bool`

HasBackdrop returns a boolean if a field has been set.

### GetRating

`func (o *SearchMovieDtoV14) GetRating() Rating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *SearchMovieDtoV14) GetRatingOk() (*Rating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *SearchMovieDtoV14) SetRating(v Rating)`

SetRating sets Rating field to given value.

### HasRating

`func (o *SearchMovieDtoV14) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetVotes

`func (o *SearchMovieDtoV14) GetVotes() Votes`

GetVotes returns the Votes field if non-nil, zero value otherwise.

### GetVotesOk

`func (o *SearchMovieDtoV14) GetVotesOk() (*Votes, bool)`

GetVotesOk returns a tuple with the Votes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVotes

`func (o *SearchMovieDtoV14) SetVotes(v Votes)`

SetVotes sets Votes field to given value.

### HasVotes

`func (o *SearchMovieDtoV14) HasVotes() bool`

HasVotes returns a boolean if a field has been set.

### GetGenres

`func (o *SearchMovieDtoV14) GetGenres() []ItemName`

GetGenres returns the Genres field if non-nil, zero value otherwise.

### GetGenresOk

`func (o *SearchMovieDtoV14) GetGenresOk() (*[]ItemName, bool)`

GetGenresOk returns a tuple with the Genres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenres

`func (o *SearchMovieDtoV14) SetGenres(v []ItemName)`

SetGenres sets Genres field to given value.

### HasGenres

`func (o *SearchMovieDtoV14) HasGenres() bool`

HasGenres returns a boolean if a field has been set.

### GetCountries

`func (o *SearchMovieDtoV14) GetCountries() []ItemName`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *SearchMovieDtoV14) GetCountriesOk() (*[]ItemName, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *SearchMovieDtoV14) SetCountries(v []ItemName)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *SearchMovieDtoV14) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetReleaseYears

`func (o *SearchMovieDtoV14) GetReleaseYears() []YearRange`

GetReleaseYears returns the ReleaseYears field if non-nil, zero value otherwise.

### GetReleaseYearsOk

`func (o *SearchMovieDtoV14) GetReleaseYearsOk() (*[]YearRange, bool)`

GetReleaseYearsOk returns a tuple with the ReleaseYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseYears

`func (o *SearchMovieDtoV14) SetReleaseYears(v []YearRange)`

SetReleaseYears sets ReleaseYears field to given value.

### HasReleaseYears

`func (o *SearchMovieDtoV14) HasReleaseYears() bool`

HasReleaseYears returns a boolean if a field has been set.

### GetIsSeries

`func (o *SearchMovieDtoV14) GetIsSeries() bool`

GetIsSeries returns the IsSeries field if non-nil, zero value otherwise.

### GetIsSeriesOk

`func (o *SearchMovieDtoV14) GetIsSeriesOk() (*bool, bool)`

GetIsSeriesOk returns a tuple with the IsSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeries

`func (o *SearchMovieDtoV14) SetIsSeries(v bool)`

SetIsSeries sets IsSeries field to given value.


### GetTicketsOnSale

`func (o *SearchMovieDtoV14) GetTicketsOnSale() bool`

GetTicketsOnSale returns the TicketsOnSale field if non-nil, zero value otherwise.

### GetTicketsOnSaleOk

`func (o *SearchMovieDtoV14) GetTicketsOnSaleOk() (*bool, bool)`

GetTicketsOnSaleOk returns a tuple with the TicketsOnSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketsOnSale

`func (o *SearchMovieDtoV14) SetTicketsOnSale(v bool)`

SetTicketsOnSale sets TicketsOnSale field to given value.


### GetTotalSeriesLength

`func (o *SearchMovieDtoV14) GetTotalSeriesLength() float32`

GetTotalSeriesLength returns the TotalSeriesLength field if non-nil, zero value otherwise.

### GetTotalSeriesLengthOk

`func (o *SearchMovieDtoV14) GetTotalSeriesLengthOk() (*float32, bool)`

GetTotalSeriesLengthOk returns a tuple with the TotalSeriesLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSeriesLength

`func (o *SearchMovieDtoV14) SetTotalSeriesLength(v float32)`

SetTotalSeriesLength sets TotalSeriesLength field to given value.


### GetSeriesLength

`func (o *SearchMovieDtoV14) GetSeriesLength() float32`

GetSeriesLength returns the SeriesLength field if non-nil, zero value otherwise.

### GetSeriesLengthOk

`func (o *SearchMovieDtoV14) GetSeriesLengthOk() (*float32, bool)`

GetSeriesLengthOk returns a tuple with the SeriesLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesLength

`func (o *SearchMovieDtoV14) SetSeriesLength(v float32)`

SetSeriesLength sets SeriesLength field to given value.


### GetRatingMpaa

`func (o *SearchMovieDtoV14) GetRatingMpaa() string`

GetRatingMpaa returns the RatingMpaa field if non-nil, zero value otherwise.

### GetRatingMpaaOk

`func (o *SearchMovieDtoV14) GetRatingMpaaOk() (*string, bool)`

GetRatingMpaaOk returns a tuple with the RatingMpaa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingMpaa

`func (o *SearchMovieDtoV14) SetRatingMpaa(v string)`

SetRatingMpaa sets RatingMpaa field to given value.


### GetAgeRating

`func (o *SearchMovieDtoV14) GetAgeRating() float32`

GetAgeRating returns the AgeRating field if non-nil, zero value otherwise.

### GetAgeRatingOk

`func (o *SearchMovieDtoV14) GetAgeRatingOk() (*float32, bool)`

GetAgeRatingOk returns a tuple with the AgeRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeRating

`func (o *SearchMovieDtoV14) SetAgeRating(v float32)`

SetAgeRating sets AgeRating field to given value.


### GetTop10

`func (o *SearchMovieDtoV14) GetTop10() float32`

GetTop10 returns the Top10 field if non-nil, zero value otherwise.

### GetTop10Ok

`func (o *SearchMovieDtoV14) GetTop10Ok() (*float32, bool)`

GetTop10Ok returns a tuple with the Top10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop10

`func (o *SearchMovieDtoV14) SetTop10(v float32)`

SetTop10 sets Top10 field to given value.

### HasTop10

`func (o *SearchMovieDtoV14) HasTop10() bool`

HasTop10 returns a boolean if a field has been set.

### SetTop10Nil

`func (o *SearchMovieDtoV14) SetTop10Nil(b bool)`

 SetTop10Nil sets the value for Top10 to be an explicit nil

### UnsetTop10
`func (o *SearchMovieDtoV14) UnsetTop10()`

UnsetTop10 ensures that no value is present for Top10, not even an explicit nil
### GetTop250

`func (o *SearchMovieDtoV14) GetTop250() float32`

GetTop250 returns the Top250 field if non-nil, zero value otherwise.

### GetTop250Ok

`func (o *SearchMovieDtoV14) GetTop250Ok() (*float32, bool)`

GetTop250Ok returns a tuple with the Top250 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop250

`func (o *SearchMovieDtoV14) SetTop250(v float32)`

SetTop250 sets Top250 field to given value.

### HasTop250

`func (o *SearchMovieDtoV14) HasTop250() bool`

HasTop250 returns a boolean if a field has been set.

### SetTop250Nil

`func (o *SearchMovieDtoV14) SetTop250Nil(b bool)`

 SetTop250Nil sets the value for Top250 to be an explicit nil

### UnsetTop250
`func (o *SearchMovieDtoV14) UnsetTop250()`

UnsetTop250 ensures that no value is present for Top250, not even an explicit nil
### GetTypeNumber

`func (o *SearchMovieDtoV14) GetTypeNumber() float32`

GetTypeNumber returns the TypeNumber field if non-nil, zero value otherwise.

### GetTypeNumberOk

`func (o *SearchMovieDtoV14) GetTypeNumberOk() (*float32, bool)`

GetTypeNumberOk returns a tuple with the TypeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeNumber

`func (o *SearchMovieDtoV14) SetTypeNumber(v float32)`

SetTypeNumber sets TypeNumber field to given value.


### GetStatus

`func (o *SearchMovieDtoV14) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchMovieDtoV14) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchMovieDtoV14) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetInternalNames

`func (o *SearchMovieDtoV14) GetInternalNames() []string`

GetInternalNames returns the InternalNames field if non-nil, zero value otherwise.

### GetInternalNamesOk

`func (o *SearchMovieDtoV14) GetInternalNamesOk() (*[]string, bool)`

GetInternalNamesOk returns a tuple with the InternalNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNames

`func (o *SearchMovieDtoV14) SetInternalNames(v []string)`

SetInternalNames sets InternalNames field to given value.


### GetInternalRating

`func (o *SearchMovieDtoV14) GetInternalRating() float32`

GetInternalRating returns the InternalRating field if non-nil, zero value otherwise.

### GetInternalRatingOk

`func (o *SearchMovieDtoV14) GetInternalRatingOk() (*float32, bool)`

GetInternalRatingOk returns a tuple with the InternalRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalRating

`func (o *SearchMovieDtoV14) SetInternalRating(v float32)`

SetInternalRating sets InternalRating field to given value.


### GetInternalVotes

`func (o *SearchMovieDtoV14) GetInternalVotes() float32`

GetInternalVotes returns the InternalVotes field if non-nil, zero value otherwise.

### GetInternalVotesOk

`func (o *SearchMovieDtoV14) GetInternalVotesOk() (*float32, bool)`

GetInternalVotesOk returns a tuple with the InternalVotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalVotes

`func (o *SearchMovieDtoV14) SetInternalVotes(v float32)`

SetInternalVotes sets InternalVotes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


