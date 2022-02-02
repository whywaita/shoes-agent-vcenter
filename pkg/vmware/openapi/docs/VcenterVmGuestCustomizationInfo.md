# VcenterVmGuestCustomizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VcenterVmGuestCustomizationInfoStatus**](VcenterVmGuestCustomizationInfoStatus.md) |  | 
**Error** | Pointer to **string** | Description of the error if the Customization.Info.status of customization operation is FAILED. This field will be unset if the status is not FAILED or there is no information available for the error. | [optional] 
**StartTime** | Pointer to **time.Time** | Time when the customization process has started inside the guest operating system. This field will be unset if the status is PENDING. | [optional] 
**EndTime** | Pointer to **time.Time** | Time when the customization process has completed inside the guest operating system. This field will be unset if the status is not SUCCEEDED or FAILED. | [optional] 

## Methods

### NewVcenterVmGuestCustomizationInfo

`func NewVcenterVmGuestCustomizationInfo(status VcenterVmGuestCustomizationInfoStatus, ) *VcenterVmGuestCustomizationInfo`

NewVcenterVmGuestCustomizationInfo instantiates a new VcenterVmGuestCustomizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestCustomizationInfoWithDefaults

`func NewVcenterVmGuestCustomizationInfoWithDefaults() *VcenterVmGuestCustomizationInfo`

NewVcenterVmGuestCustomizationInfoWithDefaults instantiates a new VcenterVmGuestCustomizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterVmGuestCustomizationInfo) GetStatus() VcenterVmGuestCustomizationInfoStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterVmGuestCustomizationInfo) GetStatusOk() (*VcenterVmGuestCustomizationInfoStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterVmGuestCustomizationInfo) SetStatus(v VcenterVmGuestCustomizationInfoStatus)`

SetStatus sets Status field to given value.


### GetError

`func (o *VcenterVmGuestCustomizationInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterVmGuestCustomizationInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterVmGuestCustomizationInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterVmGuestCustomizationInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartTime

`func (o *VcenterVmGuestCustomizationInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VcenterVmGuestCustomizationInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VcenterVmGuestCustomizationInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VcenterVmGuestCustomizationInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *VcenterVmGuestCustomizationInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VcenterVmGuestCustomizationInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VcenterVmGuestCustomizationInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VcenterVmGuestCustomizationInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


