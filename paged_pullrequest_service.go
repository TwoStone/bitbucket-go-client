package bitbucket

import "context"

func (a *PullRequestsApiService) GetPullRequests(ctx context.Context, projectKey string, repositorySlug string) ([]PullRequest, error) {
	pullRequests := []PullRequest{}

	morePages := true
	start := int32(0)
	limit := DefaultPageLimit

	for morePages {
		result, _, err := a.GetPullRequestsPaged(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()
		if err != nil {
			return nil, err
		}

		pullRequests = append(pullRequests, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return pullRequests, nil
}
