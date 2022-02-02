# VcenterVmHardwareDiskCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterVmHardwareDiskHostBusAdapterType**](VcenterVmHardwareDiskHostBusAdapterType.md) |  | [optional] 
**Ide** | Pointer to [**VcenterVmHardwareIdeAddressSpec**](VcenterVmHardwareIdeAddressSpec.md) |  | [optional] 
**Scsi** | Pointer to [**VcenterVmHardwareScsiAddressSpec**](VcenterVmHardwareScsiAddressSpec.md) |  | [optional] 
**Sata** | Pointer to [**VcenterVmHardwareSataAddressSpec**](VcenterVmHardwareSataAddressSpec.md) |  | [optional] 
**Nvme** | Pointer to [**VcenterVmHardwareNvmeAddressSpec**](VcenterVmHardwareNvmeAddressSpec.md) |  | [optional] 
**Backing** | Pointer to [**VcenterVmHardwareDiskBackingSpec**](VcenterVmHardwareDiskBackingSpec.md) |  | [optional] 
**NewVmdk** | Pointer to [**VcenterVmHardwareDiskVmdkCreateSpec**](VcenterVmHardwareDiskVmdkCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterVmHardwareDiskCreateSpec

`func NewVcenterVmHardwareDiskCreateSpec() *VcenterVmHardwareDiskCreateSpec`

NewVcenterVmHardwareDiskCreateSpec instantiates a new VcenterVmHardwareDiskCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskCreateSpecWithDefaults

`func NewVcenterVmHardwareDiskCreateSpecWithDefaults() *VcenterVmHardwareDiskCreateSpec`

NewVcenterVmHardwareDiskCreateSpecWithDefaults instantiates a new VcenterVmHardwareDiskCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareDiskCreateSpec) GetType() VcenterVmHardwareDiskHostBusAdapterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetTypeOk() (*VcenterVmHardwareDiskHostBusAdapterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareDiskCreateSpec) SetType(v VcenterVmHardwareDiskHostBusAdapterType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterVmHardwareDiskCreateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIde

`func (o *VcenterVmHardwareDiskCreateSpec) GetIde() VcenterVmHardwareIdeAddressSpec`

GetIde returns the Ide field if non-nil, zero value otherwise.

### GetIdeOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetIdeOk() (*VcenterVmHardwareIdeAddressSpec, bool)`

GetIdeOk returns a tuple with the Ide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde

`func (o *VcenterVmHardwareDiskCreateSpec) SetIde(v VcenterVmHardwareIdeAddressSpec)`

SetIde sets Ide field to given value.

### HasIde

`func (o *VcenterVmHardwareDiskCreateSpec) HasIde() bool`

HasIde returns a boolean if a field has been set.

### GetScsi

`func (o *VcenterVmHardwareDiskCreateSpec) GetScsi() VcenterVmHardwareScsiAddressSpec`

GetScsi returns the Scsi field if non-nil, zero value otherwise.

### GetScsiOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetScsiOk() (*VcenterVmHardwareScsiAddressSpec, bool)`

GetScsiOk returns a tuple with the Scsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi

`func (o *VcenterVmHardwareDiskCreateSpec) SetScsi(v VcenterVmHardwareScsiAddressSpec)`

SetScsi sets Scsi field to given value.

### HasScsi

`func (o *VcenterVmHardwareDiskCreateSpec) HasScsi() bool`

HasScsi returns a boolean if a field has been set.

### GetSata

`func (o *VcenterVmHardwareDiskCreateSpec) GetSata() VcenterVmHardwareSataAddressSpec`

GetSata returns the Sata field if non-nil, zero value otherwise.

### GetSataOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetSataOk() (*VcenterVmHardwareSataAddressSpec, bool)`

GetSataOk returns a tuple with the Sata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata

`func (o *VcenterVmHardwareDiskCreateSpec) SetSata(v VcenterVmHardwareSataAddressSpec)`

SetSata sets Sata field to given value.

### HasSata

`func (o *VcenterVmHardwareDiskCreateSpec) HasSata() bool`

HasSata returns a boolean if a field has been set.

### GetNvme

`func (o *VcenterVmHardwareDiskCreateSpec) GetNvme() VcenterVmHardwareNvmeAddressSpec`

GetNvme returns the Nvme field if non-nil, zero value otherwise.

### GetNvmeOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetNvmeOk() (*VcenterVmHardwareNvmeAddressSpec, bool)`

GetNvmeOk returns a tuple with the Nvme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNvme

`func (o *VcenterVmHardwareDiskCreateSpec) SetNvme(v VcenterVmHardwareNvmeAddressSpec)`

SetNvme sets Nvme field to given value.

### HasNvme

`func (o *VcenterVmHardwareDiskCreateSpec) HasNvme() bool`

HasNvme returns a boolean if a field has been set.

### GetBacking

`func (o *VcenterVmHardwareDiskCreateSpec) GetBacking() VcenterVmHardwareDiskBackingSpec`

GetBacking returns the Backing field if non-nil, zero value otherwise.

### GetBackingOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetBackingOk() (*VcenterVmHardwareDiskBackingSpec, bool)`

GetBackingOk returns a tuple with the Backing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacking

`func (o *VcenterVmHardwareDiskCreateSpec) SetBacking(v VcenterVmHardwareDiskBackingSpec)`

SetBacking sets Backing field to given value.

### HasBacking

`func (o *VcenterVmHardwareDiskCreateSpec) HasBacking() bool`

HasBacking returns a boolean if a field has been set.

### GetNewVmdk

`func (o *VcenterVmHardwareDiskCreateSpec) GetNewVmdk() VcenterVmHardwareDiskVmdkCreateSpec`

GetNewVmdk returns the NewVmdk field if non-nil, zero value otherwise.

### GetNewVmdkOk

`func (o *VcenterVmHardwareDiskCreateSpec) GetNewVmdkOk() (*VcenterVmHardwareDiskVmdkCreateSpec, bool)`

GetNewVmdkOk returns a tuple with the NewVmdk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVmdk

`func (o *VcenterVmHardwareDiskCreateSpec) SetNewVmdk(v VcenterVmHardwareDiskVmdkCreateSpec)`

SetNewVmdk sets NewVmdk field to given value.

### HasNewVmdk

`func (o *VcenterVmHardwareDiskCreateSpec) HasNewVmdk() bool`

HasNewVmdk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


