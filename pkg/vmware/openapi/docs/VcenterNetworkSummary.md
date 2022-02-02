# VcenterNetworkSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Identifier of the network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network. | 
**Name** | **string** | Name of the network. | 
**Type** | [**VcenterNetworkType**](VcenterNetworkType.md) |  | 

## Methods

### NewVcenterNetworkSummary

`func NewVcenterNetworkSummary(network string, name string, type_ VcenterNetworkType, ) *VcenterNetworkSummary`

NewVcenterNetworkSummary instantiates a new VcenterNetworkSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNetworkSummaryWithDefaults

`func NewVcenterNetworkSummaryWithDefaults() *VcenterNetworkSummary`

NewVcenterNetworkSummaryWithDefaults instantiates a new VcenterNetworkSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *VcenterNetworkSummary) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VcenterNetworkSummary) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VcenterNetworkSummary) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetName

`func (o *VcenterNetworkSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterNetworkSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterNetworkSummary) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VcenterNetworkSummary) GetType() VcenterNetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterNetworkSummary) GetTypeOk() (*VcenterNetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterNetworkSummary) SetType(v VcenterNetworkType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


