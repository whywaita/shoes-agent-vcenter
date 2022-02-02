# VapiStdErrorsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Stack of one or more localizable messages for human {@term error} consumers. &lt;p&gt; The message at the top of the stack (first in the list) describes the {@term error} from the perspective of the {@term operation} the client invoked. Each subsequent message in the stack describes the \&quot;cause\&quot; of the prior message. | 
**Data** | Pointer to **map[string]interface{}** | Data to facilitate clients responding to the {@term operation} reporting a standard {@term error} to indicating that it was unable to complete successfully. &lt;p&gt; {@term Operations} may provide data that clients can use when responding to {@term errors}.  Since the data that clients need may be specific to the context of the {@term operation} reporting the {@term error}, different {@term operations} that report the same {@term error} may provide different data in the {@term error}.  The documentation for each each {@term operation} will describe what, if any, data it provides for each {@term error} it reports. The {@link ArgumentLocations}, {@link FileLocations}, and {@link TransientIndication} {@term structures} are intended as possible values for this {@term field}.  {@link vapi.std.DynamicID} may also be useful as a value for this {@term field} (although that is not its primary purpose).  Some {@term services} may provide their own specific {@term structures} for use as the value of this {@term field} when reporting {@term errors} from their {@term operations}. | [optional] 
**ErrorType** | Pointer to [**VapiStdErrorsErrorType**](VapiStdErrorsErrorType.md) |  | [optional] 

## Methods

### NewVapiStdErrorsError

`func NewVapiStdErrorsError(messages []VapiStdLocalizableMessage, ) *VapiStdErrorsError`

NewVapiStdErrorsError instantiates a new VapiStdErrorsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVapiStdErrorsErrorWithDefaults

`func NewVapiStdErrorsErrorWithDefaults() *VapiStdErrorsError`

NewVapiStdErrorsErrorWithDefaults instantiates a new VapiStdErrorsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *VapiStdErrorsError) GetMessages() []VapiStdLocalizableMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VapiStdErrorsError) GetMessagesOk() (*[]VapiStdLocalizableMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VapiStdErrorsError) SetMessages(v []VapiStdLocalizableMessage)`

SetMessages sets Messages field to given value.


### GetData

`func (o *VapiStdErrorsError) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VapiStdErrorsError) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VapiStdErrorsError) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *VapiStdErrorsError) HasData() bool`

HasData returns a boolean if a field has been set.

### GetErrorType

`func (o *VapiStdErrorsError) GetErrorType() VapiStdErrorsErrorType`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *VapiStdErrorsError) GetErrorTypeOk() (*VapiStdErrorsErrorType, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *VapiStdErrorsError) SetErrorType(v VapiStdErrorsErrorType)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *VapiStdErrorsError) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


