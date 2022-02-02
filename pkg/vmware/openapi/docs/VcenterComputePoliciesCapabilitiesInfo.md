# VcenterComputePoliciesCapabilitiesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the capability. | 
**Description** | **string** | Description of the capability. | 
**CreateSpecType** | **string** | Identifier of the {@term structure} used to create a policy based on this capability. See {@link vcenter.compute.Policies#create}. | 
**InfoType** | **string** | Identifier of the {@term structure} returned when retrieving information about a policy based on this capability. See {@link vcenter.compute.Policies#get}. | 

## Methods

### NewVcenterComputePoliciesCapabilitiesInfo

`func NewVcenterComputePoliciesCapabilitiesInfo(name string, description string, createSpecType string, infoType string, ) *VcenterComputePoliciesCapabilitiesInfo`

NewVcenterComputePoliciesCapabilitiesInfo instantiates a new VcenterComputePoliciesCapabilitiesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterComputePoliciesCapabilitiesInfoWithDefaults

`func NewVcenterComputePoliciesCapabilitiesInfoWithDefaults() *VcenterComputePoliciesCapabilitiesInfo`

NewVcenterComputePoliciesCapabilitiesInfoWithDefaults instantiates a new VcenterComputePoliciesCapabilitiesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterComputePoliciesCapabilitiesInfo) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterComputePoliciesCapabilitiesInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreateSpecType

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetCreateSpecType() string`

GetCreateSpecType returns the CreateSpecType field if non-nil, zero value otherwise.

### GetCreateSpecTypeOk

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetCreateSpecTypeOk() (*string, bool)`

GetCreateSpecTypeOk returns a tuple with the CreateSpecType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSpecType

`func (o *VcenterComputePoliciesCapabilitiesInfo) SetCreateSpecType(v string)`

SetCreateSpecType sets CreateSpecType field to given value.


### GetInfoType

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetInfoType() string`

GetInfoType returns the InfoType field if non-nil, zero value otherwise.

### GetInfoTypeOk

`func (o *VcenterComputePoliciesCapabilitiesInfo) GetInfoTypeOk() (*string, bool)`

GetInfoTypeOk returns a tuple with the InfoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoType

`func (o *VcenterComputePoliciesCapabilitiesInfo) SetInfoType(v string)`

SetInfoType sets InfoType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


