# VcenterCryptoManagerKmsProvidersImportSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to ***os.File** | Configuration to import. Currently this is required. Other import methods may be supported in the future. | [optional] 
**Password** | Pointer to **string** | Password to decrypt the configuration to import. If unset or empty, configuration to import must be unencrypted. | [optional] 
**Constraints** | Pointer to [**VcenterCryptoManagerKmsProvidersConstraintsSpec**](VcenterCryptoManagerKmsProvidersConstraintsSpec.md) |  | [optional] 
**DryRun** | Pointer to **bool** | Whether to perform a trial import without actuallly creating a provider. If unset, a new provider will be created. | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersImportSpec

`func NewVcenterCryptoManagerKmsProvidersImportSpec() *VcenterCryptoManagerKmsProvidersImportSpec`

NewVcenterCryptoManagerKmsProvidersImportSpec instantiates a new VcenterCryptoManagerKmsProvidersImportSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersImportSpecWithDefaults

`func NewVcenterCryptoManagerKmsProvidersImportSpecWithDefaults() *VcenterCryptoManagerKmsProvidersImportSpec`

NewVcenterCryptoManagerKmsProvidersImportSpecWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersImportSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetConfig() *os.File`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetConfigOk() (**os.File, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) SetConfig(v *os.File)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetPassword

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetConstraints() VcenterCryptoManagerKmsProvidersConstraintsSpec`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetConstraintsOk() (*VcenterCryptoManagerKmsProvidersConstraintsSpec, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) SetConstraints(v VcenterCryptoManagerKmsProvidersConstraintsSpec)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDryRun

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *VcenterCryptoManagerKmsProvidersImportSpec) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


