# Review

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**MovieId** | Pointer to **float32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Review** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**AuthorId** | Pointer to **float32** |  | [optional] 
**UserRating** | **float32** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewReview

`func NewReview(userRating float32, updatedAt time.Time, createdAt time.Time, ) *Review`

NewReview instantiates a new Review object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWithDefaults

`func NewReviewWithDefaults() *Review`

NewReviewWithDefaults instantiates a new Review object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Review) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Review) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Review) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Review) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMovieId

`func (o *Review) GetMovieId() float32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *Review) GetMovieIdOk() (*float32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *Review) SetMovieId(v float32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *Review) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.

### GetTitle

`func (o *Review) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Review) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Review) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Review) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *Review) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Review) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Review) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Review) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReview

`func (o *Review) GetReview() string`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *Review) GetReviewOk() (*string, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *Review) SetReview(v string)`

SetReview sets Review field to given value.

### HasReview

`func (o *Review) HasReview() bool`

HasReview returns a boolean if a field has been set.

### GetDate

`func (o *Review) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Review) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Review) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Review) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetAuthor

`func (o *Review) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Review) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Review) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Review) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorId

`func (o *Review) GetAuthorId() float32`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Review) GetAuthorIdOk() (*float32, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Review) SetAuthorId(v float32)`

SetAuthorId sets AuthorId field to given value.

### HasAuthorId

`func (o *Review) HasAuthorId() bool`

HasAuthorId returns a boolean if a field has been set.

### GetUserRating

`func (o *Review) GetUserRating() float32`

GetUserRating returns the UserRating field if non-nil, zero value otherwise.

### GetUserRatingOk

`func (o *Review) GetUserRatingOk() (*float32, bool)`

GetUserRatingOk returns a tuple with the UserRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRating

`func (o *Review) SetUserRating(v float32)`

SetUserRating sets UserRating field to given value.


### GetUpdatedAt

`func (o *Review) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Review) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Review) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Review) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Review) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Review) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


