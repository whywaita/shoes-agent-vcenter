# VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]VcenterNamespaceManagementLoadBalancersServer**](VcenterNamespaceManagementLoadBalancersServer.md) | Servers is a list of the addresses for the data plane API servers used to configure HAProxy. Note: This field cannot be updated and should match existing list of servers. | 
**Username** | **string** | An administrator user name for accessing the HAProxy Data Plane API server. | 
**Password** | **string** | The password for the administrator user. | 
**CertificateAuthorityChain** | **string** | CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server. | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec

`func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec(servers []VcenterNamespaceManagementLoadBalancersServer, username string, password string, certificateAuthorityChain string, ) *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec`

NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec`

NewVcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetServers() []VcenterNamespaceManagementLoadBalancersServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetServersOk() (*[]VcenterNamespaceManagementLoadBalancersServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) SetServers(v []VcenterNamespaceManagementLoadBalancersServer)`

SetServers sets Servers field to given value.


### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigSetSpec) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


