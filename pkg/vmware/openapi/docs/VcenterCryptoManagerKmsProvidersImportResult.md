# VcenterCryptoManagerKmsProvidersImportResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Provider identifier When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.crypto_manager.kms.provider. | 
**Type** | [**VcenterCryptoManagerKmsProvidersType**](VcenterCryptoManagerKmsProvidersType.md) |  | 
**NativeInfo** | Pointer to [**VcenterCryptoManagerKmsProvidersNativeProviderInfo**](VcenterCryptoManagerKmsProvidersNativeProviderInfo.md) |  | [optional] 
**ExportTime** | **time.Time** | Time when the provider was exported | 
**Constraints** | Pointer to [**VcenterCryptoManagerKmsProvidersConstraints**](VcenterCryptoManagerKmsProvidersConstraints.md) |  | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersImportResult

`func NewVcenterCryptoManagerKmsProvidersImportResult(provider string, type_ VcenterCryptoManagerKmsProvidersType, exportTime time.Time, ) *VcenterCryptoManagerKmsProvidersImportResult`

NewVcenterCryptoManagerKmsProvidersImportResult instantiates a new VcenterCryptoManagerKmsProvidersImportResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersImportResultWithDefaults

`func NewVcenterCryptoManagerKmsProvidersImportResultWithDefaults() *VcenterCryptoManagerKmsProvidersImportResult`

NewVcenterCryptoManagerKmsProvidersImportResultWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersImportResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterCryptoManagerKmsProvidersImportResult) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetType() VcenterCryptoManagerKmsProvidersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetTypeOk() (*VcenterCryptoManagerKmsProvidersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterCryptoManagerKmsProvidersImportResult) SetType(v VcenterCryptoManagerKmsProvidersType)`

SetType sets Type field to given value.


### GetNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetNativeInfo() VcenterCryptoManagerKmsProvidersNativeProviderInfo`

GetNativeInfo returns the NativeInfo field if non-nil, zero value otherwise.

### GetNativeInfoOk

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetNativeInfoOk() (*VcenterCryptoManagerKmsProvidersNativeProviderInfo, bool)`

GetNativeInfoOk returns a tuple with the NativeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersImportResult) SetNativeInfo(v VcenterCryptoManagerKmsProvidersNativeProviderInfo)`

SetNativeInfo sets NativeInfo field to given value.

### HasNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersImportResult) HasNativeInfo() bool`

HasNativeInfo returns a boolean if a field has been set.

### GetExportTime

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetExportTime() time.Time`

GetExportTime returns the ExportTime field if non-nil, zero value otherwise.

### GetExportTimeOk

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetExportTimeOk() (*time.Time, bool)`

GetExportTimeOk returns a tuple with the ExportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTime

`func (o *VcenterCryptoManagerKmsProvidersImportResult) SetExportTime(v time.Time)`

SetExportTime sets ExportTime field to given value.


### GetConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetConstraints() VcenterCryptoManagerKmsProvidersConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *VcenterCryptoManagerKmsProvidersImportResult) GetConstraintsOk() (*VcenterCryptoManagerKmsProvidersConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportResult) SetConstraints(v VcenterCryptoManagerKmsProvidersConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *VcenterCryptoManagerKmsProvidersImportResult) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


