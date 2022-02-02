# VcenterResourcePoolSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePool** | **string** | Identifier of the resource pool. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | 
**Name** | **string** | Name of the resource pool. | 

## Methods

### NewVcenterResourcePoolSummary

`func NewVcenterResourcePoolSummary(resourcePool string, name string, ) *VcenterResourcePoolSummary`

NewVcenterResourcePoolSummary instantiates a new VcenterResourcePoolSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolSummaryWithDefaults

`func NewVcenterResourcePoolSummaryWithDefaults() *VcenterResourcePoolSummary`

NewVcenterResourcePoolSummaryWithDefaults instantiates a new VcenterResourcePoolSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePool

`func (o *VcenterResourcePoolSummary) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterResourcePoolSummary) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterResourcePoolSummary) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.


### GetName

`func (o *VcenterResourcePoolSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterResourcePoolSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterResourcePoolSummary) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


