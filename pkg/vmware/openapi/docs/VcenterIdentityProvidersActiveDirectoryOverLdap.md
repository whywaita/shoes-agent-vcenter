# VcenterIdentityProvidersActiveDirectoryOverLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserName** | **string** | User name to connect to the active directory server. | 
**Password** | **string** | Password to connect to the active directory server. | 
**UsersBaseDn** | **string** | Base distinguished name for users | 
**GroupsBaseDn** | **string** | Base distinguished name for groups | 
**ServerEndpoints** | **[]string** | Active directory server endpoints. At least one active directory server endpoint must be set. | 
**CertChain** | Pointer to [**VcenterCertificateManagementX509CertChain**](VcenterCertificateManagementX509CertChain.md) |  | [optional] 

## Methods

### NewVcenterIdentityProvidersActiveDirectoryOverLdap

`func NewVcenterIdentityProvidersActiveDirectoryOverLdap(userName string, password string, usersBaseDn string, groupsBaseDn string, serverEndpoints []string, ) *VcenterIdentityProvidersActiveDirectoryOverLdap`

NewVcenterIdentityProvidersActiveDirectoryOverLdap instantiates a new VcenterIdentityProvidersActiveDirectoryOverLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersActiveDirectoryOverLdapWithDefaults

`func NewVcenterIdentityProvidersActiveDirectoryOverLdapWithDefaults() *VcenterIdentityProvidersActiveDirectoryOverLdap`

NewVcenterIdentityProvidersActiveDirectoryOverLdapWithDefaults instantiates a new VcenterIdentityProvidersActiveDirectoryOverLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserName

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetPassword

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsersBaseDn

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetUsersBaseDn() string`

GetUsersBaseDn returns the UsersBaseDn field if non-nil, zero value otherwise.

### GetUsersBaseDnOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetUsersBaseDnOk() (*string, bool)`

GetUsersBaseDnOk returns a tuple with the UsersBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersBaseDn

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetUsersBaseDn(v string)`

SetUsersBaseDn sets UsersBaseDn field to given value.


### GetGroupsBaseDn

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetGroupsBaseDn() string`

GetGroupsBaseDn returns the GroupsBaseDn field if non-nil, zero value otherwise.

### GetGroupsBaseDnOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetGroupsBaseDnOk() (*string, bool)`

GetGroupsBaseDnOk returns a tuple with the GroupsBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsBaseDn

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetGroupsBaseDn(v string)`

SetGroupsBaseDn sets GroupsBaseDn field to given value.


### GetServerEndpoints

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetServerEndpoints() []string`

GetServerEndpoints returns the ServerEndpoints field if non-nil, zero value otherwise.

### GetServerEndpointsOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetServerEndpointsOk() (*[]string, bool)`

GetServerEndpointsOk returns a tuple with the ServerEndpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEndpoints

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetServerEndpoints(v []string)`

SetServerEndpoints sets ServerEndpoints field to given value.


### GetCertChain

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetCertChain() VcenterCertificateManagementX509CertChain`

GetCertChain returns the CertChain field if non-nil, zero value otherwise.

### GetCertChainOk

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) GetCertChainOk() (*VcenterCertificateManagementX509CertChain, bool)`

GetCertChainOk returns a tuple with the CertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChain

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) SetCertChain(v VcenterCertificateManagementX509CertChain)`

SetCertChain sets CertChain field to given value.

### HasCertChain

`func (o *VcenterIdentityProvidersActiveDirectoryOverLdap) HasCertChain() bool`

HasCertChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


