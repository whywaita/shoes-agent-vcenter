# VcenterNamespacesInstancesMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Severity** | [**VcenterNamespacesInstancesMessageMessageSeverity**](VcenterNamespacesInstancesMessageMessageSeverity.md) |  | 
**Details** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterNamespacesInstancesMessage

`func NewVcenterNamespacesInstancesMessage(severity VcenterNamespacesInstancesMessageMessageSeverity, ) *VcenterNamespacesInstancesMessage`

NewVcenterNamespacesInstancesMessage instantiates a new VcenterNamespacesInstancesMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesMessageWithDefaults

`func NewVcenterNamespacesInstancesMessageWithDefaults() *VcenterNamespacesInstancesMessage`

NewVcenterNamespacesInstancesMessageWithDefaults instantiates a new VcenterNamespacesInstancesMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeverity

`func (o *VcenterNamespacesInstancesMessage) GetSeverity() VcenterNamespacesInstancesMessageMessageSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *VcenterNamespacesInstancesMessage) GetSeverityOk() (*VcenterNamespacesInstancesMessageMessageSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *VcenterNamespacesInstancesMessage) SetSeverity(v VcenterNamespacesInstancesMessageMessageSeverity)`

SetSeverity sets Severity field to given value.


### GetDetails

`func (o *VcenterNamespacesInstancesMessage) GetDetails() VapiStdLocalizableMessage`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VcenterNamespacesInstancesMessage) GetDetailsOk() (*VapiStdLocalizableMessage, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VcenterNamespacesInstancesMessage) SetDetails(v VapiStdLocalizableMessage)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VcenterNamespacesInstancesMessage) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


