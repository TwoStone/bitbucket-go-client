# Commits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** |  | 
**Limit** | Pointer to **int32** |  | 
**Start** | Pointer to **int32** |  | 
**IsLastPage** | Pointer to **bool** |  | 
**NextPageStart** | Pointer to **int32** |  | [optional] 
**Values** | Pointer to [**[]Commit**](commit.md) |  | 
**AuthorCount** | Pointer to **int32** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommits

`func NewCommits(size int32, limit int32, start int32, isLastPage bool, values []Commit, ) *Commits`

NewCommits instantiates a new Commits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitsWithDefaults

`func NewCommitsWithDefaults() *Commits`

NewCommitsWithDefaults instantiates a new Commits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *Commits) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Commits) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Commits) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLimit

`func (o *Commits) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Commits) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Commits) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetStart

`func (o *Commits) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Commits) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Commits) SetStart(v int32)`

SetStart sets Start field to given value.


### GetIsLastPage

`func (o *Commits) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *Commits) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *Commits) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.


### GetNextPageStart

`func (o *Commits) GetNextPageStart() int32`

GetNextPageStart returns the NextPageStart field if non-nil, zero value otherwise.

### GetNextPageStartOk

`func (o *Commits) GetNextPageStartOk() (*int32, bool)`

GetNextPageStartOk returns a tuple with the NextPageStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageStart

`func (o *Commits) SetNextPageStart(v int32)`

SetNextPageStart sets NextPageStart field to given value.

### HasNextPageStart

`func (o *Commits) HasNextPageStart() bool`

HasNextPageStart returns a boolean if a field has been set.

### GetValues

`func (o *Commits) GetValues() []Commit`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Commits) GetValuesOk() (*[]Commit, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Commits) SetValues(v []Commit)`

SetValues sets Values field to given value.


### GetAuthorCount

`func (o *Commits) GetAuthorCount() int32`

GetAuthorCount returns the AuthorCount field if non-nil, zero value otherwise.

### GetAuthorCountOk

`func (o *Commits) GetAuthorCountOk() (*int32, bool)`

GetAuthorCountOk returns a tuple with the AuthorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorCount

`func (o *Commits) SetAuthorCount(v int32)`

SetAuthorCount sets AuthorCount field to given value.

### HasAuthorCount

`func (o *Commits) HasAuthorCount() bool`

HasAuthorCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *Commits) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *Commits) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *Commits) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *Commits) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


