# VcenterVmHardwareAdapterScsiInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Device label. | 
**Type** | [**VcenterVmHardwareAdapterScsiType**](VcenterVmHardwareAdapterScsiType.md) |  | 
**Scsi** | [**VcenterVmHardwareScsiAddressInfo**](VcenterVmHardwareScsiAddressInfo.md) |  | 
**PciSlotNumber** | Pointer to **int64** | Address of the SCSI adapter on the PCI bus. If the PCI address is invalid, the server will change it when the VM is started or as the device is hot added. May be unset if the virtual machine has never been powered on since the adapter was created. | [optional] 
**Sharing** | [**VcenterVmHardwareAdapterScsiSharing**](VcenterVmHardwareAdapterScsiSharing.md) |  | 

## Methods

### NewVcenterVmHardwareAdapterScsiInfo

`func NewVcenterVmHardwareAdapterScsiInfo(label string, type_ VcenterVmHardwareAdapterScsiType, scsi VcenterVmHardwareScsiAddressInfo, sharing VcenterVmHardwareAdapterScsiSharing, ) *VcenterVmHardwareAdapterScsiInfo`

NewVcenterVmHardwareAdapterScsiInfo instantiates a new VcenterVmHardwareAdapterScsiInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterScsiInfoWithDefaults

`func NewVcenterVmHardwareAdapterScsiInfoWithDefaults() *VcenterVmHardwareAdapterScsiInfo`

NewVcenterVmHardwareAdapterScsiInfoWithDefaults instantiates a new VcenterVmHardwareAdapterScsiInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *VcenterVmHardwareAdapterScsiInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *VcenterVmHardwareAdapterScsiInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *VcenterVmHardwareAdapterScsiInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *VcenterVmHardwareAdapterScsiInfo) GetType() VcenterVmHardwareAdapterScsiType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareAdapterScsiInfo) GetTypeOk() (*VcenterVmHardwareAdapterScsiType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareAdapterScsiInfo) SetType(v VcenterVmHardwareAdapterScsiType)`

SetType sets Type field to given value.


### GetScsi

`func (o *VcenterVmHardwareAdapterScsiInfo) GetScsi() VcenterVmHardwareScsiAddressInfo`

GetScsi returns the Scsi field if non-nil, zero value otherwise.

### GetScsiOk

`func (o *VcenterVmHardwareAdapterScsiInfo) GetScsiOk() (*VcenterVmHardwareScsiAddressInfo, bool)`

GetScsiOk returns a tuple with the Scsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi

`func (o *VcenterVmHardwareAdapterScsiInfo) SetScsi(v VcenterVmHardwareScsiAddressInfo)`

SetScsi sets Scsi field to given value.


### GetPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiInfo) GetPciSlotNumber() int64`

GetPciSlotNumber returns the PciSlotNumber field if non-nil, zero value otherwise.

### GetPciSlotNumberOk

`func (o *VcenterVmHardwareAdapterScsiInfo) GetPciSlotNumberOk() (*int64, bool)`

GetPciSlotNumberOk returns a tuple with the PciSlotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiInfo) SetPciSlotNumber(v int64)`

SetPciSlotNumber sets PciSlotNumber field to given value.

### HasPciSlotNumber

`func (o *VcenterVmHardwareAdapterScsiInfo) HasPciSlotNumber() bool`

HasPciSlotNumber returns a boolean if a field has been set.

### GetSharing

`func (o *VcenterVmHardwareAdapterScsiInfo) GetSharing() VcenterVmHardwareAdapterScsiSharing`

GetSharing returns the Sharing field if non-nil, zero value otherwise.

### GetSharingOk

`func (o *VcenterVmHardwareAdapterScsiInfo) GetSharingOk() (*VcenterVmHardwareAdapterScsiSharing, bool)`

GetSharingOk returns a tuple with the Sharing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharing

`func (o *VcenterVmHardwareAdapterScsiInfo) SetSharing(v VcenterVmHardwareAdapterScsiSharing)`

SetSharing sets Sharing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


