package bitbucket

import "context"

const bitbucketDefaultPageLimit = int32(25) //nolint:gomnd

type Service struct {
	*DefaultApiService
}

func NewService(api *DefaultApiService) *Service {
	return &Service{DefaultApiService: api}
}

func (a *Service) SearchAllRepositories(ctx context.Context) ([]Repository, error) {
	repositories := []Repository{}

	start := int32(0)
	limit := bitbucketDefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.SearchRepositories(ctx).Limit(limit).Start(start).Execute()
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

func (a *Service) GetAllProjects(ctx context.Context) ([]Project, error) {
	projects := []Project{}

	start := int32(0)
	limit := bitbucketDefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetProjects(ctx).Limit(limit).Start(start).Execute()
		if err != nil {
			return nil, err
		}

		projects = append(projects, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return projects, nil
}

func (a *Service) GetAllRepositories(ctx context.Context, projectKey string) ([]Repository, error) {
	repositories := []Repository{}

	start := int32(0)
	limit := bitbucketDefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetRepositories(ctx, projectKey).Limit(limit).Start(start).Execute()
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

func (a *Service) GetAllBranches(ctx context.Context, projectKey string, repositorySlug string) ([]Branch, error) {
	branches := []Branch{}

	start := int32(0)
	limit := bitbucketDefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetBranches(ctx, projectKey, repositorySlug).Limit(limit).Start(start).Execute()
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
