# VcenterVmHardwareAdapterNvmeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adapter** | **string** | Identifier of the virtual NVMe adapter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.NvmeAdapter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.NvmeAdapter. | 

## Methods

### NewVcenterVmHardwareAdapterNvmeSummary

`func NewVcenterVmHardwareAdapterNvmeSummary(adapter string, ) *VcenterVmHardwareAdapterNvmeSummary`

NewVcenterVmHardwareAdapterNvmeSummary instantiates a new VcenterVmHardwareAdapterNvmeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareAdapterNvmeSummaryWithDefaults

`func NewVcenterVmHardwareAdapterNvmeSummaryWithDefaults() *VcenterVmHardwareAdapterNvmeSummary`

NewVcenterVmHardwareAdapterNvmeSummaryWithDefaults instantiates a new VcenterVmHardwareAdapterNvmeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdapter

`func (o *VcenterVmHardwareAdapterNvmeSummary) GetAdapter() string`

GetAdapter returns the Adapter field if non-nil, zero value otherwise.

### GetAdapterOk

`func (o *VcenterVmHardwareAdapterNvmeSummary) GetAdapterOk() (*string, bool)`

GetAdapterOk returns a tuple with the Adapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdapter

`func (o *VcenterVmHardwareAdapterNvmeSummary) SetAdapter(v string)`

SetAdapter sets Adapter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


