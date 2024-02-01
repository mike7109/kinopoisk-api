# VideoTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trailers** | Pointer to [**[]Video**](Video.md) |  | [optional] 
**Teasers** | [**[]Video**](Video.md) |  | 

## Methods

### NewVideoTypes

`func NewVideoTypes(teasers []Video, ) *VideoTypes`

NewVideoTypes instantiates a new VideoTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoTypesWithDefaults

`func NewVideoTypesWithDefaults() *VideoTypes`

NewVideoTypesWithDefaults instantiates a new VideoTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrailers

`func (o *VideoTypes) GetTrailers() []Video`

GetTrailers returns the Trailers field if non-nil, zero value otherwise.

### GetTrailersOk

`func (o *VideoTypes) GetTrailersOk() (*[]Video, bool)`

GetTrailersOk returns a tuple with the Trailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailers

`func (o *VideoTypes) SetTrailers(v []Video)`

SetTrailers sets Trailers field to given value.

### HasTrailers

`func (o *VideoTypes) HasTrailers() bool`

HasTrailers returns a boolean if a field has been set.

### GetTeasers

`func (o *VideoTypes) GetTeasers() []Video`

GetTeasers returns the Teasers field if non-nil, zero value otherwise.

### GetTeasersOk

`func (o *VideoTypes) GetTeasersOk() (*[]Video, bool)`

GetTeasersOk returns a tuple with the Teasers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeasers

`func (o *VideoTypes) SetTeasers(v []Video)`

SetTeasers sets Teasers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


