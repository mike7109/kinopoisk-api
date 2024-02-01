# MovieInPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**AlternativeName** | Pointer to **NullableString** |  | [optional] 
**Rating** | Pointer to **NullableFloat32** |  | [optional] 
**General** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EnProfession** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMovieInPerson

`func NewMovieInPerson(id float32, ) *MovieInPerson`

NewMovieInPerson instantiates a new MovieInPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMovieInPersonWithDefaults

`func NewMovieInPersonWithDefaults() *MovieInPerson`

NewMovieInPersonWithDefaults instantiates a new MovieInPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MovieInPerson) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MovieInPerson) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MovieInPerson) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *MovieInPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MovieInPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MovieInPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MovieInPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MovieInPerson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MovieInPerson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAlternativeName

`func (o *MovieInPerson) GetAlternativeName() string`

GetAlternativeName returns the AlternativeName field if non-nil, zero value otherwise.

### GetAlternativeNameOk

`func (o *MovieInPerson) GetAlternativeNameOk() (*string, bool)`

GetAlternativeNameOk returns a tuple with the AlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternativeName

`func (o *MovieInPerson) SetAlternativeName(v string)`

SetAlternativeName sets AlternativeName field to given value.

### HasAlternativeName

`func (o *MovieInPerson) HasAlternativeName() bool`

HasAlternativeName returns a boolean if a field has been set.

### SetAlternativeNameNil

`func (o *MovieInPerson) SetAlternativeNameNil(b bool)`

 SetAlternativeNameNil sets the value for AlternativeName to be an explicit nil

### UnsetAlternativeName
`func (o *MovieInPerson) UnsetAlternativeName()`

UnsetAlternativeName ensures that no value is present for AlternativeName, not even an explicit nil
### GetRating

`func (o *MovieInPerson) GetRating() float32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *MovieInPerson) GetRatingOk() (*float32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *MovieInPerson) SetRating(v float32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *MovieInPerson) HasRating() bool`

HasRating returns a boolean if a field has been set.

### SetRatingNil

`func (o *MovieInPerson) SetRatingNil(b bool)`

 SetRatingNil sets the value for Rating to be an explicit nil

### UnsetRating
`func (o *MovieInPerson) UnsetRating()`

UnsetRating ensures that no value is present for Rating, not even an explicit nil
### GetGeneral

`func (o *MovieInPerson) GetGeneral() bool`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *MovieInPerson) GetGeneralOk() (*bool, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *MovieInPerson) SetGeneral(v bool)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *MovieInPerson) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### SetGeneralNil

`func (o *MovieInPerson) SetGeneralNil(b bool)`

 SetGeneralNil sets the value for General to be an explicit nil

### UnsetGeneral
`func (o *MovieInPerson) UnsetGeneral()`

UnsetGeneral ensures that no value is present for General, not even an explicit nil
### GetDescription

`func (o *MovieInPerson) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MovieInPerson) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MovieInPerson) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MovieInPerson) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MovieInPerson) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MovieInPerson) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnProfession

`func (o *MovieInPerson) GetEnProfession() string`

GetEnProfession returns the EnProfession field if non-nil, zero value otherwise.

### GetEnProfessionOk

`func (o *MovieInPerson) GetEnProfessionOk() (*string, bool)`

GetEnProfessionOk returns a tuple with the EnProfession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnProfession

`func (o *MovieInPerson) SetEnProfession(v string)`

SetEnProfession sets EnProfession field to given value.

### HasEnProfession

`func (o *MovieInPerson) HasEnProfession() bool`

HasEnProfession returns a boolean if a field has been set.

### SetEnProfessionNil

`func (o *MovieInPerson) SetEnProfessionNil(b bool)`

 SetEnProfessionNil sets the value for EnProfession to be an explicit nil

### UnsetEnProfession
`func (o *MovieInPerson) UnsetEnProfession()`

UnsetEnProfession ensures that no value is present for EnProfession, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


