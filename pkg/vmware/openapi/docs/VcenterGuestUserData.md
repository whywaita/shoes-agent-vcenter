# VcenterGuestUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComputerName** | [**VcenterGuestHostnameGenerator**](VcenterGuestHostnameGenerator.md) |  | 
**FullName** | **string** | Full name of the end user. Note that this is not the username but full name specified in \&quot;Firstname Lastname\&quot; format. | 
**Organization** | **string** | Name of the organization that owns the computer. | 
**ProductKey** | **string** | The product Key to use for activating Windows guest operating system. | 

## Methods

### NewVcenterGuestUserData

`func NewVcenterGuestUserData(computerName VcenterGuestHostnameGenerator, fullName string, organization string, productKey string, ) *VcenterGuestUserData`

NewVcenterGuestUserData instantiates a new VcenterGuestUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestUserDataWithDefaults

`func NewVcenterGuestUserDataWithDefaults() *VcenterGuestUserData`

NewVcenterGuestUserDataWithDefaults instantiates a new VcenterGuestUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComputerName

`func (o *VcenterGuestUserData) GetComputerName() VcenterGuestHostnameGenerator`

GetComputerName returns the ComputerName field if non-nil, zero value otherwise.

### GetComputerNameOk

`func (o *VcenterGuestUserData) GetComputerNameOk() (*VcenterGuestHostnameGenerator, bool)`

GetComputerNameOk returns a tuple with the ComputerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputerName

`func (o *VcenterGuestUserData) SetComputerName(v VcenterGuestHostnameGenerator)`

SetComputerName sets ComputerName field to given value.


### GetFullName

`func (o *VcenterGuestUserData) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VcenterGuestUserData) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VcenterGuestUserData) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetOrganization

`func (o *VcenterGuestUserData) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VcenterGuestUserData) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VcenterGuestUserData) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetProductKey

`func (o *VcenterGuestUserData) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *VcenterGuestUserData) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *VcenterGuestUserData) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


