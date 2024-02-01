# PersonAwardDocsResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docs** | [**[]PersonAward**](PersonAward.md) |  | 
**Total** | **float32** | Общее количество результатов | 
**Limit** | **float32** | Количество результатов на странице | 
**Page** | **float32** | Текущая страница | 
**Pages** | **float32** | Сколько страниц всего | 

## Methods

### NewPersonAwardDocsResponseDto

`func NewPersonAwardDocsResponseDto(docs []PersonAward, total float32, limit float32, page float32, pages float32, ) *PersonAwardDocsResponseDto`

NewPersonAwardDocsResponseDto instantiates a new PersonAwardDocsResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonAwardDocsResponseDtoWithDefaults

`func NewPersonAwardDocsResponseDtoWithDefaults() *PersonAwardDocsResponseDto`

NewPersonAwardDocsResponseDtoWithDefaults instantiates a new PersonAwardDocsResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocs

`func (o *PersonAwardDocsResponseDto) GetDocs() []PersonAward`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *PersonAwardDocsResponseDto) GetDocsOk() (*[]PersonAward, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *PersonAwardDocsResponseDto) SetDocs(v []PersonAward)`

SetDocs sets Docs field to given value.


### GetTotal

`func (o *PersonAwardDocsResponseDto) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PersonAwardDocsResponseDto) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PersonAwardDocsResponseDto) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *PersonAwardDocsResponseDto) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PersonAwardDocsResponseDto) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PersonAwardDocsResponseDto) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *PersonAwardDocsResponseDto) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PersonAwardDocsResponseDto) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PersonAwardDocsResponseDto) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *PersonAwardDocsResponseDto) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PersonAwardDocsResponseDto) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PersonAwardDocsResponseDto) SetPages(v float32)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


