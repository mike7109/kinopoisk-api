# Nomination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Award** | [**NominationAward**](NominationAward.md) |  | 
**Title** | **string** |  | 

## Methods

### NewNomination

`func NewNomination(award NominationAward, title string, ) *Nomination`

NewNomination instantiates a new Nomination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNominationWithDefaults

`func NewNominationWithDefaults() *Nomination`

NewNominationWithDefaults instantiates a new Nomination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAward

`func (o *Nomination) GetAward() NominationAward`

GetAward returns the Award field if non-nil, zero value otherwise.

### GetAwardOk

`func (o *Nomination) GetAwardOk() (*NominationAward, bool)`

GetAwardOk returns a tuple with the Award field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAward

`func (o *Nomination) SetAward(v NominationAward)`

SetAward sets Award field to given value.


### GetTitle

`func (o *Nomination) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Nomination) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Nomination) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


