# Studio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SubType** | **string** |  | 
**Title** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Movies** | Pointer to [**MovieFromStudio**](MovieFromStudio.md) |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewStudio

`func NewStudio(id string, subType string, title string, updatedAt time.Time, createdAt time.Time, ) *Studio`

NewStudio instantiates a new Studio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudioWithDefaults

`func NewStudioWithDefaults() *Studio`

NewStudioWithDefaults instantiates a new Studio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Studio) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Studio) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Studio) SetId(v string)`

SetId sets Id field to given value.


### GetSubType

`func (o *Studio) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Studio) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Studio) SetSubType(v string)`

SetSubType sets SubType field to given value.


### GetTitle

`func (o *Studio) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Studio) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Studio) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *Studio) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Studio) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Studio) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Studio) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMovies

`func (o *Studio) GetMovies() MovieFromStudio`

GetMovies returns the Movies field if non-nil, zero value otherwise.

### GetMoviesOk

`func (o *Studio) GetMoviesOk() (*MovieFromStudio, bool)`

GetMoviesOk returns a tuple with the Movies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovies

`func (o *Studio) SetMovies(v MovieFromStudio)`

SetMovies sets Movies field to given value.

### HasMovies

`func (o *Studio) HasMovies() bool`

HasMovies returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Studio) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Studio) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Studio) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Studio) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Studio) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Studio) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


