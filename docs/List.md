# List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** |  | 
**Slug** | **string** |  | 
**MoviesCount** | **float32** |  | 
**Cover** | [**ShortImage**](ShortImage.md) |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewList

`func NewList(category string, slug string, moviesCount float32, cover ShortImage, name string, updatedAt time.Time, createdAt time.Time, ) *List`

NewList instantiates a new List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWithDefaults

`func NewListWithDefaults() *List`

NewListWithDefaults instantiates a new List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *List) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *List) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *List) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetSlug

`func (o *List) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *List) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *List) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetMoviesCount

`func (o *List) GetMoviesCount() float32`

GetMoviesCount returns the MoviesCount field if non-nil, zero value otherwise.

### GetMoviesCountOk

`func (o *List) GetMoviesCountOk() (*float32, bool)`

GetMoviesCountOk returns a tuple with the MoviesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoviesCount

`func (o *List) SetMoviesCount(v float32)`

SetMoviesCount sets MoviesCount field to given value.


### GetCover

`func (o *List) GetCover() ShortImage`

GetCover returns the Cover field if non-nil, zero value otherwise.

### GetCoverOk

`func (o *List) GetCoverOk() (*ShortImage, bool)`

GetCoverOk returns a tuple with the Cover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCover

`func (o *List) SetCover(v ShortImage)`

SetCover sets Cover field to given value.


### GetName

`func (o *List) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *List) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *List) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *List) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *List) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *List) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *List) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *List) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *List) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


