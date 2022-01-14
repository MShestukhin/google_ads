package tasks

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewAdCampaignsLoader(
	api google_api.GoogleApiService,
	campaignRepo *repositories.AdCampaignsRepository,
) *AdCampaignsLoader {
	return &AdCampaignsLoader{
		googleApi:    api,
		campaignRepo: campaignRepo,
	}
}

type AdCampaignsLoader struct {
	googleApi    google_api.GoogleApiService
	campaignRepo *repositories.AdCampaignsRepository
}

func (l *AdCampaignsLoader) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var ctx = context.TODO()
	resp, err := l.googleApi.GetCampaigns(ctx, &google_api.GetCampaignsRequest{
		CustomerId: args.AdAccountId,
		Token:      args.Token,
	})

	if err != nil {
		return err
	}

	var (
		campaigns   []*models.AdCampaign
		campaignIds []int64
	)

	for _, c := range resp.Campaign {
		campaign := &models.AdCampaign{}
		campaign.FromApi(c)

		campaigns = append(campaigns, campaign)
		campaignIds = append(campaignIds, campaign.Id)
	}

	_, err = l.campaignRepo.Save(campaigns)
	if err != nil {
		return err
	}

	if err = l.campaignRepo.AttachToBrand(args.BrandId, campaignIds); err != nil {
		return err
	}

	return nil
}
