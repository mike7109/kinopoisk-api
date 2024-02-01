# Fees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**World** | Pointer to [**CurrencyValue**](CurrencyValue.md) |  | [optional] 
**Russia** | Pointer to [**CurrencyValue**](CurrencyValue.md) |  | [optional] 
**Usa** | Pointer to [**CurrencyValue**](CurrencyValue.md) |  | [optional] 

## Methods

### NewFees

`func NewFees() *Fees`

NewFees instantiates a new Fees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeesWithDefaults

`func NewFeesWithDefaults() *Fees`

NewFeesWithDefaults instantiates a new Fees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorld

`func (o *Fees) GetWorld() CurrencyValue`

GetWorld returns the World field if non-nil, zero value otherwise.

### GetWorldOk

`func (o *Fees) GetWorldOk() (*CurrencyValue, bool)`

GetWorldOk returns a tuple with the World field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorld

`func (o *Fees) SetWorld(v CurrencyValue)`

SetWorld sets World field to given value.

### HasWorld

`func (o *Fees) HasWorld() bool`

HasWorld returns a boolean if a field has been set.

### GetRussia

`func (o *Fees) GetRussia() CurrencyValue`

GetRussia returns the Russia field if non-nil, zero value otherwise.

### GetRussiaOk

`func (o *Fees) GetRussiaOk() (*CurrencyValue, bool)`

GetRussiaOk returns a tuple with the Russia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRussia

`func (o *Fees) SetRussia(v CurrencyValue)`

SetRussia sets Russia field to given value.

### HasRussia

`func (o *Fees) HasRussia() bool`

HasRussia returns a boolean if a field has been set.

### GetUsa

`func (o *Fees) GetUsa() CurrencyValue`

GetUsa returns the Usa field if non-nil, zero value otherwise.

### GetUsaOk

`func (o *Fees) GetUsaOk() (*CurrencyValue, bool)`

GetUsaOk returns a tuple with the Usa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsa

`func (o *Fees) SetUsa(v CurrencyValue)`

SetUsa sets Usa field to given value.

### HasUsa

`func (o *Fees) HasUsa() bool`

HasUsa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


