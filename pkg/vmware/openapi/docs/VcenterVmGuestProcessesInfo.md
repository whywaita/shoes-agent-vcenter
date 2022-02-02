# VcenterVmGuestProcessesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The process name. | 
**Owner** | **string** | The process owner. | 
**Command** | **string** | The full command line of the process. | 
**Started** | **time.Time** | The start time of the process. | 
**Finished** | Pointer to **time.Time** | If the process was started using Processes.create then the process completion time will be available if queried within 5 minutes after it completes. Set if the process was started with Processes.create and has recently exited. | [optional] 
**ExitCode** | Pointer to **int64** | If the process was started using Processes.create then the process exit code will be available if queried within 5 minutes after it completes. Set if the process was started with Processes.create and has recently exited. | [optional] 

## Methods

### NewVcenterVmGuestProcessesInfo

`func NewVcenterVmGuestProcessesInfo(name string, owner string, command string, started time.Time, ) *VcenterVmGuestProcessesInfo`

NewVcenterVmGuestProcessesInfo instantiates a new VcenterVmGuestProcessesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestProcessesInfoWithDefaults

`func NewVcenterVmGuestProcessesInfoWithDefaults() *VcenterVmGuestProcessesInfo`

NewVcenterVmGuestProcessesInfoWithDefaults instantiates a new VcenterVmGuestProcessesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmGuestProcessesInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmGuestProcessesInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmGuestProcessesInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *VcenterVmGuestProcessesInfo) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VcenterVmGuestProcessesInfo) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VcenterVmGuestProcessesInfo) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetCommand

`func (o *VcenterVmGuestProcessesInfo) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *VcenterVmGuestProcessesInfo) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *VcenterVmGuestProcessesInfo) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetStarted

`func (o *VcenterVmGuestProcessesInfo) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *VcenterVmGuestProcessesInfo) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *VcenterVmGuestProcessesInfo) SetStarted(v time.Time)`

SetStarted sets Started field to given value.


### GetFinished

`func (o *VcenterVmGuestProcessesInfo) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *VcenterVmGuestProcessesInfo) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *VcenterVmGuestProcessesInfo) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *VcenterVmGuestProcessesInfo) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetExitCode

`func (o *VcenterVmGuestProcessesInfo) GetExitCode() int64`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *VcenterVmGuestProcessesInfo) GetExitCodeOk() (*int64, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *VcenterVmGuestProcessesInfo) SetExitCode(v int64)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *VcenterVmGuestProcessesInfo) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


