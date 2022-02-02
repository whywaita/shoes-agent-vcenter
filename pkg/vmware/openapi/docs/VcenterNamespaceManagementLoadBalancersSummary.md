# VcenterNamespaceManagementLoadBalancersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An DNS compliant identifier for a load balancer, which can be used to query or configure the load balancer properties. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.LoadBalancerConfig. | 
**Provider** | [**VcenterNamespaceManagementLoadBalancersProvider**](VcenterNamespaceManagementLoadBalancersProvider.md) |  | 

## Methods

### NewVcenterNamespaceManagementLoadBalancersSummary

`func NewVcenterNamespaceManagementLoadBalancersSummary(id string, provider VcenterNamespaceManagementLoadBalancersProvider, ) *VcenterNamespaceManagementLoadBalancersSummary`

NewVcenterNamespaceManagementLoadBalancersSummary instantiates a new VcenterNamespaceManagementLoadBalancersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementLoadBalancersSummaryWithDefaults

`func NewVcenterNamespaceManagementLoadBalancersSummaryWithDefaults() *VcenterNamespaceManagementLoadBalancersSummary`

NewVcenterNamespaceManagementLoadBalancersSummaryWithDefaults instantiates a new VcenterNamespaceManagementLoadBalancersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterNamespaceManagementLoadBalancersSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterNamespaceManagementLoadBalancersSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterNamespaceManagementLoadBalancersSummary) SetId(v string)`

SetId sets Id field to given value.


### GetProvider

`func (o *VcenterNamespaceManagementLoadBalancersSummary) GetProvider() VcenterNamespaceManagementLoadBalancersProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterNamespaceManagementLoadBalancersSummary) GetProviderOk() (*VcenterNamespaceManagementLoadBalancersProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterNamespaceManagementLoadBalancersSummary) SetProvider(v VcenterNamespaceManagementLoadBalancersProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


