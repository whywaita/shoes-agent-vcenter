# VcenterDatastoreInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the datastore. | 
**Type** | [**VcenterDatastoreType**](VcenterDatastoreType.md) |  | 
**Accessible** | **bool** | Whether or not this datastore is accessible. | 
**FreeSpace** | Pointer to **int64** | Available space of this datastore, in bytes.   The server periodically updates this value.  This field will be unset if the available space of this datastore is not known. | [optional] 
**MultipleHostAccess** | **bool** | Whether or not more than one host in the datacenter has been configured with access to the datastore. | 
**ThinProvisioningSupported** | **bool** | Whether or not the datastore supports thin provisioning on a per file basis. When thin provisioning is used, backing storage is lazily allocated. | 

## Methods

### NewVcenterDatastoreInfo

`func NewVcenterDatastoreInfo(name string, type_ VcenterDatastoreType, accessible bool, multipleHostAccess bool, thinProvisioningSupported bool, ) *VcenterDatastoreInfo`

NewVcenterDatastoreInfo instantiates a new VcenterDatastoreInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatastoreInfoWithDefaults

`func NewVcenterDatastoreInfoWithDefaults() *VcenterDatastoreInfo`

NewVcenterDatastoreInfoWithDefaults instantiates a new VcenterDatastoreInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterDatastoreInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterDatastoreInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterDatastoreInfo) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VcenterDatastoreInfo) GetType() VcenterDatastoreType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterDatastoreInfo) GetTypeOk() (*VcenterDatastoreType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterDatastoreInfo) SetType(v VcenterDatastoreType)`

SetType sets Type field to given value.


### GetAccessible

`func (o *VcenterDatastoreInfo) GetAccessible() bool`

GetAccessible returns the Accessible field if non-nil, zero value otherwise.

### GetAccessibleOk

`func (o *VcenterDatastoreInfo) GetAccessibleOk() (*bool, bool)`

GetAccessibleOk returns a tuple with the Accessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessible

`func (o *VcenterDatastoreInfo) SetAccessible(v bool)`

SetAccessible sets Accessible field to given value.


### GetFreeSpace

`func (o *VcenterDatastoreInfo) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *VcenterDatastoreInfo) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *VcenterDatastoreInfo) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *VcenterDatastoreInfo) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetMultipleHostAccess

`func (o *VcenterDatastoreInfo) GetMultipleHostAccess() bool`

GetMultipleHostAccess returns the MultipleHostAccess field if non-nil, zero value otherwise.

### GetMultipleHostAccessOk

`func (o *VcenterDatastoreInfo) GetMultipleHostAccessOk() (*bool, bool)`

GetMultipleHostAccessOk returns a tuple with the MultipleHostAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleHostAccess

`func (o *VcenterDatastoreInfo) SetMultipleHostAccess(v bool)`

SetMultipleHostAccess sets MultipleHostAccess field to given value.


### GetThinProvisioningSupported

`func (o *VcenterDatastoreInfo) GetThinProvisioningSupported() bool`

GetThinProvisioningSupported returns the ThinProvisioningSupported field if non-nil, zero value otherwise.

### GetThinProvisioningSupportedOk

`func (o *VcenterDatastoreInfo) GetThinProvisioningSupportedOk() (*bool, bool)`

GetThinProvisioningSupportedOk returns a tuple with the ThinProvisioningSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinProvisioningSupported

`func (o *VcenterDatastoreInfo) SetThinProvisioningSupported(v bool)`

SetThinProvisioningSupported sets ThinProvisioningSupported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


