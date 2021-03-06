/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VcenterCertificateManagementVcenterTlsInfo struct for VcenterCertificateManagementVcenterTlsInfo
type VcenterCertificateManagementVcenterTlsInfo struct {
	// Version (version number) value from the certificate.
	Version int64 `json:"version"`
	// SerialNumber value from the certificate.
	SerialNumber string `json:"serial_number"`
	// Signature algorithm name from the certificate.
	SignatureAlgorithm string `json:"signature_algorithm"`
	// Issuer (issuer distinguished name) value from the certificate.
	IssuerDn string `json:"issuer_dn"`
	// validFrom specify the start date of the certificate.
	ValidFrom time.Time `json:"valid_from"`
	// validTo specify the end date of the certificate.
	ValidTo time.Time `json:"valid_to"`
	// Subject (subject distinguished name) value from the certificate.
	SubjectDn string `json:"subject_dn"`
	// Thumbprint value from the certificate.
	Thumbprint string `json:"thumbprint"`
	// Certificate constraints isCA from the critical BasicConstraints extension, (OID = 2.5.29.19).
	IsCA bool `json:"is_CA"`
	// Certificate constraints path length from the critical BasicConstraints extension, (OID = 2.5.29.19).
	PathLengthConstraint int64 `json:"path_length_constraint"`
	// Collection of keyusage contained in the certificate.
	KeyUsage []string `json:"key_usage"`
	// Collection of extended keyusage that contains details for which the certificate can be used for.
	ExtendedKeyUsage []string `json:"extended_key_usage"`
	// Collection of subject alternative names.
	SubjectAlternativeName []string `json:"subject_alternative_name"`
	// Collection of authority information access URI.
	AuthorityInformationAccessUri []string `json:"authority_information_access_uri"`
	// TLS certificate in PEM format.
	Cert string `json:"cert"`
}

// NewVcenterCertificateManagementVcenterTlsInfo instantiates a new VcenterCertificateManagementVcenterTlsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCertificateManagementVcenterTlsInfo(version int64, serialNumber string, signatureAlgorithm string, issuerDn string, validFrom time.Time, validTo time.Time, subjectDn string, thumbprint string, isCA bool, pathLengthConstraint int64, keyUsage []string, extendedKeyUsage []string, subjectAlternativeName []string, authorityInformationAccessUri []string, cert string) *VcenterCertificateManagementVcenterTlsInfo {
	this := VcenterCertificateManagementVcenterTlsInfo{}
	this.Version = version
	this.SerialNumber = serialNumber
	this.SignatureAlgorithm = signatureAlgorithm
	this.IssuerDn = issuerDn
	this.ValidFrom = validFrom
	this.ValidTo = validTo
	this.SubjectDn = subjectDn
	this.Thumbprint = thumbprint
	this.IsCA = isCA
	this.PathLengthConstraint = pathLengthConstraint
	this.KeyUsage = keyUsage
	this.ExtendedKeyUsage = extendedKeyUsage
	this.SubjectAlternativeName = subjectAlternativeName
	this.AuthorityInformationAccessUri = authorityInformationAccessUri
	this.Cert = cert
	return &this
}

// NewVcenterCertificateManagementVcenterTlsInfoWithDefaults instantiates a new VcenterCertificateManagementVcenterTlsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCertificateManagementVcenterTlsInfoWithDefaults() *VcenterCertificateManagementVcenterTlsInfo {
	this := VcenterCertificateManagementVcenterTlsInfo{}
	return &this
}

// GetVersion returns the Version field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetVersionOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetVersion(v int64) {
	o.Version = v
}

// GetSerialNumber returns the SerialNumber field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSerialNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSignatureAlgorithm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SignatureAlgorithm, true
}

// SetSignatureAlgorithm sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = v
}

// GetIssuerDn returns the IssuerDn field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetIssuerDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssuerDn
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetIssuerDnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IssuerDn, true
}

// SetIssuerDn sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetIssuerDn(v string) {
	o.IssuerDn = v
}

// GetValidFrom returns the ValidFrom field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidFromOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidFrom, true
}

// SetValidFrom sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetValidFrom(v time.Time) {
	o.ValidFrom = v
}

// GetValidTo returns the ValidTo field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidToOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ValidTo, true
}

// SetValidTo sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetValidTo(v time.Time) {
	o.ValidTo = v
}

// GetSubjectDn returns the SubjectDn field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubjectDn
}

// GetSubjectDnOk returns a tuple with the SubjectDn field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectDnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubjectDn, true
}

// SetSubjectDn sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetSubjectDn(v string) {
	o.SubjectDn = v
}

// GetThumbprint returns the Thumbprint field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetThumbprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Thumbprint
}

