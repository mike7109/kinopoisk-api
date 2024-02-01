# PossibleValueDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Значение по которому нужно делать запрос в базу данных | 
**Slug** | **string** | Вспомогательное значение | 

## Methods

### NewPossibleValueDto

`func NewPossibleValueDto(name string, slug string, ) *PossibleValueDto`

NewPossibleValueDto instantiates a new PossibleValueDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPossibleValueDtoWithDefaults

`func NewPossibleValueDtoWithDefaults() *PossibleValueDto`

NewPossibleValueDtoWithDefaults instantiates a new PossibleValueDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PossibleValueDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PossibleValueDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PossibleValueDto) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PossibleValueDto) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PossibleValueDto) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PossibleValueDto) SetSlug(v string)`

SetSlug sets Slug field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


