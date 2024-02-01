# Person

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**EnName** | Pointer to **NullableString** |  | [optional] 
**Photo** | Pointer to **NullableString** |  | [optional] 
**Sex** | Pointer to **NullableString** |  | [optional] 
**Growth** | Pointer to **NullableFloat32** |  | [optional] 
**Birthday** | Pointer to **NullableString** |  | [optional] 
**Death** | Pointer to **NullableString** |  | [optional] 
**Age** | Pointer to **NullableFloat32** |  | [optional] 
**BirthPlace** | Pointer to [**[]BirthPlace**](BirthPlace.md) |  | [optional] 
**DeathPlace** | Pointer to [**[]DeathPlace**](DeathPlace.md) |  | [optional] 
**Spouses** | Pointer to [**Spouses**](Spouses.md) |  | [optional] 
**CountAwards** | Pointer to **float32** |  | [optional] 
**Profession** | Pointer to [**[]Profession**](Profession.md) |  | [optional] 
**Facts** | Pointer to [**[]FactInPerson**](FactInPerson.md) |  | [optional] 
**Movies** | Pointer to [**[]MovieInPerson**](MovieInPerson.md) |  | [optional] 
**UpdatedAt** | **time.Time** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewPerson

`func NewPerson(id float32, updatedAt time.Time, createdAt time.Time, ) *Person`

NewPerson instantiates a new Person object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonWithDefaults

`func NewPersonWithDefaults() *Person`

NewPersonWithDefaults instantiates a new Person object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Person) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Person) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Person) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Person) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Person) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Person) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Person) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Person) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Person) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnName

`func (o *Person) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *Person) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *Person) SetEnName(v string)`

SetEnName sets EnName field to given value.

### HasEnName

`func (o *Person) HasEnName() bool`

HasEnName returns a boolean if a field has been set.

### SetEnNameNil

`func (o *Person) SetEnNameNil(b bool)`

 SetEnNameNil sets the value for EnName to be an explicit nil

### UnsetEnName
`func (o *Person) UnsetEnName()`

UnsetEnName ensures that no value is present for EnName, not even an explicit nil
### GetPhoto

`func (o *Person) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Person) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Person) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Person) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *Person) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *Person) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetSex

`func (o *Person) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *Person) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *Person) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *Person) HasSex() bool`

HasSex returns a boolean if a field has been set.

### SetSexNil

`func (o *Person) SetSexNil(b bool)`

 SetSexNil sets the value for Sex to be an explicit nil

### UnsetSex
`func (o *Person) UnsetSex()`

UnsetSex ensures that no value is present for Sex, not even an explicit nil
### GetGrowth

`func (o *Person) GetGrowth() float32`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *Person) GetGrowthOk() (*float32, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *Person) SetGrowth(v float32)`

SetGrowth sets Growth field to given value.

### HasGrowth

`func (o *Person) HasGrowth() bool`

HasGrowth returns a boolean if a field has been set.

### SetGrowthNil

`func (o *Person) SetGrowthNil(b bool)`

 SetGrowthNil sets the value for Growth to be an explicit nil

### UnsetGrowth
`func (o *Person) UnsetGrowth()`

UnsetGrowth ensures that no value is present for Growth, not even an explicit nil
### GetBirthday

`func (o *Person) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *Person) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *Person) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.

### HasBirthday

`func (o *Person) HasBirthday() bool`

HasBirthday returns a boolean if a field has been set.

### SetBirthdayNil

`func (o *Person) SetBirthdayNil(b bool)`

 SetBirthdayNil sets the value for Birthday to be an explicit nil

### UnsetBirthday
`func (o *Person) UnsetBirthday()`

UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
### GetDeath

`func (o *Person) GetDeath() string`

GetDeath returns the Death field if non-nil, zero value otherwise.

### GetDeathOk

`func (o *Person) GetDeathOk() (*string, bool)`

GetDeathOk returns a tuple with the Death field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeath

`func (o *Person) SetDeath(v string)`

SetDeath sets Death field to given value.

### HasDeath

`func (o *Person) HasDeath() bool`

HasDeath returns a boolean if a field has been set.

### SetDeathNil

`func (o *Person) SetDeathNil(b bool)`

 SetDeathNil sets the value for Death to be an explicit nil

### UnsetDeath
`func (o *Person) UnsetDeath()`

UnsetDeath ensures that no value is present for Death, not even an explicit nil
### GetAge

`func (o *Person) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Person) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Person) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *Person) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *Person) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *Person) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetBirthPlace

`func (o *Person) GetBirthPlace() []BirthPlace`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *Person) GetBirthPlaceOk() (*[]BirthPlace, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *Person) SetBirthPlace(v []BirthPlace)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *Person) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### GetDeathPlace

`func (o *Person) GetDeathPlace() []DeathPlace`

GetDeathPlace returns the DeathPlace field if non-nil, zero value otherwise.

### GetDeathPlaceOk

`func (o *Person) GetDeathPlaceOk() (*[]DeathPlace, bool)`

GetDeathPlaceOk returns a tuple with the DeathPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathPlace

`func (o *Person) SetDeathPlace(v []DeathPlace)`

SetDeathPlace sets DeathPlace field to given value.

### HasDeathPlace

`func (o *Person) HasDeathPlace() bool`

HasDeathPlace returns a boolean if a field has been set.

### GetSpouses

`func (o *Person) GetSpouses() Spouses`

GetSpouses returns the Spouses field if non-nil, zero value otherwise.

### GetSpousesOk

`func (o *Person) GetSpousesOk() (*Spouses, bool)`

GetSpousesOk returns a tuple with the Spouses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouses

`func (o *Person) SetSpouses(v Spouses)`

SetSpouses sets Spouses field to given value.

### HasSpouses

`func (o *Person) HasSpouses() bool`

HasSpouses returns a boolean if a field has been set.

### GetCountAwards

`func (o *Person) GetCountAwards() float32`

GetCountAwards returns the CountAwards field if non-nil, zero value otherwise.

### GetCountAwardsOk

`func (o *Person) GetCountAwardsOk() (*float32, bool)`

GetCountAwardsOk returns a tuple with the CountAwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAwards

`func (o *Person) SetCountAwards(v float32)`

SetCountAwards sets CountAwards field to given value.

### HasCountAwards

`func (o *Person) HasCountAwards() bool`

HasCountAwards returns a boolean if a field has been set.

### GetProfession

`func (o *Person) GetProfession() []Profession`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *Person) GetProfessionOk() (*[]Profession, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *Person) SetProfession(v []Profession)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *Person) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetFacts

`func (o *Person) GetFacts() []FactInPerson`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *Person) GetFactsOk() (*[]FactInPerson, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *Person) SetFacts(v []FactInPerson)`

SetFacts sets Facts field to given value.

### HasFacts

`func (o *Person) HasFacts() bool`

HasFacts returns a boolean if a field has been set.

### GetMovies

`func (o *Person) GetMovies() []MovieInPerson`

GetMovies returns the Movies field if non-nil, zero value otherwise.

### GetMoviesOk

`func (o *Person) GetMoviesOk() (*[]MovieInPerson, bool)`

GetMoviesOk returns a tuple with the Movies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovies

`func (o *Person) SetMovies(v []MovieInPerson)`

SetMovies sets Movies field to given value.

### HasMovies

`func (o *Person) HasMovies() bool`

HasMovies returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Person) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Person) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Person) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatedAt

`func (o *Person) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Person) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Person) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


