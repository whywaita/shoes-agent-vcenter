# VcenterCertificateManagementVcenterTlsSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | **string** | Certificate string in PEM format. | 
**Key** | Pointer to **string** | Private key string in PEM format. If unset the private key from the certificate store will be used. It is required when replacing the certificate with a third party signed certificate. | [optional] 
**RootCert** | Pointer to **string** | Third party Root CA certificate in PEM format. If unset the new third party root CA certificate will not be added to the trust store. It is required when replacing the certificate with a third party signed certificate if the root certificate of the third party is not already a trusted root. | [optional] 

## Methods

### NewVcenterCertificateManagementVcenterTlsSpec

`func NewVcenterCertificateManagementVcenterTlsSpec(cert string, ) *VcenterCertificateManagementVcenterTlsSpec`

NewVcenterCertificateManagementVcenterTlsSpec instantiates a new VcenterCertificateManagementVcenterTlsSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterTlsSpecWithDefaults

`func NewVcenterCertificateManagementVcenterTlsSpecWithDefaults() *VcenterCertificateManagementVcenterTlsSpec`

NewVcenterCertificateManagementVcenterTlsSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterTlsSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *VcenterCertificateManagementVcenterTlsSpec) SetCert(v string)`

SetCert sets Cert field to given value.


### GetKey

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *VcenterCertificateManagementVcenterTlsSpec) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *VcenterCertificateManagementVcenterTlsSpec) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRootCert

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetRootCert() string`

GetRootCert returns the RootCert field if non-nil, zero value otherwise.

### GetRootCertOk

`func (o *VcenterCertificateManagementVcenterTlsSpec) GetRootCertOk() (*string, bool)`

GetRootCertOk returns a tuple with the RootCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCert

`func (o *VcenterCertificateManagementVcenterTlsSpec) SetRootCert(v string)`

SetRootCert sets RootCert field to given value.

### HasRootCert

`func (o *VcenterCertificateManagementVcenterTlsSpec) HasRootCert() bool`

HasRootCert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


