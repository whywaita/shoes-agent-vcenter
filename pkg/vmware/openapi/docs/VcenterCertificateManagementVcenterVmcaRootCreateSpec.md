# VcenterCertificateManagementVcenterVmcaRootCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySize** | Pointer to **int64** | The size of the key to be used for public and private key generation. If unset the key size will be 2048. | [optional] 
**CommonName** | Pointer to **string** | The common name of the host for which certificate is generated. If unset the common name will be the primary network identifier (PNID) of the vCenter Virtual Server Appliance (VCSA). | [optional] 
**Organization** | Pointer to **string** | Organization field in certificate subject. If unset the organization will be &#39;VMware&#39;. | [optional] 
**OrganizationUnit** | Pointer to **string** | Organization unit field in certificate subject. If unset the organization unit will be &#39;VMware Engineering&#39;. | [optional] 
**Locality** | Pointer to **string** | Locality field in certificate subject. If unset the locality will be &#39;Palo Alto&#39;. | [optional] 
**StateOrProvince** | Pointer to **string** | State field in certificate subject. If unset the state will be &#39;California&#39;. | [optional] 
**Country** | Pointer to **string** | Country field in certificate subject. If unset the country will be &#39;US&#39;. | [optional] 
**EmailAddress** | Pointer to **string** | Email field in Certificate extensions. If unset the emailAddress will be &#39;email@acme.com&#39;. | [optional] 
**SubjectAltName** | Pointer to **[]string** | SubjectAltName is list of Dns Names and Ip addresses. If unset PNID of host will be used as IPAddress or Hostname for certificate generation. | [optional] 

## Methods

### NewVcenterCertificateManagementVcenterVmcaRootCreateSpec

`func NewVcenterCertificateManagementVcenterVmcaRootCreateSpec() *VcenterCertificateManagementVcenterVmcaRootCreateSpec`

NewVcenterCertificateManagementVcenterVmcaRootCreateSpec instantiates a new VcenterCertificateManagementVcenterVmcaRootCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterVmcaRootCreateSpecWithDefaults

`func NewVcenterCertificateManagementVcenterVmcaRootCreateSpecWithDefaults() *VcenterCertificateManagementVcenterVmcaRootCreateSpec`

NewVcenterCertificateManagementVcenterVmcaRootCreateSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterVmcaRootCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySize

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetKeySize() int64`

GetKeySize returns the KeySize field if non-nil, zero value otherwise.

### GetKeySizeOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetKeySizeOk() (*int64, bool)`

GetKeySizeOk returns a tuple with the KeySize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySize

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetKeySize(v int64)`

SetKeySize sets KeySize field to given value.

### HasKeySize

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasKeySize() bool`

HasKeySize returns a boolean if a field has been set.

### GetCommonName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.

### HasOrganizationUnit

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasOrganizationUnit() bool`

HasOrganizationUnit returns a boolean if a field has been set.

### GetLocality

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetLocality(v string)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### GetStateOrProvince

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetStateOrProvince() string`

GetStateOrProvince returns the StateOrProvince field if non-nil, zero value otherwise.

### GetStateOrProvinceOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetStateOrProvinceOk() (*string, bool)`

GetStateOrProvinceOk returns a tuple with the StateOrProvince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateOrProvince

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetStateOrProvince(v string)`

SetStateOrProvince sets StateOrProvince field to given value.

### HasStateOrProvince

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasStateOrProvince() bool`

HasStateOrProvince returns a boolean if a field has been set.

### GetCountry

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmailAddress

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetSubjectAltName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetSubjectAltName() []string`

GetSubjectAltName returns the SubjectAltName field if non-nil, zero value otherwise.

### GetSubjectAltNameOk

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetSubjectAltNameOk() (*[]string, bool)`

GetSubjectAltNameOk returns a tuple with the SubjectAltName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetSubjectAltName(v []string)`

SetSubjectAltName sets SubjectAltName field to given value.

### HasSubjectAltName

`func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasSubjectAltName() bool`

HasSubjectAltName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


