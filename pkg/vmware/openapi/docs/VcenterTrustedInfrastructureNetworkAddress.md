# VcenterTrustedInfrastructureNetworkAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The IP address or DNS resolvable name of the service. | 
**Port** | Pointer to **int64** | The port of the service. If unset, port 443 will be used. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureNetworkAddress

`func NewVcenterTrustedInfrastructureNetworkAddress(hostname string, ) *VcenterTrustedInfrastructureNetworkAddress`

NewVcenterTrustedInfrastructureNetworkAddress instantiates a new VcenterTrustedInfrastructureNetworkAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureNetworkAddressWithDefaults

`func NewVcenterTrustedInfrastructureNetworkAddressWithDefaults() *VcenterTrustedInfrastructureNetworkAddress`

NewVcenterTrustedInfrastructureNetworkAddressWithDefaults instantiates a new VcenterTrustedInfrastructureNetworkAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterTrustedInfrastructureNetworkAddress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterTrustedInfrastructureNetworkAddress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterTrustedInfrastructureNetworkAddress) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *VcenterTrustedInfrastructureNetworkAddress) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterTrustedInfrastructureNetworkAddress) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterTrustedInfrastructureNetworkAddress) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VcenterTrustedInfrastructureNetworkAddress) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


