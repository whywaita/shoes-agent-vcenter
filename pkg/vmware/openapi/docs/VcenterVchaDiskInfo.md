# VcenterVchaDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | **string** | The identifier of the datastore to put all the virtual disks on. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore:VCenter. | 
**DatastoreName** | **string** | The name of the datastore. | 

## Methods

### NewVcenterVchaDiskInfo

`func NewVcenterVchaDiskInfo(datastore string, datastoreName string, ) *VcenterVchaDiskInfo`

NewVcenterVchaDiskInfo instantiates a new VcenterVchaDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaDiskInfoWithDefaults

`func NewVcenterVchaDiskInfoWithDefaults() *VcenterVchaDiskInfo`

NewVcenterVchaDiskInfoWithDefaults instantiates a new VcenterVchaDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVchaDiskInfo) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVchaDiskInfo) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVchaDiskInfo) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetDatastoreName

`func (o *VcenterVchaDiskInfo) GetDatastoreName() string`

GetDatastoreName returns the DatastoreName field if non-nil, zero value otherwise.

### GetDatastoreNameOk

`func (o *VcenterVchaDiskInfo) GetDatastoreNameOk() (*string, bool)`

GetDatastoreNameOk returns a tuple with the DatastoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreName

`func (o *VcenterVchaDiskInfo) SetDatastoreName(v string)`

SetDatastoreName sets DatastoreName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


