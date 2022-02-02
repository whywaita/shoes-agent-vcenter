# VcenterDeploymentRemotePscSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PscHostname** | **string** | The IP address or DNS resolvable name of the remote PSC to which this configuring vCenter Server will be registered to. | 
**HttpsPort** | Pointer to **int64** | The HTTPS port of the external PSC appliance. If unset, port 443 will be used. | [optional] 
**SsoAdminPassword** | **string** | The SSO administrator account password. | 
**SslThumbprint** | Pointer to **string** | SHA1 thumbprint of the server SSL certificate will be used for verification when ssl_verify field is set to true. This field is only relevant if RemotePscSpec.ssl-verify is unset or has the value true. | [optional] 
**SslVerify** | Pointer to **bool** | SSL verification should be enabled or disabled. If RemotePscSpec.ssl-verify is true and and RemotePscSpec.ssl-thumbprint is unset, the CA certificate will be used for verification. If RemotePscSpec.ssl-verify is true and RemotePscSpec.ssl-thumbprint is set then the thumbprint will be used for verification. No verification will be performed if RemotePscSpec.ssl-verify value is set to false. If unset, RemotePscSpec.ssl-verify true will be used. | [optional] 

## Methods

### NewVcenterDeploymentRemotePscSpec

`func NewVcenterDeploymentRemotePscSpec(pscHostname string, ssoAdminPassword string, ) *VcenterDeploymentRemotePscSpec`

NewVcenterDeploymentRemotePscSpec instantiates a new VcenterDeploymentRemotePscSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentRemotePscSpecWithDefaults

`func NewVcenterDeploymentRemotePscSpecWithDefaults() *VcenterDeploymentRemotePscSpec`

NewVcenterDeploymentRemotePscSpecWithDefaults instantiates a new VcenterDeploymentRemotePscSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPscHostname

`func (o *VcenterDeploymentRemotePscSpec) GetPscHostname() string`

GetPscHostname returns the PscHostname field if non-nil, zero value otherwise.

### GetPscHostnameOk

`func (o *VcenterDeploymentRemotePscSpec) GetPscHostnameOk() (*string, bool)`

GetPscHostnameOk returns a tuple with the PscHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPscHostname

`func (o *VcenterDeploymentRemotePscSpec) SetPscHostname(v string)`

SetPscHostname sets PscHostname field to given value.


### GetHttpsPort

`func (o *VcenterDeploymentRemotePscSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentRemotePscSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentRemotePscSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentRemotePscSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSsoAdminPassword

`func (o *VcenterDeploymentRemotePscSpec) GetSsoAdminPassword() string`

GetSsoAdminPassword returns the SsoAdminPassword field if non-nil, zero value otherwise.

### GetSsoAdminPasswordOk

`func (o *VcenterDeploymentRemotePscSpec) GetSsoAdminPasswordOk() (*string, bool)`

GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminPassword

`func (o *VcenterDeploymentRemotePscSpec) SetSsoAdminPassword(v string)`

SetSsoAdminPassword sets SsoAdminPassword field to given value.


### GetSslThumbprint

`func (o *VcenterDeploymentRemotePscSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterDeploymentRemotePscSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterDeploymentRemotePscSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterDeploymentRemotePscSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetSslVerify

`func (o *VcenterDeploymentRemotePscSpec) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *VcenterDeploymentRemotePscSpec) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *VcenterDeploymentRemotePscSpec) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *VcenterDeploymentRemotePscSpec) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


