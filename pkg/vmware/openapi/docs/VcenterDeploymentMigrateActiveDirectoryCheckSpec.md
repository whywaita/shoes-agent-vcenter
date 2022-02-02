# VcenterDeploymentMigrateActiveDirectoryCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsServers** | **[]string** | IP addresses of the DNS servers of the Active Directory server. | 
**Domain** | **string** | The domain name of the Active Directory server to which the migrated vCenter Server appliance should be joined. | 
**Username** | **string** | Active Directory user that has permission to join the Active Directory after the vCenter Server is migrated to appliance. | 
**Password** | **string** | Active Directory user password that has permission to join the Active Directory after the vCenter Server is migrated to appliance. | 

## Methods

### NewVcenterDeploymentMigrateActiveDirectoryCheckSpec

`func NewVcenterDeploymentMigrateActiveDirectoryCheckSpec(dnsServers []string, domain string, username string, password string, ) *VcenterDeploymentMigrateActiveDirectoryCheckSpec`

NewVcenterDeploymentMigrateActiveDirectoryCheckSpec instantiates a new VcenterDeploymentMigrateActiveDirectoryCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentMigrateActiveDirectoryCheckSpecWithDefaults

`func NewVcenterDeploymentMigrateActiveDirectoryCheckSpecWithDefaults() *VcenterDeploymentMigrateActiveDirectoryCheckSpec`

NewVcenterDeploymentMigrateActiveDirectoryCheckSpecWithDefaults instantiates a new VcenterDeploymentMigrateActiveDirectoryCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsServers

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### GetDomain

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetUsername

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *VcenterDeploymentMigrateActiveDirectoryCheckSpec) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


