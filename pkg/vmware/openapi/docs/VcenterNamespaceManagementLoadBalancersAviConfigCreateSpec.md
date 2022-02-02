# VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**VcenterNamespaceManagementLoadBalancersServer**](VcenterNamespaceManagementLoadBalancersServer.md) |  | 
**Username** | **string** | An administrator user name for accessing the Avi Controller. | 
**Password** | **string** | The password for the administrator user. | 
**CertificateAuthorityChain** | **string** | CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server. | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec

`func NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec(server VcenterNamespaceManagementLoadBalancersServer, username string, password string, certificateAuthorityChain string, ) *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec`

NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpec instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec`

NewVcenterNamespaceManagementLoadBalancersAviConfigCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetServer() VcenterNamespaceManagementLoadBalancersServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetServerOk() (*VcenterNamespaceManagementLoadBalancersServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetServer(v VcenterNamespaceManagementLoadBalancersServer)`

SetServer sets Server field to given value.


### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigCreateSpec) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


