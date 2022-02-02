# VcenterNamespaceManagementHostsConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamespacesSupported** | **bool** | True if vSphere Namespace feature is supported in this VC. | 
**NamespacesLicensed** | **bool** | True if vSphere Namespace feature is licensed on any hosts in this VC. | 

## Methods

### NewVcenterNamespaceManagementHostsConfigInfo

`func NewVcenterNamespaceManagementHostsConfigInfo(namespacesSupported bool, namespacesLicensed bool, ) *VcenterNamespaceManagementHostsConfigInfo`

NewVcenterNamespaceManagementHostsConfigInfo instantiates a new VcenterNamespaceManagementHostsConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementHostsConfigInfoWithDefaults

`func NewVcenterNamespaceManagementHostsConfigInfoWithDefaults() *VcenterNamespaceManagementHostsConfigInfo`

NewVcenterNamespaceManagementHostsConfigInfoWithDefaults instantiates a new VcenterNamespaceManagementHostsConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespacesSupported

`func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesSupported() bool`

GetNamespacesSupported returns the NamespacesSupported field if non-nil, zero value otherwise.

### GetNamespacesSupportedOk

`func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesSupportedOk() (*bool, bool)`

GetNamespacesSupportedOk returns a tuple with the NamespacesSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesSupported

`func (o *VcenterNamespaceManagementHostsConfigInfo) SetNamespacesSupported(v bool)`

SetNamespacesSupported sets NamespacesSupported field to given value.


### GetNamespacesLicensed

`func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesLicensed() bool`

GetNamespacesLicensed returns the NamespacesLicensed field if non-nil, zero value otherwise.

### GetNamespacesLicensedOk

`func (o *VcenterNamespaceManagementHostsConfigInfo) GetNamespacesLicensedOk() (*bool, bool)`

GetNamespacesLicensedOk returns a tuple with the NamespacesLicensed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespacesLicensed

`func (o *VcenterNamespaceManagementHostsConfigInfo) SetNamespacesLicensed(v bool)`

SetNamespacesLicensed sets NamespacesLicensed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


