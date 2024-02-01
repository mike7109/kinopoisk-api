# MeiliPersonEntityV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1** | **map[string]interface{}** |  | 
**Id** | **float32** |  | 
**Name** | **string** |  | 
**EnName** | **string** |  | 
**Photo** | **string** |  | 
**Sex** | **string** |  | 
**Birthday** | **string** |  | 
**Death** | **string** |  | 
**Age** | **float32** |  | 
**BirthPlace** | Pointer to [**[]BirthPlace**](BirthPlace.md) |  | [optional] 
**DeathPlace** | Pointer to [**[]DeathPlace**](DeathPlace.md) |  | [optional] 
**Profession** | Pointer to [**[]Profession**](Profession.md) |  | [optional] 
**Growth** | **float32** |  | 

## Methods

### NewMeiliPersonEntityV14

`func NewMeiliPersonEntityV14(var1 map[string]interface{}, id float32, name string, enName string, photo string, sex string, birthday string, death string, age float32, growth float32, ) *MeiliPersonEntityV14`

NewMeiliPersonEntityV14 instantiates a new MeiliPersonEntityV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeiliPersonEntityV14WithDefaults

`func NewMeiliPersonEntityV14WithDefaults() *MeiliPersonEntityV14`

NewMeiliPersonEntityV14WithDefaults instantiates a new MeiliPersonEntityV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1

`func (o *MeiliPersonEntityV14) GetVar1() map[string]interface{}`

GetVar1 returns the Var1 field if non-nil, zero value otherwise.

### GetVar1Ok

`func (o *MeiliPersonEntityV14) GetVar1Ok() (*map[string]interface{}, bool)`

GetVar1Ok returns a tuple with the Var1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1

`func (o *MeiliPersonEntityV14) SetVar1(v map[string]interface{})`

SetVar1 sets Var1 field to given value.


### GetId

`func (o *MeiliPersonEntityV14) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MeiliPersonEntityV14) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MeiliPersonEntityV14) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *MeiliPersonEntityV14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MeiliPersonEntityV14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MeiliPersonEntityV14) SetName(v string)`

SetName sets Name field to given value.


### GetEnName

`func (o *MeiliPersonEntityV14) GetEnName() string`

GetEnName returns the EnName field if non-nil, zero value otherwise.

### GetEnNameOk

`func (o *MeiliPersonEntityV14) GetEnNameOk() (*string, bool)`

GetEnNameOk returns a tuple with the EnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnName

`func (o *MeiliPersonEntityV14) SetEnName(v string)`

SetEnName sets EnName field to given value.


### GetPhoto

`func (o *MeiliPersonEntityV14) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *MeiliPersonEntityV14) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *MeiliPersonEntityV14) SetPhoto(v string)`

SetPhoto sets Photo field to given value.


### GetSex

`func (o *MeiliPersonEntityV14) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *MeiliPersonEntityV14) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *MeiliPersonEntityV14) SetSex(v string)`

SetSex sets Sex field to given value.


### GetBirthday

`func (o *MeiliPersonEntityV14) GetBirthday() string`

GetBirthday returns the Birthday field if non-nil, zero value otherwise.

### GetBirthdayOk

`func (o *MeiliPersonEntityV14) GetBirthdayOk() (*string, bool)`

GetBirthdayOk returns a tuple with the Birthday field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthday

`func (o *MeiliPersonEntityV14) SetBirthday(v string)`

SetBirthday sets Birthday field to given value.


### GetDeath

`func (o *MeiliPersonEntityV14) GetDeath() string`

GetDeath returns the Death field if non-nil, zero value otherwise.

### GetDeathOk

`func (o *MeiliPersonEntityV14) GetDeathOk() (*string, bool)`

GetDeathOk returns a tuple with the Death field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeath

`func (o *MeiliPersonEntityV14) SetDeath(v string)`

SetDeath sets Death field to given value.


### GetAge

`func (o *MeiliPersonEntityV14) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *MeiliPersonEntityV14) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *MeiliPersonEntityV14) SetAge(v float32)`

SetAge sets Age field to given value.


### GetBirthPlace

`func (o *MeiliPersonEntityV14) GetBirthPlace() []BirthPlace`

GetBirthPlace returns the BirthPlace field if non-nil, zero value otherwise.

### GetBirthPlaceOk

`func (o *MeiliPersonEntityV14) GetBirthPlaceOk() (*[]BirthPlace, bool)`

GetBirthPlaceOk returns a tuple with the BirthPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthPlace

`func (o *MeiliPersonEntityV14) SetBirthPlace(v []BirthPlace)`

SetBirthPlace sets BirthPlace field to given value.

### HasBirthPlace

`func (o *MeiliPersonEntityV14) HasBirthPlace() bool`

HasBirthPlace returns a boolean if a field has been set.

### GetDeathPlace

`func (o *MeiliPersonEntityV14) GetDeathPlace() []DeathPlace`

GetDeathPlace returns the DeathPlace field if non-nil, zero value otherwise.

### GetDeathPlaceOk

`func (o *MeiliPersonEntityV14) GetDeathPlaceOk() (*[]DeathPlace, bool)`

GetDeathPlaceOk returns a tuple with the DeathPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeathPlace

`func (o *MeiliPersonEntityV14) SetDeathPlace(v []DeathPlace)`

SetDeathPlace sets DeathPlace field to given value.

### HasDeathPlace

`func (o *MeiliPersonEntityV14) HasDeathPlace() bool`

HasDeathPlace returns a boolean if a field has been set.

### GetProfession

`func (o *MeiliPersonEntityV14) GetProfession() []Profession`

GetProfession returns the Profession field if non-nil, zero value otherwise.

### GetProfessionOk

`func (o *MeiliPersonEntityV14) GetProfessionOk() (*[]Profession, bool)`

GetProfessionOk returns a tuple with the Profession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfession

`func (o *MeiliPersonEntityV14) SetProfession(v []Profession)`

SetProfession sets Profession field to given value.

### HasProfession

`func (o *MeiliPersonEntityV14) HasProfession() bool`

HasProfession returns a boolean if a field has been set.

### GetGrowth

`func (o *MeiliPersonEntityV14) GetGrowth() float32`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *MeiliPersonEntityV14) GetGrowthOk() (*float32, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *MeiliPersonEntityV14) SetGrowth(v float32)`

SetGrowth sets Growth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


