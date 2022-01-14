package tasks

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewCampagnsInsightLoader(
	api google_api.GoogleApiService,
	repo *repositories.CampaignInsightRepository,
) *CampagnsInsightLoader {
	return &CampagnsInsightLoader{
		googleApi: api,
		repo:      repo,
	}
}

type CampagnsInsightLoader struct {
	googleApi google_api.GoogleApiService
	repo      *repositories.CampaignInsightRepository
}

func (l *CampagnsInsightLoader) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var ctx = context.TODO()
	resp, err := l.googleApi.GetCampaignCriteria(ctx, &google_api.GetCampaignCriteriaRequest{
		CustomerId: args.AdAccountId,
		Token:      args.Token,
	})

	if err != nil {
		return err
	}

	insights := make([]*models.CampaignInsight, 0, len(resp.CampaignCriteria))
	for _, insight := range resp.CampaignCriteria {
		insights = append(insights, &models.CampaignInsight{
			CriterionId:  insight.CriterionId,
			ResourceName: insight.ResourceName,
			Campaign:     insight.Campaign,
			DisplayName:  insight.DisplayName,
			BidModifier:  insight.BidModifier,
			Negative:     insight.Negative,
		})
	}

	_, err = l.repo.Create(insights)

	return nil
}
