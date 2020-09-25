package bitbucket

import "context"

func (a *ProjectsApiService) GetProjects(ctx context.Context) ([]Project, error) {
	projects := []Project{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetProjectsPaged(ctx).Limit(limit).Start(start).Execute()
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
