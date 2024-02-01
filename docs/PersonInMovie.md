# PersonInMovie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableFloat32** | Id персоны с кинопоиска | [optional] 
**Photo** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**EnName** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**Profession** | **string** |  | 
**EnProfession** | **string** |  | 

## Methods

### NewPersonInMovie

`func NewPersonInMovie(description string, profession string, enProfession string, ) *PersonInMovie`

NewPersonInMovie instantiates a new PersonInMovie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonInMovieWithDefaults

`func NewPersonInMovieWithDefaults() *PersonInMovie`

NewPersonInMovieWithDefaults instantiates a new PersonInMovie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PersonInMovie) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PersonInMovie) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PersonInMovie) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *PersonInMovie) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PersonInMovie) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PersonInMovie) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPhoto

`func (o *PersonInMovie) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *PersonInMovie) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *PersonInMovie) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *PersonInMovie) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *PersonInMovie) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *PersonInMovie) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetName

`func (o *PersonInMovie) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PersonInMovie) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PersonInMovie) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PersonInMovie) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PersonInMovie) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PersonInMovie) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnName

`func (o *PersonInMovie) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *PersonInMovie) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *PersonInMovie) SetEnName(v string)`

SetEnName sets EnName field to given value.

### HasEnName

`func (o *PersonInMovie) HasEnName() bool`

HasEnName returns a boolean if a field has been set.

### SetEnNameNil

`func (o *PersonInMovie) SetEnNameNil(b bool)`

 SetEnNameNil sets the value for EnName to be an explicit nil

### UnsetEnName
`func (o *PersonInMovie) UnsetEnName()`

UnsetEnName ensures that no value is present for EnName, not even an explicit nil
### GetDescription

`func (o *PersonInMovie) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PersonInMovie) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PersonInMovie) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetProfession

`func (o *PersonInMovie) GetProfession() string`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *PersonInMovie) GetProfessionOk() (*string, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *PersonInMovie) SetProfession(v string)`

SetProfession sets Profession field to given value.


### GetEnProfession

`func (o *PersonInMovie) GetEnProfession() string`

GetEnProfession returns the EnProfession field if non-nil, zero value otherwise.

### GetEnProfessionOk

`func (o *PersonInMovie) GetEnProfessionOk() (*string, bool)`

GetEnProfessionOk returns a tuple with the EnProfession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnProfession

`func (o *PersonInMovie) SetEnProfession(v string)`

SetEnProfession sets EnProfession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


