# VcenterVmGuestIdentityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**VcenterVmGuestOS**](VcenterVmGuestOS.md) |  | 
**Family** | [**VcenterVmGuestOSFamily**](VcenterVmGuestOSFamily.md) |  | 
**FullName** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 
**HostName** | **string** | Hostname of the guest operating system. | 
**IpAddress** | Pointer to **string** | IP address assigned by the guest operating system. If unset the guest does not have an IP address. | [optional] 

## Methods

### NewVcenterVmGuestIdentityInfo

`func NewVcenterVmGuestIdentityInfo(name VcenterVmGuestOS, family VcenterVmGuestOSFamily, fullName VapiStdLocalizableMessage, hostName string, ) *VcenterVmGuestIdentityInfo`

NewVcenterVmGuestIdentityInfo instantiates a new VcenterVmGuestIdentityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestIdentityInfoWithDefaults

`func NewVcenterVmGuestIdentityInfoWithDefaults() *VcenterVmGuestIdentityInfo`

NewVcenterVmGuestIdentityInfoWithDefaults instantiates a new VcenterVmGuestIdentityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmGuestIdentityInfo) GetName() VcenterVmGuestOS`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmGuestIdentityInfo) GetNameOk() (*VcenterVmGuestOS, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmGuestIdentityInfo) SetName(v VcenterVmGuestOS)`

SetName sets Name field to given value.


### GetFamily

`func (o *VcenterVmGuestIdentityInfo) GetFamily() VcenterVmGuestOSFamily`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *VcenterVmGuestIdentityInfo) GetFamilyOk() (*VcenterVmGuestOSFamily, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *VcenterVmGuestIdentityInfo) SetFamily(v VcenterVmGuestOSFamily)`

SetFamily sets Family field to given value.


### GetFullName

`func (o *VcenterVmGuestIdentityInfo) GetFullName() VapiStdLocalizableMessage`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VcenterVmGuestIdentityInfo) GetFullNameOk() (*VapiStdLocalizableMessage, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VcenterVmGuestIdentityInfo) SetFullName(v VapiStdLocalizableMessage)`

SetFullName sets FullName field to given value.


### GetHostName

`func (o *VcenterVmGuestIdentityInfo) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VcenterVmGuestIdentityInfo) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VcenterVmGuestIdentityInfo) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetIpAddress

`func (o *VcenterVmGuestIdentityInfo) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VcenterVmGuestIdentityInfo) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VcenterVmGuestIdentityInfo) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *VcenterVmGuestIdentityInfo) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


