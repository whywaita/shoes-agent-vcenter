# VcenterDeploymentQuestionQuestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the question raised. | 
**Question** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 
**Type** | [**VcenterDeploymentQuestionQuestionType**](VcenterDeploymentQuestionQuestionType.md) |  | 
**DefaultAnswer** | **string** | Default answer value. | 
**PossibleAnswers** | **[]string** | Possible answers values. | 

## Methods

### NewVcenterDeploymentQuestionQuestion

`func NewVcenterDeploymentQuestionQuestion(id string, question VapiStdLocalizableMessage, type_ VcenterDeploymentQuestionQuestionType, defaultAnswer string, possibleAnswers []string, ) *VcenterDeploymentQuestionQuestion`

NewVcenterDeploymentQuestionQuestion instantiates a new VcenterDeploymentQuestionQuestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentQuestionQuestionWithDefaults

`func NewVcenterDeploymentQuestionQuestionWithDefaults() *VcenterDeploymentQuestionQuestion`

NewVcenterDeploymentQuestionQuestionWithDefaults instantiates a new VcenterDeploymentQuestionQuestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterDeploymentQuestionQuestion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterDeploymentQuestionQuestion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterDeploymentQuestionQuestion) SetId(v string)`

SetId sets Id field to given value.


### GetQuestion

`func (o *VcenterDeploymentQuestionQuestion) GetQuestion() VapiStdLocalizableMessage`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *VcenterDeploymentQuestionQuestion) GetQuestionOk() (*VapiStdLocalizableMessage, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *VcenterDeploymentQuestionQuestion) SetQuestion(v VapiStdLocalizableMessage)`

SetQuestion sets Question field to given value.


### GetType

`func (o *VcenterDeploymentQuestionQuestion) GetType() VcenterDeploymentQuestionQuestionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterDeploymentQuestionQuestion) GetTypeOk() (*VcenterDeploymentQuestionQuestionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterDeploymentQuestionQuestion) SetType(v VcenterDeploymentQuestionQuestionType)`

SetType sets Type field to given value.


### GetDefaultAnswer

`func (o *VcenterDeploymentQuestionQuestion) GetDefaultAnswer() string`

GetDefaultAnswer returns the DefaultAnswer field if non-nil, zero value otherwise.

### GetDefaultAnswerOk

`func (o *VcenterDeploymentQuestionQuestion) GetDefaultAnswerOk() (*string, bool)`

GetDefaultAnswerOk returns a tuple with the DefaultAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAnswer

`func (o *VcenterDeploymentQuestionQuestion) SetDefaultAnswer(v string)`

SetDefaultAnswer sets DefaultAnswer field to given value.


### GetPossibleAnswers

`func (o *VcenterDeploymentQuestionQuestion) GetPossibleAnswers() []string`

GetPossibleAnswers returns the PossibleAnswers field if non-nil, zero value otherwise.

### GetPossibleAnswersOk

`func (o *VcenterDeploymentQuestionQuestion) GetPossibleAnswersOk() (*[]string, bool)`

GetPossibleAnswersOk returns a tuple with the PossibleAnswers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleAnswers

`func (o *VcenterDeploymentQuestionQuestion) SetPossibleAnswers(v []string)`

SetPossibleAnswers sets PossibleAnswers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


