package tasks

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewAdsLoader(
	api google_api.GoogleApiService,
	adsRepo *repositories.AdsRepository,
) *AdsLoader {
	return &AdsLoader{
		googleApi: api,
		adsRepo:   adsRepo,
	}
}

type AdsLoader struct {
	googleApi google_api.GoogleApiService
	adsRepo   *repositories.AdsRepository
}

func (l *AdsLoader) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var ctx = context.TODO()
	resp, err := l.googleApi.GetAdGroupsAd(ctx, &google_api.GetAdGroupsAdRequest{
		CustomerId: args.AdAccountId,
		Token:      args.Token,
	})

	if err != nil {
		return err
	}

	var ads []*models.Ads

	for _, apiAd := range resp.Ads {
		ad := &models.Ads{}
		ad.FromApi(apiAd)
		ads = append(ads, ad)
	}

	if err = l.adsRepo.Save(ads); err != nil {
		return err
	}
	return nil
}
