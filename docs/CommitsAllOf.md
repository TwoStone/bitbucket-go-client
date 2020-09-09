# CommitsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]Commit**](commit.md) |  | 
**AuthorCount** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommitsAllOf

`func NewCommitsAllOf(values []Commit, ) *CommitsAllOf`

NewCommitsAllOf instantiates a new CommitsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsAllOfWithDefaults

`func NewCommitsAllOfWithDefaults() *CommitsAllOf`

NewCommitsAllOfWithDefaults instantiates a new CommitsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *CommitsAllOf) GetValues() []Commit`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CommitsAllOf) GetValuesOk() (*[]Commit, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CommitsAllOf) SetValues(v []Commit)`

SetValues sets Values field to given value.


### GetAuthorCount

`func (o *CommitsAllOf) GetAuthorCount() int32`

GetAuthorCount returns the AuthorCount field if non-nil, zero value otherwise.

### GetAuthorCountOk

`func (o *CommitsAllOf) GetAuthorCountOk() (*int32, bool)`

GetAuthorCountOk returns a tuple with the AuthorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorCount

`func (o *CommitsAllOf) SetAuthorCount(v int32)`

SetAuthorCount sets AuthorCount field to given value.

### HasAuthorCount

`func (o *CommitsAllOf) HasAuthorCount() bool`

HasAuthorCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *CommitsAllOf) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CommitsAllOf) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CommitsAllOf) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CommitsAllOf) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


