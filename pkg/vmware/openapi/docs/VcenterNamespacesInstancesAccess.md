# VcenterNamespacesInstancesAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectType** | [**VcenterNamespacesAccessSubjectType**](VcenterNamespacesAccessSubjectType.md) |  | 
**Subject** | **string** | Name of the subject. | 
**Domain** | **string** | Domain of the subject. | 
**Role** | [**VcenterNamespacesAccessRole**](VcenterNamespacesAccessRole.md) |  | 

## Methods

### NewVcenterNamespacesInstancesAccess

`func NewVcenterNamespacesInstancesAccess(subjectType VcenterNamespacesAccessSubjectType, subject string, domain string, role VcenterNamespacesAccessRole, ) *VcenterNamespacesInstancesAccess`

NewVcenterNamespacesInstancesAccess instantiates a new VcenterNamespacesInstancesAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesAccessWithDefaults

`func NewVcenterNamespacesInstancesAccessWithDefaults() *VcenterNamespacesInstancesAccess`

NewVcenterNamespacesInstancesAccessWithDefaults instantiates a new VcenterNamespacesInstancesAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectType

`func (o *VcenterNamespacesInstancesAccess) GetSubjectType() VcenterNamespacesAccessSubjectType`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *VcenterNamespacesInstancesAccess) GetSubjectTypeOk() (*VcenterNamespacesAccessSubjectType, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *VcenterNamespacesInstancesAccess) SetSubjectType(v VcenterNamespacesAccessSubjectType)`

SetSubjectType sets SubjectType field to given value.


### GetSubject

`func (o *VcenterNamespacesInstancesAccess) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *VcenterNamespacesInstancesAccess) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *VcenterNamespacesInstancesAccess) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDomain

`func (o *VcenterNamespacesInstancesAccess) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterNamespacesInstancesAccess) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterNamespacesInstancesAccess) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRole

`func (o *VcenterNamespacesInstancesAccess) GetRole() VcenterNamespacesAccessRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *VcenterNamespacesInstancesAccess) GetRoleOk() (*VcenterNamespacesAccessRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *VcenterNamespacesInstancesAccess) SetRole(v VcenterNamespacesAccessRole)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


