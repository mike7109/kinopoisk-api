# Premiere

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** |  | [optional] 
**World** | Pointer to **NullableTime** | Для более релевантного поиска, используйте интервал дат 01.02.2022-01.02.2023 | [optional] 
**Russia** | Pointer to **NullableTime** | Для более релевантного поиска, используйте интервал дат 01.02.2022-01.02.2023 | [optional] 
**Digital** | Pointer to **NullableString** |  | [optional] 
**Cinema** | Pointer to **NullableTime** | Для более релевантного поиска, используйте интервал дат 01.02.2022-01.02.2023 | [optional] 
**Bluray** | **string** |  | 
**Dvd** | **string** |  | 

## Methods

### NewPremiere

`func NewPremiere(bluray string, dvd string, ) *Premiere`

NewPremiere instantiates a new Premiere object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPremiereWithDefaults

`func NewPremiereWithDefaults() *Premiere`

NewPremiereWithDefaults instantiates a new Premiere object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *Premiere) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Premiere) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Premiere) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Premiere) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Premiere) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Premiere) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetWorld

`func (o *Premiere) GetWorld() time.Time`

GetWorld returns the World field if non-nil, zero value otherwise.

### GetWorldOk

`func (o *Premiere) GetWorldOk() (*time.Time, bool)`

GetWorldOk returns a tuple with the World field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorld

`func (o *Premiere) SetWorld(v time.Time)`

SetWorld sets World field to given value.

### HasWorld

`func (o *Premiere) HasWorld() bool`

HasWorld returns a boolean if a field has been set.

### SetWorldNil

`func (o *Premiere) SetWorldNil(b bool)`

 SetWorldNil sets the value for World to be an explicit nil

### UnsetWorld
`func (o *Premiere) UnsetWorld()`

UnsetWorld ensures that no value is present for World, not even an explicit nil
### GetRussia

`func (o *Premiere) GetRussia() time.Time`

GetRussia returns the Russia field if non-nil, zero value otherwise.

### GetRussiaOk

`func (o *Premiere) GetRussiaOk() (*time.Time, bool)`

GetRussiaOk returns a tuple with the Russia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRussia

`func (o *Premiere) SetRussia(v time.Time)`

SetRussia sets Russia field to given value.

### HasRussia

`func (o *Premiere) HasRussia() bool`

HasRussia returns a boolean if a field has been set.

### SetRussiaNil

`func (o *Premiere) SetRussiaNil(b bool)`

 SetRussiaNil sets the value for Russia to be an explicit nil

### UnsetRussia
`func (o *Premiere) UnsetRussia()`

UnsetRussia ensures that no value is present for Russia, not even an explicit nil
### GetDigital

`func (o *Premiere) GetDigital() string`

GetDigital returns the Digital field if non-nil, zero value otherwise.

### GetDigitalOk

`func (o *Premiere) GetDigitalOk() (*string, bool)`

GetDigitalOk returns a tuple with the Digital field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigital

`func (o *Premiere) SetDigital(v string)`

SetDigital sets Digital field to given value.

### HasDigital

`func (o *Premiere) HasDigital() bool`

HasDigital returns a boolean if a field has been set.

### SetDigitalNil

`func (o *Premiere) SetDigitalNil(b bool)`

 SetDigitalNil sets the value for Digital to be an explicit nil

### UnsetDigital
`func (o *Premiere) UnsetDigital()`

UnsetDigital ensures that no value is present for Digital, not even an explicit nil
### GetCinema

`func (o *Premiere) GetCinema() time.Time`

GetCinema returns the Cinema field if non-nil, zero value otherwise.

### GetCinemaOk

`func (o *Premiere) GetCinemaOk() (*time.Time, bool)`

GetCinemaOk returns a tuple with the Cinema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCinema

`func (o *Premiere) SetCinema(v time.Time)`

SetCinema sets Cinema field to given value.

### HasCinema

`func (o *Premiere) HasCinema() bool`

HasCinema returns a boolean if a field has been set.

### SetCinemaNil

`func (o *Premiere) SetCinemaNil(b bool)`

 SetCinemaNil sets the value for Cinema to be an explicit nil

### UnsetCinema
`func (o *Premiere) UnsetCinema()`

UnsetCinema ensures that no value is present for Cinema, not even an explicit nil
### GetBluray

`func (o *Premiere) GetBluray() string`

GetBluray returns the Bluray field if non-nil, zero value otherwise.

### GetBlurayOk

`func (o *Premiere) GetBlurayOk() (*string, bool)`

GetBlurayOk returns a tuple with the Bluray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBluray

`func (o *Premiere) SetBluray(v string)`

SetBluray sets Bluray field to given value.


### GetDvd

`func (o *Premiere) GetDvd() string`

GetDvd returns the Dvd field if non-nil, zero value otherwise.

### GetDvdOk

`func (o *Premiere) GetDvdOk() (*string, bool)`

GetDvdOk returns a tuple with the Dvd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvd

`func (o *Premiere) SetDvd(v string)`

SetDvd sets Dvd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


