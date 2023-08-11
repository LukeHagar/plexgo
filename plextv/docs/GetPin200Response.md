# GetPin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **interface{}** | PinID for use with authentication | [optional] 
**Code** | Pointer to **interface{}** |  | [optional] 
**Product** | Pointer to **interface{}** |  | [optional] 
**Trusted** | Pointer to **interface{}** |  | [optional] 
**Qr** | Pointer to **interface{}** | a link to a QR code hosted on plex.tv  The QR code redirects to the relevant &#x60;plex.tv/link&#x60; authentication page Which then prompts the user for the 4 Digit Link Pin  | [optional] 
**ClientIdentifier** | Pointer to **interface{}** |  | [optional] 
**Location** | Pointer to [**GetPin200ResponseLocation**](GetPin200ResponseLocation.md) |  | [optional] 
**ExpiresIn** | Pointer to **interface{}** |  | [optional] 
**CreatedAt** | Pointer to **interface{}** |  | [optional] 
**ExpiresAt** | Pointer to **interface{}** |  | [optional] 
**AuthToken** | Pointer to **interface{}** |  | [optional] 
**NewRegistration** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewGetPin200Response

`func NewGetPin200Response() *GetPin200Response`

NewGetPin200Response instantiates a new GetPin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPin200ResponseWithDefaults

`func NewGetPin200ResponseWithDefaults() *GetPin200Response`

NewGetPin200ResponseWithDefaults instantiates a new GetPin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPin200Response) GetId() interface{}`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPin200Response) GetIdOk() (*interface{}, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPin200Response) SetId(v interface{})`

SetId sets Id field to given value.

### HasId

`func (o *GetPin200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetPin200Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetPin200Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCode

`func (o *GetPin200Response) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPin200Response) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPin200Response) SetCode(v interface{})`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPin200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *GetPin200Response) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *GetPin200Response) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetProduct

`func (o *GetPin200Response) GetProduct() interface{}`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetPin200Response) GetProductOk() (*interface{}, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetPin200Response) SetProduct(v interface{})`

SetProduct sets Product field to given value.

### HasProduct

`func (o *GetPin200Response) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### SetProductNil

`func (o *GetPin200Response) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *GetPin200Response) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetTrusted

`func (o *GetPin200Response) GetTrusted() interface{}`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *GetPin200Response) GetTrustedOk() (*interface{}, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *GetPin200Response) SetTrusted(v interface{})`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *GetPin200Response) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.

### SetTrustedNil

`func (o *GetPin200Response) SetTrustedNil(b bool)`

 SetTrustedNil sets the value for Trusted to be an explicit nil

### UnsetTrusted
`func (o *GetPin200Response) UnsetTrusted()`

UnsetTrusted ensures that no value is present for Trusted, not even an explicit nil
### GetQr

`func (o *GetPin200Response) GetQr() interface{}`

GetQr returns the Qr field if non-nil, zero value otherwise.

### GetQrOk

`func (o *GetPin200Response) GetQrOk() (*interface{}, bool)`

GetQrOk returns a tuple with the Qr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQr

`func (o *GetPin200Response) SetQr(v interface{})`

SetQr sets Qr field to given value.

### HasQr

`func (o *GetPin200Response) HasQr() bool`

HasQr returns a boolean if a field has been set.

### SetQrNil

`func (o *GetPin200Response) SetQrNil(b bool)`

 SetQrNil sets the value for Qr to be an explicit nil

### UnsetQr
`func (o *GetPin200Response) UnsetQr()`

UnsetQr ensures that no value is present for Qr, not even an explicit nil
### GetClientIdentifier

`func (o *GetPin200Response) GetClientIdentifier() interface{}`

GetClientIdentifier returns the ClientIdentifier field if non-nil, zero value otherwise.

### GetClientIdentifierOk

`func (o *GetPin200Response) GetClientIdentifierOk() (*interface{}, bool)`

GetClientIdentifierOk returns a tuple with the ClientIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdentifier

`func (o *GetPin200Response) SetClientIdentifier(v interface{})`

SetClientIdentifier sets ClientIdentifier field to given value.

### HasClientIdentifier

`func (o *GetPin200Response) HasClientIdentifier() bool`

HasClientIdentifier returns a boolean if a field has been set.

### SetClientIdentifierNil

`func (o *GetPin200Response) SetClientIdentifierNil(b bool)`

 SetClientIdentifierNil sets the value for ClientIdentifier to be an explicit nil

### UnsetClientIdentifier
`func (o *GetPin200Response) UnsetClientIdentifier()`

UnsetClientIdentifier ensures that no value is present for ClientIdentifier, not even an explicit nil
### GetLocation

`func (o *GetPin200Response) GetLocation() GetPin200ResponseLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetPin200Response) GetLocationOk() (*GetPin200ResponseLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetPin200Response) SetLocation(v GetPin200ResponseLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetPin200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetPin200Response) GetExpiresIn() interface{}`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetPin200Response) GetExpiresInOk() (*interface{}, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetPin200Response) SetExpiresIn(v interface{})`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetPin200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### SetExpiresInNil

`func (o *GetPin200Response) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *GetPin200Response) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
### GetCreatedAt

`func (o *GetPin200Response) GetCreatedAt() interface{}`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetPin200Response) GetCreatedAtOk() (*interface{}, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetPin200Response) SetCreatedAt(v interface{})`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetPin200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *GetPin200Response) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *GetPin200Response) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiresAt

`func (o *GetPin200Response) GetExpiresAt() interface{}`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetPin200Response) GetExpiresAtOk() (*interface{}, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetPin200Response) SetExpiresAt(v interface{})`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetPin200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *GetPin200Response) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *GetPin200Response) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetAuthToken

`func (o *GetPin200Response) GetAuthToken() interface{}`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *GetPin200Response) GetAuthTokenOk() (*interface{}, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *GetPin200Response) SetAuthToken(v interface{})`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *GetPin200Response) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### SetAuthTokenNil

`func (o *GetPin200Response) SetAuthTokenNil(b bool)`

 SetAuthTokenNil sets the value for AuthToken to be an explicit nil

### UnsetAuthToken
`func (o *GetPin200Response) UnsetAuthToken()`

UnsetAuthToken ensures that no value is present for AuthToken, not even an explicit nil
### GetNewRegistration

`func (o *GetPin200Response) GetNewRegistration() interface{}`

GetNewRegistration returns the NewRegistration field if non-nil, zero value otherwise.

### GetNewRegistrationOk

`func (o *GetPin200Response) GetNewRegistrationOk() (*interface{}, bool)`

GetNewRegistrationOk returns a tuple with the NewRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRegistration

`func (o *GetPin200Response) SetNewRegistration(v interface{})`

SetNewRegistration sets NewRegistration field to given value.

### HasNewRegistration

`func (o *GetPin200Response) HasNewRegistration() bool`

HasNewRegistration returns a boolean if a field has been set.

### SetNewRegistrationNil

`func (o *GetPin200Response) SetNewRegistrationNil(b bool)`

 SetNewRegistrationNil sets the value for NewRegistration to be an explicit nil

### UnsetNewRegistration
`func (o *GetPin200Response) UnsetNewRegistration()`

UnsetNewRegistration ensures that no value is present for NewRegistration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


