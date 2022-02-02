# VcenterCertificateManagementVcenterTlsReplaceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySize** | Pointer to **int64** | The size of the key to be used for public and private key generation. If unset the key size will be &#39;2048&#39;. | [optional] 
**CommonName** | Pointer to **string** | The common name of the host for which certificate is generated If unset will default to PNID of host. | [optional] 
**Organization** | **string** | Organization field in certificate subject | 
**OrganizationUnit** | **string** | Organization unit field in certificate subject | 
**Locality** | **string** | Locality field in certificate subject | 
**StateOrProvince** | **string** | State field in certificate subject | 
**Country** | **string** | Country field in certificate subject | 
**EmailAddress** | **string** | Email field in Certificate extensions | 
**SubjectAltName** | Pointer to **[]string** | SubjectAltName is list of Dns Names and Ip addresses If unset PNID of host will be used as IPAddress or Hostname for certificate generation . | [optional] 

## Methods

### NewVcenterCertificateManagementVcenterTlsReplaceSpec

`func NewVcenterCertificateManagementVcenterTlsReplaceSpec(organization string, organizationUnit string, locality string, stateOrProvince string, country string, emailAddress string, ) *VcenterCertificateManagementVcenterTlsReplaceSpec`

NewVcenterCertificateManagementVcenterTlsReplaceSpec instantiates a new VcenterCertificateManagementVcenterTlsReplaceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterTlsReplaceSpecWithDefaults

`func NewVcenterCertificateManagementVcenterTlsReplaceSpecWithDefaults() *VcenterCertificateManagementVcenterTlsReplaceSpec`

NewVcenterCertificateManagementVcenterTlsReplaceSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterTlsReplaceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySize

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetCommonName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### GetLocality

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetStateOrProvince

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.


### GetCountry

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmailAddress

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetSubjectAltName() []string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) GetSubjectAltNameOk() (*[]string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) SetSubjectAltName(v []string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsReplaceSpec) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


