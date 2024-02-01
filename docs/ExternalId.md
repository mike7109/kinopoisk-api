# ExternalId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KpHD** | Pointer to **NullableString** | ID из kinopoisk HD | [optional] 
**Imdb** | Pointer to **NullableString** |  | [optional] 
**Tmdb** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewExternalId

`func NewExternalId() *ExternalId`

NewExternalId instantiates a new ExternalId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIdWithDefaults

`func NewExternalIdWithDefaults() *ExternalId`

NewExternalIdWithDefaults instantiates a new ExternalId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKpHD

`func (o *ExternalId) GetKpHD() string`

GetKpHD returns the KpHD field if non-nil, zero value otherwise.

### GetKpHDOk

`func (o *ExternalId) GetKpHDOk() (*string, bool)`

GetKpHDOk returns a tuple with the KpHD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKpHD

`func (o *ExternalId) SetKpHD(v string)`

SetKpHD sets KpHD field to given value.

### HasKpHD

`func (o *ExternalId) HasKpHD() bool`

HasKpHD returns a boolean if a field has been set.

### SetKpHDNil

`func (o *ExternalId) SetKpHDNil(b bool)`

 SetKpHDNil sets the value for KpHD to be an explicit nil

### UnsetKpHD
`func (o *ExternalId) UnsetKpHD()`

UnsetKpHD ensures that no value is present for KpHD, not even an explicit nil
### GetImdb

`func (o *ExternalId) GetImdb() string`

GetImdb returns the Imdb field if non-nil, zero value otherwise.

### GetImdbOk

`func (o *ExternalId) GetImdbOk() (*string, bool)`

GetImdbOk returns a tuple with the Imdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImdb

`func (o *ExternalId) SetImdb(v string)`

SetImdb sets Imdb field to given value.

### HasImdb

`func (o *ExternalId) HasImdb() bool`

HasImdb returns a boolean if a field has been set.

### SetImdbNil

`func (o *ExternalId) SetImdbNil(b bool)`

 SetImdbNil sets the value for Imdb to be an explicit nil

### UnsetImdb
`func (o *ExternalId) UnsetImdb()`

UnsetImdb ensures that no value is present for Imdb, not even an explicit nil
### GetTmdb

`func (o *ExternalId) GetTmdb() float32`

GetTmdb returns the Tmdb field if non-nil, zero value otherwise.

### GetTmdbOk

`func (o *ExternalId) GetTmdbOk() (*float32, bool)`

GetTmdbOk returns a tuple with the Tmdb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTmdb

`func (o *ExternalId) SetTmdb(v float32)`

SetTmdb sets Tmdb field to given value.

### HasTmdb

`func (o *ExternalId) HasTmdb() bool`

HasTmdb returns a boolean if a field has been set.

### SetTmdbNil

`func (o *ExternalId) SetTmdbNil(b bool)`

 SetTmdbNil sets the value for Tmdb to be an explicit nil

### UnsetTmdb
`func (o *ExternalId) UnsetTmdb()`

UnsetTmdb ensures that no value is present for Tmdb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


