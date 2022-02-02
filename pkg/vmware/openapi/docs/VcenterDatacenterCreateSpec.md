# VcenterDatacenterCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the datacenter to be created. | 
**Folder** | Pointer to **string** | Datacenter folder in which the new datacenter should be created. This field is currently required. In the future, if this field is unset, the system will attempt to choose a suitable folder for the datacenter; if a folder cannot be chosen, the datacenter creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 

## Methods

### NewVcenterDatacenterCreateSpec

`func NewVcenterDatacenterCreateSpec(name string, ) *VcenterDatacenterCreateSpec`

NewVcenterDatacenterCreateSpec instantiates a new VcenterDatacenterCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatacenterCreateSpecWithDefaults

`func NewVcenterDatacenterCreateSpecWithDefaults() *VcenterDatacenterCreateSpec`

NewVcenterDatacenterCreateSpecWithDefaults instantiates a new VcenterDatacenterCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterDatacenterCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterDatacenterCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterDatacenterCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetFolder

`func (o *VcenterDatacenterCreateSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterDatacenterCreateSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterDatacenterCreateSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterDatacenterCreateSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


