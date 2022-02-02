# VcenterVmGuestFilesystemFilesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchPattern** | Pointer to **string** | The perl-compatible regular expression used to filter the returned files. If unset the pattern &#39;.*&#39; (match everything) is used. | [optional] 

## Methods

### NewVcenterVmGuestFilesystemFilesFilterSpec

`func NewVcenterVmGuestFilesystemFilesFilterSpec() *VcenterVmGuestFilesystemFilesFilterSpec`

NewVcenterVmGuestFilesystemFilesFilterSpec instantiates a new VcenterVmGuestFilesystemFilesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestFilesystemFilesFilterSpecWithDefaults

`func NewVcenterVmGuestFilesystemFilesFilterSpecWithDefaults() *VcenterVmGuestFilesystemFilesFilterSpec`

NewVcenterVmGuestFilesystemFilesFilterSpecWithDefaults instantiates a new VcenterVmGuestFilesystemFilesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchPattern

`func (o *VcenterVmGuestFilesystemFilesFilterSpec) GetMatchPattern() string`

GetMatchPattern returns the MatchPattern field if non-nil, zero value otherwise.

### GetMatchPatternOk

`func (o *VcenterVmGuestFilesystemFilesFilterSpec) GetMatchPatternOk() (*string, bool)`

GetMatchPatternOk returns a tuple with the MatchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchPattern

`func (o *VcenterVmGuestFilesystemFilesFilterSpec) SetMatchPattern(v string)`

SetMatchPattern sets MatchPattern field to given value.

### HasMatchPattern

`func (o *VcenterVmGuestFilesystemFilesFilterSpec) HasMatchPattern() bool`

HasMatchPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


