# VcenterVmGuestFilesystemFilesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmGuestFilesystemFilesType**](VcenterVmGuestFilesystemFilesType.md) |  | 
**Size** | **int64** | The file size in bytes. | 
**Attributes** | [**VcenterVmGuestFilesystemFilesFileAttributesInfo**](VcenterVmGuestFilesystemFilesFileAttributesInfo.md) |  | 

## Methods

### NewVcenterVmGuestFilesystemFilesInfo

`func NewVcenterVmGuestFilesystemFilesInfo(type_ VcenterVmGuestFilesystemFilesType, size int64, attributes VcenterVmGuestFilesystemFilesFileAttributesInfo, ) *VcenterVmGuestFilesystemFilesInfo`

NewVcenterVmGuestFilesystemFilesInfo instantiates a new VcenterVmGuestFilesystemFilesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesInfoWithDefaults

`func NewVcenterVmGuestFilesystemFilesInfoWithDefaults() *VcenterVmGuestFilesystemFilesInfo`

NewVcenterVmGuestFilesystemFilesInfoWithDefaults instantiates a new VcenterVmGuestFilesystemFilesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmGuestFilesystemFilesInfo) GetType() VcenterVmGuestFilesystemFilesType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmGuestFilesystemFilesInfo) GetTypeOk() (*VcenterVmGuestFilesystemFilesType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmGuestFilesystemFilesInfo) SetType(v VcenterVmGuestFilesystemFilesType)`

SetType sets Type field to given value.


### GetSize

`func (o *VcenterVmGuestFilesystemFilesInfo) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VcenterVmGuestFilesystemFilesInfo) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VcenterVmGuestFilesystemFilesInfo) SetSize(v int64)`

SetSize sets Size field to given value.


### GetAttributes

`func (o *VcenterVmGuestFilesystemFilesInfo) GetAttributes() VcenterVmGuestFilesystemFilesFileAttributesInfo`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *VcenterVmGuestFilesystemFilesInfo) GetAttributesOk() (*VcenterVmGuestFilesystemFilesFileAttributesInfo, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *VcenterVmGuestFilesystemFilesInfo) SetAttributes(v VcenterVmGuestFilesystemFilesFileAttributesInfo)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


