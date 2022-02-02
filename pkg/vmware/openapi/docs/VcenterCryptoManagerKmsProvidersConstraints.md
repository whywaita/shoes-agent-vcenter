# VcenterCryptoManagerKmsProvidersConstraints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TpmRequired** | **bool** | Determines if a provider is restricted to hosts with TPM 2.0 capability. | 

## Methods

### NewVcenterCryptoManagerKmsProvidersConstraints

`func NewVcenterCryptoManagerKmsProvidersConstraints(tpmRequired bool, ) *VcenterCryptoManagerKmsProvidersConstraints`

NewVcenterCryptoManagerKmsProvidersConstraints instantiates a new VcenterCryptoManagerKmsProvidersConstraints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersConstraintsWithDefaults

`func NewVcenterCryptoManagerKmsProvidersConstraintsWithDefaults() *VcenterCryptoManagerKmsProvidersConstraints`

NewVcenterCryptoManagerKmsProvidersConstraintsWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersConstraints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTpmRequired

`func (o *VcenterCryptoManagerKmsProvidersConstraints) GetTpmRequired() bool`

GetTpmRequired returns the TpmRequired field if non-nil, zero value otherwise.

### GetTpmRequiredOk

`func (o *VcenterCryptoManagerKmsProvidersConstraints) GetTpmRequiredOk() (*bool, bool)`

GetTpmRequiredOk returns a tuple with the TpmRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmRequired

`func (o *VcenterCryptoManagerKmsProvidersConstraints) SetTpmRequired(v bool)`

SetTpmRequired sets TpmRequired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


