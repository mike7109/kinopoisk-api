# WatchabilityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Logo** | [**Logo**](Logo.md) |  | 
**Url** | **string** |  | 

## Methods

### NewWatchabilityItem

`func NewWatchabilityItem(logo Logo, url string, ) *WatchabilityItem`

NewWatchabilityItem instantiates a new WatchabilityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWatchabilityItemWithDefaults

`func NewWatchabilityItemWithDefaults() *WatchabilityItem`

NewWatchabilityItemWithDefaults instantiates a new WatchabilityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WatchabilityItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WatchabilityItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WatchabilityItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WatchabilityItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *WatchabilityItem) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *WatchabilityItem) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLogo

`func (o *WatchabilityItem) GetLogo() Logo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *WatchabilityItem) GetLogoOk() (*Logo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *WatchabilityItem) SetLogo(v Logo)`

SetLogo sets Logo field to given value.


### GetUrl

`func (o *WatchabilityItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WatchabilityItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WatchabilityItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


