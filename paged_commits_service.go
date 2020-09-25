package bitbucket

import "context"

func (a CommitsApiService) GetCommits(ctx context.Context, projectKey string, repositorySlug string) ([]Commit, error) {
	commits := []Commit{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetCommitsPaged(ctx, projectKey, repositorySlug).Limit(limit).Start(start).Execute()
		if err != nil {
			return nil, err
		}

		commits = append(commits, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return commits, nil
}
