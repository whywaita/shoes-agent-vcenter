# VcenterNamespaceManagementSupervisorServicesContentCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content of a Supervisor Service version, which shall be base64 encoded. If unset, the content shall be provided separately. In the current release, this field is required, otherwise InvalidArgument will be thrown. | [optional] 

## Methods

### NewVcenterNamespaceManagementSupervisorServicesContentCheckSpec

`func NewVcenterNamespaceManagementSupervisorServicesContentCheckSpec() *VcenterNamespaceManagementSupervisorServicesContentCheckSpec`

NewVcenterNamespaceManagementSupervisorServicesContentCheckSpec instantiates a new VcenterNamespaceManagementSupervisorServicesContentCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupervisorServicesContentCheckSpecWithDefaults

`func NewVcenterNamespaceManagementSupervisorServicesContentCheckSpecWithDefaults() *VcenterNamespaceManagementSupervisorServicesContentCheckSpec`

NewVcenterNamespaceManagementSupervisorServicesContentCheckSpecWithDefaults instantiates a new VcenterNamespaceManagementSupervisorServicesContentCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *VcenterNamespaceManagementSupervisorServicesContentCheckSpec) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *VcenterNamespaceManagementSupervisorServicesContentCheckSpec) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *VcenterNamespaceManagementSupervisorServicesContentCheckSpec) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *VcenterNamespaceManagementSupervisorServicesContentCheckSpec) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


