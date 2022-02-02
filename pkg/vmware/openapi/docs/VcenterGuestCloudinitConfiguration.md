# VcenterGuestCloudinitConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **string** | Metadata includes the network, instance id and hostname that cloud-init processes to configure the virtual machine. It is in json or yaml format. The max size of the metadata is 524288 bytes. See https://cloudinit.readthedocs.io/en/latest/topics/datasources/ovf.html about supported meta data formats. | 
**Userdata** | Pointer to **string** | Userdata is the user customized content that cloud-init processes to configure the virtual machine. See https://cloudinit.readthedocs.io/en/latest/topics/format.html about user data formats. See https://cloudinit.readthedocs.io/en/latest/topics/modules.html# about user data modules. The max size of the userdata is 524288 bytes. If unset, no cloud-init user data will be used as part of the cloud-init configuration. | [optional] 

## Methods

### NewVcenterGuestCloudinitConfiguration

`func NewVcenterGuestCloudinitConfiguration(metadata string, ) *VcenterGuestCloudinitConfiguration`

NewVcenterGuestCloudinitConfiguration instantiates a new VcenterGuestCloudinitConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCloudinitConfigurationWithDefaults

`func NewVcenterGuestCloudinitConfigurationWithDefaults() *VcenterGuestCloudinitConfiguration`

NewVcenterGuestCloudinitConfigurationWithDefaults instantiates a new VcenterGuestCloudinitConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *VcenterGuestCloudinitConfiguration) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VcenterGuestCloudinitConfiguration) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VcenterGuestCloudinitConfiguration) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.


### GetUserdata

`func (o *VcenterGuestCloudinitConfiguration) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *VcenterGuestCloudinitConfiguration) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *VcenterGuestCloudinitConfiguration) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *VcenterGuestCloudinitConfiguration) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


