# VcenterOvfLibraryItemCreateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Succeeded** | **bool** | Whether the {@name LibraryItem#create} {@term operation} completed successfully. | 
**OvfLibraryItemId** | Pointer to **string** | Identifier of the created or updated library item. | [optional] 
**Error** | Pointer to [**VcenterOvfLibraryItemResultInfo**](VcenterOvfLibraryItemResultInfo.md) |  | [optional] 

## Methods

### NewVcenterOvfLibraryItemCreateResult

`func NewVcenterOvfLibraryItemCreateResult(succeeded bool, ) *VcenterOvfLibraryItemCreateResult`

NewVcenterOvfLibraryItemCreateResult instantiates a new VcenterOvfLibraryItemCreateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemCreateResultWithDefaults

`func NewVcenterOvfLibraryItemCreateResultWithDefaults() *VcenterOvfLibraryItemCreateResult`

NewVcenterOvfLibraryItemCreateResultWithDefaults instantiates a new VcenterOvfLibraryItemCreateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSucceeded

`func (o *VcenterOvfLibraryItemCreateResult) GetSucceeded() bool`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *VcenterOvfLibraryItemCreateResult) GetSucceededOk() (*bool, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *VcenterOvfLibraryItemCreateResult) SetSucceeded(v bool)`

SetSucceeded sets Succeeded field to given value.


### GetOvfLibraryItemId

`func (o *VcenterOvfLibraryItemCreateResult) GetOvfLibraryItemId() string`

GetOvfLibraryItemId returns the OvfLibraryItemId field if non-nil, zero value otherwise.

### GetOvfLibraryItemIdOk

`func (o *VcenterOvfLibraryItemCreateResult) GetOvfLibraryItemIdOk() (*string, bool)`

GetOvfLibraryItemIdOk returns a tuple with the OvfLibraryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOvfLibraryItemId

`func (o *VcenterOvfLibraryItemCreateResult) SetOvfLibraryItemId(v string)`

SetOvfLibraryItemId sets OvfLibraryItemId field to given value.

### HasOvfLibraryItemId

`func (o *VcenterOvfLibraryItemCreateResult) HasOvfLibraryItemId() bool`

HasOvfLibraryItemId returns a boolean if a field has been set.

### GetError

`func (o *VcenterOvfLibraryItemCreateResult) GetError() VcenterOvfLibraryItemResultInfo`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterOvfLibraryItemCreateResult) GetErrorOk() (*VcenterOvfLibraryItemResultInfo, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterOvfLibraryItemCreateResult) SetError(v VcenterOvfLibraryItemResultInfo)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterOvfLibraryItemCreateResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


