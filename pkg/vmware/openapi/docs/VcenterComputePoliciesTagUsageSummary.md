# VcenterComputePoliciesTagUsageSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | **string** | Identifier of the policy that uses the tag indicated by {@link #tag}. | 
**TagType** | **string** | Identifier of the tag type used by the policy indicated by {@link #policy}. | 
**Tag** | **string** | Identifier of the tag used by the policy indicated by {@link #policy}. | 
**TagName** | **string** | Name of the tag used by the policy indicated by {@link #policy}. | 
**CategoryName** | **string** | Name of the category that has {@link #tag}. | 

## Methods

### NewVcenterComputePoliciesTagUsageSummary

`func NewVcenterComputePoliciesTagUsageSummary(policy string, tagType string, tag string, tagName string, categoryName string, ) *VcenterComputePoliciesTagUsageSummary`

NewVcenterComputePoliciesTagUsageSummary instantiates a new VcenterComputePoliciesTagUsageSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterComputePoliciesTagUsageSummaryWithDefaults

`func NewVcenterComputePoliciesTagUsageSummaryWithDefaults() *VcenterComputePoliciesTagUsageSummary`

NewVcenterComputePoliciesTagUsageSummaryWithDefaults instantiates a new VcenterComputePoliciesTagUsageSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VcenterComputePoliciesTagUsageSummary) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VcenterComputePoliciesTagUsageSummary) GetPolicyOk() (*string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VcenterComputePoliciesTagUsageSummary) SetPolicy(v string)`

SetPolicy sets Policy field to given value.


### GetTagType

`func (o *VcenterComputePoliciesTagUsageSummary) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *VcenterComputePoliciesTagUsageSummary) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *VcenterComputePoliciesTagUsageSummary) SetTagType(v string)`

SetTagType sets TagType field to given value.


### GetTag

`func (o *VcenterComputePoliciesTagUsageSummary) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *VcenterComputePoliciesTagUsageSummary) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *VcenterComputePoliciesTagUsageSummary) SetTag(v string)`

SetTag sets Tag field to given value.


### GetTagName

`func (o *VcenterComputePoliciesTagUsageSummary) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *VcenterComputePoliciesTagUsageSummary) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *VcenterComputePoliciesTagUsageSummary) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetCategoryName

`func (o *VcenterComputePoliciesTagUsageSummary) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *VcenterComputePoliciesTagUsageSummary) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *VcenterComputePoliciesTagUsageSummary) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


