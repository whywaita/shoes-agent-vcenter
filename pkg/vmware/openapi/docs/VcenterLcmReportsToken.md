# VcenterLcmReportsToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name \&quot;session-id\&quot;. | 
**Expiry** | **time.Time** | Expiry time of the token | 

## Methods

### NewVcenterLcmReportsToken

`func NewVcenterLcmReportsToken(token string, expiry time.Time, ) *VcenterLcmReportsToken`

NewVcenterLcmReportsToken instantiates a new VcenterLcmReportsToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmReportsTokenWithDefaults

`func NewVcenterLcmReportsTokenWithDefaults() *VcenterLcmReportsToken`

NewVcenterLcmReportsTokenWithDefaults instantiates a new VcenterLcmReportsToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *VcenterLcmReportsToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VcenterLcmReportsToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VcenterLcmReportsToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiry

`func (o *VcenterLcmReportsToken) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *VcenterLcmReportsToken) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *VcenterLcmReportsToken) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


