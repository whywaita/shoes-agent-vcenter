# VcenterContentRegistriesHarborGarbageCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterContentRegistriesRecurrence**](VcenterContentRegistriesRecurrence.md) |  | 
**DayOfWeek** | Pointer to [**VcenterContentRegistriesDayOfWeek**](VcenterContentRegistriesDayOfWeek.md) |  | [optional] 
**Hour** | Pointer to **int64** | Hour at which garbage collection should run. | [optional] 
**Minute** | Pointer to **int64** | Minute at which garbage collection should run. | [optional] 

## Methods

### NewVcenterContentRegistriesHarborGarbageCollection

`func NewVcenterContentRegistriesHarborGarbageCollection(type_ VcenterContentRegistriesRecurrence, ) *VcenterContentRegistriesHarborGarbageCollection`

NewVcenterContentRegistriesHarborGarbageCollection instantiates a new VcenterContentRegistriesHarborGarbageCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborGarbageCollectionWithDefaults

`func NewVcenterContentRegistriesHarborGarbageCollectionWithDefaults() *VcenterContentRegistriesHarborGarbageCollection`

NewVcenterContentRegistriesHarborGarbageCollectionWithDefaults instantiates a new VcenterContentRegistriesHarborGarbageCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetType() VcenterContentRegistriesRecurrence`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetTypeOk() (*VcenterContentRegistriesRecurrence, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterContentRegistriesHarborGarbageCollection) SetType(v VcenterContentRegistriesRecurrence)`

SetType sets Type field to given value.


### GetDayOfWeek

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetDayOfWeek() VcenterContentRegistriesDayOfWeek`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetDayOfWeekOk() (*VcenterContentRegistriesDayOfWeek, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *VcenterContentRegistriesHarborGarbageCollection) SetDayOfWeek(v VcenterContentRegistriesDayOfWeek)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *VcenterContentRegistriesHarborGarbageCollection) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetHour() int64`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetHourOk() (*int64, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *VcenterContentRegistriesHarborGarbageCollection) SetHour(v int64)`

SetHour sets Hour field to given value.

### HasHour

`func (o *VcenterContentRegistriesHarborGarbageCollection) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetMinute() int64`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *VcenterContentRegistriesHarborGarbageCollection) GetMinuteOk() (*int64, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *VcenterContentRegistriesHarborGarbageCollection) SetMinute(v int64)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *VcenterContentRegistriesHarborGarbageCollection) HasMinute() bool`

HasMinute returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


