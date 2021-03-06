/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VcenterCertificateManagementVcenterVmcaRootCreateSpec struct for VcenterCertificateManagementVcenterVmcaRootCreateSpec
type VcenterCertificateManagementVcenterVmcaRootCreateSpec struct {
	// The size of the key to be used for public and private key generation. If unset the key size will be 2048.
	KeySize *int64 `json:"key_size,omitempty"`
	// The common name of the host for which certificate is generated. If unset the common name will be the primary network identifier (PNID) of the vCenter Virtual Server Appliance (VCSA).
	CommonName *string `json:"common_name,omitempty"`
	// Organization field in certificate subject. If unset the organization will be 'VMware'.
	Organization *string `json:"organization,omitempty"`
	// Organization unit field in certificate subject. If unset the organization unit will be 'VMware Engineering'.
	OrganizationUnit *string `json:"organization_unit,omitempty"`
	// Locality field in certificate subject. If unset the locality will be 'Palo Alto'.
	Locality *string `json:"locality,omitempty"`
	// State field in certificate subject. If unset the state will be 'California'.
	StateOrProvince *string `json:"state_or_province,omitempty"`
	// Country field in certificate subject. If unset the country will be 'US'.
	Country *string `json:"country,omitempty"`
	// Email field in Certificate extensions. If unset the emailAddress will be 'email@acme.com'.
	EmailAddress *string `json:"email_address,omitempty"`
	// SubjectAltName is list of Dns Names and Ip addresses. If unset PNID of host will be used as IPAddress or Hostname for certificate generation.
	SubjectAltName *[]string `json:"subject_alt_name,omitempty"`
}

// NewVcenterCertificateManagementVcenterVmcaRootCreateSpec instantiates a new VcenterCertificateManagementVcenterVmcaRootCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterVmcaRootCreateSpec() *VcenterCertificateManagementVcenterVmcaRootCreateSpec {
	this := VcenterCertificateManagementVcenterVmcaRootCreateSpec{}
	return &this
}

// NewVcenterCertificateManagementVcenterVmcaRootCreateSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterVmcaRootCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterVmcaRootCreateSpecWithDefaults() *VcenterCertificateManagementVcenterVmcaRootCreateSpec {
	this := VcenterCertificateManagementVcenterVmcaRootCreateSpec{}
	return &this
}

// GetKeySize returns the KeySize field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetKeySize() int64 {
	if o == nil || o.KeySize == nil {
		var ret int64
		return ret
	}
	return *o.KeySize
}

// GetKeySizeOk returns a tuple with the KeySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetKeySizeOk() (*int64, bool) {
	if o == nil || o.KeySize == nil {
		return nil, false
	}
	return o.KeySize, true
}

// HasKeySize returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasKeySize() bool {
	if o != nil && o.KeySize != nil {
		return true
	}

	return false
}

// SetKeySize gets a reference to the given int64 and assigns it to the KeySize field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetKeySize(v int64) {
	o.KeySize = &v
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCommonName() string {
	if o == nil || o.CommonName == nil {
		var ret string
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCommonNameOk() (*string, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given string and assigns it to the CommonName field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetCommonName(v string) {
	o.CommonName = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetOrganization(v string) {
	o.Organization = &v
}

// GetOrganizationUnit returns the OrganizationUnit field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationUnit() string {
	if o == nil || o.OrganizationUnit == nil {
		var ret string
		return ret
	}
	return *o.OrganizationUnit
}

// GetOrganizationUnitOk returns a tuple with the OrganizationUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetOrganizationUnitOk() (*string, bool) {
	if o == nil || o.OrganizationUnit == nil {
		return nil, false
	}
	return o.OrganizationUnit, true
}

// HasOrganizationUnit returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasOrganizationUnit() bool {
	if o != nil && o.OrganizationUnit != nil {
		return true
	}

	return false
}

// SetOrganizationUnit gets a reference to the given string and assigns it to the OrganizationUnit field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetOrganizationUnit(v string) {
	o.OrganizationUnit = &v
}

// GetLocality returns the Locality field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetLocality() string {
	if o == nil || o.Locality == nil {
		var ret string
		return ret
	}
	return *o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetLocalityOk() (*string, bool) {
	if o == nil || o.Locality == nil {
		return nil, false
	}
	return o.Locality, true
}

// HasLocality returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasLocality() bool {
	if o != nil && o.Locality != nil {
		return true
	}

	return false
}

// SetLocality gets a reference to the given string and assigns it to the Locality field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetLocality(v string) {
	o.Locality = &v
}

// GetStateOrProvince returns the StateOrProvince field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetStateOrProvince() string {
	if o == nil || o.StateOrProvince == nil {
		var ret string
		return ret
	}
	return *o.StateOrProvince
}

// GetStateOrProvinceOk returns a tuple with the StateOrProvince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetStateOrProvinceOk() (*string, bool) {
	if o == nil || o.StateOrProvince == nil {
		return nil, false
	}
	return o.StateOrProvince, true
}

// HasStateOrProvince returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasStateOrProvince() bool {
	if o != nil && o.StateOrProvince != nil {
		return true
	}

	return false
}

// SetStateOrProvince gets a reference to the given string and assigns it to the StateOrProvince field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetStateOrProvince(v string) {
	o.StateOrProvince = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetCountry(v string) {
	o.Country = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetSubjectAltName returns the SubjectAltName field value if set, zero value otherwise.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetSubjectAltName() []string {
	if o == nil || o.SubjectAltName == nil {
		var ret []string
		return ret
	}
	return *o.SubjectAltName
}

// GetSubjectAltNameOk returns a tuple with the SubjectAltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) GetSubjectAltNameOk() (*[]string, bool) {
	if o == nil || o.SubjectAltName == nil {
		return nil, false
	}
	return o.SubjectAltName, true
}

// HasSubjectAltName returns a boolean if a field has been set.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) HasSubjectAltName() bool {
	if o != nil && o.SubjectAltName != nil {
		return true
	}

	return false
}

// SetSubjectAltName gets a reference to the given []string and assigns it to the SubjectAltName field.
func (o *VcenterCertificateManagementVcenterVmcaRootCreateSpec) SetSubjectAltName(v []string) {
	o.SubjectAltName = &v
}

func (o VcenterCertificateManagementVcenterVmcaRootCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.KeySize != nil {
		toSerialize["key_size"] = o.KeySize
	}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.OrganizationUnit != nil {
		toSerialize["organization_unit"] = o.OrganizationUnit
	}
	if o.Locality != nil {
		toSerialize["locality"] = o.Locality
	}
	if o.StateOrProvince != nil {
		toSerialize["state_or_province"] = o.StateOrProvince
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.SubjectAltName != nil {
		toSerialize["subject_alt_name"] = o.SubjectAltName
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec struct {
	value *VcenterCertificateManagementVcenterVmcaRootCreateSpec
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) Get() *VcenterCertificateManagementVcenterVmcaRootCreateSpec {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) Set(val *VcenterCertificateManagementVcenterVmcaRootCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterVmcaRootCreateSpec(val *VcenterCertificateManagementVcenterVmcaRootCreateSpec) *NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec {
	return &NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterVmcaRootCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


