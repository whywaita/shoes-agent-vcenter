# VcenterDeploymentLocationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The IP address or DNS resolvable name of the container. | 
**HttpsPort** | Pointer to **int64** | The HTTPS port of the container. If unset, port 443 will be used. | [optional] 
**SslThumbprint** | Pointer to **string** | SHA1 thumbprint of the server SSL certificate will be used for verification. This field is only relevant if LocationSpec.ssl-verify is unset or has the value true. | [optional] 
**SslVerify** | Pointer to **bool** | SSL verification should be enabled or disabled. If LocationSpec.ssl-verify is true and and LocationSpec.ssl-thumbprint is unset, the CA certificate will be used for verification. If LocationSpec.ssl-verify is true and LocationSpec.ssl-thumbprint is set then the thumbprint will be used for verification. No verification will be performed if LocationSpec.ssl-verify value is set to false. If unset, ssl_verify true will be used. | [optional] 
**Username** | **string** | The administrator account on the host. | 
**Password** | **string** | The administrator account password. | 

## Methods

### NewVcenterDeploymentLocationSpec

`func NewVcenterDeploymentLocationSpec(hostname string, username string, password string, ) *VcenterDeploymentLocationSpec`

NewVcenterDeploymentLocationSpec instantiates a new VcenterDeploymentLocationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentLocationSpecWithDefaults

`func NewVcenterDeploymentLocationSpecWithDefaults() *VcenterDeploymentLocationSpec`

NewVcenterDeploymentLocationSpecWithDefaults instantiates a new VcenterDeploymentLocationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterDeploymentLocationSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterDeploymentLocationSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterDeploymentLocationSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetHttpsPort

`func (o *VcenterDeploymentLocationSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentLocationSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentLocationSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentLocationSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSslThumbprint

`func (o *VcenterDeploymentLocationSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterDeploymentLocationSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterDeploymentLocationSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterDeploymentLocationSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetSslVerify

`func (o *VcenterDeploymentLocationSpec) GetSslVerify() bool`

GetSslVerify returns the SslVerify field if non-nil, zero value otherwise.

### GetSslVerifyOk

`func (o *VcenterDeploymentLocationSpec) GetSslVerifyOk() (*bool, bool)`

GetSslVerifyOk returns a tuple with the SslVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerify

`func (o *VcenterDeploymentLocationSpec) SetSslVerify(v bool)`

SetSslVerify sets SslVerify field to given value.

### HasSslVerify

`func (o *VcenterDeploymentLocationSpec) HasSslVerify() bool`

HasSslVerify returns a boolean if a field has been set.

### GetUsername

`func (o *VcenterDeploymentLocationSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterDeploymentLocationSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterDeploymentLocationSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterDeploymentLocationSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterDeploymentLocationSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterDeploymentLocationSpec) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


