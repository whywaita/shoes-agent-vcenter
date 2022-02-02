# VcenterNamespaceManagementLoadBalancersHAProxyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]VcenterNamespaceManagementLoadBalancersServer**](VcenterNamespaceManagementLoadBalancersServer.md) | A list of the addresses for the DataPlane API servers used to configure HAProxy. | 
**Username** | **string** | An administrator user name for accessing the HAProxy Data Plane API server. | 
**CertificateAuthorityChain** | **string** | PEM-encoded CA certificate chain which is used to verify x509 certificates received from the server. | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersHAProxyInfo

`func NewVcenterNamespaceManagementLoadBalancersHAProxyInfo(servers []VcenterNamespaceManagementLoadBalancersServer, username string, certificateAuthorityChain string, ) *VcenterNamespaceManagementLoadBalancersHAProxyInfo`

NewVcenterNamespaceManagementLoadBalancersHAProxyInfo instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersHAProxyInfoWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersHAProxyInfoWithDefaults() *VcenterNamespaceManagementLoadBalancersHAProxyInfo`

NewVcenterNamespaceManagementLoadBalancersHAProxyInfoWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetServers() []VcenterNamespaceManagementLoadBalancersServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetServersOk() (*[]VcenterNamespaceManagementLoadBalancersServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) SetServers(v []VcenterNamespaceManagementLoadBalancersServer)`

SetServers sets Servers field to given value.


### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyInfo) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


