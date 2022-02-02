# VcenterGuestCloudConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterGuestCloudConfigurationType**](VcenterGuestCloudConfigurationType.md) |  | 
**Cloudinit** | Pointer to [**VcenterGuestCloudinitConfiguration**](VcenterGuestCloudinitConfiguration.md) |  | [optional] 

## Methods

### NewVcenterGuestCloudConfiguration

`func NewVcenterGuestCloudConfiguration(type_ VcenterGuestCloudConfigurationType, ) *VcenterGuestCloudConfiguration`

NewVcenterGuestCloudConfiguration instantiates a new VcenterGuestCloudConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCloudConfigurationWithDefaults

`func NewVcenterGuestCloudConfigurationWithDefaults() *VcenterGuestCloudConfiguration`

NewVcenterGuestCloudConfigurationWithDefaults instantiates a new VcenterGuestCloudConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterGuestCloudConfiguration) GetType() VcenterGuestCloudConfigurationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterGuestCloudConfiguration) GetTypeOk() (*VcenterGuestCloudConfigurationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterGuestCloudConfiguration) SetType(v VcenterGuestCloudConfigurationType)`

SetType sets Type field to given value.


### GetCloudinit

`func (o *VcenterGuestCloudConfiguration) GetCloudinit() VcenterGuestCloudinitConfiguration`

GetCloudinit returns the Cloudinit field if non-nil, zero value otherwise.

### GetCloudinitOk

`func (o *VcenterGuestCloudConfiguration) GetCloudinitOk() (*VcenterGuestCloudinitConfiguration, bool)`

GetCloudinitOk returns a tuple with the Cloudinit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudinit

`func (o *VcenterGuestCloudConfiguration) SetCloudinit(v VcenterGuestCloudinitConfiguration)`

SetCloudinit sets Cloudinit field to given value.

### HasCloudinit

`func (o *VcenterGuestCloudConfiguration) HasCloudinit() bool`

HasCloudinit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


