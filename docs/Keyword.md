# Keyword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Title** | **string** |  | 
**Movies** | Pointer to [**MovieFromKeyword**](MovieFromKeyword.md) |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewKeyword

`func NewKeyword(id float32, title string, updatedAt time.Time, createdAt time.Time, ) *Keyword`

NewKeyword instantiates a new Keyword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordWithDefaults

`func NewKeywordWithDefaults() *Keyword`

NewKeywordWithDefaults instantiates a new Keyword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Keyword) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Keyword) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Keyword) SetId(v float32)`

SetId sets Id field to given value.


### GetTitle

`func (o *Keyword) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Keyword) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Keyword) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMovies

`func (o *Keyword) GetMovies() MovieFromKeyword`

GetMovies returns the Movies field if non-nil, zero value otherwise.

### GetMoviesOk

`func (o *Keyword) GetMoviesOk() (*MovieFromKeyword, bool)`

GetMoviesOk returns a tuple with the Movies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovies

`func (o *Keyword) SetMovies(v MovieFromKeyword)`

SetMovies sets Movies field to given value.

### HasMovies

`func (o *Keyword) HasMovies() bool`

HasMovies returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Keyword) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Keyword) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Keyword) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Keyword) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Keyword) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Keyword) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


