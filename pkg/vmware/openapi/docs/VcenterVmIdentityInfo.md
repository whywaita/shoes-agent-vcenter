# VcenterVmIdentityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Virtual machine name. | 
**BiosUuid** | **string** | 128-bit SMBIOS UUID of a virtual machine represented as a hexadecimal string in \&quot;12345678-abcd-1234-cdef-123456789abc\&quot; format. | 
**InstanceUuid** | **string** | VirtualCenter-specific 128-bit UUID of a virtual machine, represented as a hexademical string. This identifier is used by VirtualCenter to uniquely identify all virtual machine instances, including those that may share the same SMBIOS UUID. | 

## Methods

### NewVcenterVmIdentityInfo

`func NewVcenterVmIdentityInfo(name string, biosUuid string, instanceUuid string, ) *VcenterVmIdentityInfo`

NewVcenterVmIdentityInfo instantiates a new VcenterVmIdentityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmIdentityInfoWithDefaults

`func NewVcenterVmIdentityInfoWithDefaults() *VcenterVmIdentityInfo`

NewVcenterVmIdentityInfoWithDefaults instantiates a new VcenterVmIdentityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmIdentityInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmIdentityInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmIdentityInfo) SetName(v string)`

SetName sets Name field to given value.


### GetBiosUuid

`func (o *VcenterVmIdentityInfo) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *VcenterVmIdentityInfo) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *VcenterVmIdentityInfo) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.


### GetInstanceUuid

`func (o *VcenterVmIdentityInfo) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *VcenterVmIdentityInfo) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *VcenterVmIdentityInfo) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


