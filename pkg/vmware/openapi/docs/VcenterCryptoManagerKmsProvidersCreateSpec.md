# VcenterCryptoManagerKmsProvidersCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Provider identifier.   A unique string provided by the client.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.crypto_manager.kms.provider. | 
**Constraints** | Pointer to [**VcenterCryptoManagerKmsProvidersConstraintsSpec**](VcenterCryptoManagerKmsProvidersConstraintsSpec.md) |  | [optional] 
**NativeSpec** | Pointer to [**VcenterCryptoManagerKmsProvidersNativeProviderCreateSpec**](VcenterCryptoManagerKmsProvidersNativeProviderCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersCreateSpec

`func NewVcenterCryptoManagerKmsProvidersCreateSpec(provider string, ) *VcenterCryptoManagerKmsProvidersCreateSpec`

NewVcenterCryptoManagerKmsProvidersCreateSpec instantiates a new VcenterCryptoManagerKmsProvidersCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersCreateSpecWithDefaults

`func NewVcenterCryptoManagerKmsProvidersCreateSpecWithDefaults() *VcenterCryptoManagerKmsProvidersCreateSpec`

NewVcenterCryptoManagerKmsProvidersCreateSpecWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetConstraints

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetConstraints() VcenterCryptoManagerKmsProvidersConstraintsSpec`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetConstraintsOk() (*VcenterCryptoManagerKmsProvidersConstraintsSpec, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) SetConstraints(v VcenterCryptoManagerKmsProvidersConstraintsSpec)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetNativeSpec

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetNativeSpec() VcenterCryptoManagerKmsProvidersNativeProviderCreateSpec`

GetNativeSpec returns the NativeSpec field if non-nil, zero value otherwise.

### GetNativeSpecOk

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) GetNativeSpecOk() (*VcenterCryptoManagerKmsProvidersNativeProviderCreateSpec, bool)`

GetNativeSpecOk returns a tuple with the NativeSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeSpec

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) SetNativeSpec(v VcenterCryptoManagerKmsProvidersNativeProviderCreateSpec)`

SetNativeSpec sets NativeSpec field to given value.

### HasNativeSpec

`func (o *VcenterCryptoManagerKmsProvidersCreateSpec) HasNativeSpec() bool`

HasNativeSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


