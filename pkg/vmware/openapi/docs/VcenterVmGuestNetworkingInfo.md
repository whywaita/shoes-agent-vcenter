# VcenterVmGuestNetworkingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsValues** | Pointer to [**VcenterVmGuestDnsAssignedValues**](VcenterVmGuestDnsAssignedValues.md) |  | [optional] 
**Dns** | Pointer to [**VcenterVmGuestDnsConfigInfo**](VcenterVmGuestDnsConfigInfo.md) |  | [optional] 

## Methods

### NewVcenterVmGuestNetworkingInfo

`func NewVcenterVmGuestNetworkingInfo() *VcenterVmGuestNetworkingInfo`

NewVcenterVmGuestNetworkingInfo instantiates a new VcenterVmGuestNetworkingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestNetworkingInfoWithDefaults

`func NewVcenterVmGuestNetworkingInfoWithDefaults() *VcenterVmGuestNetworkingInfo`

NewVcenterVmGuestNetworkingInfoWithDefaults instantiates a new VcenterVmGuestNetworkingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsValues

`func (o *VcenterVmGuestNetworkingInfo) GetDnsValues() VcenterVmGuestDnsAssignedValues`

GetDnsValues returns the DnsValues field if non-nil, zero value otherwise.

### GetDnsValuesOk

`func (o *VcenterVmGuestNetworkingInfo) GetDnsValuesOk() (*VcenterVmGuestDnsAssignedValues, bool)`

GetDnsValuesOk returns a tuple with the DnsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValues

`func (o *VcenterVmGuestNetworkingInfo) SetDnsValues(v VcenterVmGuestDnsAssignedValues)`

SetDnsValues sets DnsValues field to given value.

### HasDnsValues

`func (o *VcenterVmGuestNetworkingInfo) HasDnsValues() bool`

HasDnsValues returns a boolean if a field has been set.

### GetDns

`func (o *VcenterVmGuestNetworkingInfo) GetDns() VcenterVmGuestDnsConfigInfo`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *VcenterVmGuestNetworkingInfo) GetDnsOk() (*VcenterVmGuestDnsConfigInfo, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *VcenterVmGuestNetworkingInfo) SetDns(v VcenterVmGuestDnsConfigInfo)`

SetDns sets Dns field to given value.

### HasDns

`func (o *VcenterVmGuestNetworkingInfo) HasDns() bool`

HasDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


