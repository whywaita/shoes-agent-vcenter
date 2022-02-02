# VcenterNamespaceManagementSoftwareClustersMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | [**VcenterNamespaceManagementSoftwareClustersMessageSeverity**](VcenterNamespaceManagementSoftwareClustersMessageSeverity.md) |  | 
**Details** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersMessage

`func NewVcenterNamespaceManagementSoftwareClustersMessage(severity VcenterNamespaceManagementSoftwareClustersMessageSeverity, details VapiStdLocalizableMessage, ) *VcenterNamespaceManagementSoftwareClustersMessage`

NewVcenterNamespaceManagementSoftwareClustersMessage instantiates a new VcenterNamespaceManagementSoftwareClustersMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersMessageWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersMessageWithDefaults() *VcenterNamespaceManagementSoftwareClustersMessage`

NewVcenterNamespaceManagementSoftwareClustersMessageWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetSeverity() VcenterNamespaceManagementSoftwareClustersMessageSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetSeverityOk() (*VcenterNamespaceManagementSoftwareClustersMessageSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) SetSeverity(v VcenterNamespaceManagementSoftwareClustersMessageSeverity)`

SetSeverity sets Severity field to given value.


### GetDetails

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetDetails() VapiStdLocalizableMessage`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetDetailsOk() (*VapiStdLocalizableMessage, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VcenterNamespaceManagementSoftwareClustersMessage) SetDetails(v VapiStdLocalizableMessage)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


