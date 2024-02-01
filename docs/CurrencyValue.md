# CurrencyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **NullableFloat32** | Сумма | [optional] 
**Currency** | Pointer to **NullableString** | Валюта | [optional] 

## Methods

### NewCurrencyValue

`func NewCurrencyValue() *CurrencyValue`

NewCurrencyValue instantiates a new CurrencyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyValueWithDefaults

`func NewCurrencyValueWithDefaults() *CurrencyValue`

NewCurrencyValueWithDefaults instantiates a new CurrencyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *CurrencyValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CurrencyValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CurrencyValue) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *CurrencyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *CurrencyValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *CurrencyValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCurrency

`func (o *CurrencyValue) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CurrencyValue) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CurrencyValue) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CurrencyValue) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CurrencyValue) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CurrencyValue) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


