# VcenterVchaConnectionSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | IP Address or DNS of the vCenter. | 
**Port** | Pointer to **int64** | Port number. If unset, port 443 will be used. | [optional] 
**SslThumbprint** | Pointer to **string** | SHA1 hash of the server SSL certificate. If unset, empty ssl thumbprint is assumed. | [optional] 
**Username** | Pointer to **string** | Username to access the server. This field is currently required. If unset, an error is returned. In the future, if this field is unset, the system will attempt to identify the user. If a user cannot be identified, then the requested operation will fail. | [optional] 
**Password** | Pointer to **string** | Password for the specified user. This field is currently required. If unset, an empty password is assumed. In the future, if this field is unset, the system will attempt to authenticate the user. If a user cannot be identified, then the requested operation will fail. | [optional] 

## Methods

### NewVcenterVchaConnectionSpec

`func NewVcenterVchaConnectionSpec(hostname string, ) *VcenterVchaConnectionSpec`

NewVcenterVchaConnectionSpec instantiates a new VcenterVchaConnectionSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaConnectionSpecWithDefaults

`func NewVcenterVchaConnectionSpecWithDefaults() *VcenterVchaConnectionSpec`

NewVcenterVchaConnectionSpecWithDefaults instantiates a new VcenterVchaConnectionSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterVchaConnectionSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterVchaConnectionSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterVchaConnectionSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *VcenterVchaConnectionSpec) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterVchaConnectionSpec) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterVchaConnectionSpec) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VcenterVchaConnectionSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSslThumbprint

`func (o *VcenterVchaConnectionSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterVchaConnectionSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterVchaConnectionSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterVchaConnectionSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetUsername

`func (o *VcenterVchaConnectionSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterVchaConnectionSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterVchaConnectionSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VcenterVchaConnectionSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *VcenterVchaConnectionSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterVchaConnectionSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterVchaConnectionSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterVchaConnectionSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


