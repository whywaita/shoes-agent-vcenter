# VcenterNamespaceManagementLoadBalancersAviInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | [**VcenterNamespaceManagementLoadBalancersServer**](VcenterNamespaceManagementLoadBalancersServer.md) |  | 
**Username** | **string** | An administrator user name for accessing the Avi Controller. | 
**CertificateAuthorityChain** | **string** | PEM-encoded CA certificate chain which is used to verify x509 certificates received from the server. | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersAviInfo

`func NewVcenterNamespaceManagementLoadBalancersAviInfo(server VcenterNamespaceManagementLoadBalancersServer, username string, certificateAuthorityChain string, ) *VcenterNamespaceManagementLoadBalancersAviInfo`

NewVcenterNamespaceManagementLoadBalancersAviInfo instantiates a new VcenterNamespaceManagementLoadBalancersAviInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersAviInfoWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersAviInfoWithDefaults() *VcenterNamespaceManagementLoadBalancersAviInfo`

NewVcenterNamespaceManagementLoadBalancersAviInfoWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersAviInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetServer() VcenterNamespaceManagementLoadBalancersServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetServerOk() (*VcenterNamespaceManagementLoadBalancersServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetServer(v VcenterNamespaceManagementLoadBalancersServer)`

SetServer sets Server field to given value.


### GetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetCertificateAuthorityChain() string`

GetCertificateAuthorityChain returns the CertificateAuthorityChain field if non-nil, zero value otherwise.

### GetCertificateAuthorityChainOk

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) GetCertificateAuthorityChainOk() (*string, bool)`

GetCertificateAuthorityChainOk returns a tuple with the CertificateAuthorityChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthorityChain

`func (o *VcenterNamespaceManagementLoadBalancersAviInfo) SetCertificateAuthorityChain(v string)`

SetCertificateAuthorityChain sets CertificateAuthorityChain field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


