# Votes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kp** | Pointer to **NullableString** |  | [optional] 
**Imdb** | Pointer to **NullableFloat32** |  | [optional] 
**Tmdb** | Pointer to **NullableFloat32** |  | [optional] 
**FilmCritics** | Pointer to **NullableFloat32** | Количество голосов кинокритиков | [optional] 
**RussianFilmCritics** | Pointer to **NullableFloat32** | Количество голосов кинокритиков из РФ | [optional] 
**Await** | Pointer to **NullableFloat32** | Количество ожидающих выхода | [optional] 

## Methods

### NewVotes

`func NewVotes() *Votes`

NewVotes instantiates a new Votes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVotesWithDefaults

`func NewVotesWithDefaults() *Votes`

NewVotesWithDefaults instantiates a new Votes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKp

`func (o *Votes) GetKp() string`

GetKp returns the Kp field if non-nil, zero value otherwise.

### GetKpOk

`func (o *Votes) GetKpOk() (*string, bool)`

GetKpOk returns a tuple with the Kp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKp

`func (o *Votes) SetKp(v string)`

SetKp sets Kp field to given value.

### HasKp

`func (o *Votes) HasKp() bool`

HasKp returns a boolean if a field has been set.

### SetKpNil

`func (o *Votes) SetKpNil(b bool)`

 SetKpNil sets the value for Kp to be an explicit nil

### UnsetKp
`func (o *Votes) UnsetKp()`

UnsetKp ensures that no value is present for Kp, not even an explicit nil
### GetImdb

`func (o *Votes) GetImdb() float32`

GetImdb returns the Imdb field if non-nil, zero value otherwise.

### GetImdbOk

`func (o *Votes) GetImdbOk() (*float32, bool)`

GetImdbOk returns a tuple with the Imdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdb

`func (o *Votes) SetImdb(v float32)`

SetImdb sets Imdb field to given value.

### HasImdb

`func (o *Votes) HasImdb() bool`

HasImdb returns a boolean if a field has been set.

### SetImdbNil

`func (o *Votes) SetImdbNil(b bool)`

 SetImdbNil sets the value for Imdb to be an explicit nil

### UnsetImdb
`func (o *Votes) UnsetImdb()`

UnsetImdb ensures that no value is present for Imdb, not even an explicit nil
### GetTmdb

`func (o *Votes) GetTmdb() float32`

GetTmdb returns the Tmdb field if non-nil, zero value otherwise.

### GetTmdbOk

`func (o *Votes) GetTmdbOk() (*float32, bool)`

GetTmdbOk returns a tuple with the Tmdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdb

`func (o *Votes) SetTmdb(v float32)`

SetTmdb sets Tmdb field to given value.

### HasTmdb

`func (o *Votes) HasTmdb() bool`

HasTmdb returns a boolean if a field has been set.

### SetTmdbNil

`func (o *Votes) SetTmdbNil(b bool)`

 SetTmdbNil sets the value for Tmdb to be an explicit nil

### UnsetTmdb
`func (o *Votes) UnsetTmdb()`

UnsetTmdb ensures that no value is present for Tmdb, not even an explicit nil
### GetFilmCritics

`func (o *Votes) GetFilmCritics() float32`

GetFilmCritics returns the FilmCritics field if non-nil, zero value otherwise.

### GetFilmCriticsOk

`func (o *Votes) GetFilmCriticsOk() (*float32, bool)`

GetFilmCriticsOk returns a tuple with the FilmCritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilmCritics

`func (o *Votes) SetFilmCritics(v float32)`

SetFilmCritics sets FilmCritics field to given value.

### HasFilmCritics

`func (o *Votes) HasFilmCritics() bool`

HasFilmCritics returns a boolean if a field has been set.

### SetFilmCriticsNil

`func (o *Votes) SetFilmCriticsNil(b bool)`

 SetFilmCriticsNil sets the value for FilmCritics to be an explicit nil

### UnsetFilmCritics
`func (o *Votes) UnsetFilmCritics()`

UnsetFilmCritics ensures that no value is present for FilmCritics, not even an explicit nil
### GetRussianFilmCritics

`func (o *Votes) GetRussianFilmCritics() float32`

GetRussianFilmCritics returns the RussianFilmCritics field if non-nil, zero value otherwise.

### GetRussianFilmCriticsOk

`func (o *Votes) GetRussianFilmCriticsOk() (*float32, bool)`

GetRussianFilmCriticsOk returns a tuple with the RussianFilmCritics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRussianFilmCritics

`func (o *Votes) SetRussianFilmCritics(v float32)`

SetRussianFilmCritics sets RussianFilmCritics field to given value.

### HasRussianFilmCritics

`func (o *Votes) HasRussianFilmCritics() bool`

HasRussianFilmCritics returns a boolean if a field has been set.

### SetRussianFilmCriticsNil

`func (o *Votes) SetRussianFilmCriticsNil(b bool)`

 SetRussianFilmCriticsNil sets the value for RussianFilmCritics to be an explicit nil

### UnsetRussianFilmCritics
`func (o *Votes) UnsetRussianFilmCritics()`

UnsetRussianFilmCritics ensures that no value is present for RussianFilmCritics, not even an explicit nil
### GetAwait

`func (o *Votes) GetAwait() float32`

GetAwait returns the Await field if non-nil, zero value otherwise.

### GetAwaitOk

`func (o *Votes) GetAwaitOk() (*float32, bool)`

GetAwaitOk returns a tuple with the Await field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwait

`func (o *Votes) SetAwait(v float32)`

SetAwait sets Await field to given value.

### HasAwait

`func (o *Votes) HasAwait() bool`

HasAwait returns a boolean if a field has been set.

### SetAwaitNil

`func (o *Votes) SetAwaitNil(b bool)`

 SetAwaitNil sets the value for Await to be an explicit nil

### UnsetAwait
`func (o *Votes) UnsetAwait()`

UnsetAwait ensures that no value is present for Await, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


