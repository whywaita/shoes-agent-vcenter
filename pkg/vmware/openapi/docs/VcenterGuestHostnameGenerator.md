# VcenterGuestHostnameGenerator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterGuestHostnameGeneratorType**](VcenterGuestHostnameGeneratorType.md) |  | 
**FixedName** | Pointer to **string** | The virtual machine name specified by the client. This field is optional and it is only relevant when the value of HostnameGenerator.type is FIXED. | [optional] 
**Prefix** | Pointer to **string** | Base prefix, to which a unique number is appended. This field is optional and it is only relevant when the value of HostnameGenerator.type is PREFIX. | [optional] 

## Methods

### NewVcenterGuestHostnameGenerator

`func NewVcenterGuestHostnameGenerator(type_ VcenterGuestHostnameGeneratorType, ) *VcenterGuestHostnameGenerator`

NewVcenterGuestHostnameGenerator instantiates a new VcenterGuestHostnameGenerator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestHostnameGeneratorWithDefaults

`func NewVcenterGuestHostnameGeneratorWithDefaults() *VcenterGuestHostnameGenerator`

NewVcenterGuestHostnameGeneratorWithDefaults instantiates a new VcenterGuestHostnameGenerator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterGuestHostnameGenerator) GetType() VcenterGuestHostnameGeneratorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterGuestHostnameGenerator) GetTypeOk() (*VcenterGuestHostnameGeneratorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterGuestHostnameGenerator) SetType(v VcenterGuestHostnameGeneratorType)`

SetType sets Type field to given value.


### GetFixedName

`func (o *VcenterGuestHostnameGenerator) GetFixedName() string`

GetFixedName returns the FixedName field if non-nil, zero value otherwise.

### GetFixedNameOk

`func (o *VcenterGuestHostnameGenerator) GetFixedNameOk() (*string, bool)`

GetFixedNameOk returns a tuple with the FixedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedName

`func (o *VcenterGuestHostnameGenerator) SetFixedName(v string)`

SetFixedName sets FixedName field to given value.

### HasFixedName

`func (o *VcenterGuestHostnameGenerator) HasFixedName() bool`

HasFixedName returns a boolean if a field has been set.

### GetPrefix

`func (o *VcenterGuestHostnameGenerator) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *VcenterGuestHostnameGenerator) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *VcenterGuestHostnameGenerator) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *VcenterGuestHostnameGenerator) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


