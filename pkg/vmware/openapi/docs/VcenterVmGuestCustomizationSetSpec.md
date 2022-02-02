# VcenterVmGuestCustomizationSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the customization specification that has be retrieved from the virtual center inventory and applied for the virtual machine. Either one of Customization.SetSpec.name or Customization.SetSpec.spec or none of them should be specified. If unset and Customization.SetSpec.spec is also unset when executing Customization.set operationg, then any pending customization for the virtual machine will be cleared. | [optional] 
**Spec** | Pointer to [**VcenterGuestCustomizationSpec**](VcenterGuestCustomizationSpec.md) |  | [optional] 

## Methods

### NewVcenterVmGuestCustomizationSetSpec

`func NewVcenterVmGuestCustomizationSetSpec() *VcenterVmGuestCustomizationSetSpec`

NewVcenterVmGuestCustomizationSetSpec instantiates a new VcenterVmGuestCustomizationSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestCustomizationSetSpecWithDefaults

`func NewVcenterVmGuestCustomizationSetSpecWithDefaults() *VcenterVmGuestCustomizationSetSpec`

NewVcenterVmGuestCustomizationSetSpecWithDefaults instantiates a new VcenterVmGuestCustomizationSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmGuestCustomizationSetSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmGuestCustomizationSetSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmGuestCustomizationSetSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterVmGuestCustomizationSetSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpec

`func (o *VcenterVmGuestCustomizationSetSpec) GetSpec() VcenterGuestCustomizationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VcenterVmGuestCustomizationSetSpec) GetSpecOk() (*VcenterGuestCustomizationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VcenterVmGuestCustomizationSetSpec) SetSpec(v VcenterGuestCustomizationSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VcenterVmGuestCustomizationSetSpec) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


