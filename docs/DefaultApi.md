# \DefaultApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseRepository**](DefaultApi.md#BrowseRepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse | browseRepository
[**BrowseRepositoryPath**](DefaultApi.md#BrowseRepositoryPath) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | browseRepositoryPath
[**GetBranches**](DefaultApi.md#GetBranches) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Your GET endpoint
[**GetProject**](DefaultApi.md#GetProject) | **Get** /rest/api/1.0/projects/{projectKey} | REST resource for working with projects
[**GetProjects**](DefaultApi.md#GetProjects) | **Get** /rest/api/1.0/projects | REST resource for working with projects
[**GetRepositories**](DefaultApi.md#GetRepositories) | **Get** /rest/api/1.0/projects/{projectKey}/repos | REST resource for working with repositories
[**GetRepository**](DefaultApi.md#GetRepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug} | REST resource for working with repositories
[**SearchRepositories**](DefaultApi.md#SearchRepositories) | **Get** /rest/api/1.0/repos | REST resource for searching through repositories



## BrowseRepository

> Directory BrowseRepository(ctx, projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()

browseRepository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    at := "at_example" // string | the commit ID or ref to retrieve the content for. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BrowseRepository(context.Background(), projectKey, repositorySlug).At(at).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BrowseRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseRepository`: Directory
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BrowseRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **at** | **string** | the commit ID or ref to retrieve the content for. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**Directory**](directory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BrowseRepositoryPath

> FileOrDirectory BrowseRepositoryPath(ctx, projectKey, repositorySlug, path).At(at).Start(start).Limit(limit).Execute()

browseRepositoryPath



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    path := "path_example" // string | 
    at := "at_example" // string | the commit ID or ref to retrieve the content for. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.BrowseRepositoryPath(context.Background(), projectKey, repositorySlug, path).At(at).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.BrowseRepositoryPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BrowseRepositoryPath`: FileOrDirectory
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.BrowseRepositoryPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**path** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBrowseRepositoryPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **at** | **string** | the commit ID or ref to retrieve the content for. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**FileOrDirectory**](fileOrDirectory.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranches

> Branches GetBranches(ctx, projectKey, repositorySlug).Base(base).Details(details).FilterText(filterText).OrderBy(orderBy).BoostMatches(boostMatches).Start(start).Limit(limit).Execute()

Your GET endpoint



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 
    base := "base_example" // string | base branch or tag to compare each branch to (for the metadata providers that uses that information) (optional)
    details := true // bool | whether to retrieve plugin-provided metadata about each branch (optional)
    filterText := "filterText_example" // string | the text to match on (optional)
    orderBy := "orderBy_example" // string | ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) (optional)
    boostMatches := true // bool | controls whether exact and prefix matches will be boosted to the top (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetBranches(context.Background(), projectKey, repositorySlug).Base(base).Details(details).FilterText(filterText).OrderBy(orderBy).BoostMatches(boostMatches).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBranches`: Branches
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **base** | **string** | base branch or tag to compare each branch to (for the metadata providers that uses that information) | 
 **details** | **bool** | whether to retrieve plugin-provided metadata about each branch | 
 **filterText** | **string** | the text to match on | 
 **orderBy** | **string** | ordering of refs either ALPHABETICAL (by name) or MODIFICATION (last updated) | 
 **boostMatches** | **bool** | controls whether exact and prefix matches will be boosted to the top | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**Branches**](branches.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, projectKey).Execute()

REST resource for working with projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetProject(context.Background(), projectKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml, 

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> Projects GetProjects(ctx).Name(name).Permission(permission).Start(start).Limit(limit).Execute()

REST resource for working with projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | name to filter by (optional)
    permission := "permission_example" // string | permission to filter by (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetProjects(context.Background(), ).Name(name).Permission(permission).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: Projects
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name to filter by | 
 **permission** | **string** | permission to filter by | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**Projects**](projects.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> Repositories GetRepositories(ctx, projectKey).Limit(limit).Start(start).Execute()

REST resource for working with repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetRepositories(context.Background(), projectKey).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepositories`: Repositories
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **start** | **int32** |  | 

### Return type

[**Repositories**](repositories.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> Repository GetRepository(ctx, projectKey, repositorySlug).Execute()

REST resource for working with repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    projectKey := "projectKey_example" // string | 
    repositorySlug := "repositorySlug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetRepository(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Repository**](repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchRepositories

> Repositories SearchRepositories(ctx).Name(name).Projectname(projectname).Permission(permission).State(state).Visibility(visibility).Start(start).Limit(limit).Execute()

REST resource for searching through repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter's value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the name parameter will be stripped. (optional)
    projectname := "projectname_example" // string | (optional) if specified, this will limit the resulting repository list to ones whose project's name matches this parameter's value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the projectname parameter will be stripped. (optional)
    permission := "permission_example" // string | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit 'read' permission level will be assumed. The currently supported explicit permission values are REPO_READ, REPO_WRITE and REPO_ADMIN. (optional)
    state := "state_example" // string | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are AVAILABLE, INITIALISING and INITIALISATION_FAILED. Available since 5.13 (optional)
    visibility := "visibility_example" // string | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are public or private. (optional)
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.SearchRepositories(context.Background(), ).Name(name).Projectname(projectname).Permission(permission).State(state).Visibility(visibility).Start(start).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchRepositories`: Repositories
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose name matches this parameter&#39;s value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the name parameter will be stripped. | 
 **projectname** | **string** | (optional) if specified, this will limit the resulting repository list to ones whose project&#39;s name matches this parameter&#39;s value. The match will be done case-insensitive and any leading and/or trailing whitespace characters on the projectname parameter will be stripped. | 
 **permission** | **string** | (optional) if specified, it must be a valid repository permission level name and will limit the resulting repository list to ones that the requesting user has the specified permission level to. If not specified, the default implicit &#39;read&#39; permission level will be assumed. The currently supported explicit permission values are REPO_READ, REPO_WRITE and REPO_ADMIN. | 
 **state** | **string** | (optional) if specified, it must be a valid repository state name and will limit the resulting repository list to ones that are in the specified state. The currently supported explicit state values are AVAILABLE, INITIALISING and INITIALISATION_FAILED. Available since 5.13 | 
 **visibility** | **string** | (optional) if specified, this will limit the resulting repository list based on the repositories visibility. Valid values are public or private. | 
 **start** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

[**Repositories**](repositories.md)

### Authorization

[usernamePassword](../README.md#usernamePassword)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

