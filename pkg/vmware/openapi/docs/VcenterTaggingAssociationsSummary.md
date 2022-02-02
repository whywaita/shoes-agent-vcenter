# VcenterTaggingAssociationsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The identifier of a tag. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: cis.tagging.Tag. When operations return a value of this structure as a result, the field will be an identifier for the resource type: cis.tagging.Tag. | 
**Object** | [**VapiStdDynamicID**](VapiStdDynamicID.md) |  | 

## Methods

### NewVcenterTaggingAssociationsSummary

`func NewVcenterTaggingAssociationsSummary(tag string, object VapiStdDynamicID, ) *VcenterTaggingAssociationsSummary`

NewVcenterTaggingAssociationsSummary instantiates a new VcenterTaggingAssociationsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTaggingAssociationsSummaryWithDefaults

`func NewVcenterTaggingAssociationsSummaryWithDefaults() *VcenterTaggingAssociationsSummary`

NewVcenterTaggingAssociationsSummaryWithDefaults instantiates a new VcenterTaggingAssociationsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *VcenterTaggingAssociationsSummary) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *VcenterTaggingAssociationsSummary) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *VcenterTaggingAssociationsSummary) SetTag(v string)`

SetTag sets Tag field to given value.


### GetObject

`func (o *VcenterTaggingAssociationsSummary) GetObject() VapiStdDynamicID`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *VcenterTaggingAssociationsSummary) GetObjectOk() (*VapiStdDynamicID, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *VcenterTaggingAssociationsSummary) SetObject(v VapiStdDynamicID)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


