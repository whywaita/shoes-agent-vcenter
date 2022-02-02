# VcenterHostCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The IP address or DNS resolvable name of the host. | 
**Port** | Pointer to **int64** | The port of the host. If unset, port 443 will be used. | [optional] 
**UserName** | **string** | The administrator account on the host. | 
**Password** | **string** | The password for the administrator account on the host. | 
**Folder** | Pointer to **string** | Host and cluster folder in which the new standalone host should be created. This field is currently required. In the future, if this field is unset, the system will attempt to choose a suitable folder for the host; if a folder cannot be chosen, the host creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 
**ThumbprintVerification** | [**VcenterHostCreateSpecThumbprintVerification**](VcenterHostCreateSpecThumbprintVerification.md) |  | 
**Thumbprint** | Pointer to **string** | The thumbprint of the SSL certificate, which the host is expected to have. The thumbprint is always computed using the SHA1 hash and is the string representation of that hash in the format: xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx:xx where, &#39;x&#39; represents a hexadecimal digit. This field is optional and it is only relevant when the value of Host.CreateSpec.thumbprint-verification is THUMBPRINT. | [optional] 
**ForceAdd** | Pointer to **bool** | Whether host should be added to the vCenter Server even if it is being managed by another vCenter Server. The original vCenterServer loses connection to the host. If unset, forceAdd is default to false. | [optional] 

## Methods

### NewVcenterHostCreateSpec

`func NewVcenterHostCreateSpec(hostname string, userName string, password string, thumbprintVerification VcenterHostCreateSpecThumbprintVerification, ) *VcenterHostCreateSpec`

NewVcenterHostCreateSpec instantiates a new VcenterHostCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHostCreateSpecWithDefaults

`func NewVcenterHostCreateSpecWithDefaults() *VcenterHostCreateSpec`

NewVcenterHostCreateSpecWithDefaults instantiates a new VcenterHostCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterHostCreateSpec) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterHostCreateSpec) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterHostCreateSpec) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetPort

`func (o *VcenterHostCreateSpec) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterHostCreateSpec) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterHostCreateSpec) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *VcenterHostCreateSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUserName

`func (o *VcenterHostCreateSpec) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VcenterHostCreateSpec) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VcenterHostCreateSpec) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *VcenterHostCreateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterHostCreateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterHostCreateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetFolder

`func (o *VcenterHostCreateSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterHostCreateSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterHostCreateSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterHostCreateSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetThumbprintVerification

`func (o *VcenterHostCreateSpec) GetThumbprintVerification() VcenterHostCreateSpecThumbprintVerification`

GetThumbprintVerification returns the ThumbprintVerification field if non-nil, zero value otherwise.

### GetThumbprintVerificationOk

`func (o *VcenterHostCreateSpec) GetThumbprintVerificationOk() (*VcenterHostCreateSpecThumbprintVerification, bool)`

GetThumbprintVerificationOk returns a tuple with the ThumbprintVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintVerification

`func (o *VcenterHostCreateSpec) SetThumbprintVerification(v VcenterHostCreateSpecThumbprintVerification)`

SetThumbprintVerification sets ThumbprintVerification field to given value.


### GetThumbprint

`func (o *VcenterHostCreateSpec) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *VcenterHostCreateSpec) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *VcenterHostCreateSpec) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *VcenterHostCreateSpec) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetForceAdd

`func (o *VcenterHostCreateSpec) GetForceAdd() bool`

GetForceAdd returns the ForceAdd field if non-nil, zero value otherwise.

### GetForceAddOk

`func (o *VcenterHostCreateSpec) GetForceAddOk() (*bool, bool)`

GetForceAddOk returns a tuple with the ForceAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceAdd

`func (o *VcenterHostCreateSpec) SetForceAdd(v bool)`

SetForceAdd sets ForceAdd field to given value.

### HasForceAdd

`func (o *VcenterHostCreateSpec) HasForceAdd() bool`

HasForceAdd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


