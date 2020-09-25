package bitbucket

import "context"

func (a *RepositoriesApiService) SearchRepositories(ctx context.Context) ([]Repository, error) {
	repositories := []Repository{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.SearchRepositoriesPaged(ctx).Limit(limit).Start(start).Execute()
		if err != nil {
			return nil, err
		}

		repositories = append(repositories, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return repositories, nil
}

func (a *RepositoriesApiService) GetRepositories(ctx context.Context, projectKey string) ([]Repository, error) {
	repositories := []Repository{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetRepositoriesPaged(ctx, projectKey).Limit(limit).Start(start).Execute()
		if err != nil {
			return nil, err
		}

		repositories = append(repositories, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return repositories, nil
}
