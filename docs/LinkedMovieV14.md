# LinkedMovieV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableFloat32** |  | [optional] 
**Rating** | [**Rating**](Rating.md) |  | 
**Year** | **float32** |  | 
**Name** | **string** |  | 
**EnName** | **string** |  | 
**AlternativeName** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Poster** | [**ShortImage**](ShortImage.md) |  | 

## Methods

### NewLinkedMovieV14

`func NewLinkedMovieV14(rating Rating, year float32, name string, enName string, alternativeName string, poster ShortImage, ) *LinkedMovieV14`

NewLinkedMovieV14 instantiates a new LinkedMovieV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedMovieV14WithDefaults

`func NewLinkedMovieV14WithDefaults() *LinkedMovieV14`

NewLinkedMovieV14WithDefaults instantiates a new LinkedMovieV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedMovieV14) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedMovieV14) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedMovieV14) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *LinkedMovieV14) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *LinkedMovieV14) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *LinkedMovieV14) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRating

`func (o *LinkedMovieV14) GetRating() Rating`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *LinkedMovieV14) GetRatingOk() (*Rating, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *LinkedMovieV14) SetRating(v Rating)`

SetRating sets Rating field to given value.


### GetYear

`func (o *LinkedMovieV14) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *LinkedMovieV14) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *LinkedMovieV14) SetYear(v float32)`

SetYear sets Year field to given value.


### GetName

`func (o *LinkedMovieV14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LinkedMovieV14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LinkedMovieV14) SetName(v string)`

SetName sets Name field to given value.


### GetEnName

`func (o *LinkedMovieV14) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *LinkedMovieV14) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *LinkedMovieV14) SetEnName(v string)`

SetEnName sets EnName field to given value.


### GetAlternativeName

`func (o *LinkedMovieV14) GetAlternativeName() string`

GetAlternativeName returns the AlternativeName field if non-nil, zero value otherwise.

### GetAlternativeNameOk

`func (o *LinkedMovieV14) GetAlternativeNameOk() (*string, bool)`

GetAlternativeNameOk returns a tuple with the AlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeName

`func (o *LinkedMovieV14) SetAlternativeName(v string)`

SetAlternativeName sets AlternativeName field to given value.


### GetType

`func (o *LinkedMovieV14) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkedMovieV14) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkedMovieV14) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LinkedMovieV14) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPoster

`func (o *LinkedMovieV14) GetPoster() ShortImage`

GetPoster returns the Poster field if non-nil, zero value otherwise.

### GetPosterOk

`func (o *LinkedMovieV14) GetPosterOk() (*ShortImage, bool)`

GetPosterOk returns a tuple with the Poster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoster

`func (o *LinkedMovieV14) SetPoster(v ShortImage)`

SetPoster sets Poster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


