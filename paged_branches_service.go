package bitbucket

import "context"

func (a *BranchesApiService) GetBranches(ctx context.Context, projectKey string, repositorySlug string) ([]Branch, error) {
	branches := []Branch{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetBranchesPaged(ctx, projectKey, repositorySlug).Limit(limit).Start(start).Execute()
		if err != nil {
			return nil, err
		}

		branches = append(branches, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return branches, nil
}