// GetThumbprintOk returns a tuple with the Thumbprint field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetThumbprintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Thumbprint, true
}

// SetThumbprint sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetThumbprint(v string) {
	o.Thumbprint = v
}

// GetIsCA returns the IsCA field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetIsCA() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCA
}

// GetIsCAOk returns a tuple with the IsCA field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetIsCAOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsCA, true
}

// SetIsCA sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetIsCA(v bool) {
	o.IsCA = v
}

// GetPathLengthConstraint returns the PathLengthConstraint field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetPathLengthConstraint() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PathLengthConstraint
}

// GetPathLengthConstraintOk returns a tuple with the PathLengthConstraint field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetPathLengthConstraintOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PathLengthConstraint, true
}

// SetPathLengthConstraint sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetPathLengthConstraint(v int64) {
	o.PathLengthConstraint = v
}

// GetKeyUsage returns the KeyUsage field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetKeyUsage() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.KeyUsage
}

// GetKeyUsageOk returns a tuple with the KeyUsage field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetKeyUsageOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyUsage, true
}

// SetKeyUsage sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetKeyUsage(v []string) {
	o.KeyUsage = v
}

// GetExtendedKeyUsage returns the ExtendedKeyUsage field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetExtendedKeyUsage() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ExtendedKeyUsage
}

// GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetExtendedKeyUsageOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExtendedKeyUsage, true
}

// SetExtendedKeyUsage sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetExtendedKeyUsage(v []string) {
	o.ExtendedKeyUsage = v
}

// GetSubjectAlternativeName returns the SubjectAlternativeName field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectAlternativeName() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectAlternativeName
}

// GetSubjectAlternativeNameOk returns a tuple with the SubjectAlternativeName field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectAlternativeNameOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubjectAlternativeName, true
}

// SetSubjectAlternativeName sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetSubjectAlternativeName(v []string) {
	o.SubjectAlternativeName = v
}

// GetAuthorityInformationAccessUri returns the AuthorityInformationAccessUri field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetAuthorityInformationAccessUri() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuthorityInformationAccessUri
}

// GetAuthorityInformationAccessUriOk returns a tuple with the AuthorityInformationAccessUri field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetAuthorityInformationAccessUriOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthorityInformationAccessUri, true
}

// SetAuthorityInformationAccessUri sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetAuthorityInformationAccessUri(v []string) {
	o.AuthorityInformationAccessUri = v
}

// GetCert returns the Cert field value
func (o *VcenterCertificateManagementVcenterTlsInfo) GetCert() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cert
}

// GetCertOk returns a tuple with the Cert field value
// and a boolean to check if the value has been set.
func (o *VcenterCertificateManagementVcenterTlsInfo) GetCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cert, true
}

// SetCert sets field value
func (o *VcenterCertificateManagementVcenterTlsInfo) SetCert(v string) {
	o.Cert = v
}

func (o VcenterCertificateManagementVcenterTlsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if true {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if true {
		toSerialize["issuer_dn"] = o.IssuerDn
	}
	if true {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if true {
		toSerialize["valid_to"] = o.ValidTo
	}
	if true {
		toSerialize["subject_dn"] = o.SubjectDn
	}
	if true {
		toSerialize["thumbprint"] = o.Thumbprint
	}
	if true {
		toSerialize["is_CA"] = o.IsCA
	}
	if true {
		toSerialize["path_length_constraint"] = o.PathLengthConstraint
	}
	if true {
		toSerialize["key_usage"] = o.KeyUsage
	}
	if true {
		toSerialize["extended_key_usage"] = o.ExtendedKeyUsage
	}
	if true {
		toSerialize["subject_alternative_name"] = o.SubjectAlternativeName
	}
	if true {
		toSerialize["authority_information_access_uri"] = o.AuthorityInformationAccessUri
	}
	if true {
		toSerialize["cert"] = o.Cert
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCertificateManagementVcenterTlsInfo struct {
	value *VcenterCertificateManagementVcenterTlsInfo
	isSet bool
}

func (v NullableVcenterCertificateManagementVcenterTlsInfo) Get() *VcenterCertificateManagementVcenterTlsInfo {
	return v.value
}

func (v *NullableVcenterCertificateManagementVcenterTlsInfo) Set(val *VcenterCertificateManagementVcenterTlsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCertificateManagementVcenterTlsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCertificateManagementVcenterTlsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCertificateManagementVcenterTlsInfo(val *VcenterCertificateManagementVcenterTlsInfo) *NullableVcenterCertificateManagementVcenterTlsInfo {
	return &NullableVcenterCertificateManagementVcenterTlsInfo{value: val, isSet: true}
}

func (v NullableVcenterCertificateManagementVcenterTlsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCertificateManagementVcenterTlsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


