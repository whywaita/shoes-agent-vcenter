# VcenterVmGuestProcessesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The absolute path to the program to start.   For Linux guest operating systems, /bin/bash is used to start the program.    For Solaris guest operating systems, if /bin/bash exists, its used to start the program, otherwise /bin/sh is used. If /bin/sh is used, then the process ID returned by Processes.create will be that of the shell used to start the program, rather than the program itself, due to the differences in how /bin/sh and /bin/bash work. This PID will still be usable for watching the process with Processes.list to find its exit code and elapsed time.    For Windows, no shell is used. Using a simple batch file instead by prepending c:\\windows\\system32\\cmd.exe /c will allow stdio redirection to work if passed in the Processes.CreateSpec.arguments parameter.  | 
**Arguments** | Pointer to **string** | The arguments to the program.   Characters which must be escaped to the shell should also be escaped in Processes.CreateSpec.arguments.    In Linux and Solaris guest operating systems, stdio redirection arguments may be used.    For Windows, stdio redirection can be added to the argments if Processes.CreateSpec.path is prefixed with c:\\windows\\system32\\cmd.exe /c.  If unset no arguments are passed to the program. | [optional] 
**WorkingDirectory** | Pointer to **string** | The absolute path of the working directory for the program to be run. VMware recommends explicitly setting the working directory for the program to be run.  If unset or is an empty string, the behavior depends on the guest operating system. For Linux guest operating systems, if unset or is an empty string, the working directory will be the home directory of the user associated with the guest authentication. For other guest operating systems, if unset, the behavior is unspecified. | [optional] 
**EnvironmentVariables** | Pointer to **map[string]string** | A map of environment variables, specified using the guest OS rules (for example PATH, c:\\bin;c:\\windows\\system32 or LD_LIBRARY_PATH,/usr/lib:/lib), to be set for the program being run. Note that these are not additions to the default environment variables; they define the complete set available to the program. If unset, the environment variables used are guest dependent defaults. | [optional] 
**StartMinimized** | Pointer to **bool** | Makes any program window start minimized in Windows operating systems. Returns an error if set for non-Windows guests. Defaults to false. | [optional] 

## Methods

### NewVcenterVmGuestProcessesCreateSpec

`func NewVcenterVmGuestProcessesCreateSpec(path string, ) *VcenterVmGuestProcessesCreateSpec`

NewVcenterVmGuestProcessesCreateSpec instantiates a new VcenterVmGuestProcessesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestProcessesCreateSpecWithDefaults

`func NewVcenterVmGuestProcessesCreateSpecWithDefaults() *VcenterVmGuestProcessesCreateSpec`

NewVcenterVmGuestProcessesCreateSpecWithDefaults instantiates a new VcenterVmGuestProcessesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *VcenterVmGuestProcessesCreateSpec) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VcenterVmGuestProcessesCreateSpec) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VcenterVmGuestProcessesCreateSpec) SetPath(v string)`

SetPath sets Path field to given value.


### GetArguments

`func (o *VcenterVmGuestProcessesCreateSpec) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *VcenterVmGuestProcessesCreateSpec) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *VcenterVmGuestProcessesCreateSpec) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *VcenterVmGuestProcessesCreateSpec) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetWorkingDirectory

`func (o *VcenterVmGuestProcessesCreateSpec) GetWorkingDirectory() string`

GetWorkingDirectory returns the WorkingDirectory field if non-nil, zero value otherwise.

### GetWorkingDirectoryOk

`func (o *VcenterVmGuestProcessesCreateSpec) GetWorkingDirectoryOk() (*string, bool)`

GetWorkingDirectoryOk returns a tuple with the WorkingDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDirectory

`func (o *VcenterVmGuestProcessesCreateSpec) SetWorkingDirectory(v string)`

SetWorkingDirectory sets WorkingDirectory field to given value.

### HasWorkingDirectory

`func (o *VcenterVmGuestProcessesCreateSpec) HasWorkingDirectory() bool`

HasWorkingDirectory returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *VcenterVmGuestProcessesCreateSpec) GetEnvironmentVariables() map[string]string`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *VcenterVmGuestProcessesCreateSpec) GetEnvironmentVariablesOk() (*map[string]string, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *VcenterVmGuestProcessesCreateSpec) SetEnvironmentVariables(v map[string]string)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *VcenterVmGuestProcessesCreateSpec) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetStartMinimized

`func (o *VcenterVmGuestProcessesCreateSpec) GetStartMinimized() bool`

GetStartMinimized returns the StartMinimized field if non-nil, zero value otherwise.

### GetStartMinimizedOk

`func (o *VcenterVmGuestProcessesCreateSpec) GetStartMinimizedOk() (*bool, bool)`

GetStartMinimizedOk returns a tuple with the StartMinimized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartMinimized

`func (o *VcenterVmGuestProcessesCreateSpec) SetStartMinimized(v bool)`

SetStartMinimized sets StartMinimized field to given value.

### HasStartMinimized

`func (o *VcenterVmGuestProcessesCreateSpec) HasStartMinimized() bool`

HasStartMinimized returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


