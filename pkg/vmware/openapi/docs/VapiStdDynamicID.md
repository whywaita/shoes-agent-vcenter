# VapiStdDynamicID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of resource being identified (for example {@code com.acme.Person}). &lt;p&gt; {@term Services} that contain {@term operations} for creating and deleting resources typically contain a {@term constant} specifying the resource type for the resources being created and deleted. The API metamodel metadata {@term services} include a {@term service} that allows retrieving all the known resource types. | 
**Id** | **string** | The identifier for a resource whose type is specified by {@link #type}. | 

## Methods

### NewVapiStdDynamicID

`func NewVapiStdDynamicID(type_ string, id string, ) *VapiStdDynamicID`

NewVapiStdDynamicID instantiates a new VapiStdDynamicID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdDynamicIDWithDefaults

`func NewVapiStdDynamicIDWithDefaults() *VapiStdDynamicID`

NewVapiStdDynamicIDWithDefaults instantiates a new VapiStdDynamicID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VapiStdDynamicID) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VapiStdDynamicID) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VapiStdDynamicID) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *VapiStdDynamicID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VapiStdDynamicID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VapiStdDynamicID) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


