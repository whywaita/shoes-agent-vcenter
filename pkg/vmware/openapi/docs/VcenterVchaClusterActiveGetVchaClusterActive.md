# VcenterVchaClusterActiveGetVchaClusterActive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcSpec** | Pointer to [**VcenterVchaCredentialsSpec**](VcenterVchaCredentialsSpec.md) |  | [optional] 
**Partial** | Pointer to **bool** | If true, then return only the information that does not require connecting to the Active vCenter Server.   If false or unset, then return all the information. If unset, then return all the information. | [optional] 

## Methods

### NewVcenterVchaClusterActiveGetVchaClusterActive

`func NewVcenterVchaClusterActiveGetVchaClusterActive() *VcenterVchaClusterActiveGetVchaClusterActive`

NewVcenterVchaClusterActiveGetVchaClusterActive instantiates a new VcenterVchaClusterActiveGetVchaClusterActive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterActiveGetVchaClusterActiveWithDefaults

`func NewVcenterVchaClusterActiveGetVchaClusterActiveWithDefaults() *VcenterVchaClusterActiveGetVchaClusterActive`

NewVcenterVchaClusterActiveGetVchaClusterActiveWithDefaults instantiates a new VcenterVchaClusterActiveGetVchaClusterActive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcSpec

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetVcSpec() VcenterVchaCredentialsSpec`

GetVcSpec returns the VcSpec field if non-nil, zero value otherwise.

### GetVcSpecOk

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool)`

GetVcSpecOk returns a tuple with the VcSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcSpec

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) SetVcSpec(v VcenterVchaCredentialsSpec)`

SetVcSpec sets VcSpec field to given value.

### HasVcSpec

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) HasVcSpec() bool`

HasVcSpec returns a boolean if a field has been set.

### GetPartial

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *VcenterVchaClusterActiveGetVchaClusterActive) HasPartial() bool`

HasPartial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


