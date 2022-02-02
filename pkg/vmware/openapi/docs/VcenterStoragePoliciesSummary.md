# VcenterStoragePoliciesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Identifier of the storage policy. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.StoragePolicy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.StoragePolicy. | 
**Name** | **string** | Name of the storage policy. | 
**Description** | **string** | Description of the storage policy. | 

## Methods

### NewVcenterStoragePoliciesSummary

`func NewVcenterStoragePoliciesSummary(policy string, name string, description string, ) *VcenterStoragePoliciesSummary`

NewVcenterStoragePoliciesSummary instantiates a new VcenterStoragePoliciesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesSummaryWithDefaults

`func NewVcenterStoragePoliciesSummaryWithDefaults() *VcenterStoragePoliciesSummary`

NewVcenterStoragePoliciesSummaryWithDefaults instantiates a new VcenterStoragePoliciesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterStoragePoliciesSummary) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterStoragePoliciesSummary) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterStoragePoliciesSummary) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetName

`func (o *VcenterStoragePoliciesSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterStoragePoliciesSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterStoragePoliciesSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VcenterStoragePoliciesSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterStoragePoliciesSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterStoragePoliciesSummary) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


