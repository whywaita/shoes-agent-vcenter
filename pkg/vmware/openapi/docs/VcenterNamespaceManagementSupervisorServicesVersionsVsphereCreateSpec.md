# VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Inline content that contains all service definition of the version in vSphere application service format, which shall be base64 encoded. The service definition here follows the vSphere application service format. | 
**TrustedProvider** | Pointer to **bool** | Whether or not the Supervisor Service version is from a trusted provider, this field must be set to false if the service version is not from a trusted provider. If it is set to be true, but the Versions.VsphereCreateSpec.content is not signed or the signature is invalid, an InvalidArgument will be thrown. If unset, the default value is true. In this case, the Versions.VsphereCreateSpec.content must be signed and will be verified. | [optional] 
**AcceptEULA** | Pointer to **bool** | Whether or not the End User License Agreement (EULA) that is specified in the Versions.VsphereCreateSpec.content is accepted. If a EULA is specified, this field must be set to be true so that the Supervisor Service version can be created. If unset, the default value is false. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec

`func NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec(content string, ) *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec`

NewVcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) SetContent(v string)`

SetContent sets Content field to given value.


### GetTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetTrustedProvider() bool`

GetTrustedProvider returns the TrustedProvider field if non-nil, zero value otherwise.

### GetTrustedProviderOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetTrustedProviderOk() (*bool, bool)`

GetTrustedProviderOk returns a tuple with the TrustedProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) SetTrustedProvider(v bool)`

SetTrustedProvider sets TrustedProvider field to given value.

### HasTrustedProvider

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) HasTrustedProvider() bool`

HasTrustedProvider returns a boolean if a field has been set.

### GetAcceptEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetAcceptEULA() bool`

GetAcceptEULA returns the AcceptEULA field if non-nil, zero value otherwise.

### GetAcceptEULAOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) GetAcceptEULAOk() (*bool, bool)`

GetAcceptEULAOk returns a tuple with the AcceptEULA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) SetAcceptEULA(v bool)`

SetAcceptEULA sets AcceptEULA field to given value.

### HasAcceptEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsVsphereCreateSpec) HasAcceptEULA() bool`

HasAcceptEULA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


