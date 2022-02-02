# VapiStdErrorsUnauthenticated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Challenge** | Pointer to **string** | Indicates the authentication challenges applicable to the target API provider. It can be used by a client to discover the correct authentication scheme to use. The exact syntax of the value is defined by the specific provider, the protocol and authentication schemes used. &lt;p&gt; For example, a provider using REST may adhere to the WWW-Authenticate HTTP header specification, RFC7235, section 4.1. In this case an example challenge value may be: SIGN realm&#x3D;\&quot;27da1358-2ba4-11e9-b210-d663bd873d93\&quot;,sts&#x3D;\&quot;http://vcenter/sso?vsphere.local\&quot;, Basic realm&#x3D;\&quot;vCenter\&quot; | [optional] 
**Messages** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Stack of one or more localizable messages for human {@term error} consumers. &lt;p&gt; The message at the top of the stack (first in the list) describes the {@term error} from the perspective of the {@term operation} the client invoked. Each subsequent message in the stack describes the \&quot;cause\&quot; of the prior message. | 
**Data** | Pointer to **map[string]interface{}** | Data to facilitate clients responding to the {@term operation} reporting a standard {@term error} to indicating that it was unable to complete successfully. &lt;p&gt; {@term Operations} may provide data that clients can use when responding to {@term errors}.  Since the data that clients need may be specific to the context of the {@term operation} reporting the {@term error}, different {@term operations} that report the same {@term error} may provide different data in the {@term error}.  The documentation for each each {@term operation} will describe what, if any, data it provides for each {@term error} it reports. The {@link ArgumentLocations}, {@link FileLocations}, and {@link TransientIndication} {@term structures} are intended as possible values for this {@term field}.  {@link vapi.std.DynamicID} may also be useful as a value for this {@term field} (although that is not its primary purpose).  Some {@term services} may provide their own specific {@term structures} for use as the value of this {@term field} when reporting {@term errors} from their {@term operations}. | [optional] 
**ErrorType** | Pointer to [**VapiStdErrorsErrorType**](VapiStdErrorsErrorType.md) |  | [optional] 

## Methods

### NewVapiStdErrorsUnauthenticated

`func NewVapiStdErrorsUnauthenticated(messages []VapiStdLocalizableMessage, ) *VapiStdErrorsUnauthenticated`

NewVapiStdErrorsUnauthenticated instantiates a new VapiStdErrorsUnauthenticated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdErrorsUnauthenticatedWithDefaults

`func NewVapiStdErrorsUnauthenticatedWithDefaults() *VapiStdErrorsUnauthenticated`

NewVapiStdErrorsUnauthenticatedWithDefaults instantiates a new VapiStdErrorsUnauthenticated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChallenge

`func (o *VapiStdErrorsUnauthenticated) GetChallenge() string`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *VapiStdErrorsUnauthenticated) GetChallengeOk() (*string, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *VapiStdErrorsUnauthenticated) SetChallenge(v string)`

SetChallenge sets Challenge field to given value.

### HasChallenge

`func (o *VapiStdErrorsUnauthenticated) HasChallenge() bool`

HasChallenge returns a boolean if a field has been set.

### GetMessages

`func (o *VapiStdErrorsUnauthenticated) GetMessages() []VapiStdLocalizableMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VapiStdErrorsUnauthenticated) GetMessagesOk() (*[]VapiStdLocalizableMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VapiStdErrorsUnauthenticated) SetMessages(v []VapiStdLocalizableMessage)`

SetMessages sets Messages field to given value.


### GetData

`func (o *VapiStdErrorsUnauthenticated) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VapiStdErrorsUnauthenticated) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VapiStdErrorsUnauthenticated) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *VapiStdErrorsUnauthenticated) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *VapiStdErrorsUnauthenticated) GetErrorType() VapiStdErrorsErrorType`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *VapiStdErrorsUnauthenticated) GetErrorTypeOk() (*VapiStdErrorsErrorType, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *VapiStdErrorsUnauthenticated) SetErrorType(v VapiStdErrorsErrorType)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *VapiStdErrorsUnauthenticated) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


