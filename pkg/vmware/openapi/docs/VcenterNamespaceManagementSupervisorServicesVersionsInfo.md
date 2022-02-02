# VcenterNamespaceManagementSupervisorServicesVersionsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | The human readable name of the Supervisor Service version. | 
**Description** | Pointer to **string** | A human-readable description of the Supervisor Service version. If unset, no description is available for the Supervisor Service version. | [optional] 
**EULA** | Pointer to **string** | The End User License Agreement (EULA) associated with the Supervisor Service version. If unset, no EULA is available for the Supervisor Service version. | [optional] 
**ContentType** | [**VcenterNamespaceManagementSupervisorServicesVersionsContentType**](VcenterNamespaceManagementSupervisorServicesVersionsContentType.md) |  | 
**Content** | Pointer to **string** | Inline content that contains base64 encoded service definition for the version. | [optional] 
**TrustVerified** | **bool** | If true, the Supervisor Service version is from trusted provider and the trust is verified. | 
**State** | [**VcenterNamespaceManagementSupervisorServicesVersionsState**](VcenterNamespaceManagementSupervisorServicesVersionsState.md) |  | 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesVersionsInfo

`func NewVcenterNamespaceManagementSupervisorServicesVersionsInfo(displayName string, contentType VcenterNamespaceManagementSupervisorServicesVersionsContentType, trustVerified bool, state VcenterNamespaceManagementSupervisorServicesVersionsState, ) *VcenterNamespaceManagementSupervisorServicesVersionsInfo`

NewVcenterNamespaceManagementSupervisorServicesVersionsInfo instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesVersionsInfoWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesVersionsInfoWithDefaults() *VcenterNamespaceManagementSupervisorServicesVersionsInfo`

NewVcenterNamespaceManagementSupervisorServicesVersionsInfoWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesVersionsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetEULA() string`

GetEULA returns the EULA field if non-nil, zero value otherwise.

### GetEULAOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetEULAOk() (*string, bool)`

GetEULAOk returns a tuple with the EULA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetEULA(v string)`

SetEULA sets EULA field to given value.

### HasEULA

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) HasEULA() bool`

HasEULA returns a boolean if a field has been set.

### GetContentType

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetContentType() VcenterNamespaceManagementSupervisorServicesVersionsContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetContentTypeOk() (*VcenterNamespaceManagementSupervisorServicesVersionsContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetContentType(v VcenterNamespaceManagementSupervisorServicesVersionsContentType)`

SetContentType sets ContentType field to given value.


### GetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetTrustVerified

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetTrustVerified() bool`

GetTrustVerified returns the TrustVerified field if non-nil, zero value otherwise.

### GetTrustVerifiedOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetTrustVerifiedOk() (*bool, bool)`

GetTrustVerifiedOk returns a tuple with the TrustVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustVerified

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetTrustVerified(v bool)`

SetTrustVerified sets TrustVerified field to given value.


### GetState

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetState() VcenterNamespaceManagementSupervisorServicesVersionsState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) GetStateOk() (*VcenterNamespaceManagementSupervisorServicesVersionsState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSupervisorServicesVersionsInfo) SetState(v VcenterNamespaceManagementSupervisorServicesVersionsState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


