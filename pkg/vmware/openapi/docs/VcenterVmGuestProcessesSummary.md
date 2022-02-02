# VcenterVmGuestProcessesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The process name. | 
**Pid** | **int64** | The process ID. | 
**Owner** | **string** | The process owner. | 
**Command** | **string** | The full command line of the process. | 
**Started** | **time.Time** | The start time of the process. | 

## Methods

### NewVcenterVmGuestProcessesSummary

`func NewVcenterVmGuestProcessesSummary(name string, pid int64, owner string, command string, started time.Time, ) *VcenterVmGuestProcessesSummary`

NewVcenterVmGuestProcessesSummary instantiates a new VcenterVmGuestProcessesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestProcessesSummaryWithDefaults

`func NewVcenterVmGuestProcessesSummaryWithDefaults() *VcenterVmGuestProcessesSummary`

NewVcenterVmGuestProcessesSummaryWithDefaults instantiates a new VcenterVmGuestProcessesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmGuestProcessesSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmGuestProcessesSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmGuestProcessesSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPid

`func (o *VcenterVmGuestProcessesSummary) GetPid() int64`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *VcenterVmGuestProcessesSummary) GetPidOk() (*int64, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *VcenterVmGuestProcessesSummary) SetPid(v int64)`

SetPid sets Pid field to given value.


### GetOwner

`func (o *VcenterVmGuestProcessesSummary) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VcenterVmGuestProcessesSummary) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VcenterVmGuestProcessesSummary) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetCommand

`func (o *VcenterVmGuestProcessesSummary) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *VcenterVmGuestProcessesSummary) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *VcenterVmGuestProcessesSummary) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetStarted

`func (o *VcenterVmGuestProcessesSummary) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *VcenterVmGuestProcessesSummary) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *VcenterVmGuestProcessesSummary) SetStarted(v time.Time)`

SetStarted sets Started field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


