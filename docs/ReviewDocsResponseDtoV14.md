# ReviewDocsResponseDtoV14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Docs** | [**[]Review**](Review.md) |  | 
**Total** | **float32** | Общее количество результатов | 
**Limit** | **float32** | Количество результатов на странице | 
**Page** | **float32** | Текущая страница | 
**Pages** | **float32** | Сколько страниц всего | 

## Methods

### NewReviewDocsResponseDtoV14

`func NewReviewDocsResponseDtoV14(docs []Review, total float32, limit float32, page float32, pages float32, ) *ReviewDocsResponseDtoV14`

NewReviewDocsResponseDtoV14 instantiates a new ReviewDocsResponseDtoV14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewDocsResponseDtoV14WithDefaults

`func NewReviewDocsResponseDtoV14WithDefaults() *ReviewDocsResponseDtoV14`

NewReviewDocsResponseDtoV14WithDefaults instantiates a new ReviewDocsResponseDtoV14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocs

`func (o *ReviewDocsResponseDtoV14) GetDocs() []Review`

GetDocs returns the Docs field if non-nil, zero value otherwise.

### GetDocsOk

`func (o *ReviewDocsResponseDtoV14) GetDocsOk() (*[]Review, bool)`

GetDocsOk returns a tuple with the Docs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocs

`func (o *ReviewDocsResponseDtoV14) SetDocs(v []Review)`

SetDocs sets Docs field to given value.


### GetTotal

`func (o *ReviewDocsResponseDtoV14) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ReviewDocsResponseDtoV14) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ReviewDocsResponseDtoV14) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetLimit

`func (o *ReviewDocsResponseDtoV14) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ReviewDocsResponseDtoV14) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ReviewDocsResponseDtoV14) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *ReviewDocsResponseDtoV14) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ReviewDocsResponseDtoV14) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ReviewDocsResponseDtoV14) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *ReviewDocsResponseDtoV14) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ReviewDocsResponseDtoV14) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ReviewDocsResponseDtoV14) SetPages(v float32)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


