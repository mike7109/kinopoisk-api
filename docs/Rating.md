# Rating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kp** | Pointer to **NullableFloat32** | Рейтинг кинопоиска | [optional] 
**Imdb** | Pointer to **NullableFloat32** | Рейтинг IMDB | [optional] 
**Tmdb** | Pointer to **NullableFloat32** | Рейтинг TMDB | [optional] 
**FilmCritics** | Pointer to **NullableFloat32** | Рейтинг кинокритиков | [optional] 
**RussianFilmCritics** | Pointer to **NullableFloat32** | Рейтинг кинокритиков из РФ | [optional] 
**Await** | Pointer to **NullableFloat32** | Рейтинг основанный на ожиданиях пользователей | [optional] 

## Methods

### NewRating

`func NewRating() *Rating`

NewRating instantiates a new Rating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingWithDefaults

`func NewRatingWithDefaults() *Rating`

NewRatingWithDefaults instantiates a new Rating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKp

`func (o *Rating) GetKp() float32`

GetKp returns the Kp field if non-nil, zero value otherwise.

### GetKpOk

`func (o *Rating) GetKpOk() (*float32, bool)`

GetKpOk returns a tuple with the Kp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKp

`func (o *Rating) SetKp(v float32)`

SetKp sets Kp field to given value.

### HasKp

`func (o *Rating) HasKp() bool`

HasKp returns a boolean if a field has been set.

### SetKpNil

`func (o *Rating) SetKpNil(b bool)`

 SetKpNil sets the value for Kp to be an explicit nil

### UnsetKp
`func (o *Rating) UnsetKp()`

UnsetKp ensures that no value is present for Kp, not even an explicit nil
### GetImdb

`func (o *Rating) GetImdb() float32`

GetImdb returns the Imdb field if non-nil, zero value otherwise.

### GetImdbOk

`func (o *Rating) GetImdbOk() (*float32, bool)`

GetImdbOk returns a tuple with the Imdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdb

`func (o *Rating) SetImdb(v float32)`

SetImdb sets Imdb field to given value.

### HasImdb

`func (o *Rating) HasImdb() bool`

HasImdb returns a boolean if a field has been set.

### SetImdbNil

`func (o *Rating) SetImdbNil(b bool)`

 SetImdbNil sets the value for Imdb to be an explicit nil

### UnsetImdb
`func (o *Rating) UnsetImdb()`

UnsetImdb ensures that no value is present for Imdb, not even an explicit nil
### GetTmdb

`func (o *Rating) GetTmdb() float32`

GetTmdb returns the Tmdb field if non-nil, zero value otherwise.

### GetTmdbOk

`func (o *Rating) GetTmdbOk() (*float32, bool)`

GetTmdbOk returns a tuple with the Tmdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdb

`func (o *Rating) SetTmdb(v float32)`

SetTmdb sets Tmdb field to given value.

### HasTmdb

`func (o *Rating) HasTmdb() bool`

HasTmdb returns a boolean if a field has been set.

### SetTmdbNil

`func (o *Rating) SetTmdbNil(b bool)`

 SetTmdbNil sets the value for Tmdb to be an explicit nil

### UnsetTmdb
`func (o *Rating) UnsetTmdb()`

UnsetTmdb ensures that no value is present for Tmdb, not even an explicit nil
### GetFilmCritics

`func (o *Rating) GetFilmCritics() float32`

GetFilmCritics returns the FilmCritics field if non-nil, zero value otherwise.

### GetFilmCriticsOk

`func (o *Rating) GetFilmCriticsOk() (*float32, bool)`

GetFilmCriticsOk returns a tuple with the FilmCritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilmCritics

`func (o *Rating) SetFilmCritics(v float32)`

SetFilmCritics sets FilmCritics field to given value.

### HasFilmCritics

`func (o *Rating) HasFilmCritics() bool`

HasFilmCritics returns a boolean if a field has been set.

### SetFilmCriticsNil

`func (o *Rating) SetFilmCriticsNil(b bool)`

 SetFilmCriticsNil sets the value for FilmCritics to be an explicit nil

### UnsetFilmCritics
`func (o *Rating) UnsetFilmCritics()`

UnsetFilmCritics ensures that no value is present for FilmCritics, not even an explicit nil
### GetRussianFilmCritics

`func (o *Rating) GetRussianFilmCritics() float32`

GetRussianFilmCritics returns the RussianFilmCritics field if non-nil, zero value otherwise.

### GetRussianFilmCriticsOk

`func (o *Rating) GetRussianFilmCriticsOk() (*float32, bool)`

GetRussianFilmCriticsOk returns a tuple with the RussianFilmCritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRussianFilmCritics

`func (o *Rating) SetRussianFilmCritics(v float32)`

SetRussianFilmCritics sets RussianFilmCritics field to given value.

### HasRussianFilmCritics

`func (o *Rating) HasRussianFilmCritics() bool`

HasRussianFilmCritics returns a boolean if a field has been set.

### SetRussianFilmCriticsNil

`func (o *Rating) SetRussianFilmCriticsNil(b bool)`

 SetRussianFilmCriticsNil sets the value for RussianFilmCritics to be an explicit nil

### UnsetRussianFilmCritics
`func (o *Rating) UnsetRussianFilmCritics()`

UnsetRussianFilmCritics ensures that no value is present for RussianFilmCritics, not even an explicit nil
### GetAwait

`func (o *Rating) GetAwait() float32`

GetAwait returns the Await field if non-nil, zero value otherwise.

### GetAwaitOk

`func (o *Rating) GetAwaitOk() (*float32, bool)`

GetAwaitOk returns a tuple with the Await field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwait

`func (o *Rating) SetAwait(v float32)`

SetAwait sets Await field to given value.

### HasAwait

`func (o *Rating) HasAwait() bool`

HasAwait returns a boolean if a field has been set.

### SetAwaitNil

`func (o *Rating) SetAwaitNil(b bool)`

 SetAwaitNil sets the value for Await to be an explicit nil

### UnsetAwait
`func (o *Rating) UnsetAwait()`

UnsetAwait ensures that no value is present for Await, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


