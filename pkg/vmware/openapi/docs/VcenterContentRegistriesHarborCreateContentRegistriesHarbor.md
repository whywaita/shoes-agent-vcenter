# VcenterContentRegistriesHarborCreateContentRegistriesHarbor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientToken** | Pointer to **string** | A unique token generated on the client for each creation request. The token should be a universally unique identifier (UUID), for example: {@code b8a2a2e3-2314-43cd-a871-6ede0f429751}. This token can be used to guarantee idempotent creation. | [optional] 
**Spec** | Pointer to [**VcenterContentRegistriesHarborCreateSpec**](VcenterContentRegistriesHarborCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterContentRegistriesHarborCreateContentRegistriesHarbor

`func NewVcenterContentRegistriesHarborCreateContentRegistriesHarbor() *VcenterContentRegistriesHarborCreateContentRegistriesHarbor`

NewVcenterContentRegistriesHarborCreateContentRegistriesHarbor instantiates a new VcenterContentRegistriesHarborCreateContentRegistriesHarbor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborCreateContentRegistriesHarborWithDefaults

`func NewVcenterContentRegistriesHarborCreateContentRegistriesHarborWithDefaults() *VcenterContentRegistriesHarborCreateContentRegistriesHarbor`

NewVcenterContentRegistriesHarborCreateContentRegistriesHarborWithDefaults instantiates a new VcenterContentRegistriesHarborCreateContentRegistriesHarbor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientToken

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) GetClientToken() string`

GetClientToken returns the ClientToken field if non-nil, zero value otherwise.

### GetClientTokenOk

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) GetClientTokenOk() (*string, bool)`

GetClientTokenOk returns a tuple with the ClientToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientToken

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) SetClientToken(v string)`

SetClientToken sets ClientToken field to given value.

### HasClientToken

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) HasClientToken() bool`

HasClientToken returns a boolean if a field has been set.

### GetSpec

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) GetSpec() VcenterContentRegistriesHarborCreateSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) GetSpecOk() (*VcenterContentRegistriesHarborCreateSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) SetSpec(v VcenterContentRegistriesHarborCreateSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *VcenterContentRegistriesHarborCreateContentRegistriesHarbor) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


