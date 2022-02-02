# VcenterHostSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Identifier of the host. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | 
**Name** | **string** | Name of the host. | 
**ConnectionState** | [**VcenterHostConnectionState**](VcenterHostConnectionState.md) |  | 
**PowerState** | Pointer to [**VcenterHostPowerState**](VcenterHostPowerState.md) |  | [optional] 

## Methods

### NewVcenterHostSummary

`func NewVcenterHostSummary(host string, name string, connectionState VcenterHostConnectionState, ) *VcenterHostSummary`

NewVcenterHostSummary instantiates a new VcenterHostSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHostSummaryWithDefaults

`func NewVcenterHostSummaryWithDefaults() *VcenterHostSummary`

NewVcenterHostSummaryWithDefaults instantiates a new VcenterHostSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *VcenterHostSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterHostSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterHostSummary) SetHost(v string)`

SetHost sets Host field to given value.


### GetName

`func (o *VcenterHostSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterHostSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterHostSummary) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionState

`func (o *VcenterHostSummary) GetConnectionState() VcenterHostConnectionState`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *VcenterHostSummary) GetConnectionStateOk() (*VcenterHostConnectionState, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *VcenterHostSummary) SetConnectionState(v VcenterHostConnectionState)`

SetConnectionState sets ConnectionState field to given value.


### GetPowerState

`func (o *VcenterHostSummary) GetPowerState() VcenterHostPowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VcenterHostSummary) GetPowerStateOk() (*VcenterHostPowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VcenterHostSummary) SetPowerState(v VcenterHostPowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VcenterHostSummary) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


