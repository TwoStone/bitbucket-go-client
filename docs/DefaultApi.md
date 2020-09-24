# \DefaultApi

All URIs are relative to *https://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BrowseRepository**](DefaultApi.md#BrowseRepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse | browseRepository
[**BrowseRepositoryPath**](DefaultApi.md#BrowseRepositoryPath) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/browse/{path} | browseRepositoryPath
[**CreateRepository**](DefaultApi.md#CreateRepository) | **Post** /rest/api/1.0/projects/{projectKey}/repos | Create repository
[**CreateWebhook**](DefaultApi.md#CreateWebhook) | **Post** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Create webhook
[**DeletePostWebhook**](DefaultApi.md#DeletePostWebhook) | **Delete** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations/{ID} | Delete post webhook
[**DeleteWebhook**](DefaultApi.md#DeleteWebhook) | **Delete** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Delete Webhook
[**GetBranches**](DefaultApi.md#GetBranches) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches | Get Branches
[**GetCommit**](DefaultApi.md#GetCommit) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId} | Get commit
[**GetCommits**](DefaultApi.md#GetCommits) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits | Get commits
[**GetDefaultBranch**](DefaultApi.md#GetDefaultBranch) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/branches/default | Get default branch
[**GetPostWebhooks**](DefaultApi.md#GetPostWebhooks) | **Get** /rest/webhook/1.0/projects/{projectKey}/repos/{repositorySlug}/configurations | Get Post Webhooks
[**GetProject**](DefaultApi.md#GetProject) | **Get** /rest/api/1.0/projects/{projectKey} | REST resource for working with projects
[**GetProjects**](DefaultApi.md#GetProjects) | **Get** /rest/api/1.0/projects | REST resource for working with projects
[**GetRepositories**](DefaultApi.md#GetRepositories) | **Get** /rest/api/1.0/projects/{projectKey}/repos | REST resource for working with repositories
[**GetRepository**](DefaultApi.md#GetRepository) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug} | REST resource for working with repositories
[**GetWebhook**](DefaultApi.md#GetWebhook) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Get Webhook
[**GetWebhooks**](DefaultApi.md#GetWebhooks) | **Get** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks | Get webhooks
[**PostBuildResult**](DefaultApi.md#PostBuildResult) | **Post** /rest/build-status/1.0/commits/{commitHash} | Post build-result
[**SearchRepositories**](DefaultApi.md#SearchRepositories) | **Get** /rest/api/1.0/repos | REST resource for searching through repositories
[**UpdateWebhook**](DefaultApi.md#UpdateWebhook) | **Put** /rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/webhooks/{webhookId} | Update webhook



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


## CreateRepository

> Repository CreateRepository(ctx, projectKey).CreateRepository(createRepository).Execute()

Create repository



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
    createRepository := openapiclient.createRepository{Name: "Name_example", ScmId: "ScmId_example", Forkable: false, DefaultBranch: "DefaultBranch_example", Public: false, Description: "Description_example"} // CreateRepository |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateRepository(context.Background(), projectKey).CreateRepository(createRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRepository** | [**CreateRepository**](CreateRepository.md) |  | 

### Return type

[**Repository**](repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx, projectKey, repositorySlug).Webhook(webhook).Execute()

Create webhook



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
    webhook := openapiclient.webhook{Id: 123, Name: "Name_example", CreateDate: 123, UpdatedDate: 123, Events: []WebhookEvent{openapiclient.webhookEvent{}), Configuration: openapiclient.webhook_configuration{Secret: "Secret_example"}, Url: "Url_example"} // Webhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateWebhook(context.Background(), projectKey, repositorySlug).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostWebhook

> DeletePostWebhook(ctx, projectKey, repositorySlug, iD).Execute()

Delete post webhook



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
    iD := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeletePostWebhook(context.Background(), projectKey, repositorySlug, iD).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeletePostWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**iD** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, projectKey, repositorySlug, webhookId).Execute()

Delete Webhook



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
    webhookId := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteWebhook(context.Background(), projectKey, repositorySlug, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

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

Get Branches



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


## GetCommit

> Commit GetCommit(ctx, projectKey, repositorySlug, commitId).Path(path).Execute()

Get commit



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
    commitId := "commitId_example" // string | 
    path := "path_example" // string | an optional path to filter the commit by. If supplied the details returned may not be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetCommit(context.Background(), projectKey, repositorySlug, commitId).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommit`: Commit
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**commitId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **path** | **string** | an optional path to filter the commit by. If supplied the details returned may not be for the specified commit. Instead, starting from the specified commit, they will be the details for the first commit affecting the specified path. | 

### Return type

[**Commit**](commit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommits

> Commits GetCommits(ctx, projectKey, repositorySlug).FollowRenames(followRenames).IgnoreMissing(ignoreMissing).Merges(merges).Path(path).Since(since).Until(until).WithCounts(withCounts).Limit(limit).Start(start).Execute()

Get commits



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
    followRenames := true // bool | if true, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. (optional) (default to false)
    ignoreMissing := true // bool | true to ignore missing commits, false otherwise (optional) (default to false)
    merges := "merges_example" // string |   if present, controls how merge commits should be filtered. Can be either exclude, to exclude merge commits, include, to include both merge commits and non-merge commits or only, to only return merge commits. (optional)
    path := "path_example" // string | an optional path to filter commits by (optional)
    since := "since_example" // string | the commit ID or ref (exclusively) to retrieve commits after (optional)
    until := "until_example" // string | the commit ID (SHA1) or ref (inclusively) to retrieve commits before (optional)
    withCounts := true // bool | optionally include the total number of commits and total number of unique authors (optional) (default to false)
    limit := 987 // int32 |  (optional)
    start := 987 // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetCommits(context.Background(), projectKey, repositorySlug).FollowRenames(followRenames).IgnoreMissing(ignoreMissing).Merges(merges).Path(path).Since(since).Until(until).WithCounts(withCounts).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCommits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommits`: Commits
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCommits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **followRenames** | **bool** | if true, the commit history of the specified file will be followed past renames. Only valid for a path to a single file. | [default to false]
 **ignoreMissing** | **bool** | true to ignore missing commits, false otherwise | [default to false]
 **merges** | **string** |   if present, controls how merge commits should be filtered. Can be either exclude, to exclude merge commits, include, to include both merge commits and non-merge commits or only, to only return merge commits. | 
 **path** | **string** | an optional path to filter commits by | 
 **since** | **string** | the commit ID or ref (exclusively) to retrieve commits after | 
 **until** | **string** | the commit ID (SHA1) or ref (inclusively) to retrieve commits before | 
 **withCounts** | **bool** | optionally include the total number of commits and total number of unique authors | [default to false]
 **limit** | **int32** |  | 
 **start** | **int32** |  | 

### Return type

[**Commits**](commits.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultBranch

> Branch GetDefaultBranch(ctx, projectKey, repositorySlug).Execute()

Get default branch



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
    resp, r, err := api_client.DefaultApi.GetDefaultBranch(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDefaultBranch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultBranch`: Branch
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDefaultBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Branch**](branch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostWebhooks

> []PostWebhook GetPostWebhooks(ctx, projectKey, repositorySlug).Execute()

Get Post Webhooks



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
    resp, r, err := api_client.DefaultApi.GetPostWebhooks(context.Background(), projectKey, repositorySlug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetPostWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPostWebhooks`: []PostWebhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetPostWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]PostWebhook**](postWebhook.md)

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


## GetWebhook

> Webhook GetWebhook(ctx, projectKey, repositorySlug, webhookId).Execute()

Get Webhook



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
    webhookId := 987 // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWebhook(context.Background(), projectKey, repositorySlug, webhookId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebhooks

> Webhooks GetWebhooks(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Event(event).Execute()

Get webhooks



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
    start := 987 // int32 |  (optional)
    limit := 987 // int32 |  (optional)
    event := "event_example" // string | list of {@link com.atlassian.webhooks.WebhookEvent} ids to filter for (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetWebhooks(context.Background(), projectKey, repositorySlug).Start(start).Limit(limit).Event(event).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhooks`: Webhooks
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWebhooks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **int32** |  | 
 **limit** | **int32** |  | 
 **event** | **string** | list of {@link com.atlassian.webhooks.WebhookEvent} ids to filter for | 

### Return type

[**Webhooks**](webhooks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostBuildResult

> PostBuildResult(ctx, commitHash).BuildResult(buildResult).Execute()

Post build-result



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
    commitHash := "commitHash_example" // string | 
    buildResult := openapiclient.buildResult{State: "State_example", Key: "Key_example", Name: "Name_example", Url: "Url_example", Description: "Description_example"} // BuildResult |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.PostBuildResult(context.Background(), commitHash).BuildResult(buildResult).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostBuildResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commitHash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBuildResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **buildResult** | [**BuildResult**](BuildResult.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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


## UpdateWebhook

> Webhook UpdateWebhook(ctx, projectKey, repositorySlug, webhookId).Webhook(webhook).Execute()

Update webhook



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
    webhookId := 987 // int32 | 
    webhook := openapiclient.webhook{Id: 123, Name: "Name_example", CreateDate: 123, UpdatedDate: 123, Events: []WebhookEvent{openapiclient.webhookEvent{}), Configuration: openapiclient.webhook_configuration{Secret: "Secret_example"}, Url: "Url_example"} // Webhook |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.UpdateWebhook(context.Background(), projectKey, repositorySlug, webhookId).Webhook(webhook).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebhook`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKey** | **string** |  | 
**repositorySlug** | **string** |  | 
**webhookId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](webhook.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

