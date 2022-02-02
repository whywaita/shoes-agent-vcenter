# VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The identifier of the Supervisor Service version. This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the &#39;-&#39; and &#39;.&#39; characters allowed anywhere except the first or last character. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.supervisor_services.Version. | 
**DisplayName** | **string** | A human readable name of the Supervisor Service version. | 
**Description** | Pointer to **string** | A human readable description of the Supervisor Service version. If unset, the description for the service version will be empty. | [optional] 
**Content** | **string** | Inline content that contains all service definition of the version, which shall be base64 encoded. The service definition here doesn&#39;t follow the vSphere application service format. | 
**TrustedProvider** | Pointer to **bool** | Whether or not the Supervisor Service version is from a trusted provider, this field must be set to false if the service version is not from a trusted provider. If it is set to be true, but the Versions.CustomCreateSpec.content is not signed or the signature is invalid, an InvalidArgument will be thrown. If unset, the default value is true. In this case, the Versions.CustomCreateSpec.content must be signed and will be verified. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec

`func NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec(version string, displayName string, content string, ) *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetContent(v string)`

SetContent sets Content field to given value.


### GetTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetTrustedProvider() bool`

GetTrustedProvider returns the TrustedProvider field if non-nil, zero value otherwise.

### GetTrustedProviderOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) GetTrustedProviderOk() (*bool, bool)`

GetTrustedProviderOk returns a tuple with the TrustedProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) SetTrustedProvider(v bool)`

SetTrustedProvider sets TrustedProvider field to given value.

### HasTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsCustomCreateSpec) HasTrustedProvider() bool`

HasTrustedProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


