# VcenterDeploymentMigrateMigrationAssistantSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpsPort** | Pointer to **int64** | The HTTPS port being used by Migration Assistant. If unset, port 9123 will be used. | [optional] 
**SslThumbprint** | **string** | SHA1 thumbprint of the Migration Assistant SSL certificate that will be used for verification. | 

## Methods

### NewVcenterDeploymentMigrateMigrationAssistantSpec

`func NewVcenterDeploymentMigrateMigrationAssistantSpec(sslThumbprint string, ) *VcenterDeploymentMigrateMigrationAssistantSpec`

NewVcenterDeploymentMigrateMigrationAssistantSpec instantiates a new VcenterDeploymentMigrateMigrationAssistantSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentMigrateMigrationAssistantSpecWithDefaults

`func NewVcenterDeploymentMigrateMigrationAssistantSpecWithDefaults() *VcenterDeploymentMigrateMigrationAssistantSpec`

NewVcenterDeploymentMigrateMigrationAssistantSpecWithDefaults instantiates a new VcenterDeploymentMigrateMigrationAssistantSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpsPort

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.

### GetSslThumbprint

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) GetSslThumbprint() string`

GetSslThumbprint returns the SslThumbprint field if non-nil, zero value otherwise.

### GetSslThumbprintOk

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) GetSslThumbprintOk() (*string, bool)`

GetSslThumbprintOk returns a tuple with the SslThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslThumbprint

`func (o *VcenterDeploymentMigrateMigrationAssistantSpec) SetSslThumbprint(v string)`

SetSslThumbprint sets SslThumbprint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


