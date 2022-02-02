# VcenterVmHardwareSerialBackingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareSerialBackingType**](VcenterVmHardwareSerialBackingType.md) |  | 
**File** | Pointer to **string** | Path of the file backing the virtual serial port. This field is optional and it is only relevant when the value of Serial.BackingSpec.type is FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device backing the virtual serial port.    If unset, the virtual serial port will be configured to automatically detect a suitable host device. | [optional] 
**Pipe** | Pointer to **string** | Name of the pipe backing the virtual serial port. This field is optional and it is only relevant when the value of Serial.BackingSpec.type is one of PIPE_SERVER or PIPE_CLIENT. | [optional] 
**NoRxLoss** | Pointer to **bool** | Flag that enables optimized data transfer over the pipe. When the value is true, the host buffers data to prevent data overrun. This allows the virtual machine to read all of the data transferred over the pipe with no data loss. If unset, defaults to false. | [optional] 
**NetworkLocation** | Pointer to **string** | URI specifying the location of the network service backing the virtual serial port.     - If Serial.BackingSpec.type is NETWORK_SERVER, this field is the location used by clients to connect to this server. The hostname part of the URI should either be empty or should specify the address of the host on which the virtual machine is running.    - If Serial.BackingSpec.type is NETWORK_CLIENT, this field is the location used by the virtual machine to connect to the remote server.   This field is optional and it is only relevant when the value of Serial.BackingSpec.type is one of NETWORK_SERVER or NETWORK_CLIENT. | [optional] 
**Proxy** | Pointer to **string** | Proxy service that provides network access to the network backing. If set, the virtual machine initiates a connection with the proxy service and forwards the traffic to the proxy. If unset, no proxy service should be used. | [optional] 

## Methods

### NewVcenterVmHardwareSerialBackingSpec

`func NewVcenterVmHardwareSerialBackingSpec(type_ VcenterVmHardwareSerialBackingType, ) *VcenterVmHardwareSerialBackingSpec`

NewVcenterVmHardwareSerialBackingSpec instantiates a new VcenterVmHardwareSerialBackingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareSerialBackingSpecWithDefaults

`func NewVcenterVmHardwareSerialBackingSpecWithDefaults() *VcenterVmHardwareSerialBackingSpec`

NewVcenterVmHardwareSerialBackingSpecWithDefaults instantiates a new VcenterVmHardwareSerialBackingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareSerialBackingSpec) GetType() VcenterVmHardwareSerialBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetTypeOk() (*VcenterVmHardwareSerialBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareSerialBackingSpec) SetType(v VcenterVmHardwareSerialBackingType)`

SetType sets Type field to given value.


### GetFile

`func (o *VcenterVmHardwareSerialBackingSpec) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *VcenterVmHardwareSerialBackingSpec) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *VcenterVmHardwareSerialBackingSpec) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareSerialBackingSpec) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareSerialBackingSpec) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareSerialBackingSpec) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetPipe

`func (o *VcenterVmHardwareSerialBackingSpec) GetPipe() string`

GetPipe returns the Pipe field if non-nil, zero value otherwise.

### GetPipeOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetPipeOk() (*string, bool)`

GetPipeOk returns a tuple with the Pipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipe

`func (o *VcenterVmHardwareSerialBackingSpec) SetPipe(v string)`

SetPipe sets Pipe field to given value.

### HasPipe

`func (o *VcenterVmHardwareSerialBackingSpec) HasPipe() bool`

HasPipe returns a boolean if a field has been set.

### GetNoRxLoss

`func (o *VcenterVmHardwareSerialBackingSpec) GetNoRxLoss() bool`

GetNoRxLoss returns the NoRxLoss field if non-nil, zero value otherwise.

### GetNoRxLossOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetNoRxLossOk() (*bool, bool)`

GetNoRxLossOk returns a tuple with the NoRxLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRxLoss

`func (o *VcenterVmHardwareSerialBackingSpec) SetNoRxLoss(v bool)`

SetNoRxLoss sets NoRxLoss field to given value.

### HasNoRxLoss

`func (o *VcenterVmHardwareSerialBackingSpec) HasNoRxLoss() bool`

HasNoRxLoss returns a boolean if a field has been set.

### GetNetworkLocation

`func (o *VcenterVmHardwareSerialBackingSpec) GetNetworkLocation() string`

GetNetworkLocation returns the NetworkLocation field if non-nil, zero value otherwise.

### GetNetworkLocationOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetNetworkLocationOk() (*string, bool)`

GetNetworkLocationOk returns a tuple with the NetworkLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLocation

`func (o *VcenterVmHardwareSerialBackingSpec) SetNetworkLocation(v string)`

SetNetworkLocation sets NetworkLocation field to given value.

### HasNetworkLocation

`func (o *VcenterVmHardwareSerialBackingSpec) HasNetworkLocation() bool`

HasNetworkLocation returns a boolean if a field has been set.

### GetProxy

`func (o *VcenterVmHardwareSerialBackingSpec) GetProxy() string`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *VcenterVmHardwareSerialBackingSpec) GetProxyOk() (*string, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *VcenterVmHardwareSerialBackingSpec) SetProxy(v string)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *VcenterVmHardwareSerialBackingSpec) HasProxy() bool`

HasProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


