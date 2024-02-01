# SeasonV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MovieId** | Pointer to **float32** |  | [optional] 
**Number** | Pointer to **float32** |  | [optional] 
**EpisodesCount** | Pointer to **float32** |  | [optional] 
**Episodes** | Pointer to [**[]EpisodeV14**](EpisodeV14.md) |  | [optional] 
**Poster** | Pointer to [**ShortImage**](ShortImage.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EnName** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnDescription** | Pointer to **string** |  | [optional] 
**AirDate** | Pointer to **string** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSeasonV14

`func NewSeasonV14(updatedAt time.Time, createdAt time.Time, ) *SeasonV14`

NewSeasonV14 instantiates a new SeasonV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonV14WithDefaults

`func NewSeasonV14WithDefaults() *SeasonV14`

NewSeasonV14WithDefaults instantiates a new SeasonV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMovieId

`func (o *SeasonV14) GetMovieId() float32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *SeasonV14) GetMovieIdOk() (*float32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *SeasonV14) SetMovieId(v float32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *SeasonV14) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetNumber

`func (o *SeasonV14) GetNumber() float32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SeasonV14) GetNumberOk() (*float32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SeasonV14) SetNumber(v float32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SeasonV14) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetEpisodesCount

`func (o *SeasonV14) GetEpisodesCount() float32`

GetEpisodesCount returns the EpisodesCount field if non-nil, zero value otherwise.

### GetEpisodesCountOk

`func (o *SeasonV14) GetEpisodesCountOk() (*float32, bool)`

GetEpisodesCountOk returns a tuple with the EpisodesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodesCount

`func (o *SeasonV14) SetEpisodesCount(v float32)`

SetEpisodesCount sets EpisodesCount field to given value.

### HasEpisodesCount

`func (o *SeasonV14) HasEpisodesCount() bool`

HasEpisodesCount returns a boolean if a field has been set.

### GetEpisodes

`func (o *SeasonV14) GetEpisodes() []EpisodeV14`

GetEpisodes returns the Episodes field if non-nil, zero value otherwise.

### GetEpisodesOk

`func (o *SeasonV14) GetEpisodesOk() (*[]EpisodeV14, bool)`

GetEpisodesOk returns a tuple with the Episodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpisodes

`func (o *SeasonV14) SetEpisodes(v []EpisodeV14)`

SetEpisodes sets Episodes field to given value.

### HasEpisodes

`func (o *SeasonV14) HasEpisodes() bool`

HasEpisodes returns a boolean if a field has been set.

### GetPoster

`func (o *SeasonV14) GetPoster() ShortImage`

GetPoster returns the Poster field if non-nil, zero value otherwise.

### GetPosterOk

`func (o *SeasonV14) GetPosterOk() (*ShortImage, bool)`

GetPosterOk returns a tuple with the Poster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoster

`func (o *SeasonV14) SetPoster(v ShortImage)`

SetPoster sets Poster field to given value.

### HasPoster

`func (o *SeasonV14) HasPoster() bool`

HasPoster returns a boolean if a field has been set.

### GetName

`func (o *SeasonV14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SeasonV14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SeasonV14) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SeasonV14) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnName

`func (o *SeasonV14) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *SeasonV14) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *SeasonV14) SetEnName(v string)`

SetEnName sets EnName field to given value.

### HasEnName

`func (o *SeasonV14) HasEnName() bool`

HasEnName returns a boolean if a field has been set.

### GetDuration

`func (o *SeasonV14) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SeasonV14) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SeasonV14) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *SeasonV14) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDescription

`func (o *SeasonV14) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SeasonV14) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SeasonV14) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SeasonV14) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnDescription

`func (o *SeasonV14) GetEnDescription() string`

GetEnDescription returns the EnDescription field if non-nil, zero value otherwise.

### GetEnDescriptionOk

`func (o *SeasonV14) GetEnDescriptionOk() (*string, bool)`

GetEnDescriptionOk returns a tuple with the EnDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnDescription

`func (o *SeasonV14) SetEnDescription(v string)`

SetEnDescription sets EnDescription field to given value.

### HasEnDescription

`func (o *SeasonV14) HasEnDescription() bool`

HasEnDescription returns a boolean if a field has been set.

### GetAirDate

`func (o *SeasonV14) GetAirDate() string`

GetAirDate returns the AirDate field if non-nil, zero value otherwise.

### GetAirDateOk

`func (o *SeasonV14) GetAirDateOk() (*string, bool)`

GetAirDateOk returns a tuple with the AirDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirDate

`func (o *SeasonV14) SetAirDate(v string)`

SetAirDate sets AirDate field to given value.

### HasAirDate

`func (o *SeasonV14) HasAirDate() bool`

HasAirDate returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SeasonV14) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SeasonV14) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SeasonV14) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *SeasonV14) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SeasonV14) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SeasonV14) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


