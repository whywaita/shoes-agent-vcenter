# VcenterCertificateManagementVcenterTlsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int64** | Version (version number) value from the certificate. | 
**SerialNumber** | **string** | SerialNumber value from the certificate. | 
**SignatureAlgorithm** | **string** | Signature algorithm name from the certificate. | 
**IssuerDn** | **string** | Issuer (issuer distinguished name) value from the certificate. | 
**ValidFrom** | **time.Time** | validFrom specify the start date of the certificate. | 
**ValidTo** | **time.Time** | validTo specify the end date of the certificate. | 
**SubjectDn** | **string** | Subject (subject distinguished name) value from the certificate. | 
**Thumbprint** | **string** | Thumbprint value from the certificate. | 
**IsCA** | **bool** | Certificate constraints isCA from the critical BasicConstraints extension, (OID &#x3D; 2.5.29.19). | 
**PathLengthConstraint** | **int64** | Certificate constraints path length from the critical BasicConstraints extension, (OID &#x3D; 2.5.29.19). | 
**KeyUsage** | **[]string** | Collection of keyusage contained in the certificate. | 
**ExtendedKeyUsage** | **[]string** | Collection of extended keyusage that contains details for which the certificate can be used for. | 
**SubjectAlternativeName** | **[]string** | Collection of subject alternative names. | 
**AuthorityInformationAccessUri** | **[]string** | Collection of authority information access URI. | 
**Cert** | **string** | TLS certificate in PEM format. | 

## Methods

### NewVcenterCertificateManagementVcenterTlsInfo

`func NewVcenterCertificateManagementVcenterTlsInfo(version int64, serialNumber string, signatureAlgorithm string, issuerDn string, validFrom time.Time, validTo time.Time, subjectDn string, thumbprint string, isCA bool, pathLengthConstraint int64, keyUsage []string, extendedKeyUsage []string, subjectAlternativeName []string, authorityInformationAccessUri []string, cert string, ) *VcenterCertificateManagementVcenterTlsInfo`

NewVcenterCertificateManagementVcenterTlsInfo instantiates a new VcenterCertificateManagementVcenterTlsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterTlsInfoWithDefaults

`func NewVcenterCertificateManagementVcenterTlsInfoWithDefaults() *VcenterCertificateManagementVcenterTlsInfo`

NewVcenterCertificateManagementVcenterTlsInfoWithDefaults instantiates a new VcenterCertificateManagementVcenterTlsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetSerialNumber

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetSignatureAlgorithm

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSignatureAlgorithm() string`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSignatureAlgorithmOk() (*string, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetSignatureAlgorithm(v string)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetIssuerDn

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetIssuerDn() string`

GetIssuerDn returns the IssuerDn field if non-nil, zero value otherwise.

### GetIssuerDnOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetIssuerDnOk() (*string, bool)`

GetIssuerDnOk returns a tuple with the IssuerDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDn

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetIssuerDn(v string)`

SetIssuerDn sets IssuerDn field to given value.


### GetValidFrom

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.


### GetSubjectDn

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectDn() string`

GetSubjectDn returns the SubjectDn field if non-nil, zero value otherwise.

### GetSubjectDnOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectDnOk() (*string, bool)`

GetSubjectDnOk returns a tuple with the SubjectDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDn

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetSubjectDn(v string)`

SetSubjectDn sets SubjectDn field to given value.


### GetThumbprint

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.


### GetIsCA

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetIsCA() bool`

GetIsCA returns the IsCA field if non-nil, zero value otherwise.

### GetIsCAOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetIsCAOk() (*bool, bool)`

GetIsCAOk returns a tuple with the IsCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCA

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetIsCA(v bool)`

SetIsCA sets IsCA field to given value.


### GetPathLengthConstraint

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetPathLengthConstraint() int64`

GetPathLengthConstraint returns the PathLengthConstraint field if non-nil, zero value otherwise.

### GetPathLengthConstraintOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetPathLengthConstraintOk() (*int64, bool)`

GetPathLengthConstraintOk returns a tuple with the PathLengthConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathLengthConstraint

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetPathLengthConstraint(v int64)`

SetPathLengthConstraint sets PathLengthConstraint field to given value.


### GetKeyUsage

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetKeyUsage() []string`

GetKeyUsage returns the KeyUsage field if non-nil, zero value otherwise.

### GetKeyUsageOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetKeyUsageOk() (*[]string, bool)`

GetKeyUsageOk returns a tuple with the KeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUsage

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetKeyUsage(v []string)`

SetKeyUsage sets KeyUsage field to given value.


### GetExtendedKeyUsage

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetExtendedKeyUsage() []string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetExtendedKeyUsageOk() (*[]string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetExtendedKeyUsage(v []string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.


### GetSubjectAlternativeName

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectAlternativeName() []string`

GetSubjectAlternativeName returns the SubjectAlternativeName field if non-nil, zero value otherwise.

### GetSubjectAlternativeNameOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetSubjectAlternativeNameOk() (*[]string, bool)`

GetSubjectAlternativeNameOk returns a tuple with the SubjectAlternativeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeName

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetSubjectAlternativeName(v []string)`

SetSubjectAlternativeName sets SubjectAlternativeName field to given value.


### GetAuthorityInformationAccessUri

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetAuthorityInformationAccessUri() []string`

GetAuthorityInformationAccessUri returns the AuthorityInformationAccessUri field if non-nil, zero value otherwise.

### GetAuthorityInformationAccessUriOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetAuthorityInformationAccessUriOk() (*[]string, bool)`

GetAuthorityInformationAccessUriOk returns a tuple with the AuthorityInformationAccessUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityInformationAccessUri

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetAuthorityInformationAccessUri(v []string)`

SetAuthorityInformationAccessUri sets AuthorityInformationAccessUri field to given value.


### GetCert

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *VcenterCertificateManagementVcenterTlsInfo) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *VcenterCertificateManagementVcenterTlsInfo) SetCert(v string)`

SetCert sets Cert field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


