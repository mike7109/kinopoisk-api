# PersonAward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nomination** | [**Nomination**](Nomination.md) |  | 
**Winning** | **bool** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 
**PersonId** | **float32** |  | 
**Movie** | [**Movie**](Movie.md) |  | 

## Methods

### NewPersonAward

`func NewPersonAward(nomination Nomination, winning bool, updatedAt time.Time, createdAt time.Time, personId float32, movie Movie, ) *PersonAward`

NewPersonAward instantiates a new PersonAward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonAwardWithDefaults

`func NewPersonAwardWithDefaults() *PersonAward`

NewPersonAwardWithDefaults instantiates a new PersonAward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNomination

`func (o *PersonAward) GetNomination() Nomination`

GetNomination returns the Nomination field if non-nil, zero value otherwise.

### GetNominationOk

`func (o *PersonAward) GetNominationOk() (*Nomination, bool)`

GetNominationOk returns a tuple with the Nomination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomination

`func (o *PersonAward) SetNomination(v Nomination)`

SetNomination sets Nomination field to given value.


### GetWinning

`func (o *PersonAward) GetWinning() bool`

GetWinning returns the Winning field if non-nil, zero value otherwise.

### GetWinningOk

`func (o *PersonAward) GetWinningOk() (*bool, bool)`

GetWinningOk returns a tuple with the Winning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinning

`func (o *PersonAward) SetWinning(v bool)`

SetWinning sets Winning field to given value.


### GetUpdatedAt

`func (o *PersonAward) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PersonAward) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PersonAward) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *PersonAward) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PersonAward) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PersonAward) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPersonId

`func (o *PersonAward) GetPersonId() float32`

GetPersonId returns the PersonId field if non-nil, zero value otherwise.

### GetPersonIdOk

`func (o *PersonAward) GetPersonIdOk() (*float32, bool)`

GetPersonIdOk returns a tuple with the PersonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonId

`func (o *PersonAward) SetPersonId(v float32)`

SetPersonId sets PersonId field to given value.


### GetMovie

`func (o *PersonAward) GetMovie() Movie`

GetMovie returns the Movie field if non-nil, zero value otherwise.

### GetMovieOk

`func (o *PersonAward) GetMovieOk() (*Movie, bool)`

GetMovieOk returns a tuple with the Movie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovie

`func (o *PersonAward) SetMovie(v Movie)`

SetMovie sets Movie field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


