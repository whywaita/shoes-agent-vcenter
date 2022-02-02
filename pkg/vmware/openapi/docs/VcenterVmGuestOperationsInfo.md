# VcenterVmGuestOperationsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuestOperationsReady** | **bool** | Guest operations availability. Whether or not the virtual machine is ready to process guest operations. | 
**InteractiveGuestOperationsReady** | **bool** | Interactive guest operations availability. Whether or not the virtual machine is ready to process interactive guest operations. | 

## Methods

### NewVcenterVmGuestOperationsInfo

`func NewVcenterVmGuestOperationsInfo(guestOperationsReady bool, interactiveGuestOperationsReady bool, ) *VcenterVmGuestOperationsInfo`

NewVcenterVmGuestOperationsInfo instantiates a new VcenterVmGuestOperationsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestOperationsInfoWithDefaults

`func NewVcenterVmGuestOperationsInfoWithDefaults() *VcenterVmGuestOperationsInfo`

NewVcenterVmGuestOperationsInfoWithDefaults instantiates a new VcenterVmGuestOperationsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuestOperationsReady

`func (o *VcenterVmGuestOperationsInfo) GetGuestOperationsReady() bool`

GetGuestOperationsReady returns the GuestOperationsReady field if non-nil, zero value otherwise.

### GetGuestOperationsReadyOk

`func (o *VcenterVmGuestOperationsInfo) GetGuestOperationsReadyOk() (*bool, bool)`

GetGuestOperationsReadyOk returns a tuple with the GuestOperationsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOperationsReady

`func (o *VcenterVmGuestOperationsInfo) SetGuestOperationsReady(v bool)`

SetGuestOperationsReady sets GuestOperationsReady field to given value.


### GetInteractiveGuestOperationsReady

`func (o *VcenterVmGuestOperationsInfo) GetInteractiveGuestOperationsReady() bool`

GetInteractiveGuestOperationsReady returns the InteractiveGuestOperationsReady field if non-nil, zero value otherwise.

### GetInteractiveGuestOperationsReadyOk

`func (o *VcenterVmGuestOperationsInfo) GetInteractiveGuestOperationsReadyOk() (*bool, bool)`

GetInteractiveGuestOperationsReadyOk returns a tuple with the InteractiveGuestOperationsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractiveGuestOperationsReady

`func (o *VcenterVmGuestOperationsInfo) SetInteractiveGuestOperationsReady(v bool)`

SetInteractiveGuestOperationsReady sets InteractiveGuestOperationsReady field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


