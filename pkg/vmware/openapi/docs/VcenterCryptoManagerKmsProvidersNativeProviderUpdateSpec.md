# VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** | Key identifier for the provider.   The key identifier is required to be a 128-bit UUID represented as a hexadecimal string in \&quot;12345678-abcd-1234-cdef-123456789abc\&quot; format.  If unset, the key identifier will remain unchanged. | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec

`func NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec() *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec`

NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec instantiates a new VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpecWithDefaults

`func NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpecWithDefaults() *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec`

NewVcenterCryptoManagerKmsProvidersNativeProviderUpdateSpecWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *VcenterCryptoManagerKmsProvidersNativeProviderUpdateSpec) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


