# VcenterGuestCustomizationSpecsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fingerprint** | **string** | The fingerprint is a unique identifier for a given version of the configuration. Each change to the configuration will update this value. A client cannot change this value. If specified when updating a specification, the changes will only be applied if the current fingerprint matches the specified fingerprint. This field can be used to guard against updates that has happened between the specification content was read and until it is applied. | 
**Spec** | [**VcenterGuestCustomizationSpec**](VcenterGuestCustomizationSpec.md) |  | 
**Description** | **string** | Description of the specification. | 
**Name** | **string** | Name of the specification. | 

## Methods

### NewVcenterGuestCustomizationSpecsSpec

`func NewVcenterGuestCustomizationSpecsSpec(fingerprint string, spec VcenterGuestCustomizationSpec, description string, name string, ) *VcenterGuestCustomizationSpecsSpec`

NewVcenterGuestCustomizationSpecsSpec instantiates a new VcenterGuestCustomizationSpecsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCustomizationSpecsSpecWithDefaults

`func NewVcenterGuestCustomizationSpecsSpecWithDefaults() *VcenterGuestCustomizationSpecsSpec`

NewVcenterGuestCustomizationSpecsSpecWithDefaults instantiates a new VcenterGuestCustomizationSpecsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFingerprint

`func (o *VcenterGuestCustomizationSpecsSpec) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *VcenterGuestCustomizationSpecsSpec) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *VcenterGuestCustomizationSpecsSpec) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetSpec

`func (o *VcenterGuestCustomizationSpecsSpec) GetSpec() VcenterGuestCustomizationSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VcenterGuestCustomizationSpecsSpec) GetSpecOk() (*VcenterGuestCustomizationSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VcenterGuestCustomizationSpecsSpec) SetSpec(v VcenterGuestCustomizationSpec)`

SetSpec sets Spec field to given value.


### GetDescription

`func (o *VcenterGuestCustomizationSpecsSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterGuestCustomizationSpecsSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterGuestCustomizationSpecsSpec) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *VcenterGuestCustomizationSpecsSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterGuestCustomizationSpecsSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterGuestCustomizationSpecsSpec) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


