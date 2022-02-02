# VcenterDeploymentReplicatedPscSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoSiteName** | Pointer to **string** | The SSO sitename that will be used in PSC replication. If unset, default-first-site will be used. | [optional] 
**PartnerHostname** | **string** | The IP address or DNS resolvable name of the partner PSC appliance. | 
**HttpsPort** | Pointer to **int64** | The HTTPS port of the external PSC appliance. If unset, port 443 will be used. | [optional] 
**SsoAdminPassword** | **string** | The SSO administrator account password. | 
**SslThumbprint** | Pointer to **string** | SHA1 thumbprint of the server SSL certificate will be used for verification. This field is only relevant if ReplicatedSpec.ssl-verify is unset or has the value true. | [optional] 
**SslVerify** | Pointer to **bool** | SSL verification should be enabled or disabled. If unset, ssl_verify true will be used. | [optional] 

## Methods

### NewVcenterDeploymentReplicatedPscSpec

`func NewVcenterDeploymentReplicatedPscSpec(partnerHostname string, ssoAdminPassword string, ) *VcenterDeploymentReplicatedPscSpec`

NewVcenterDeploymentReplicatedPscSpec instantiates a new VcenterDeploymentReplicatedPscSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentReplicatedPscSpecWithDefaults

`func NewVcenterDeploymentReplicatedPscSpecWithDefaults() *VcenterDeploymentReplicatedPscSpec`

NewVcenterDeploymentReplicatedPscSpecWithDefaults instantiates a new VcenterDeploymentReplicatedPscSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoSiteName

`func (o *VcenterDeploymentReplicatedPscSpec) GetSsoSiteName() string`

GetSsoSiteName returns the SsoSiteName field if non-nil, zero value otherwise.

### GetSsoSiteNameOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetSsoSiteNameOk() (*string, bool)`

GetSsoSiteNameOk returns a tuple with the SsoSiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoSiteName

`func (o *VcenterDeploymentReplicatedPscSpec) SetSsoSiteName(v string)`

SetSsoSiteName sets SsoSiteName field to given value.

### HasSsoSiteName

`func (o *VcenterDeploymentReplicatedPscSpec) HasSsoSiteName() bool`

HasSsoSiteName returns a boolean if a field has been set.

### GetPartnerHostname

`func (o *VcenterDeploymentReplicatedPscSpec) GetPartnerHostname() string`

GetPartnerHostname returns the PartnerHostname field if non-nil, zero value otherwise.

### GetPartnerHostnameOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetPartnerHostnameOk() (*string, bool)`

GetPartnerHostnameOk returns a tuple with the PartnerHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerHostname

`func (o *VcenterDeploymentReplicatedPscSpec) SetPartnerHostname(v string)`

SetPartnerHostname sets PartnerHostname field to given value.


### GetHttpsPort

`func (o *VcenterDeploymentReplicatedPscSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentReplicatedPscSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentReplicatedPscSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSsoAdminPassword

`func (o *VcenterDeploymentReplicatedPscSpec) GetSsoAdminPassword() string`

GetSsoAdminPassword returns the SsoAdminPassword field if non-nil, zero value otherwise.

### GetSsoAdminPasswordOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetSsoAdminPasswordOk() (*string, bool)`

GetSsoAdminPasswordOk returns a tuple with the SsoAdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoAdminPassword

`func (o *VcenterDeploymentReplicatedPscSpec) SetSsoAdminPassword(v string)`

SetSsoAdminPassword sets SsoAdminPassword field to given value.


### GetSslThumbprint

`func (o *VcenterDeploymentReplicatedPscSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterDeploymentReplicatedPscSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterDeploymentReplicatedPscSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetSslVerify

`func (o *VcenterDeploymentReplicatedPscSpec) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *VcenterDeploymentReplicatedPscSpec) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *VcenterDeploymentReplicatedPscSpec) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *VcenterDeploymentReplicatedPscSpec) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


