package bitbucket

import "context"

func (a WebhookApiService) GetWebhooks(ctx context.Context, projectKey, repositorySlug string) ([]Webhook, error) {
	webhooks := []Webhook{}

	start := int32(0)
	limit := DefaultPageLimit
	morePages := true

	for morePages {
		result, _, err := a.GetWebhooksPaged(ctx, projectKey, repositorySlug).Start(start).Limit(limit).Execute()
		if err != nil {
			return nil, err
		}

		webhooks = append(webhooks, result.Values...)

		morePages = !result.IsLastPage
		if morePages {
			start = *result.NextPageStart
		}
	}

	return webhooks, nil
}
