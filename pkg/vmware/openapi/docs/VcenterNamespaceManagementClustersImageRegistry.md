# VcenterNamespaceManagementClustersImageRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | IP address or the hostname of container image registry. | 
**Port** | Pointer to **int64** | Port number of the container image registry. If unset, defaults to 443. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersImageRegistry

`func NewVcenterNamespaceManagementClustersImageRegistry(hostname string, ) *VcenterNamespaceManagementClustersImageRegistry`

NewVcenterNamespaceManagementClustersImageRegistry instantiates a new VcenterNamespaceManagementClustersImageRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersImageRegistryWithDefaults

`func NewVcenterNamespaceManagementClustersImageRegistryWithDefaults() *VcenterNamespaceManagementClustersImageRegistry`

NewVcenterNamespaceManagementClustersImageRegistryWithDefaults instantiates a new VcenterNamespaceManagementClustersImageRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterNamespaceManagementClustersImageRegistry) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterNamespaceManagementClustersImageRegistry) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterNamespaceManagementClustersImageRegistry) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *VcenterNamespaceManagementClustersImageRegistry) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterNamespaceManagementClustersImageRegistry) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterNamespaceManagementClustersImageRegistry) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VcenterNamespaceManagementClustersImageRegistry) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


