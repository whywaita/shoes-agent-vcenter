# VcenterDeploymentStandalonePscSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoSiteName** | Pointer to **string** | The SSO site name used for this PSC. If unset, default-first-site will be used. | [optional] 
**SsoAdminPassword** | **string** | The SSO administrator account password. | 
**SsoDomainName** | Pointer to **string** | The SSO domain name to be used to configure this appliance. If unset, vsphere.local will be used. | [optional] 

## Methods

### NewVcenterDeploymentStandalonePscSpec

`func NewVcenterDeploymentStandalonePscSpec(ssoAdminPassword string, ) *VcenterDeploymentStandalonePscSpec`

NewVcenterDeploymentStandalonePscSpec instantiates a new VcenterDeploymentStandalonePscSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentStandalonePscSpecWithDefaults

`func NewVcenterDeploymentStandalonePscSpecWithDefaults() *VcenterDeploymentStandalonePscSpec`

NewVcenterDeploymentStandalonePscSpecWithDefaults instantiates a new VcenterDeploymentStandalonePscSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoSiteName

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoSiteName() string`

GetSsoSiteName returns the SsoSiteName field if non-nil, zero value otherwise.

### GetSsoSiteNameOk

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoSiteNameOk() (*string, bool)`

GetSsoSiteNameOk returns a tuple with the SsoSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSiteName

`func (o *VcenterDeploymentStandalonePscSpec) SetSsoSiteName(v string)`

SetSsoSiteName sets SsoSiteName field to given value.

### HasSsoSiteName

`func (o *VcenterDeploymentStandalonePscSpec) HasSsoSiteName() bool`

HasSsoSiteName returns a boolean if a field has been set.

### GetSsoAdminPassword

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoAdminPassword() string`

GetSsoAdminPassword returns the SsoAdminPassword field if non-nil, zero value otherwise.

### GetSsoAdminPasswordOk

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoAdminPasswordOk() (*string, bool)`

GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminPassword

`func (o *VcenterDeploymentStandalonePscSpec) SetSsoAdminPassword(v string)`

SetSsoAdminPassword sets SsoAdminPassword field to given value.


### GetSsoDomainName

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoDomainName() string`

GetSsoDomainName returns the SsoDomainName field if non-nil, zero value otherwise.

### GetSsoDomainNameOk

`func (o *VcenterDeploymentStandalonePscSpec) GetSsoDomainNameOk() (*string, bool)`

GetSsoDomainNameOk returns a tuple with the SsoDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomainName

`func (o *VcenterDeploymentStandalonePscSpec) SetSsoDomainName(v string)`

SetSsoDomainName sets SsoDomainName field to given value.

### HasSsoDomainName

`func (o *VcenterDeploymentStandalonePscSpec) HasSsoDomainName() bool`

HasSsoDomainName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


