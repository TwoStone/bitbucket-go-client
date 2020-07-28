# Branches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Size** | Pointer to **int32** |  | 
**Limit** | Pointer to **int32** |  | 
**Start** | Pointer to **int32** |  | 
**IsLastPage** | Pointer to **bool** |  | 
**NextPageStart** | Pointer to **int32** |  | [optional] 
**Values** | Pointer to [**[]Branch**](branch.md) |  | [optional] 

## Methods

### NewBranches

`func NewBranches(size int32, limit int32, start int32, isLastPage bool, ) *Branches`

NewBranches instantiates a new Branches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchesWithDefaults

`func NewBranchesWithDefaults() *Branches`

NewBranchesWithDefaults instantiates a new Branches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSize

`func (o *Branches) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Branches) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Branches) SetSize(v int32)`

SetSize sets Size field to given value.


### GetLimit

`func (o *Branches) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Branches) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Branches) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetStart

`func (o *Branches) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Branches) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Branches) SetStart(v int32)`

SetStart sets Start field to given value.


### GetIsLastPage

`func (o *Branches) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *Branches) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *Branches) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.


### GetNextPageStart

`func (o *Branches) GetNextPageStart() int32`

GetNextPageStart returns the NextPageStart field if non-nil, zero value otherwise.

### GetNextPageStartOk

`func (o *Branches) GetNextPageStartOk() (*int32, bool)`

GetNextPageStartOk returns a tuple with the NextPageStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageStart

`func (o *Branches) SetNextPageStart(v int32)`

SetNextPageStart sets NextPageStart field to given value.

### HasNextPageStart

`func (o *Branches) HasNextPageStart() bool`

HasNextPageStart returns a boolean if a field has been set.

### GetValues

`func (o *Branches) GetValues() []Branch`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Branches) GetValuesOk() (*[]Branch, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Branches) SetValues(v []Branch)`

SetValues sets Values field to given value.

### HasValues

`func (o *Branches) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

