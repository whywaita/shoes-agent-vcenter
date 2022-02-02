# VcenterDatacenterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the datacenter. | 
**DatastoreFolder** | **string** | The root datastore folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | 
**HostFolder** | **string** | The root host and cluster folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | 
**NetworkFolder** | **string** | The root network folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | 
**VmFolder** | **string** | The root virtual machine folder associated with the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | 

## Methods

### NewVcenterDatacenterInfo

`func NewVcenterDatacenterInfo(name string, datastoreFolder string, hostFolder string, networkFolder string, vmFolder string, ) *VcenterDatacenterInfo`

NewVcenterDatacenterInfo instantiates a new VcenterDatacenterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatacenterInfoWithDefaults

`func NewVcenterDatacenterInfoWithDefaults() *VcenterDatacenterInfo`

NewVcenterDatacenterInfoWithDefaults instantiates a new VcenterDatacenterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterDatacenterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterDatacenterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterDatacenterInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDatastoreFolder

`func (o *VcenterDatacenterInfo) GetDatastoreFolder() string`

GetDatastoreFolder returns the DatastoreFolder field if non-nil, zero value otherwise.

### GetDatastoreFolderOk

`func (o *VcenterDatacenterInfo) GetDatastoreFolderOk() (*string, bool)`

GetDatastoreFolderOk returns a tuple with the DatastoreFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreFolder

`func (o *VcenterDatacenterInfo) SetDatastoreFolder(v string)`

SetDatastoreFolder sets DatastoreFolder field to given value.


### GetHostFolder

`func (o *VcenterDatacenterInfo) GetHostFolder() string`

GetHostFolder returns the HostFolder field if non-nil, zero value otherwise.

### GetHostFolderOk

`func (o *VcenterDatacenterInfo) GetHostFolderOk() (*string, bool)`

GetHostFolderOk returns a tuple with the HostFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostFolder

`func (o *VcenterDatacenterInfo) SetHostFolder(v string)`

SetHostFolder sets HostFolder field to given value.


### GetNetworkFolder

`func (o *VcenterDatacenterInfo) GetNetworkFolder() string`

GetNetworkFolder returns the NetworkFolder field if non-nil, zero value otherwise.

### GetNetworkFolderOk

`func (o *VcenterDatacenterInfo) GetNetworkFolderOk() (*string, bool)`

GetNetworkFolderOk returns a tuple with the NetworkFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFolder

`func (o *VcenterDatacenterInfo) SetNetworkFolder(v string)`

SetNetworkFolder sets NetworkFolder field to given value.


### GetVmFolder

`func (o *VcenterDatacenterInfo) GetVmFolder() string`

GetVmFolder returns the VmFolder field if non-nil, zero value otherwise.

### GetVmFolderOk

`func (o *VcenterDatacenterInfo) GetVmFolderOk() (*string, bool)`

GetVmFolderOk returns a tuple with the VmFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolder

`func (o *VcenterDatacenterInfo) SetVmFolder(v string)`

SetVmFolder sets VmFolder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


