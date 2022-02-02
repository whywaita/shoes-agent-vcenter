# VcenterStoragePoliciesComplianceFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**[]VcenterStoragePoliciesComplianceStatus**](VcenterStoragePoliciesComplianceStatus.md) | Compliance Status that a virtual machine must have to match the filter. | 

## Methods

### NewVcenterStoragePoliciesComplianceFilterSpec

`func NewVcenterStoragePoliciesComplianceFilterSpec(status []VcenterStoragePoliciesComplianceStatus, ) *VcenterStoragePoliciesComplianceFilterSpec`

NewVcenterStoragePoliciesComplianceFilterSpec instantiates a new VcenterStoragePoliciesComplianceFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesComplianceFilterSpecWithDefaults

`func NewVcenterStoragePoliciesComplianceFilterSpecWithDefaults() *VcenterStoragePoliciesComplianceFilterSpec`

NewVcenterStoragePoliciesComplianceFilterSpecWithDefaults instantiates a new VcenterStoragePoliciesComplianceFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterStoragePoliciesComplianceFilterSpec) GetStatus() []VcenterStoragePoliciesComplianceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterStoragePoliciesComplianceFilterSpec) GetStatusOk() (*[]VcenterStoragePoliciesComplianceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterStoragePoliciesComplianceFilterSpec) SetStatus(v []VcenterStoragePoliciesComplianceStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


