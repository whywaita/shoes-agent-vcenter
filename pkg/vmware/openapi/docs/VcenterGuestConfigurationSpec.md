# VcenterGuestConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WindowsConfig** | Pointer to [**VcenterGuestWindowsConfiguration**](VcenterGuestWindowsConfiguration.md) |  | [optional] 
**LinuxConfig** | Pointer to [**VcenterGuestLinuxConfiguration**](VcenterGuestLinuxConfiguration.md) |  | [optional] 
**CloudConfig** | Pointer to [**VcenterGuestCloudConfiguration**](VcenterGuestCloudConfiguration.md) |  | [optional] 

## Methods

### NewVcenterGuestConfigurationSpec

`func NewVcenterGuestConfigurationSpec() *VcenterGuestConfigurationSpec`

NewVcenterGuestConfigurationSpec instantiates a new VcenterGuestConfigurationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestConfigurationSpecWithDefaults

`func NewVcenterGuestConfigurationSpecWithDefaults() *VcenterGuestConfigurationSpec`

NewVcenterGuestConfigurationSpecWithDefaults instantiates a new VcenterGuestConfigurationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindowsConfig

`func (o *VcenterGuestConfigurationSpec) GetWindowsConfig() VcenterGuestWindowsConfiguration`

GetWindowsConfig returns the WindowsConfig field if non-nil, zero value otherwise.

### GetWindowsConfigOk

`func (o *VcenterGuestConfigurationSpec) GetWindowsConfigOk() (*VcenterGuestWindowsConfiguration, bool)`

GetWindowsConfigOk returns a tuple with the WindowsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsConfig

`func (o *VcenterGuestConfigurationSpec) SetWindowsConfig(v VcenterGuestWindowsConfiguration)`

SetWindowsConfig sets WindowsConfig field to given value.

### HasWindowsConfig

`func (o *VcenterGuestConfigurationSpec) HasWindowsConfig() bool`

HasWindowsConfig returns a boolean if a field has been set.

### GetLinuxConfig

`func (o *VcenterGuestConfigurationSpec) GetLinuxConfig() VcenterGuestLinuxConfiguration`

GetLinuxConfig returns the LinuxConfig field if non-nil, zero value otherwise.

### GetLinuxConfigOk

`func (o *VcenterGuestConfigurationSpec) GetLinuxConfigOk() (*VcenterGuestLinuxConfiguration, bool)`

GetLinuxConfigOk returns a tuple with the LinuxConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxConfig

`func (o *VcenterGuestConfigurationSpec) SetLinuxConfig(v VcenterGuestLinuxConfiguration)`

SetLinuxConfig sets LinuxConfig field to given value.

### HasLinuxConfig

`func (o *VcenterGuestConfigurationSpec) HasLinuxConfig() bool`

HasLinuxConfig returns a boolean if a field has been set.

### GetCloudConfig

`func (o *VcenterGuestConfigurationSpec) GetCloudConfig() VcenterGuestCloudConfiguration`

GetCloudConfig returns the CloudConfig field if non-nil, zero value otherwise.

### GetCloudConfigOk

`func (o *VcenterGuestConfigurationSpec) GetCloudConfigOk() (*VcenterGuestCloudConfiguration, bool)`

GetCloudConfigOk returns a tuple with the CloudConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudConfig

`func (o *VcenterGuestConfigurationSpec) SetCloudConfig(v VcenterGuestCloudConfiguration)`

SetCloudConfig sets CloudConfig field to given value.

### HasCloudConfig

`func (o *VcenterGuestConfigurationSpec) HasCloudConfig() bool`

HasCloudConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


