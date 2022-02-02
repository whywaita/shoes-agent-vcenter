# VcenterNamespaceManagementClusterSizeInfoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumSupportedPods** | **int64** | The maximum number of supported pods. | 
**NumSupportedServices** | **int64** | The maximum number of supported services. | 
**DefaultServiceCidr** | [**VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) |  | 
**DefaultPodCidr** | [**VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) |  | 
**MasterVmInfo** | [**VcenterNamespaceManagementClusterSizeInfoVmInfo**](VcenterNamespaceManagementClusterSizeInfoVmInfo.md) |  | 
**WorkerVmInfo** | Pointer to [**VcenterNamespaceManagementClusterSizeInfoVmInfo**](VcenterNamespaceManagementClusterSizeInfoVmInfo.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementClusterSizeInfoInfo

`func NewVcenterNamespaceManagementClusterSizeInfoInfo(numSupportedPods int64, numSupportedServices int64, defaultServiceCidr VcenterNamespaceManagementIpv4Cidr, defaultPodCidr VcenterNamespaceManagementIpv4Cidr, masterVmInfo VcenterNamespaceManagementClusterSizeInfoVmInfo, ) *VcenterNamespaceManagementClusterSizeInfoInfo`

NewVcenterNamespaceManagementClusterSizeInfoInfo instantiates a new VcenterNamespaceManagementClusterSizeInfoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClusterSizeInfoInfoWithDefaults

`func NewVcenterNamespaceManagementClusterSizeInfoInfoWithDefaults() *VcenterNamespaceManagementClusterSizeInfoInfo`

NewVcenterNamespaceManagementClusterSizeInfoInfoWithDefaults instantiates a new VcenterNamespaceManagementClusterSizeInfoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumSupportedPods

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetNumSupportedPods() int64`

GetNumSupportedPods returns the NumSupportedPods field if non-nil, zero value otherwise.

### GetNumSupportedPodsOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetNumSupportedPodsOk() (*int64, bool)`

GetNumSupportedPodsOk returns a tuple with the NumSupportedPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSupportedPods

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetNumSupportedPods(v int64)`

SetNumSupportedPods sets NumSupportedPods field to given value.


### GetNumSupportedServices

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetNumSupportedServices() int64`

GetNumSupportedServices returns the NumSupportedServices field if non-nil, zero value otherwise.

### GetNumSupportedServicesOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetNumSupportedServicesOk() (*int64, bool)`

GetNumSupportedServicesOk returns a tuple with the NumSupportedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSupportedServices

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetNumSupportedServices(v int64)`

SetNumSupportedServices sets NumSupportedServices field to given value.


### GetDefaultServiceCidr

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetDefaultServiceCidr() VcenterNamespaceManagementIpv4Cidr`

GetDefaultServiceCidr returns the DefaultServiceCidr field if non-nil, zero value otherwise.

### GetDefaultServiceCidrOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetDefaultServiceCidrOk() (*VcenterNamespaceManagementIpv4Cidr, bool)`

GetDefaultServiceCidrOk returns a tuple with the DefaultServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultServiceCidr

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetDefaultServiceCidr(v VcenterNamespaceManagementIpv4Cidr)`

SetDefaultServiceCidr sets DefaultServiceCidr field to given value.


### GetDefaultPodCidr

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetDefaultPodCidr() VcenterNamespaceManagementIpv4Cidr`

GetDefaultPodCidr returns the DefaultPodCidr field if non-nil, zero value otherwise.

### GetDefaultPodCidrOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetDefaultPodCidrOk() (*VcenterNamespaceManagementIpv4Cidr, bool)`

GetDefaultPodCidrOk returns a tuple with the DefaultPodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPodCidr

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetDefaultPodCidr(v VcenterNamespaceManagementIpv4Cidr)`

SetDefaultPodCidr sets DefaultPodCidr field to given value.


### GetMasterVmInfo

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetMasterVmInfo() VcenterNamespaceManagementClusterSizeInfoVmInfo`

GetMasterVmInfo returns the MasterVmInfo field if non-nil, zero value otherwise.

### GetMasterVmInfoOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetMasterVmInfoOk() (*VcenterNamespaceManagementClusterSizeInfoVmInfo, bool)`

GetMasterVmInfoOk returns a tuple with the MasterVmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterVmInfo

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetMasterVmInfo(v VcenterNamespaceManagementClusterSizeInfoVmInfo)`

SetMasterVmInfo sets MasterVmInfo field to given value.


### GetWorkerVmInfo

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetWorkerVmInfo() VcenterNamespaceManagementClusterSizeInfoVmInfo`

GetWorkerVmInfo returns the WorkerVmInfo field if non-nil, zero value otherwise.

### GetWorkerVmInfoOk

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) GetWorkerVmInfoOk() (*VcenterNamespaceManagementClusterSizeInfoVmInfo, bool)`

GetWorkerVmInfoOk returns a tuple with the WorkerVmInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerVmInfo

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) SetWorkerVmInfo(v VcenterNamespaceManagementClusterSizeInfoVmInfo)`

SetWorkerVmInfo sets WorkerVmInfo field to given value.

### HasWorkerVmInfo

`func (o *VcenterNamespaceManagementClusterSizeInfoInfo) HasWorkerVmInfo() bool`

HasWorkerVmInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


