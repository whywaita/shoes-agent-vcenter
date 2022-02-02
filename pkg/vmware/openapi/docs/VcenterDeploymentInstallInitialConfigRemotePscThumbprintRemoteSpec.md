# VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The IP address or DNS resolvable name of the remote PSC. | 
**HttpsPort** | Pointer to **int64** | The HTTPS port of the remote PSC. If unset, port 443 will be used. | [optional] 

## Methods

### NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec

`func NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec(address string, ) *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec`

NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec instantiates a new VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpecWithDefaults

`func NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpecWithDefaults() *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec`

NewVcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpecWithDefaults instantiates a new VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetHttpsPort

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) GetHttpsPort() int64`

GetHttpsPort returns the HttpsPort field if non-nil, zero value otherwise.

### GetHttpsPortOk

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) GetHttpsPortOk() (*int64, bool)`

GetHttpsPortOk returns a tuple with the HttpsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPort

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) SetHttpsPort(v int64)`

SetHttpsPort sets HttpsPort field to given value.

### HasHttpsPort

`func (o *VcenterDeploymentInstallInitialConfigRemotePscThumbprintRemoteSpec) HasHttpsPort() bool`

HasHttpsPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


