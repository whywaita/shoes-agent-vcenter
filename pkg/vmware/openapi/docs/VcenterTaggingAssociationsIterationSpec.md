# VcenterTaggingAssociationsIterationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Marker** | Pointer to **string** | Marker is an opaque token that allows the caller to request the next page of tag associations. If unset or empty, first page of tag associations will be returned. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.tagging.associations.Marker. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.tagging.associations.Marker. | [optional] 

## Methods

### NewVcenterTaggingAssociationsIterationSpec

`func NewVcenterTaggingAssociationsIterationSpec() *VcenterTaggingAssociationsIterationSpec`

NewVcenterTaggingAssociationsIterationSpec instantiates a new VcenterTaggingAssociationsIterationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTaggingAssociationsIterationSpecWithDefaults

`func NewVcenterTaggingAssociationsIterationSpecWithDefaults() *VcenterTaggingAssociationsIterationSpec`

NewVcenterTaggingAssociationsIterationSpecWithDefaults instantiates a new VcenterTaggingAssociationsIterationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarker

`func (o *VcenterTaggingAssociationsIterationSpec) GetMarker() string`

GetMarker returns the Marker field if non-nil, zero value otherwise.

### GetMarkerOk

`func (o *VcenterTaggingAssociationsIterationSpec) GetMarkerOk() (*string, bool)`

GetMarkerOk returns a tuple with the Marker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarker

`func (o *VcenterTaggingAssociationsIterationSpec) SetMarker(v string)`

SetMarker sets Marker field to given value.

### HasMarker

`func (o *VcenterTaggingAssociationsIterationSpec) HasMarker() bool`

HasMarker returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


