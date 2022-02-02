# VcenterDatastoreSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | **string** | Identifier of the datastore. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | 
**Name** | **string** | Name of the datastore. | 
**Type** | [**VcenterDatastoreType**](VcenterDatastoreType.md) |  | 
**FreeSpace** | Pointer to **int64** | Available space of this datastore, in bytes.   The server periodically updates this value.  This field will be unset if the available space of this datastore is not known. | [optional] 
**Capacity** | Pointer to **int64** | Capacity of this datastore, in bytes.   The server periodically updates this value.  This field will be unset if the capacity of this datastore is not known. | [optional] 

## Methods

### NewVcenterDatastoreSummary

`func NewVcenterDatastoreSummary(datastore string, name string, type_ VcenterDatastoreType, ) *VcenterDatastoreSummary`

NewVcenterDatastoreSummary instantiates a new VcenterDatastoreSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDatastoreSummaryWithDefaults

`func NewVcenterDatastoreSummaryWithDefaults() *VcenterDatastoreSummary`

NewVcenterDatastoreSummaryWithDefaults instantiates a new VcenterDatastoreSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterDatastoreSummary) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterDatastoreSummary) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterDatastoreSummary) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetName

`func (o *VcenterDatastoreSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterDatastoreSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterDatastoreSummary) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VcenterDatastoreSummary) GetType() VcenterDatastoreType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterDatastoreSummary) GetTypeOk() (*VcenterDatastoreType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterDatastoreSummary) SetType(v VcenterDatastoreType)`

SetType sets Type field to given value.


### GetFreeSpace

`func (o *VcenterDatastoreSummary) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *VcenterDatastoreSummary) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *VcenterDatastoreSummary) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.

### HasFreeSpace

`func (o *VcenterDatastoreSummary) HasFreeSpace() bool`

HasFreeSpace returns a boolean if a field has been set.

### GetCapacity

`func (o *VcenterDatastoreSummary) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterDatastoreSummary) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterDatastoreSummary) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VcenterDatastoreSummary) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


