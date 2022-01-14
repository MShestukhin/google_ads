package tasks

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewAdGroupsLoader(
	api google_api.GoogleApiService,
	groupRepo *repositories.AdGroupsRepository,
) *AdGroupLoader {
	return &AdGroupLoader{
		googleApi: api,
		groupRepo: groupRepo,
	}
}

type AdGroupLoader struct {
	googleApi google_api.GoogleApiService
	groupRepo *repositories.AdGroupsRepository
}

func (l *AdGroupLoader) Run(_ <-chan core.Action, job *work.Job) error {
	args := LoadArgs{}
	args.FromArgs(job.Args)
	var ctx = context.TODO()
	resp, err := l.googleApi.GetAdGroups(ctx, &google_api.GetAdGroupsRequest{
		CustomerId: args.AdAccountId,
		Token:      args.Token,
	})

	if err != nil {
		return err
	}

	var groups []*models.AdGroup

	for _, c := range resp.AdGroup {
		group := &models.AdGroup{}
		group.FromApi(c)

		groups = append(groups, group)
	}

	_, err = l.groupRepo.Save(groups)
	if err != nil {
		return err
	}
	return nil
}
