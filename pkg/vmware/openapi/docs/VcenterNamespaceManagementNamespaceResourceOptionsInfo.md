# VcenterNamespaceManagementNamespaceResourceOptionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateResourceQuotaType** | **string** | Identifier of the structure used to set resource quotas on the namespace. See vcenter.namespaces.Instances#create and vcenter.namespaces.Instances#set. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vapi.structure. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vapi.structure. | 
**UpdateResourceQuotaType** | **string** | Identifier of the structure used to update resource quotas on the namespace. See vcenter.namespaces.Instances#update. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vapi.structure. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vapi.structure. | 

## Methods

### NewVcenterNamespaceManagementNamespaceResourceOptionsInfo

`func NewVcenterNamespaceManagementNamespaceResourceOptionsInfo(createResourceQuotaType string, updateResourceQuotaType string, ) *VcenterNamespaceManagementNamespaceResourceOptionsInfo`

NewVcenterNamespaceManagementNamespaceResourceOptionsInfo instantiates a new VcenterNamespaceManagementNamespaceResourceOptionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementNamespaceResourceOptionsInfoWithDefaults

`func NewVcenterNamespaceManagementNamespaceResourceOptionsInfoWithDefaults() *VcenterNamespaceManagementNamespaceResourceOptionsInfo`

NewVcenterNamespaceManagementNamespaceResourceOptionsInfoWithDefaults instantiates a new VcenterNamespaceManagementNamespaceResourceOptionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateResourceQuotaType

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) GetCreateResourceQuotaType() string`

GetCreateResourceQuotaType returns the CreateResourceQuotaType field if non-nil, zero value otherwise.

### GetCreateResourceQuotaTypeOk

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) GetCreateResourceQuotaTypeOk() (*string, bool)`

GetCreateResourceQuotaTypeOk returns a tuple with the CreateResourceQuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateResourceQuotaType

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) SetCreateResourceQuotaType(v string)`

SetCreateResourceQuotaType sets CreateResourceQuotaType field to given value.


### GetUpdateResourceQuotaType

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) GetUpdateResourceQuotaType() string`

GetUpdateResourceQuotaType returns the UpdateResourceQuotaType field if non-nil, zero value otherwise.

### GetUpdateResourceQuotaTypeOk

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) GetUpdateResourceQuotaTypeOk() (*string, bool)`

GetUpdateResourceQuotaTypeOk returns a tuple with the UpdateResourceQuotaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateResourceQuotaType

`func (o *VcenterNamespaceManagementNamespaceResourceOptionsInfo) SetUpdateResourceQuotaType(v string)`

SetUpdateResourceQuotaType sets UpdateResourceQuotaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


