# VcenterVchaClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigState** | Pointer to [**VcenterVchaClusterConfigState**](VcenterVchaClusterConfigState.md) |  | [optional] 
**Node1** | Pointer to [**VcenterVchaClusterNodeInfo**](VcenterVchaClusterNodeInfo.md) |  | [optional] 
**Node2** | Pointer to [**VcenterVchaClusterNodeInfo**](VcenterVchaClusterNodeInfo.md) |  | [optional] 
**Witness** | Pointer to [**VcenterVchaClusterWitnessInfo**](VcenterVchaClusterWitnessInfo.md) |  | [optional] 
**Mode** | Pointer to [**VcenterVchaClusterClusterMode**](VcenterVchaClusterClusterMode.md) |  | [optional] 
**HealthState** | Pointer to [**VcenterVchaClusterClusterState**](VcenterVchaClusterClusterState.md) |  | [optional] 
**HealthException** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Health warning messages if the health information is unavailable. If unset, then the cluster is in a healthy state. | [optional] 
**HealthWarnings** | Pointer to [**[]VcenterVchaClusterErrorCondition**](VcenterVchaClusterErrorCondition.md) | A collection of messages describing the reason for a non-healthy Cluster. If unset, then the cluster is in a healthy state. | [optional] 
**ManualFailoverAllowed** | Pointer to **bool** | Specifies if manual failover is allowed. If unset, then the cluster state healthy and manual failover allowance in accordance with the cluster mode. | [optional] 
**AutoFailoverAllowed** | Pointer to **bool** | Specifies if automatic failover is allowed. If unset, then the cluster state healthy and automatic failover allowance in accordance with the cluster mode. | [optional] 

## Methods

### NewVcenterVchaClusterInfo

`func NewVcenterVchaClusterInfo() *VcenterVchaClusterInfo`

NewVcenterVchaClusterInfo instantiates a new VcenterVchaClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterInfoWithDefaults

`func NewVcenterVchaClusterInfoWithDefaults() *VcenterVchaClusterInfo`

NewVcenterVchaClusterInfoWithDefaults instantiates a new VcenterVchaClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigState

`func (o *VcenterVchaClusterInfo) GetConfigState() VcenterVchaClusterConfigState`

GetConfigState returns the ConfigState field if non-nil, zero value otherwise.

### GetConfigStateOk

`func (o *VcenterVchaClusterInfo) GetConfigStateOk() (*VcenterVchaClusterConfigState, bool)`

GetConfigStateOk returns a tuple with the ConfigState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigState

`func (o *VcenterVchaClusterInfo) SetConfigState(v VcenterVchaClusterConfigState)`

SetConfigState sets ConfigState field to given value.

### HasConfigState

`func (o *VcenterVchaClusterInfo) HasConfigState() bool`

HasConfigState returns a boolean if a field has been set.

### GetNode1

`func (o *VcenterVchaClusterInfo) GetNode1() VcenterVchaClusterNodeInfo`

GetNode1 returns the Node1 field if non-nil, zero value otherwise.

### GetNode1Ok

`func (o *VcenterVchaClusterInfo) GetNode1Ok() (*VcenterVchaClusterNodeInfo, bool)`

GetNode1Ok returns a tuple with the Node1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode1

`func (o *VcenterVchaClusterInfo) SetNode1(v VcenterVchaClusterNodeInfo)`

SetNode1 sets Node1 field to given value.

### HasNode1

`func (o *VcenterVchaClusterInfo) HasNode1() bool`

HasNode1 returns a boolean if a field has been set.

### GetNode2

`func (o *VcenterVchaClusterInfo) GetNode2() VcenterVchaClusterNodeInfo`

GetNode2 returns the Node2 field if non-nil, zero value otherwise.

### GetNode2Ok

`func (o *VcenterVchaClusterInfo) GetNode2Ok() (*VcenterVchaClusterNodeInfo, bool)`

GetNode2Ok returns a tuple with the Node2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode2

`func (o *VcenterVchaClusterInfo) SetNode2(v VcenterVchaClusterNodeInfo)`

SetNode2 sets Node2 field to given value.

### HasNode2

`func (o *VcenterVchaClusterInfo) HasNode2() bool`

HasNode2 returns a boolean if a field has been set.

### GetWitness

`func (o *VcenterVchaClusterInfo) GetWitness() VcenterVchaClusterWitnessInfo`

GetWitness returns the Witness field if non-nil, zero value otherwise.

### GetWitnessOk

`func (o *VcenterVchaClusterInfo) GetWitnessOk() (*VcenterVchaClusterWitnessInfo, bool)`

