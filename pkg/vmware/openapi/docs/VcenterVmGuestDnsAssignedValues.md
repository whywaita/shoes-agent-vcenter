# VcenterVmGuestDnsAssignedValues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostName** | **string** | The host name portion of DNS name. For example, \&quot;esx01\&quot; part of esx01.example.com. | 
**DomainName** | **string** | The domain name portion of the DNS name. \&quot;example.com\&quot; part of esx01.example.com. | 

## Methods

### NewVcenterVmGuestDnsAssignedValues

`func NewVcenterVmGuestDnsAssignedValues(hostName string, domainName string, ) *VcenterVmGuestDnsAssignedValues`

NewVcenterVmGuestDnsAssignedValues instantiates a new VcenterVmGuestDnsAssignedValues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestDnsAssignedValuesWithDefaults

`func NewVcenterVmGuestDnsAssignedValuesWithDefaults() *VcenterVmGuestDnsAssignedValues`

NewVcenterVmGuestDnsAssignedValuesWithDefaults instantiates a new VcenterVmGuestDnsAssignedValues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostName

`func (o *VcenterVmGuestDnsAssignedValues) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *VcenterVmGuestDnsAssignedValues) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *VcenterVmGuestDnsAssignedValues) SetHostName(v string)`

SetHostName sets HostName field to given value.


### GetDomainName

`func (o *VcenterVmGuestDnsAssignedValues) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VcenterVmGuestDnsAssignedValues) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VcenterVmGuestDnsAssignedValues) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


