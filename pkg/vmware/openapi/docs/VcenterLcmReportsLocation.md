# VcenterLcmReportsLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | **string** | Report Download URI. | 
**DownloadFileToken** | [**VcenterLcmReportsToken**](VcenterLcmReportsToken.md) |  | 

## Methods

### NewVcenterLcmReportsLocation

`func NewVcenterLcmReportsLocation(uri string, downloadFileToken VcenterLcmReportsToken, ) *VcenterLcmReportsLocation`

NewVcenterLcmReportsLocation instantiates a new VcenterLcmReportsLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmReportsLocationWithDefaults

`func NewVcenterLcmReportsLocationWithDefaults() *VcenterLcmReportsLocation`

NewVcenterLcmReportsLocationWithDefaults instantiates a new VcenterLcmReportsLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *VcenterLcmReportsLocation) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *VcenterLcmReportsLocation) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *VcenterLcmReportsLocation) SetUri(v string)`

SetUri sets Uri field to given value.


### GetDownloadFileToken

`func (o *VcenterLcmReportsLocation) GetDownloadFileToken() VcenterLcmReportsToken`

GetDownloadFileToken returns the DownloadFileToken field if non-nil, zero value otherwise.

### GetDownloadFileTokenOk

`func (o *VcenterLcmReportsLocation) GetDownloadFileTokenOk() (*VcenterLcmReportsToken, bool)`

GetDownloadFileTokenOk returns a tuple with the DownloadFileToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadFileToken

`func (o *VcenterLcmReportsLocation) SetDownloadFileToken(v VcenterLcmReportsToken)`

SetDownloadFileToken sets DownloadFileToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