GetWitnessOk returns a tuple with the Witness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWitness

`func (o *VcenterVchaClusterInfo) SetWitness(v VcenterVchaClusterWitnessInfo)`

SetWitness sets Witness field to given value.

### HasWitness

`func (o *VcenterVchaClusterInfo) HasWitness() bool`

HasWitness returns a boolean if a field has been set.

### GetMode

`func (o *VcenterVchaClusterInfo) GetMode() VcenterVchaClusterClusterMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *VcenterVchaClusterInfo) GetModeOk() (*VcenterVchaClusterClusterMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *VcenterVchaClusterInfo) SetMode(v VcenterVchaClusterClusterMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *VcenterVchaClusterInfo) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetHealthState

`func (o *VcenterVchaClusterInfo) GetHealthState() VcenterVchaClusterClusterState`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *VcenterVchaClusterInfo) GetHealthStateOk() (*VcenterVchaClusterClusterState, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *VcenterVchaClusterInfo) SetHealthState(v VcenterVchaClusterClusterState)`

SetHealthState sets HealthState field to given value.

### HasHealthState

`func (o *VcenterVchaClusterInfo) HasHealthState() bool`

HasHealthState returns a boolean if a field has been set.

### GetHealthException

`func (o *VcenterVchaClusterInfo) GetHealthException() []VapiStdLocalizableMessage`

GetHealthException returns the HealthException field if non-nil, zero value otherwise.

### GetHealthExceptionOk

`func (o *VcenterVchaClusterInfo) GetHealthExceptionOk() (*[]VapiStdLocalizableMessage, bool)`

GetHealthExceptionOk returns a tuple with the HealthException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthException

`func (o *VcenterVchaClusterInfo) SetHealthException(v []VapiStdLocalizableMessage)`

SetHealthException sets HealthException field to given value.

### HasHealthException

`func (o *VcenterVchaClusterInfo) HasHealthException() bool`

HasHealthException returns a boolean if a field has been set.

### GetHealthWarnings

`func (o *VcenterVchaClusterInfo) GetHealthWarnings() []VcenterVchaClusterErrorCondition`

GetHealthWarnings returns the HealthWarnings field if non-nil, zero value otherwise.

### GetHealthWarningsOk

`func (o *VcenterVchaClusterInfo) GetHealthWarningsOk() (*[]VcenterVchaClusterErrorCondition, bool)`

GetHealthWarningsOk returns a tuple with the HealthWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthWarnings

`func (o *VcenterVchaClusterInfo) SetHealthWarnings(v []VcenterVchaClusterErrorCondition)`

SetHealthWarnings sets HealthWarnings field to given value.

### HasHealthWarnings

`func (o *VcenterVchaClusterInfo) HasHealthWarnings() bool`

HasHealthWarnings returns a boolean if a field has been set.

### GetManualFailoverAllowed

`func (o *VcenterVchaClusterInfo) GetManualFailoverAllowed() bool`

GetManualFailoverAllowed returns the ManualFailoverAllowed field if non-nil, zero value otherwise.

### GetManualFailoverAllowedOk

`func (o *VcenterVchaClusterInfo) GetManualFailoverAllowedOk() (*bool, bool)`

GetManualFailoverAllowedOk returns a tuple with the ManualFailoverAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualFailoverAllowed

`func (o *VcenterVchaClusterInfo) SetManualFailoverAllowed(v bool)`

SetManualFailoverAllowed sets ManualFailoverAllowed field to given value.

### HasManualFailoverAllowed

`func (o *VcenterVchaClusterInfo) HasManualFailoverAllowed() bool`

HasManualFailoverAllowed returns a boolean if a field has been set.

### GetAutoFailoverAllowed

`func (o *VcenterVchaClusterInfo) GetAutoFailoverAllowed() bool`

GetAutoFailoverAllowed returns the AutoFailoverAllowed field if non-nil, zero value otherwise.

### GetAutoFailoverAllowedOk

`func (o *VcenterVchaClusterInfo) GetAutoFailoverAllowedOk() (*bool, bool)`

GetAutoFailoverAllowedOk returns a tuple with the AutoFailoverAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFailoverAllowed

`func (o *VcenterVchaClusterInfo) SetAutoFailoverAllowed(v bool)`

SetAutoFailoverAllowed sets AutoFailoverAllowed field to given value.

### HasAutoFailoverAllowed

`func (o *VcenterVchaClusterInfo) HasAutoFailoverAllowed() bool`

HasAutoFailoverAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


