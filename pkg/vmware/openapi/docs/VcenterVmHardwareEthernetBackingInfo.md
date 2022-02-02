# VcenterVmHardwareEthernetBackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareEthernetBackingType**](VcenterVmHardwareEthernetBackingType.md) |  | 
**Network** | Pointer to **string** | Identifier of the network backing the virtual Ethernet adapter. If unset, the identifier of the network backing could not be determined. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | [optional] 
**NetworkName** | Pointer to **string** | Name of the standard portgroup backing the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingInfo.type is STANDARD_PORTGROUP. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device backing the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingInfo.type is HOST_DEVICE. | [optional] 
**DistributedSwitchUuid** | Pointer to **string** | UUID of the distributed virtual switch that backs the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingInfo.type is DISTRIBUTED_PORTGROUP. | [optional] 
**DistributedPort** | Pointer to **string** | Key of the distributed virtual port that backs the virtual Ethernet adapter. This field will be unset if the virtual Ethernet device is not bound to a distributed virtual port; this can happen if the virtual machine is powered off or the virtual Ethernet device is not connected. | [optional] 
**ConnectionCookie** | Pointer to **int64** | Server-generated cookie that identifies the connection to the port. This ookie may be used to verify that the virtual machine is the rightful owner of the port. This field will be unset if the virtual Ethernet device is not bound to a distributed virtual port; this can happen if the virtual machine is powered off or the virtual Ethernet device is not connected. | [optional] 
**OpaqueNetworkType** | Pointer to **string** | Type of the opaque network that backs the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingInfo.type is OPAQUE_NETWORK. | [optional] 
**OpaqueNetworkId** | Pointer to **string** | Identifier of the opaque network that backs the virtual Ethernet adapter. This field is optional and it is only relevant when the value of Ethernet.BackingInfo.type is OPAQUE_NETWORK. | [optional] 

## Methods

### NewVcenterVmHardwareEthernetBackingInfo

`func NewVcenterVmHardwareEthernetBackingInfo(type_ VcenterVmHardwareEthernetBackingType, ) *VcenterVmHardwareEthernetBackingInfo`

NewVcenterVmHardwareEthernetBackingInfo instantiates a new VcenterVmHardwareEthernetBackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareEthernetBackingInfoWithDefaults

`func NewVcenterVmHardwareEthernetBackingInfoWithDefaults() *VcenterVmHardwareEthernetBackingInfo`

NewVcenterVmHardwareEthernetBackingInfoWithDefaults instantiates a new VcenterVmHardwareEthernetBackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareEthernetBackingInfo) GetType() VcenterVmHardwareEthernetBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetTypeOk() (*VcenterVmHardwareEthernetBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareEthernetBackingInfo) SetType(v VcenterVmHardwareEthernetBackingType)`

SetType sets Type field to given value.


### GetNetwork

`func (o *VcenterVmHardwareEthernetBackingInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterVmHardwareEthernetBackingInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VcenterVmHardwareEthernetBackingInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkName

`func (o *VcenterVmHardwareEthernetBackingInfo) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *VcenterVmHardwareEthernetBackingInfo) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *VcenterVmHardwareEthernetBackingInfo) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareEthernetBackingInfo) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareEthernetBackingInfo) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareEthernetBackingInfo) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetDistributedSwitchUuid

`func (o *VcenterVmHardwareEthernetBackingInfo) GetDistributedSwitchUuid() string`

GetDistributedSwitchUuid returns the DistributedSwitchUuid field if non-nil, zero value otherwise.

### GetDistributedSwitchUuidOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetDistributedSwitchUuidOk() (*string, bool)`

GetDistributedSwitchUuidOk returns a tuple with the DistributedSwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedSwitchUuid

`func (o *VcenterVmHardwareEthernetBackingInfo) SetDistributedSwitchUuid(v string)`

SetDistributedSwitchUuid sets DistributedSwitchUuid field to given value.

### HasDistributedSwitchUuid

`func (o *VcenterVmHardwareEthernetBackingInfo) HasDistributedSwitchUuid() bool`

HasDistributedSwitchUuid returns a boolean if a field has been set.

### GetDistributedPort

`func (o *VcenterVmHardwareEthernetBackingInfo) GetDistributedPort() string`

GetDistributedPort returns the DistributedPort field if non-nil, zero value otherwise.

### GetDistributedPortOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetDistributedPortOk() (*string, bool)`

GetDistributedPortOk returns a tuple with the DistributedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedPort

`func (o *VcenterVmHardwareEthernetBackingInfo) SetDistributedPort(v string)`

SetDistributedPort sets DistributedPort field to given value.

### HasDistributedPort

`func (o *VcenterVmHardwareEthernetBackingInfo) HasDistributedPort() bool`

HasDistributedPort returns a boolean if a field has been set.

### GetConnectionCookie

`func (o *VcenterVmHardwareEthernetBackingInfo) GetConnectionCookie() int64`

GetConnectionCookie returns the ConnectionCookie field if non-nil, zero value otherwise.

### GetConnectionCookieOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetConnectionCookieOk() (*int64, bool)`

GetConnectionCookieOk returns a tuple with the ConnectionCookie field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCookie

`func (o *VcenterVmHardwareEthernetBackingInfo) SetConnectionCookie(v int64)`

SetConnectionCookie sets ConnectionCookie field to given value.

### HasConnectionCookie

`func (o *VcenterVmHardwareEthernetBackingInfo) HasConnectionCookie() bool`

HasConnectionCookie returns a boolean if a field has been set.

### GetOpaqueNetworkType

`func (o *VcenterVmHardwareEthernetBackingInfo) GetOpaqueNetworkType() string`

GetOpaqueNetworkType returns the OpaqueNetworkType field if non-nil, zero value otherwise.

### GetOpaqueNetworkTypeOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetOpaqueNetworkTypeOk() (*string, bool)`

GetOpaqueNetworkTypeOk returns a tuple with the OpaqueNetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueNetworkType

`func (o *VcenterVmHardwareEthernetBackingInfo) SetOpaqueNetworkType(v string)`

SetOpaqueNetworkType sets OpaqueNetworkType field to given value.

### HasOpaqueNetworkType

`func (o *VcenterVmHardwareEthernetBackingInfo) HasOpaqueNetworkType() bool`

HasOpaqueNetworkType returns a boolean if a field has been set.

### GetOpaqueNetworkId

`func (o *VcenterVmHardwareEthernetBackingInfo) GetOpaqueNetworkId() string`

GetOpaqueNetworkId returns the OpaqueNetworkId field if non-nil, zero value otherwise.

### GetOpaqueNetworkIdOk

`func (o *VcenterVmHardwareEthernetBackingInfo) GetOpaqueNetworkIdOk() (*string, bool)`

GetOpaqueNetworkIdOk returns a tuple with the OpaqueNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueNetworkId

`func (o *VcenterVmHardwareEthernetBackingInfo) SetOpaqueNetworkId(v string)`

SetOpaqueNetworkId sets OpaqueNetworkId field to given value.

### HasOpaqueNetworkId

`func (o *VcenterVmHardwareEthernetBackingInfo) HasOpaqueNetworkId() bool`

HasOpaqueNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


