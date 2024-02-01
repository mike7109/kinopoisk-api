# SearchPersonResponseDtoV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docs** | [**[]MeiliPersonEntityV14**](MeiliPersonEntityV14.md) |  | 
**Total** | **float32** | Общее количество результатов | 
**Limit** | **float32** | Количество результатов на странице | 
**Page** | **float32** | Текущая страница | 
**Pages** | **float32** | Сколько страниц всего | 

## Methods

### NewSearchPersonResponseDtoV14

`func NewSearchPersonResponseDtoV14(docs []MeiliPersonEntityV14, total float32, limit float32, page float32, pages float32, ) *SearchPersonResponseDtoV14`

NewSearchPersonResponseDtoV14 instantiates a new SearchPersonResponseDtoV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchPersonResponseDtoV14WithDefaults

`func NewSearchPersonResponseDtoV14WithDefaults() *SearchPersonResponseDtoV14`

NewSearchPersonResponseDtoV14WithDefaults instantiates a new SearchPersonResponseDtoV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocs

`func (o *SearchPersonResponseDtoV14) GetDocs() []MeiliPersonEntityV14`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *SearchPersonResponseDtoV14) GetDocsOk() (*[]MeiliPersonEntityV14, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *SearchPersonResponseDtoV14) SetDocs(v []MeiliPersonEntityV14)`

SetDocs sets Docs field to given value.


### GetTotal

`func (o *SearchPersonResponseDtoV14) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SearchPersonResponseDtoV14) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SearchPersonResponseDtoV14) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *SearchPersonResponseDtoV14) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SearchPersonResponseDtoV14) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SearchPersonResponseDtoV14) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *SearchPersonResponseDtoV14) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchPersonResponseDtoV14) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchPersonResponseDtoV14) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *SearchPersonResponseDtoV14) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *SearchPersonResponseDtoV14) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *SearchPersonResponseDtoV14) SetPages(v float32)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


