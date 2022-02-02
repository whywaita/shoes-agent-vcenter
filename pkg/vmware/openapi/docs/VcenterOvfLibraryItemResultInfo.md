# VcenterOvfLibraryItemResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]VcenterOvfOvfError**](VcenterOvfOvfError.md) | Errors reported by the {@name LibraryItem#create} or {@name LibraryItem#deploy} {@term operation}. These errors would have prevented the {@name LibraryItem#create} or {@name LibraryItem#deploy} {@term operation} from completing successfully. | 
**Warnings** | [**[]VcenterOvfOvfWarning**](VcenterOvfOvfWarning.md) | Warnings reported by the {@name LibraryItem#create} or {@name LibraryItem#deploy} {@term operation}. These warnings would not have prevented the {@name LibraryItem#create} or {@name LibraryItem#deploy} {@term operation} from completing successfully, but there might be issues that warrant attention. | 
**Information** | [**[]VcenterOvfOvfInfo**](VcenterOvfOvfInfo.md) | Information messages reported by the {@name LibraryItem#create} or {@name LibraryItem#deploy} {@term operation}. For example, a non-required parameter was ignored. | 

## Methods

### NewVcenterOvfLibraryItemResultInfo

`func NewVcenterOvfLibraryItemResultInfo(errors []VcenterOvfOvfError, warnings []VcenterOvfOvfWarning, information []VcenterOvfOvfInfo, ) *VcenterOvfLibraryItemResultInfo`

NewVcenterOvfLibraryItemResultInfo instantiates a new VcenterOvfLibraryItemResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemResultInfoWithDefaults

`func NewVcenterOvfLibraryItemResultInfoWithDefaults() *VcenterOvfLibraryItemResultInfo`

NewVcenterOvfLibraryItemResultInfoWithDefaults instantiates a new VcenterOvfLibraryItemResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *VcenterOvfLibraryItemResultInfo) GetErrors() []VcenterOvfOvfError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VcenterOvfLibraryItemResultInfo) GetErrorsOk() (*[]VcenterOvfOvfError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VcenterOvfLibraryItemResultInfo) SetErrors(v []VcenterOvfOvfError)`

SetErrors sets Errors field to given value.


### GetWarnings

`func (o *VcenterOvfLibraryItemResultInfo) GetWarnings() []VcenterOvfOvfWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VcenterOvfLibraryItemResultInfo) GetWarningsOk() (*[]VcenterOvfOvfWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VcenterOvfLibraryItemResultInfo) SetWarnings(v []VcenterOvfOvfWarning)`

SetWarnings sets Warnings field to given value.


### GetInformation

`func (o *VcenterOvfLibraryItemResultInfo) GetInformation() []VcenterOvfOvfInfo`

GetInformation returns the Information field if non-nil, zero value otherwise.

### GetInformationOk

`func (o *VcenterOvfLibraryItemResultInfo) GetInformationOk() (*[]VcenterOvfOvfInfo, bool)`

GetInformationOk returns a tuple with the Information field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformation

`func (o *VcenterOvfLibraryItemResultInfo) SetInformation(v []VcenterOvfOvfInfo)`

SetInformation sets Information field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


