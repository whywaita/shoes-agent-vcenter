# VcenterNamespacesInstancesVMServiceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentLibraries** | Pointer to **[]string** | Set of content libraries for use by the VM Service. The content libraries specified should exist in vSphere inventory. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: content.Library. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: content.Library. | [optional] 
**VmClasses** | Pointer to **[]string** | Set of VirtualMachineClasses for use by the VM Service. The class names specified here should exist in vSphere inventory. If this field is empty in an updated specification, all VirtualMachineClasses that are currently associated with the namespace will be disassociated from it.   NOTE: Any change in virtual machine classes associated with the namespace will not impact existing VMs.  This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.namespace_management.VirtualMachineClass. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.namespace_management.VirtualMachineClass. | [optional] 

## Methods

### NewVcenterNamespacesInstancesVMServiceSpec

`func NewVcenterNamespacesInstancesVMServiceSpec() *VcenterNamespacesInstancesVMServiceSpec`

NewVcenterNamespacesInstancesVMServiceSpec instantiates a new VcenterNamespacesInstancesVMServiceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesVMServiceSpecWithDefaults

`func NewVcenterNamespacesInstancesVMServiceSpecWithDefaults() *VcenterNamespacesInstancesVMServiceSpec`

NewVcenterNamespacesInstancesVMServiceSpecWithDefaults instantiates a new VcenterNamespacesInstancesVMServiceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentLibraries

`func (o *VcenterNamespacesInstancesVMServiceSpec) GetContentLibraries() []string`

GetContentLibraries returns the ContentLibraries field if non-nil, zero value otherwise.

### GetContentLibrariesOk

`func (o *VcenterNamespacesInstancesVMServiceSpec) GetContentLibrariesOk() (*[]string, bool)`

GetContentLibrariesOk returns a tuple with the ContentLibraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLibraries

`func (o *VcenterNamespacesInstancesVMServiceSpec) SetContentLibraries(v []string)`

SetContentLibraries sets ContentLibraries field to given value.

### HasContentLibraries

`func (o *VcenterNamespacesInstancesVMServiceSpec) HasContentLibraries() bool`

HasContentLibraries returns a boolean if a field has been set.

### GetVmClasses

`func (o *VcenterNamespacesInstancesVMServiceSpec) GetVmClasses() []string`

GetVmClasses returns the VmClasses field if non-nil, zero value otherwise.

### GetVmClassesOk

`func (o *VcenterNamespacesInstancesVMServiceSpec) GetVmClassesOk() (*[]string, bool)`

GetVmClassesOk returns a tuple with the VmClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmClasses

`func (o *VcenterNamespacesInstancesVMServiceSpec) SetVmClasses(v []string)`

SetVmClasses sets VmClasses field to given value.

### HasVmClasses

`func (o *VcenterNamespacesInstancesVMServiceSpec) HasVmClasses() bool`

HasVmClasses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


