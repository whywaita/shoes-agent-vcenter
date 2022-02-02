# VcenterVmPowerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**VcenterVmPowerState**](VcenterVmPowerState.md) |  | 
**CleanPowerOff** | Pointer to **bool** | Flag indicating whether the virtual machine was powered off cleanly. This field may be used to detect that the virtual machine crashed unexpectedly and should be restarted. This field is optional and it is only relevant when the value of Power.Info.state is POWERED_OFF. | [optional] 

## Methods

### NewVcenterVmPowerInfo

`func NewVcenterVmPowerInfo(state VcenterVmPowerState, ) *VcenterVmPowerInfo`

NewVcenterVmPowerInfo instantiates a new VcenterVmPowerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmPowerInfoWithDefaults

`func NewVcenterVmPowerInfoWithDefaults() *VcenterVmPowerInfo`

NewVcenterVmPowerInfoWithDefaults instantiates a new VcenterVmPowerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VcenterVmPowerInfo) GetState() VcenterVmPowerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmPowerInfo) GetStateOk() (*VcenterVmPowerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmPowerInfo) SetState(v VcenterVmPowerState)`

SetState sets State field to given value.


### GetCleanPowerOff

`func (o *VcenterVmPowerInfo) GetCleanPowerOff() bool`

GetCleanPowerOff returns the CleanPowerOff field if non-nil, zero value otherwise.

### GetCleanPowerOffOk

`func (o *VcenterVmPowerInfo) GetCleanPowerOffOk() (*bool, bool)`

GetCleanPowerOffOk returns a tuple with the CleanPowerOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanPowerOff

`func (o *VcenterVmPowerInfo) SetCleanPowerOff(v bool)`

SetCleanPowerOff sets CleanPowerOff field to given value.

### HasCleanPowerOff

`func (o *VcenterVmPowerInfo) HasCleanPowerOff() bool`

HasCleanPowerOff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


