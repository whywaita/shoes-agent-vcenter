# VcenterDeploymentStandaloneSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoAdminPassword** | **string** | The SSO administrator account password. | 
**SsoDomainName** | Pointer to **string** | The SSO domain name to be used to configure this appliance. If unset, vsphere.local will be used. | [optional] 

## Methods

### NewVcenterDeploymentStandaloneSpec

`func NewVcenterDeploymentStandaloneSpec(ssoAdminPassword string, ) *VcenterDeploymentStandaloneSpec`

NewVcenterDeploymentStandaloneSpec instantiates a new VcenterDeploymentStandaloneSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentStandaloneSpecWithDefaults

`func NewVcenterDeploymentStandaloneSpecWithDefaults() *VcenterDeploymentStandaloneSpec`

NewVcenterDeploymentStandaloneSpecWithDefaults instantiates a new VcenterDeploymentStandaloneSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoAdminPassword

`func (o *VcenterDeploymentStandaloneSpec) GetSsoAdminPassword() string`

GetSsoAdminPassword returns the SsoAdminPassword field if non-nil, zero value otherwise.

### GetSsoAdminPasswordOk

`func (o *VcenterDeploymentStandaloneSpec) GetSsoAdminPasswordOk() (*string, bool)`

GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminPassword

`func (o *VcenterDeploymentStandaloneSpec) SetSsoAdminPassword(v string)`

SetSsoAdminPassword sets SsoAdminPassword field to given value.


### GetSsoDomainName

`func (o *VcenterDeploymentStandaloneSpec) GetSsoDomainName() string`

GetSsoDomainName returns the SsoDomainName field if non-nil, zero value otherwise.

### GetSsoDomainNameOk

`func (o *VcenterDeploymentStandaloneSpec) GetSsoDomainNameOk() (*string, bool)`

GetSsoDomainNameOk returns a tuple with the SsoDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomainName

`func (o *VcenterDeploymentStandaloneSpec) SetSsoDomainName(v string)`

SetSsoDomainName sets SsoDomainName field to given value.

### HasSsoDomainName

`func (o *VcenterDeploymentStandaloneSpec) HasSsoDomainName() bool`

HasSsoDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


