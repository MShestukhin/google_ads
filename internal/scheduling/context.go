package scheduling

import (
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/client"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/scheduling/tasks"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/scheduling/utils"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"gorm.io/gorm"
)

// WorkContext lists and initializes all kinds of jobs.
type WorkContext struct {
	acts <-chan core.Action
	db   *gorm.DB
	cl   client.Client
	pool *redis.Pool
}

func (c *WorkContext) LoadAdCampaigns(job *work.Job) error {
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	repo := repositories.NewAdCampaignsRepository(c.db)

	j := tasks.NewAdCampaignsLoader(googleApi, repo)

	return j.Run(c.acts, job)
}

func (c *WorkContext) LoadAdGroups(job *work.Job) error {
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	repo := repositories.NewAdGroupsRepository(c.db)
	j := tasks.NewAdGroupsLoader(googleApi, repo)

	return j.Run(c.acts, job)
}

func (c *WorkContext) LoadAdGroupAds(job *work.Job) error {
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	repo := repositories.NewAdsRepository(c.db)
	j := tasks.NewAdsLoader(googleApi, repo)

	return j.Run(c.acts, job)
}

func (c *WorkContext) SyncCampaignsInsight(job *work.Job) error {
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	repo := repositories.NewCampaignInsightRepository(c.db)
	j := tasks.NewCampagnsInsightLoader(googleApi, repo)

	return j.Run(c.acts, job)
}

func (c *WorkContext) SyncAdGroupsInsight(job *work.Job) error {
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	repo := repositories.NewAdSetInsightRepository(c.db)
	j := tasks.NewAdGroupInsightLoader(googleApi, repo)

	return j.Run(c.acts, job)
}

func (c *WorkContext) SyncCampaigns(job *work.Job) error {

	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, c.cl)
	campaignRepo := repositories.NewAdCampaignsRepository(c.db)

	adCampaignLoader := tasks.NewAdCampaignsLoader(googleApi, campaignRepo)

	groupRepo := repositories.NewAdGroupsRepository(c.db)
	adGroupLoader := tasks.NewAdGroupsLoader(googleApi, groupRepo)

	adsRepo := repositories.NewAdsRepository(c.db)
	adsLoader := tasks.NewAdsLoader(googleApi, adsRepo)

	jobFactory := utils.NewJobFactory()

	j := tasks.NewAdCampaignsSyncer(adCampaignLoader, adGroupLoader, adsLoader, jobFactory)

	return j.Run(c.acts, job)
}
