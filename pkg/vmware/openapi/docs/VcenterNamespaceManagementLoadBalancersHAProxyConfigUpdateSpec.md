# VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | An administrator user name for accessing the HAProxy Data Plane API server. If unset, the existing username will not be modified. | [optional] 
**Password** | Pointer to **string** | The password for the administrator user. If unset, the existing password will not be modified. | [optional] 
**CertificateAuthorityChain** | Pointer to **string** | CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server. If unset, the existing PEM-encoded CA chain will not be modified. | [optional] 

## Methods

### NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec

`func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec() *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec`

NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec`

NewVcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.

### HasCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersHAProxyConfigUpdateSpec) HasCertificateAuthorityChain() bool`

HasCertificateAuthorityChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


