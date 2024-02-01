# UnauthorizedErrorResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **float32** |  | 
**Message** | **string** |  | 
**Error** | **string** |  | 

## Methods

### NewUnauthorizedErrorResponseDto

`func NewUnauthorizedErrorResponseDto(statusCode float32, message string, error_ string, ) *UnauthorizedErrorResponseDto`

NewUnauthorizedErrorResponseDto instantiates a new UnauthorizedErrorResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorResponseDtoWithDefaults

`func NewUnauthorizedErrorResponseDtoWithDefaults() *UnauthorizedErrorResponseDto`

NewUnauthorizedErrorResponseDtoWithDefaults instantiates a new UnauthorizedErrorResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *UnauthorizedErrorResponseDto) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UnauthorizedErrorResponseDto) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UnauthorizedErrorResponseDto) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.


### GetMessage

`func (o *UnauthorizedErrorResponseDto) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedErrorResponseDto) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedErrorResponseDto) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetError

`func (o *UnauthorizedErrorResponseDto) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UnauthorizedErrorResponseDto) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UnauthorizedErrorResponseDto) SetError(v string)`

SetError sets Error field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


