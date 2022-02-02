# VcenterNamespaceManagementLoadBalancersAviConfigSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**VcenterNamespaceManagementLoadBalancersServer**](VcenterNamespaceManagementLoadBalancersServer.md) |  | 
**Username** | **string** | An administrator user name for accessing the Avi Controller. | 
**Password** | **string** | The password for the administrator user. | 
**CertificateAuthorityChain** | **string** | CertificateAuthorityChain contains PEM-encoded CA chain which is used to verify x509 certificates received from the server. | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpec

`func NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpec(server VcenterNamespaceManagementLoadBalancersServer, username string, password string, certificateAuthorityChain string, ) *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec`

NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpec instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpecWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpecWithDefaults() *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec`

NewVcenterNamespaceManagementLoadBalancersAviConfigSetSpecWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersAviConfigSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetServer() VcenterNamespaceManagementLoadBalancersServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetServerOk() (*VcenterNamespaceManagementLoadBalancersServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) SetServer(v VcenterNamespaceManagementLoadBalancersServer)`

SetServer sets Server field to given value.


### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviConfigSetSpec) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


