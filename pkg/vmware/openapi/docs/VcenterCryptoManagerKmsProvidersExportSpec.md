# VcenterCryptoManagerKmsProvidersExportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Provider identifier When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.crypto_manager.kms.provider. | 
**Password** | Pointer to **string** | Password used to encrypt the exported configuration. If unset or empty, the configuration will not be encrypted. | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersExportSpec

`func NewVcenterCryptoManagerKmsProvidersExportSpec(provider string, ) *VcenterCryptoManagerKmsProvidersExportSpec`

NewVcenterCryptoManagerKmsProvidersExportSpec instantiates a new VcenterCryptoManagerKmsProvidersExportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersExportSpecWithDefaults

`func NewVcenterCryptoManagerKmsProvidersExportSpecWithDefaults() *VcenterCryptoManagerKmsProvidersExportSpec`

NewVcenterCryptoManagerKmsProvidersExportSpecWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersExportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetPassword

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterCryptoManagerKmsProvidersExportSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


