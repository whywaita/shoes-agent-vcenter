# VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType**](VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType.md) |  | 
**Service** | Pointer to **string** | The service&#39;s unique ID. This field is optional and it is only relevant when the value of Services.CreateSpec.type is SERVICE. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. | [optional] 
**TrustAuthorityCluster** | Pointer to **string** | The attestation cluster&#39;s unique ID. This field is optional and it is only relevant when the value of Services.CreateSpec.type is CLUSTER. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec(type_ VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType, ) *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetType() VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetTypeOk() (*VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) SetType(v VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpecSourceType)`

SetType sets Type field to given value.


### GetService

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.

### HasTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesCreateSpec) HasTrustAuthorityCluster() bool`

HasTrustAuthorityCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


