# VcenterTaggingAssociationsListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]VcenterTaggingAssociationsSummary**](VcenterTaggingAssociationsSummary.md) | List of tag associations. | 
**Marker** | Pointer to **string** | Marker is an opaque data structure that allows the caller to request the next page of tag associations. If unset or empty, there are no more tag associations to request. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.tagging.associations.Marker. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.tagging.associations.Marker. | [optional] 
**Status** | [**VcenterTaggingAssociationsLastIterationStatus**](VcenterTaggingAssociationsLastIterationStatus.md) |  | 

## Methods

### NewVcenterTaggingAssociationsListResult

`func NewVcenterTaggingAssociationsListResult(associations []VcenterTaggingAssociationsSummary, status VcenterTaggingAssociationsLastIterationStatus, ) *VcenterTaggingAssociationsListResult`

NewVcenterTaggingAssociationsListResult instantiates a new VcenterTaggingAssociationsListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTaggingAssociationsListResultWithDefaults

`func NewVcenterTaggingAssociationsListResultWithDefaults() *VcenterTaggingAssociationsListResult`

NewVcenterTaggingAssociationsListResultWithDefaults instantiates a new VcenterTaggingAssociationsListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *VcenterTaggingAssociationsListResult) GetAssociations() []VcenterTaggingAssociationsSummary`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *VcenterTaggingAssociationsListResult) GetAssociationsOk() (*[]VcenterTaggingAssociationsSummary, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *VcenterTaggingAssociationsListResult) SetAssociations(v []VcenterTaggingAssociationsSummary)`

SetAssociations sets Associations field to given value.


### GetMarker

`func (o *VcenterTaggingAssociationsListResult) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *VcenterTaggingAssociationsListResult) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *VcenterTaggingAssociationsListResult) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *VcenterTaggingAssociationsListResult) HasMarker() bool`

HasMarker returns a boolean if a field has been set.

### GetStatus

`func (o *VcenterTaggingAssociationsListResult) GetStatus() VcenterTaggingAssociationsLastIterationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterTaggingAssociationsListResult) GetStatusOk() (*VcenterTaggingAssociationsLastIterationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterTaggingAssociationsListResult) SetStatus(v VcenterTaggingAssociationsLastIterationStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


