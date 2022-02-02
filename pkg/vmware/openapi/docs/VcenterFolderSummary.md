# VcenterFolderSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | **string** | Identifier of the folder. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | 
**Name** | **string** | Name of the vCenter Server folder. | 
**Type** | [**VcenterFolderType**](VcenterFolderType.md) |  | 

## Methods

### NewVcenterFolderSummary

`func NewVcenterFolderSummary(folder string, name string, type_ VcenterFolderType, ) *VcenterFolderSummary`

NewVcenterFolderSummary instantiates a new VcenterFolderSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterFolderSummaryWithDefaults

`func NewVcenterFolderSummaryWithDefaults() *VcenterFolderSummary`

NewVcenterFolderSummaryWithDefaults instantiates a new VcenterFolderSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterFolderSummary) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterFolderSummary) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterFolderSummary) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetName

`func (o *VcenterFolderSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterFolderSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterFolderSummary) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VcenterFolderSummary) GetType() VcenterFolderType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterFolderSummary) GetTypeOk() (*VcenterFolderType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterFolderSummary) SetType(v VcenterFolderType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


