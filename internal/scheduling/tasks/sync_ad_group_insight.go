package tasks

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewAdGroupInsightLoader(
	api google_api.GoogleApiService,
	repo *repositories.AdGroupsInsightRepository,
) *AdGroupInsightLoader {
	return &AdGroupInsightLoader{
		googleApi: api,
		repo:      repo,
	}
}

type AdGroupInsightLoader struct {
	googleApi google_api.GoogleApiService
	repo      *repositories.AdGroupsInsightRepository
}

func (l *AdGroupInsightLoader) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var ctx = context.TODO()
	resp, err := l.googleApi.GetAdGroupsCriteria(ctx, &google_api.GetAdGroupsCriteriaRequest{
		CustomerId: args.AdAccountId,
		Token:      args.Token,
	})

	if err != nil {
		return err
	}

	insights := make([]*models.AdGroupInsight, 0, len(resp.AdGroupsCriteria))
	for _, insight := range resp.AdGroupsCriteria {
		insights = append(insights, &models.AdGroupInsight{
			Keyword:      insight.Keyword,
			AdGroupName:  insight.AdGroupName,
			CampaignName: insight.CampaignName,
			Impressions:  insight.Impressions,
			Clicks:       insight.Clicks,
			Ctr:          insight.Ctr,
			AverageCpc:   insight.AverageCpc,
		})
	}

	_, err = l.repo.Create(insights)

	return err
}
