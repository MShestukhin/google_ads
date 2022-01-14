package tasks

import (
	"github.com/gocraft/work"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
)

type SyncAdCampaignsArgs struct {
	ContextArgs   `mapstructure:",squash"`
	IsIncremental bool `json:"isIncremental,omitempty" mapstructure:"isIncremental,omitempty"`
}

func (a *SyncAdCampaignsArgs) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	_ = mapstructure.Decode(a, &m)
	return m
}

func (a *SyncAdCampaignsArgs) fromArgs(m map[string]interface{}) {
	_ = mapstructure.Decode(m, a)
}

func (a *SyncAdCampaignsArgs) validate() error {
	switch {
	case a.AdAccountId == "":
		return errors.New("ad_account_id is missing")
	case a.BrandId == 0:
		return errors.New("brand_id is missing")
	default:
		return nil
	}
}

func NewAdCampaignsSyncer(
	adCampaignsLoader *AdCampaignsLoader,
	adGroupsLoader *AdGroupLoader,
	adsLoader *AdsLoader,
	jobFactory JobFactory,
) *AdCampaignsSyncer {
	return &AdCampaignsSyncer{
		adCampaignsLoader: adCampaignsLoader,
		adGroupsLoader:    adGroupsLoader,
		adsLoader:         adsLoader,
		jobFactory:        jobFactory,
	}
}

type AdCampaignsSyncer struct {
	adCampaignsLoader *AdCampaignsLoader
	adGroupsLoader    *AdGroupLoader
	adsLoader         *AdsLoader
	jobFactory        JobFactory
}

func (s *AdCampaignsSyncer) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var err error
	if err = s.loadCampaignsIfNeeded(args, job); err != nil {
		return err
	}
	if err = s.loadAdGroupsIfNeeded(args, job); err != nil {
		return err
	}
	if err = s.loadAdsIfNeeded(args, job); err != nil {
		return err
	}

	return nil
}

func (s *AdCampaignsSyncer) loadCampaignsIfNeeded(args LoadArgs, job *work.Job) error {

	campaignArgs := SyncAdCampaignsArgs{
		ContextArgs: ContextArgs{
			AdAccountId: args.AdAccountId,
			BrandId:     args.BrandId,
			Token:       args.Token,
		},
		IsIncremental: true,
	}
	subJob := s.jobFactory.New(core.TaskLoadAdCampaigns, campaignArgs.ToMap())

	if err := s.adCampaignsLoader.Run(nil, subJob); err != nil {
		return err
	}

	return nil
}

func (s *AdCampaignsSyncer) loadAdGroupsIfNeeded(args LoadArgs, job *work.Job) error {
	campaignArgs := SyncAdCampaignsArgs{
		ContextArgs: ContextArgs{
			AdAccountId: args.AdAccountId,
			BrandId:     args.BrandId,
			Token:       args.Token,
		},
		IsIncremental: true,
	}
	subJob := s.jobFactory.New(core.TaskLoadAdGroups, campaignArgs.ToMap())

	if err := s.adGroupsLoader.Run(nil, subJob); err != nil {
		return err
	}

	return nil
}

func (s *AdCampaignsSyncer) loadAdsIfNeeded(args LoadArgs, job *work.Job) error {
	campaignArgs := SyncAdCampaignsArgs{
		ContextArgs: ContextArgs{
			AdAccountId: args.AdAccountId,
			BrandId:     args.BrandId,
			Token:       args.Token,
		},
		IsIncremental: true,
	}
	subJob := s.jobFactory.New(core.TaskLoadAds, campaignArgs.ToMap())

	if err := s.adsLoader.Run(nil, subJob); err != nil {
		return err
	}

	return nil
}
