# VcenterGuestDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterGuestDomainType**](VcenterGuestDomainType.md) |  | 
**Workgroup** | Pointer to **string** | The workgroup that the virtual machine should join. This field is optional and it is only relevant when the value of Domain.type is WORKGROUP. | [optional] 
**Domain** | Pointer to **string** | The domain to which the virtual machine should be joined. This field is optional and it is only relevant when the value of Domain.type is DOMAIN. | [optional] 
**DomainUsername** | Pointer to **string** | The domain user that has permission to join the domain after virtual machine is joined. This field is optional and it is only relevant when the value of Domain.type is DOMAIN. | [optional] 
**DomainPassword** | Pointer to **string** | The domain user password that has permission to join the Domain.domain-username after customization. This field is optional and it is only relevant when the value of Domain.type is DOMAIN. | [optional] 

## Methods

### NewVcenterGuestDomain

`func NewVcenterGuestDomain(type_ VcenterGuestDomainType, ) *VcenterGuestDomain`

NewVcenterGuestDomain instantiates a new VcenterGuestDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestDomainWithDefaults

`func NewVcenterGuestDomainWithDefaults() *VcenterGuestDomain`

NewVcenterGuestDomainWithDefaults instantiates a new VcenterGuestDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterGuestDomain) GetType() VcenterGuestDomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterGuestDomain) GetTypeOk() (*VcenterGuestDomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterGuestDomain) SetType(v VcenterGuestDomainType)`

SetType sets Type field to given value.


### GetWorkgroup

`func (o *VcenterGuestDomain) GetWorkgroup() string`

GetWorkgroup returns the Workgroup field if non-nil, zero value otherwise.

### GetWorkgroupOk

`func (o *VcenterGuestDomain) GetWorkgroupOk() (*string, bool)`

GetWorkgroupOk returns a tuple with the Workgroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkgroup

`func (o *VcenterGuestDomain) SetWorkgroup(v string)`

SetWorkgroup sets Workgroup field to given value.

### HasWorkgroup

`func (o *VcenterGuestDomain) HasWorkgroup() bool`

HasWorkgroup returns a boolean if a field has been set.

### GetDomain

`func (o *VcenterGuestDomain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterGuestDomain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterGuestDomain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *VcenterGuestDomain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDomainUsername

`func (o *VcenterGuestDomain) GetDomainUsername() string`

GetDomainUsername returns the DomainUsername field if non-nil, zero value otherwise.

### GetDomainUsernameOk

`func (o *VcenterGuestDomain) GetDomainUsernameOk() (*string, bool)`

GetDomainUsernameOk returns a tuple with the DomainUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainUsername

`func (o *VcenterGuestDomain) SetDomainUsername(v string)`

SetDomainUsername sets DomainUsername field to given value.

### HasDomainUsername

`func (o *VcenterGuestDomain) HasDomainUsername() bool`

HasDomainUsername returns a boolean if a field has been set.

### GetDomainPassword

`func (o *VcenterGuestDomain) GetDomainPassword() string`

GetDomainPassword returns the DomainPassword field if non-nil, zero value otherwise.

### GetDomainPasswordOk

`func (o *VcenterGuestDomain) GetDomainPasswordOk() (*string, bool)`

GetDomainPasswordOk returns a tuple with the DomainPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainPassword

`func (o *VcenterGuestDomain) SetDomainPassword(v string)`

SetDomainPassword sets DomainPassword field to given value.

### HasDomainPassword

`func (o *VcenterGuestDomain) HasDomainPassword() bool`

HasDomainPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


