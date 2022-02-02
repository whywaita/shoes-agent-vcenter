# VcenterHvcLinksSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | **string** | Unique identifier for the link. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**DisplayName** | **string** | The display name is set to the domain name which was set during create. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 

## Methods

### NewVcenterHvcLinksSummary

`func NewVcenterHvcLinksSummary(link string, displayName string, ) *VcenterHvcLinksSummary`

NewVcenterHvcLinksSummary instantiates a new VcenterHvcLinksSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHvcLinksSummaryWithDefaults

`func NewVcenterHvcLinksSummaryWithDefaults() *VcenterHvcLinksSummary`

NewVcenterHvcLinksSummaryWithDefaults instantiates a new VcenterHvcLinksSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *VcenterHvcLinksSummary) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *VcenterHvcLinksSummary) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *VcenterHvcLinksSummary) SetLink(v string)`

SetLink sets Link field to given value.


### GetDisplayName

`func (o *VcenterHvcLinksSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterHvcLinksSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterHvcLinksSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


