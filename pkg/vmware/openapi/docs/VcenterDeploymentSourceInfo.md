# VcenterDeploymentSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The IP address or DNS resolvable name of the source vCenter Server. | 
**Version** | **string** | Source vCenter Server version. | 
**DeploymentType** | [**VcenterDeploymentApplianceType**](VcenterDeploymentApplianceType.md) |  | 
**DeploymentSize** | [**VcenterDeploymentApplianceSize**](VcenterDeploymentApplianceSize.md) |  | 
**SsoDomainName** | **string** | The SSO domain name of the source vCenter Server. | 
**ActiveDirectoryDomain** | **string** | The domain name of the Active Directory server to which the source vCenter Server is joined. | 
**DnsServers** | **[]string** | IP addresses of the DNS servers of the Active Directory server. | 
**DataMigrationInfo** | Pointer to [**VcenterDeploymentDataMigrationInfo**](VcenterDeploymentDataMigrationInfo.md) |  | [optional] 

## Methods

### NewVcenterDeploymentSourceInfo

`func NewVcenterDeploymentSourceInfo(hostname string, version string, deploymentType VcenterDeploymentApplianceType, deploymentSize VcenterDeploymentApplianceSize, ssoDomainName string, activeDirectoryDomain string, dnsServers []string, ) *VcenterDeploymentSourceInfo`

NewVcenterDeploymentSourceInfo instantiates a new VcenterDeploymentSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentSourceInfoWithDefaults

`func NewVcenterDeploymentSourceInfoWithDefaults() *VcenterDeploymentSourceInfo`

NewVcenterDeploymentSourceInfoWithDefaults instantiates a new VcenterDeploymentSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterDeploymentSourceInfo) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterDeploymentSourceInfo) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterDeploymentSourceInfo) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetVersion

`func (o *VcenterDeploymentSourceInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterDeploymentSourceInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterDeploymentSourceInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDeploymentType

`func (o *VcenterDeploymentSourceInfo) GetDeploymentType() VcenterDeploymentApplianceType`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *VcenterDeploymentSourceInfo) GetDeploymentTypeOk() (*VcenterDeploymentApplianceType, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *VcenterDeploymentSourceInfo) SetDeploymentType(v VcenterDeploymentApplianceType)`

SetDeploymentType sets DeploymentType field to given value.


### GetDeploymentSize

`func (o *VcenterDeploymentSourceInfo) GetDeploymentSize() VcenterDeploymentApplianceSize`

GetDeploymentSize returns the DeploymentSize field if non-nil, zero value otherwise.

### GetDeploymentSizeOk

`func (o *VcenterDeploymentSourceInfo) GetDeploymentSizeOk() (*VcenterDeploymentApplianceSize, bool)`

GetDeploymentSizeOk returns a tuple with the DeploymentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentSize

`func (o *VcenterDeploymentSourceInfo) SetDeploymentSize(v VcenterDeploymentApplianceSize)`

SetDeploymentSize sets DeploymentSize field to given value.


### GetSsoDomainName

`func (o *VcenterDeploymentSourceInfo) GetSsoDomainName() string`

GetSsoDomainName returns the SsoDomainName field if non-nil, zero value otherwise.

### GetSsoDomainNameOk

`func (o *VcenterDeploymentSourceInfo) GetSsoDomainNameOk() (*string, bool)`

GetSsoDomainNameOk returns a tuple with the SsoDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoDomainName

`func (o *VcenterDeploymentSourceInfo) SetSsoDomainName(v string)`

SetSsoDomainName sets SsoDomainName field to given value.


### GetActiveDirectoryDomain

`func (o *VcenterDeploymentSourceInfo) GetActiveDirectoryDomain() string`

GetActiveDirectoryDomain returns the ActiveDirectoryDomain field if non-nil, zero value otherwise.

### GetActiveDirectoryDomainOk

`func (o *VcenterDeploymentSourceInfo) GetActiveDirectoryDomainOk() (*string, bool)`

GetActiveDirectoryDomainOk returns a tuple with the ActiveDirectoryDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryDomain

`func (o *VcenterDeploymentSourceInfo) SetActiveDirectoryDomain(v string)`

SetActiveDirectoryDomain sets ActiveDirectoryDomain field to given value.


### GetDnsServers

`func (o *VcenterDeploymentSourceInfo) GetDnsServers() []string`

GetDnsServers returns the DnsServers field if non-nil, zero value otherwise.

### GetDnsServersOk

`func (o *VcenterDeploymentSourceInfo) GetDnsServersOk() (*[]string, bool)`

GetDnsServersOk returns a tuple with the DnsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServers

`func (o *VcenterDeploymentSourceInfo) SetDnsServers(v []string)`

SetDnsServers sets DnsServers field to given value.


### GetDataMigrationInfo

`func (o *VcenterDeploymentSourceInfo) GetDataMigrationInfo() VcenterDeploymentDataMigrationInfo`

GetDataMigrationInfo returns the DataMigrationInfo field if non-nil, zero value otherwise.

### GetDataMigrationInfoOk

`func (o *VcenterDeploymentSourceInfo) GetDataMigrationInfoOk() (*VcenterDeploymentDataMigrationInfo, bool)`

GetDataMigrationInfoOk returns a tuple with the DataMigrationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMigrationInfo

`func (o *VcenterDeploymentSourceInfo) SetDataMigrationInfo(v VcenterDeploymentDataMigrationInfo)`

SetDataMigrationInfo sets DataMigrationInfo field to given value.

### HasDataMigrationInfo

`func (o *VcenterDeploymentSourceInfo) HasDataMigrationInfo() bool`

HasDataMigrationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


