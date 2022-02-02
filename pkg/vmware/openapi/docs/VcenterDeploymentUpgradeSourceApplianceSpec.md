# VcenterDeploymentUpgradeSourceApplianceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The IP address or DNS resolvable name of the source appliance. | 
**HttpsPort** | Pointer to **int64** | The HTTPS port of the source appliance. If unset, port 443 will be used. | [optional] 
**SslThumbprint** | Pointer to **string** | SHA1 thumbprint of the server SSL certificate will be used for verification. This field is only relevant if Upgrade.SourceApplianceSpec.ssl-verify is unset or has the value true. | [optional] 
**SslVerify** | Pointer to **bool** | SSL verification should be enabled or disabled for the source appliance validations. By default it is enabled and will use SSL certificate for verification. If thumbprint is provided, will use thumbprint for the verification. If unset, ssl_verify true will be used. | [optional] 
**SsoAdminUsername** | **string** | The SSO administrator account on the source appliance. | 
**SsoAdminPassword** | **string** | The SSO administrator account password. | 
**RootPassword** | **string** | The password of the root user on the source appliance. | 
**SshVerify** | Pointer to **bool** | Appliance SSH verification should be enabled or disabled. By default it is disabled and will not use any verification. If thumbprint is provided, thumbprint verification will be performed. If unset, ssh_verify true will be used. | [optional] 
**SshThumbprint** | Pointer to **string** | MD5 thumbprint of the server SSH key will be used for verification. This field is only relevant if Upgrade.SourceApplianceSpec.ssh-verify is unset or has the value true. | [optional] 

## Methods

### NewVcenterDeploymentUpgradeSourceApplianceSpec

`func NewVcenterDeploymentUpgradeSourceApplianceSpec(hostname string, ssoAdminUsername string, ssoAdminPassword string, rootPassword string, ) *VcenterDeploymentUpgradeSourceApplianceSpec`

NewVcenterDeploymentUpgradeSourceApplianceSpec instantiates a new VcenterDeploymentUpgradeSourceApplianceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentUpgradeSourceApplianceSpecWithDefaults

`func NewVcenterDeploymentUpgradeSourceApplianceSpecWithDefaults() *VcenterDeploymentUpgradeSourceApplianceSpec`

NewVcenterDeploymentUpgradeSourceApplianceSpecWithDefaults instantiates a new VcenterDeploymentUpgradeSourceApplianceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHttpsPort

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSslThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetSslVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.

### GetSsoAdminUsername

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSsoAdminUsername() string`

GetSsoAdminUsername returns the SsoAdminUsername field if non-nil, zero value otherwise.

### GetSsoAdminUsernameOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSsoAdminUsernameOk() (*string, bool)`

GetSsoAdminUsernameOk returns a tuple with the SsoAdminUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminUsername

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSsoAdminUsername(v string)`

SetSsoAdminUsername sets SsoAdminUsername field to given value.


### GetSsoAdminPassword

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSsoAdminPassword() string`

GetSsoAdminPassword returns the SsoAdminPassword field if non-nil, zero value otherwise.

### GetSsoAdminPasswordOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSsoAdminPasswordOk() (*string, bool)`

GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminPassword

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSsoAdminPassword(v string)`

SetSsoAdminPassword sets SsoAdminPassword field to given value.


### GetRootPassword

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.


### GetSshVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSshVerify() bool`

GetSshVerify returns the SshVerify field if non-nil, zero value otherwise.

### GetSshVerifyOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSshVerifyOk() (*bool, bool)`

GetSshVerifyOk returns a tuple with the SshVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSshVerify(v bool)`

SetSshVerify sets SshVerify field to given value.

### HasSshVerify

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) HasSshVerify() bool`

HasSshVerify returns a boolean if a field has been set.

### GetSshThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSshThumbprint() string`

GetSshThumbprint returns the SshThumbprint field if non-nil, zero value otherwise.

### GetSshThumbprintOk

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) GetSshThumbprintOk() (*string, bool)`

GetSshThumbprintOk returns a tuple with the SshThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) SetSshThumbprint(v string)`

SetSshThumbprint sets SshThumbprint field to given value.

### HasSshThumbprint

`func (o *VcenterDeploymentUpgradeSourceApplianceSpec) HasSshThumbprint() bool`

HasSshThumbprint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


