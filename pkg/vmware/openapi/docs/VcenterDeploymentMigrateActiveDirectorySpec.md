# VcenterDeploymentMigrateActiveDirectorySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain name of the Active Directory server to which the migrated vCenter Server appliance should be joined. | 
**Username** | **string** | Active Directory user that has permission to join the Active Directory after the vCenter Server is migrated to appliance. | 
**Password** | **string** | Active Directory user password that has permission to join the Active Directory after the vCenter Server is migrated to appliance. | 

## Methods

### NewVcenterDeploymentMigrateActiveDirectorySpec

`func NewVcenterDeploymentMigrateActiveDirectorySpec(domain string, username string, password string, ) *VcenterDeploymentMigrateActiveDirectorySpec`

NewVcenterDeploymentMigrateActiveDirectorySpec instantiates a new VcenterDeploymentMigrateActiveDirectorySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentMigrateActiveDirectorySpecWithDefaults

`func NewVcenterDeploymentMigrateActiveDirectorySpecWithDefaults() *VcenterDeploymentMigrateActiveDirectorySpec`

NewVcenterDeploymentMigrateActiveDirectorySpecWithDefaults instantiates a new VcenterDeploymentMigrateActiveDirectorySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUsername

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterDeploymentMigrateActiveDirectorySpec) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


