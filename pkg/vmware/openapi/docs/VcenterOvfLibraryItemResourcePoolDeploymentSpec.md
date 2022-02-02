# VcenterOvfLibraryItemResourcePoolDeploymentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name assigned to the deployed target virtual machine or virtual appliance. | [optional] 
**Annotation** | Pointer to **string** | Annotation assigned to the deployed target virtual machine or virtual appliance. | [optional] 
**AcceptAllEULA** | **bool** | Whether to accept all End User License Agreements. See {@link OvfSummary#eulas}. | 
**NetworkMappings** | Pointer to [**[]VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings**](VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings.md) | Specification of the target network to use for sections of type ovf:NetworkSection in the OVF descriptor. The key in the {@term map} is the section identifier of the ovf:NetworkSection section in the OVF descriptor and the value is the target network to be used for deployment. | [optional] 
**StorageMappings** | Pointer to [**[]VcenterOvfLibraryItemResourcePoolDeploymentSpecStorageMappings**](VcenterOvfLibraryItemResourcePoolDeploymentSpecStorageMappings.md) | Specification of the target storage to use for sections of type vmw:StorageGroupSection in the OVF descriptor. The key in the {@term map} is the section identifier of the ovf:StorageGroupSection section in the OVF descriptor and the value is the target storage specification to be used for deployment. See {@link StorageGroupMapping}. | [optional] 
**StorageProvisioning** | Pointer to [**VcenterOvfDiskProvisioningType**](VcenterOvfDiskProvisioningType.md) |  | [optional] 
**StorageProfileId** | Pointer to **string** | Default storage profile to use for all sections of type vmw:StorageSection in the OVF descriptor. | [optional] 
**Locale** | Pointer to **string** | The locale to use for parsing the OVF descriptor. | [optional] 
**Flags** | Pointer to **[]string** | Flags to be use for deployment. The supported flag values can be obtained using {@link ImportFlag#list}. | [optional] 
**AdditionalParameters** | Pointer to **[]map[string]interface{}** | Additional OVF parameters that may be needed for the deployment. Additional OVF parameters may be required by the OVF descriptor of the OVF package in the library item. Examples of OVF parameters that can be specified through this {@term field} include, but are not limited to: &lt;ul&gt; &lt;li&gt;{@link DeploymentOptionParams}&lt;/li&gt; &lt;li&gt;{@link ExtraConfigParams}&lt;/li&gt; &lt;li&gt;{@link IpAllocationParams}&lt;/li&gt; &lt;li&gt;{@link PropertyParams}&lt;/li&gt; &lt;li&gt;{@link ScaleOutParams}&lt;/li&gt; &lt;li&gt;{@link VcenterExtensionParams}&lt;/li&gt; &lt;/ul&gt; | [optional] 
**DefaultDatastoreId** | Pointer to **string** | Default datastore to use for all sections of type vmw:StorageSection in the OVF descriptor. | [optional] 

## Methods

### NewVcenterOvfLibraryItemResourcePoolDeploymentSpec

`func NewVcenterOvfLibraryItemResourcePoolDeploymentSpec(acceptAllEULA bool, ) *VcenterOvfLibraryItemResourcePoolDeploymentSpec`

NewVcenterOvfLibraryItemResourcePoolDeploymentSpec instantiates a new VcenterOvfLibraryItemResourcePoolDeploymentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemResourcePoolDeploymentSpecWithDefaults

`func NewVcenterOvfLibraryItemResourcePoolDeploymentSpecWithDefaults() *VcenterOvfLibraryItemResourcePoolDeploymentSpec`

NewVcenterOvfLibraryItemResourcePoolDeploymentSpecWithDefaults instantiates a new VcenterOvfLibraryItemResourcePoolDeploymentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAnnotation

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAnnotation() string`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAnnotationOk() (*string, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetAnnotation(v string)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetAcceptAllEULA

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAcceptAllEULA() bool`

GetAcceptAllEULA returns the AcceptAllEULA field if non-nil, zero value otherwise.

### GetAcceptAllEULAOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAcceptAllEULAOk() (*bool, bool)`

GetAcceptAllEULAOk returns a tuple with the AcceptAllEULA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptAllEULA

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetAcceptAllEULA(v bool)`

SetAcceptAllEULA sets AcceptAllEULA field to given value.


### GetNetworkMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetNetworkMappings() []VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings`

GetNetworkMappings returns the NetworkMappings field if non-nil, zero value otherwise.

### GetNetworkMappingsOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetNetworkMappingsOk() (*[]VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings, bool)`

GetNetworkMappingsOk returns a tuple with the NetworkMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetNetworkMappings(v []VcenterOvfLibraryItemResourcePoolDeploymentSpecNetworkMappings)`

SetNetworkMappings sets NetworkMappings field to given value.

### HasNetworkMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasNetworkMappings() bool`

HasNetworkMappings returns a boolean if a field has been set.

### GetStorageMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageMappings() []VcenterOvfLibraryItemResourcePoolDeploymentSpecStorageMappings`

GetStorageMappings returns the StorageMappings field if non-nil, zero value otherwise.

### GetStorageMappingsOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageMappingsOk() (*[]VcenterOvfLibraryItemResourcePoolDeploymentSpecStorageMappings, bool)`

GetStorageMappingsOk returns a tuple with the StorageMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetStorageMappings(v []VcenterOvfLibraryItemResourcePoolDeploymentSpecStorageMappings)`

SetStorageMappings sets StorageMappings field to given value.

### HasStorageMappings

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasStorageMappings() bool`

HasStorageMappings returns a boolean if a field has been set.

### GetStorageProvisioning

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageProvisioning() VcenterOvfDiskProvisioningType`

GetStorageProvisioning returns the StorageProvisioning field if non-nil, zero value otherwise.

### GetStorageProvisioningOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageProvisioningOk() (*VcenterOvfDiskProvisioningType, bool)`

GetStorageProvisioningOk returns a tuple with the StorageProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvisioning

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetStorageProvisioning(v VcenterOvfDiskProvisioningType)`

SetStorageProvisioning sets StorageProvisioning field to given value.

### HasStorageProvisioning

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasStorageProvisioning() bool`

HasStorageProvisioning returns a boolean if a field has been set.

### GetStorageProfileId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageProfileId() string`

GetStorageProfileId returns the StorageProfileId field if non-nil, zero value otherwise.

### GetStorageProfileIdOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetStorageProfileIdOk() (*string, bool)`

GetStorageProfileIdOk returns a tuple with the StorageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetStorageProfileId(v string)`

SetStorageProfileId sets StorageProfileId field to given value.

### HasStorageProfileId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasStorageProfileId() bool`

HasStorageProfileId returns a boolean if a field has been set.

### GetLocale

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetFlags

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAdditionalParameters() []map[string]interface{}`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetAdditionalParametersOk() (*[]map[string]interface{}, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetAdditionalParameters(v []map[string]interface{})`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.

### GetDefaultDatastoreId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetDefaultDatastoreId() string`

GetDefaultDatastoreId returns the DefaultDatastoreId field if non-nil, zero value otherwise.

### GetDefaultDatastoreIdOk

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) GetDefaultDatastoreIdOk() (*string, bool)`

GetDefaultDatastoreIdOk returns a tuple with the DefaultDatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) SetDefaultDatastoreId(v string)`

SetDefaultDatastoreId sets DefaultDatastoreId field to given value.

### HasDefaultDatastoreId

`func (o *VcenterOvfLibraryItemResourcePoolDeploymentSpec) HasDefaultDatastoreId() bool`

HasDefaultDatastoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


