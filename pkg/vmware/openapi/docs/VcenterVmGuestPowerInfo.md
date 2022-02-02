# VcenterVmGuestPowerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**VcenterVmGuestPowerState**](VcenterVmGuestPowerState.md) |  | 
**OperationsReady** | **bool** | Flag indicating if the virtual machine is ready to process soft power operations. | 

## Methods

### NewVcenterVmGuestPowerInfo

`func NewVcenterVmGuestPowerInfo(state VcenterVmGuestPowerState, operationsReady bool, ) *VcenterVmGuestPowerInfo`

NewVcenterVmGuestPowerInfo instantiates a new VcenterVmGuestPowerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestPowerInfoWithDefaults

`func NewVcenterVmGuestPowerInfoWithDefaults() *VcenterVmGuestPowerInfo`

NewVcenterVmGuestPowerInfoWithDefaults instantiates a new VcenterVmGuestPowerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VcenterVmGuestPowerInfo) GetState() VcenterVmGuestPowerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterVmGuestPowerInfo) GetStateOk() (*VcenterVmGuestPowerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterVmGuestPowerInfo) SetState(v VcenterVmGuestPowerState)`

SetState sets State field to given value.


### GetOperationsReady

`func (o *VcenterVmGuestPowerInfo) GetOperationsReady() bool`

GetOperationsReady returns the OperationsReady field if non-nil, zero value otherwise.

### GetOperationsReadyOk

`func (o *VcenterVmGuestPowerInfo) GetOperationsReadyOk() (*bool, bool)`

GetOperationsReadyOk returns a tuple with the OperationsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationsReady

`func (o *VcenterVmGuestPowerInfo) SetOperationsReady(v bool)`

SetOperationsReady sets OperationsReady field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


