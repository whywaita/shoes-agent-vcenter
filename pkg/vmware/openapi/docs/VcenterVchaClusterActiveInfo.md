# VcenterVchaClusterActiveInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Management** | [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | 
**Ha** | Pointer to [**VcenterVchaIpSpec**](VcenterVchaIpSpec.md) |  | [optional] 
**Placement** | Pointer to [**VcenterVchaPlacementInfo**](VcenterVchaPlacementInfo.md) |  | [optional] 

## Methods

### NewVcenterVchaClusterActiveInfo

`func NewVcenterVchaClusterActiveInfo(management VcenterVchaIpSpec, ) *VcenterVchaClusterActiveInfo`

NewVcenterVchaClusterActiveInfo instantiates a new VcenterVchaClusterActiveInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterActiveInfoWithDefaults

`func NewVcenterVchaClusterActiveInfoWithDefaults() *VcenterVchaClusterActiveInfo`

NewVcenterVchaClusterActiveInfoWithDefaults instantiates a new VcenterVchaClusterActiveInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagement

`func (o *VcenterVchaClusterActiveInfo) GetManagement() VcenterVchaIpSpec`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *VcenterVchaClusterActiveInfo) GetManagementOk() (*VcenterVchaIpSpec, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *VcenterVchaClusterActiveInfo) SetManagement(v VcenterVchaIpSpec)`

SetManagement sets Management field to given value.


### GetHa

`func (o *VcenterVchaClusterActiveInfo) GetHa() VcenterVchaIpSpec`

GetHa returns the Ha field if non-nil, zero value otherwise.

### GetHaOk

`func (o *VcenterVchaClusterActiveInfo) GetHaOk() (*VcenterVchaIpSpec, bool)`

GetHaOk returns a tuple with the Ha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHa

`func (o *VcenterVchaClusterActiveInfo) SetHa(v VcenterVchaIpSpec)`

SetHa sets Ha field to given value.

### HasHa

`func (o *VcenterVchaClusterActiveInfo) HasHa() bool`

HasHa returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVchaClusterActiveInfo) GetPlacement() VcenterVchaPlacementInfo`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVchaClusterActiveInfo) GetPlacementOk() (*VcenterVchaPlacementInfo, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVchaClusterActiveInfo) SetPlacement(v VcenterVchaPlacementInfo)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVchaClusterActiveInfo) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


