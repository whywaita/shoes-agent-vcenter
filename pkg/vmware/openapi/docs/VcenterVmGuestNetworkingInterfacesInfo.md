# VcenterVmGuestNetworkingInterfacesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsValues** | Pointer to [**VcenterVmGuestDnsAssignedValues**](VcenterVmGuestDnsAssignedValues.md) |  | [optional] 
**MacAddress** | Pointer to **string** | MAC address of the adapter. If unset then not supported by the Guest OS. | [optional] 
**Dns** | Pointer to [**VcenterVmGuestDnsConfigInfo**](VcenterVmGuestDnsConfigInfo.md) |  | [optional] 
**Ip** | Pointer to [**VcenterVmGuestNetworkingInterfacesIpConfigInfo**](VcenterVmGuestNetworkingInterfacesIpConfigInfo.md) |  | [optional] 
**WinsServers** | Pointer to **[]string** | The IP addresses of any WINS name servers for the adapter. If unset then not supported by the Guest OS. | [optional] 
**Nic** | Pointer to **string** | Link to the corresponding virtual device. If unset then the interface is not backed by a virtual device. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Ethernet. | [optional] 

## Methods

### NewVcenterVmGuestNetworkingInterfacesInfo

`func NewVcenterVmGuestNetworkingInterfacesInfo() *VcenterVmGuestNetworkingInterfacesInfo`

NewVcenterVmGuestNetworkingInterfacesInfo instantiates a new VcenterVmGuestNetworkingInterfacesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestNetworkingInterfacesInfoWithDefaults

`func NewVcenterVmGuestNetworkingInterfacesInfoWithDefaults() *VcenterVmGuestNetworkingInterfacesInfo`

NewVcenterVmGuestNetworkingInterfacesInfoWithDefaults instantiates a new VcenterVmGuestNetworkingInterfacesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsValues

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetDnsValues() VcenterVmGuestDnsAssignedValues`

GetDnsValues returns the DnsValues field if non-nil, zero value otherwise.

### GetDnsValuesOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetDnsValuesOk() (*VcenterVmGuestDnsAssignedValues, bool)`

GetDnsValuesOk returns a tuple with the DnsValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsValues

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetDnsValues(v VcenterVmGuestDnsAssignedValues)`

SetDnsValues sets DnsValues field to given value.

### HasDnsValues

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasDnsValues() bool`

HasDnsValues returns a boolean if a field has been set.

### GetMacAddress

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetDns

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetDns() VcenterVmGuestDnsConfigInfo`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetDnsOk() (*VcenterVmGuestDnsConfigInfo, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetDns(v VcenterVmGuestDnsConfigInfo)`

SetDns sets Dns field to given value.

### HasDns

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetIp

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetIp() VcenterVmGuestNetworkingInterfacesIpConfigInfo`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetIpOk() (*VcenterVmGuestNetworkingInterfacesIpConfigInfo, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetIp(v VcenterVmGuestNetworkingInterfacesIpConfigInfo)`

SetIp sets Ip field to given value.

### HasIp

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetWinsServers

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetWinsServers() []string`

GetWinsServers returns the WinsServers field if non-nil, zero value otherwise.

### GetWinsServersOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetWinsServersOk() (*[]string, bool)`

GetWinsServersOk returns a tuple with the WinsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinsServers

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetWinsServers(v []string)`

SetWinsServers sets WinsServers field to given value.

### HasWinsServers

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasWinsServers() bool`

HasWinsServers returns a boolean if a field has been set.

### GetNic

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetNic() string`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *VcenterVmGuestNetworkingInterfacesInfo) GetNicOk() (*string, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *VcenterVmGuestNetworkingInterfacesInfo) SetNic(v string)`

SetNic sets Nic field to given value.

### HasNic

`func (o *VcenterVmGuestNetworkingInterfacesInfo) HasNic() bool`

HasNic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


