# VcenterGuestCustomizationSpecsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastModified** | **time.Time** | Time when the specification was last modified. | 
**Spec** | [**VcenterGuestCustomizationSpecsSpec**](VcenterGuestCustomizationSpecsSpec.md) |  | 

## Methods

### NewVcenterGuestCustomizationSpecsInfo

`func NewVcenterGuestCustomizationSpecsInfo(lastModified time.Time, spec VcenterGuestCustomizationSpecsSpec, ) *VcenterGuestCustomizationSpecsInfo`

NewVcenterGuestCustomizationSpecsInfo instantiates a new VcenterGuestCustomizationSpecsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCustomizationSpecsInfoWithDefaults

`func NewVcenterGuestCustomizationSpecsInfoWithDefaults() *VcenterGuestCustomizationSpecsInfo`

NewVcenterGuestCustomizationSpecsInfoWithDefaults instantiates a new VcenterGuestCustomizationSpecsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastModified

`func (o *VcenterGuestCustomizationSpecsInfo) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VcenterGuestCustomizationSpecsInfo) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VcenterGuestCustomizationSpecsInfo) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.


### GetSpec

`func (o *VcenterGuestCustomizationSpecsInfo) GetSpec() VcenterGuestCustomizationSpecsSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VcenterGuestCustomizationSpecsInfo) GetSpecOk() (*VcenterGuestCustomizationSpecsSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VcenterGuestCustomizationSpecsInfo) SetSpec(v VcenterGuestCustomizationSpecsSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


