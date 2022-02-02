# VcenterCertificateManagementVcenterTlsCsrSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySize** | Pointer to **int64** | keySize will take 2048 bits if not modified. | [optional] 
**CommonName** | Pointer to **string** | commonName will take PNID if not modified. | [optional] 
**Organization** | **string** | Organization field in certificate subject | 
**OrganizationUnit** | **string** | Organization unit field in certificate subject | 
**Locality** | **string** | Locality field in certificate subject | 
**StateOrProvince** | **string** | State field in certificate subject | 
**Country** | **string** | Country field in certificate subject | 
**EmailAddress** | **string** | Email field in Certificate extensions | 
**SubjectAltName** | Pointer to **[]string** | subjectAltName is list of Dns Names and Ip addresses | [optional] 

## Methods

### NewVcenterCertificateManagementVcenterTlsCsrSpec

`func NewVcenterCertificateManagementVcenterTlsCsrSpec(organization string, organizationUnit string, locality string, stateOrProvince string, country string, emailAddress string, ) *VcenterCertificateManagementVcenterTlsCsrSpec`

NewVcenterCertificateManagementVcenterTlsCsrSpec instantiates a new VcenterCertificateManagementVcenterTlsCsrSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterTlsCsrSpecWithDefaults

`func NewVcenterCertificateManagementVcenterTlsCsrSpecWithDefaults() *VcenterCertificateManagementVcenterTlsCsrSpec`

NewVcenterCertificateManagementVcenterTlsCsrSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterTlsCsrSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySize

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetCommonName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.


### GetLocality

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetStateOrProvince

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.


### GetCountry

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmailAddress

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetSubjectAltName() []string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) GetSubjectAltNameOk() (*[]string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) SetSubjectAltName(v []string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *VcenterCertificateManagementVcenterTlsCsrSpec) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


