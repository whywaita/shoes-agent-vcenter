# VcenterSystemConfigPscRegistrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IP address or DNS resolvable name of the PSC this appliance is registered with. | 
**HttpsPort** | **int64** | The HTTPs port used by the external PSC. | 
**SsoDomain** | **string** | The Single Sign-On domain name of the external PSC. | 

## Methods

### NewVcenterSystemConfigPscRegistrationInfo

`func NewVcenterSystemConfigPscRegistrationInfo(address string, httpsPort int64, ssoDomain string, ) *VcenterSystemConfigPscRegistrationInfo`

NewVcenterSystemConfigPscRegistrationInfo instantiates a new VcenterSystemConfigPscRegistrationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterSystemConfigPscRegistrationInfoWithDefaults

`func NewVcenterSystemConfigPscRegistrationInfoWithDefaults() *VcenterSystemConfigPscRegistrationInfo`

NewVcenterSystemConfigPscRegistrationInfoWithDefaults instantiates a new VcenterSystemConfigPscRegistrationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterSystemConfigPscRegistrationInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterSystemConfigPscRegistrationInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterSystemConfigPscRegistrationInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetHttpsPort

`func (o *VcenterSystemConfigPscRegistrationInfo) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterSystemConfigPscRegistrationInfo) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterSystemConfigPscRegistrationInfo) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.


### GetSsoDomain

`func (o *VcenterSystemConfigPscRegistrationInfo) GetSsoDomain() string`

GetSsoDomain returns the SsoDomain field if non-nil, zero value otherwise.

### GetSsoDomainOk

`func (o *VcenterSystemConfigPscRegistrationInfo) GetSsoDomainOk() (*string, bool)`

GetSsoDomainOk returns a tuple with the SsoDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomain

`func (o *VcenterSystemConfigPscRegistrationInfo) SetSsoDomain(v string)`

SetSsoDomain sets SsoDomain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


