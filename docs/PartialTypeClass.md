# PartialTypeClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nomination** | Pointer to [**Nomination**](Nomination.md) |  | [optional] 
**Winning** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**MovieId** | Pointer to **float32** |  | [optional] 

## Methods

### NewPartialTypeClass

`func NewPartialTypeClass() *PartialTypeClass`

NewPartialTypeClass instantiates a new PartialTypeClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialTypeClassWithDefaults

`func NewPartialTypeClassWithDefaults() *PartialTypeClass`

NewPartialTypeClassWithDefaults instantiates a new PartialTypeClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomination

`func (o *PartialTypeClass) GetNomination() Nomination`

GetNomination returns the Nomination field if non-nil, zero value otherwise.

### GetNominationOk

`func (o *PartialTypeClass) GetNominationOk() (*Nomination, bool)`

GetNominationOk returns a tuple with the Nomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomination

`func (o *PartialTypeClass) SetNomination(v Nomination)`

SetNomination sets Nomination field to given value.

### HasNomination

`func (o *PartialTypeClass) HasNomination() bool`

HasNomination returns a boolean if a field has been set.

### GetWinning

`func (o *PartialTypeClass) GetWinning() bool`

GetWinning returns the Winning field if non-nil, zero value otherwise.

### GetWinningOk

`func (o *PartialTypeClass) GetWinningOk() (*bool, bool)`

GetWinningOk returns a tuple with the Winning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinning

`func (o *PartialTypeClass) SetWinning(v bool)`

SetWinning sets Winning field to given value.

### HasWinning

`func (o *PartialTypeClass) HasWinning() bool`

HasWinning returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PartialTypeClass) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PartialTypeClass) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PartialTypeClass) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PartialTypeClass) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PartialTypeClass) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartialTypeClass) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartialTypeClass) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PartialTypeClass) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMovieId

`func (o *PartialTypeClass) GetMovieId() float32`

GetMovieId returns the MovieId field if non-nil, zero value otherwise.

### GetMovieIdOk

`func (o *PartialTypeClass) GetMovieIdOk() (*float32, bool)`

GetMovieIdOk returns a tuple with the MovieId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovieId

`func (o *PartialTypeClass) SetMovieId(v float32)`

SetMovieId sets MovieId field to given value.

### HasMovieId

`func (o *PartialTypeClass) HasMovieId() bool`

HasMovieId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


