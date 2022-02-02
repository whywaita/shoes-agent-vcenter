# VcenterResourcePoolSharesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**VcenterResourcePoolSharesInfoLevel**](VcenterResourcePoolSharesInfoLevel.md) |  | 
**Shares** | Pointer to **int64** | When ResourcePool.SharesInfo.level is set to CUSTOM, it is the number of shares allocated. Otherwise, this value is ignored.   There is no unit for this value. It is a relative measure based on the settings for other resource pools.  This field is optional and it is only relevant when the value of ResourcePool.SharesInfo.level is CUSTOM. | [optional] 

## Methods

### NewVcenterResourcePoolSharesInfo

`func NewVcenterResourcePoolSharesInfo(level VcenterResourcePoolSharesInfoLevel, ) *VcenterResourcePoolSharesInfo`

NewVcenterResourcePoolSharesInfo instantiates a new VcenterResourcePoolSharesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolSharesInfoWithDefaults

`func NewVcenterResourcePoolSharesInfoWithDefaults() *VcenterResourcePoolSharesInfo`

NewVcenterResourcePoolSharesInfoWithDefaults instantiates a new VcenterResourcePoolSharesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *VcenterResourcePoolSharesInfo) GetLevel() VcenterResourcePoolSharesInfoLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *VcenterResourcePoolSharesInfo) GetLevelOk() (*VcenterResourcePoolSharesInfoLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *VcenterResourcePoolSharesInfo) SetLevel(v VcenterResourcePoolSharesInfoLevel)`

SetLevel sets Level field to given value.


### GetShares

`func (o *VcenterResourcePoolSharesInfo) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VcenterResourcePoolSharesInfo) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VcenterResourcePoolSharesInfo) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VcenterResourcePoolSharesInfo) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


