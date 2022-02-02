# VcenterNamespaceManagementSupportBundleToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name \&quot;wcp-support-bundle-token\&quot;. | 
**Expiry** | **string** | Expiry time of the token | 

## Methods

### NewVcenterNamespaceManagementSupportBundleToken

`func NewVcenterNamespaceManagementSupportBundleToken(token string, expiry string, ) *VcenterNamespaceManagementSupportBundleToken`

NewVcenterNamespaceManagementSupportBundleToken instantiates a new VcenterNamespaceManagementSupportBundleToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSupportBundleTokenWithDefaults

`func NewVcenterNamespaceManagementSupportBundleTokenWithDefaults() *VcenterNamespaceManagementSupportBundleToken`

NewVcenterNamespaceManagementSupportBundleTokenWithDefaults instantiates a new VcenterNamespaceManagementSupportBundleToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *VcenterNamespaceManagementSupportBundleToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VcenterNamespaceManagementSupportBundleToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VcenterNamespaceManagementSupportBundleToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiry

`func (o *VcenterNamespaceManagementSupportBundleToken) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *VcenterNamespaceManagementSupportBundleToken) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *VcenterNamespaceManagementSupportBundleToken) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


