# VcenterGuestLinuxConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | [**VcenterGuestHostnameGenerator**](VcenterGuestHostnameGenerator.md) |  | 
**Domain** | **string** | The fully qualified domain name. | 
**TimeZone** | Pointer to **string** | The case-sensitive time zone, such as Europe/Sofia. Valid time zone values are based on the tz (time zone) database used by Linux. The values are strings (string) in the form \&quot;Area/Location,\&quot; in which Area is a continent or ocean name, and Location is the city, island, or other regional designation.   See the https://kb.vmware.com/kb/2145518 for a list of supported time zones for different versions in Linux.  If unset, time zone is not modified inside guest operating system. | [optional] 
**ScriptText** | Pointer to **string** | The script to run before and after Linux guest customization.  The max size of the script is 1500 bytes. As long as the script (shell, perl, python...) has the right \&quot;#!\&quot; in the header, it is supported. The caller should not assume any environment variables when the script is run.   The script is invoked by the customization engine using the command line: 1) with argument \&quot;precustomization\&quot; before customization, 2) with argument \&quot;postcustomization\&quot; after customization. The script should parse this argument and implement pre-customization or post-customization task code details in the corresponding block.    A Linux shell script example:     #!/bin/sh  if [ x$1 &#x3D;&#x3D; x\&quot;precustomization\&quot; ]; then  echo \&quot;Do Precustomization tasks\&quot;  #code for pre-customization actions...  elif [ x$1 &#x3D;&#x3D; x\&quot;postcustomization\&quot; ]; then  echo \&quot;Do Postcustomization tasks\&quot;  #code for post-customization actions...  fi    If unset, no script will be executed. | [optional] 

## Methods

### NewVcenterGuestLinuxConfiguration

`func NewVcenterGuestLinuxConfiguration(hostname VcenterGuestHostnameGenerator, domain string, ) *VcenterGuestLinuxConfiguration`

NewVcenterGuestLinuxConfiguration instantiates a new VcenterGuestLinuxConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestLinuxConfigurationWithDefaults

`func NewVcenterGuestLinuxConfigurationWithDefaults() *VcenterGuestLinuxConfiguration`

NewVcenterGuestLinuxConfigurationWithDefaults instantiates a new VcenterGuestLinuxConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *VcenterGuestLinuxConfiguration) GetHostname() VcenterGuestHostnameGenerator`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *VcenterGuestLinuxConfiguration) GetHostnameOk() (*VcenterGuestHostnameGenerator, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *VcenterGuestLinuxConfiguration) SetHostname(v VcenterGuestHostnameGenerator)`

SetHostname sets Hostname field to given value.


### GetDomain

`func (o *VcenterGuestLinuxConfiguration) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *VcenterGuestLinuxConfiguration) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *VcenterGuestLinuxConfiguration) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetTimeZone

`func (o *VcenterGuestLinuxConfiguration) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *VcenterGuestLinuxConfiguration) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *VcenterGuestLinuxConfiguration) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *VcenterGuestLinuxConfiguration) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetScriptText

`func (o *VcenterGuestLinuxConfiguration) GetScriptText() string`

GetScriptText returns the ScriptText field if non-nil, zero value otherwise.

### GetScriptTextOk

`func (o *VcenterGuestLinuxConfiguration) GetScriptTextOk() (*string, bool)`

GetScriptTextOk returns a tuple with the ScriptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptText

`func (o *VcenterGuestLinuxConfiguration) SetScriptText(v string)`

SetScriptText sets ScriptText field to given value.

### HasScriptText

`func (o *VcenterGuestLinuxConfiguration) HasScriptText() bool`

HasScriptText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


