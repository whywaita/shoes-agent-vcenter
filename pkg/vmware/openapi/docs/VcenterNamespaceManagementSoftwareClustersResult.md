# VcenterNamespaceManagementSoftwareClustersResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Res** | [**VcenterNamespaceManagementSoftwareClustersResultRes**](VcenterNamespaceManagementSoftwareClustersResultRes.md) |  | 
**Exception** | Pointer to **string** | Exception when cluster pre-check failed during upgrade invocation. This field is optional and it is only relevant when the value of Clusters.Result.res is REJECTED. | [optional] 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersResult

`func NewVcenterNamespaceManagementSoftwareClustersResult(res VcenterNamespaceManagementSoftwareClustersResultRes, ) *VcenterNamespaceManagementSoftwareClustersResult`

NewVcenterNamespaceManagementSoftwareClustersResult instantiates a new VcenterNamespaceManagementSoftwareClustersResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersResultWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersResultWithDefaults() *VcenterNamespaceManagementSoftwareClustersResult`

NewVcenterNamespaceManagementSoftwareClustersResultWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRes

`func (o *VcenterNamespaceManagementSoftwareClustersResult) GetRes() VcenterNamespaceManagementSoftwareClustersResultRes`

GetRes returns the Res field if non-nil, zero value otherwise.

### GetResOk

`func (o *VcenterNamespaceManagementSoftwareClustersResult) GetResOk() (*VcenterNamespaceManagementSoftwareClustersResultRes, bool)`

GetResOk returns a tuple with the Res field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRes

`func (o *VcenterNamespaceManagementSoftwareClustersResult) SetRes(v VcenterNamespaceManagementSoftwareClustersResultRes)`

SetRes sets Res field to given value.


### GetException

`func (o *VcenterNamespaceManagementSoftwareClustersResult) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *VcenterNamespaceManagementSoftwareClustersResult) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *VcenterNamespaceManagementSoftwareClustersResult) SetException(v string)`

SetException sets Exception field to given value.

### HasException

`func (o *VcenterNamespaceManagementSoftwareClustersResult) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


