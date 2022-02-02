# VcenterDatacenterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datacenter** | **string** | Identifier of the datacenter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datacenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datacenter. | 
**Name** | **string** | Name of the datacenter. | 

## Methods

### NewVcenterDatacenterSummary

`func NewVcenterDatacenterSummary(datacenter string, name string, ) *VcenterDatacenterSummary`

NewVcenterDatacenterSummary instantiates a new VcenterDatacenterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatacenterSummaryWithDefaults

`func NewVcenterDatacenterSummaryWithDefaults() *VcenterDatacenterSummary`

NewVcenterDatacenterSummaryWithDefaults instantiates a new VcenterDatacenterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenter

`func (o *VcenterDatacenterSummary) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *VcenterDatacenterSummary) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *VcenterDatacenterSummary) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.


### GetName

`func (o *VcenterDatacenterSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterDatacenterSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterDatacenterSummary) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


