# VcenterHvcLinksCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PscHostname** | **string** | The PSC hostname for the domain to be linked. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**Port** | Pointer to **string** | The HTTPS port of the PSC to be linked. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | [optional] 
**DomainName** | **string** | The domain to which the PSC belongs. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**Username** | **string** | The administrator username of the PSC. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**Password** | **string** | The administrator password of the PSC. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**SslThumbprint** | Pointer to **string** | The ssl thumbprint of the server. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | [optional] 
**AdminGroups** | Pointer to **[]string** | List of groups to be added to enable administrator access to. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | [optional] 

## Methods

### NewVcenterHvcLinksCreateSpec

`func NewVcenterHvcLinksCreateSpec(pscHostname string, domainName string, username string, password string, ) *VcenterHvcLinksCreateSpec`

NewVcenterHvcLinksCreateSpec instantiates a new VcenterHvcLinksCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHvcLinksCreateSpecWithDefaults

`func NewVcenterHvcLinksCreateSpecWithDefaults() *VcenterHvcLinksCreateSpec`

NewVcenterHvcLinksCreateSpecWithDefaults instantiates a new VcenterHvcLinksCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPscHostname

`func (o *VcenterHvcLinksCreateSpec) GetPscHostname() string`

GetPscHostname returns the PscHostname field if non-nil, zero value otherwise.

### GetPscHostnameOk

`func (o *VcenterHvcLinksCreateSpec) GetPscHostnameOk() (*string, bool)`

GetPscHostnameOk returns a tuple with the PscHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPscHostname

`func (o *VcenterHvcLinksCreateSpec) SetPscHostname(v string)`

SetPscHostname sets PscHostname field to given value.


### GetPort

`func (o *VcenterHvcLinksCreateSpec) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterHvcLinksCreateSpec) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterHvcLinksCreateSpec) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *VcenterHvcLinksCreateSpec) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDomainName

`func (o *VcenterHvcLinksCreateSpec) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VcenterHvcLinksCreateSpec) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VcenterHvcLinksCreateSpec) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetUsername

`func (o *VcenterHvcLinksCreateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterHvcLinksCreateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterHvcLinksCreateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterHvcLinksCreateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterHvcLinksCreateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterHvcLinksCreateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSslThumbprint

`func (o *VcenterHvcLinksCreateSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterHvcLinksCreateSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterHvcLinksCreateSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.

### HasSslThumbprint

`func (o *VcenterHvcLinksCreateSpec) HasSslThumbprint() bool`

HasSslThumbprint returns a boolean if a field has been set.

### GetAdminGroups

`func (o *VcenterHvcLinksCreateSpec) GetAdminGroups() []string`

GetAdminGroups returns the AdminGroups field if non-nil, zero value otherwise.

### GetAdminGroupsOk

`func (o *VcenterHvcLinksCreateSpec) GetAdminGroupsOk() (*[]string, bool)`

GetAdminGroupsOk returns a tuple with the AdminGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminGroups

`func (o *VcenterHvcLinksCreateSpec) SetAdminGroups(v []string)`

SetAdminGroups sets AdminGroups field to given value.

### HasAdminGroups

`func (o *VcenterHvcLinksCreateSpec) HasAdminGroups() bool`

HasAdminGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


