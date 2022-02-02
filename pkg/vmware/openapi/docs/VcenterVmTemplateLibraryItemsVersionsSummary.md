# VcenterVmTemplateLibraryItemsVersionsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The version of the library item. | 
**VmTemplate** | **string** | Identifier of the virtual machine template associated with the library item version. This {@term field} is the managed object identifier used to identify the virtual machine template in the vSphere Management (SOAP) API. | 

## Methods

### NewVcenterVmTemplateLibraryItemsVersionsSummary

`func NewVcenterVmTemplateLibraryItemsVersionsSummary(version string, vmTemplate string, ) *VcenterVmTemplateLibraryItemsVersionsSummary`

NewVcenterVmTemplateLibraryItemsVersionsSummary instantiates a new VcenterVmTemplateLibraryItemsVersionsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsVersionsSummaryWithDefaults

`func NewVcenterVmTemplateLibraryItemsVersionsSummaryWithDefaults() *VcenterVmTemplateLibraryItemsVersionsSummary`

NewVcenterVmTemplateLibraryItemsVersionsSummaryWithDefaults instantiates a new VcenterVmTemplateLibraryItemsVersionsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVmTemplate

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) GetVmTemplate() string`

GetVmTemplate returns the VmTemplate field if non-nil, zero value otherwise.

### GetVmTemplateOk

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) GetVmTemplateOk() (*string, bool)`

GetVmTemplateOk returns a tuple with the VmTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplate

`func (o *VcenterVmTemplateLibraryItemsVersionsSummary) SetVmTemplate(v string)`

SetVmTemplate sets VmTemplate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


