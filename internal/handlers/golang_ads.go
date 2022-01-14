package handlers

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/scheduling/tasks"
	pb "gitlab.somin.ai/analytics/platform/services/google_ads/pb/google_ads"
	"gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
)

func NewHandler(
	locker core.TaskLocker,
	scheduler core.TaskScheduler,
	tokenRepo *repositories.TokenRepository,
	cl client.Client,
) *tasksHandler {
	return &tasksHandler{
		locker:    locker,
		scheduler: scheduler,
		tokenRepo: tokenRepo,
		cl:        cl,
	}
}

type tasksHandler struct {
	locker    core.TaskLocker
	results   core.TaskResulter
	scheduler core.TaskScheduler
	tokenRepo *repositories.TokenRepository
	cl        client.Client
}

func (h tasksHandler) SyncCampaignCriteria(ctx context.Context, request *pb.SyncCampaignCriteriaRequest, identifier *pb.TaskIdentifier) error {
	token, err := h.tokenRepo.GetToken(request.BrandId)
	if err != nil {
		//h.tokenRepo.Delete(request.BrandId)
		return err
	}

	args := tasks.SyncAdCampaignsArgs{
		ContextArgs: tasks.ContextArgs{
			AdAccountId: request.AdAccountId,
			BrandId:     request.BrandId,
			Token:       token.Token,
		},
	}

	jobId, err := h.scheduler.Execute(core.TaskArgs{
		Type:           core.TaskLoadCampaignsInsights,
		CompositeKey:   map[string]interface{}{"ad_account_id": request.AdAccountId, "brand_id": request.BrandId},
		Data:           args.ToMap(),
		RelatedObjType: "sync_campaigns_criteria",
		RelatedObjId:   "1",
		IsDependent:    false,
	})
	identifier.Id = jobId
	return err
}

func (h tasksHandler) SyncAdGroupCriteria(ctx context.Context, request *pb.SyncAdGroupRequest, identifier *pb.TaskIdentifier) error {
	token, err := h.tokenRepo.GetToken(request.BrandId)
	if err != nil {
		//h.tokenRepo.Delete(request.BrandId)
		return err
	}

	args := tasks.SyncAdCampaignsArgs{
		ContextArgs: tasks.ContextArgs{
			AdAccountId: request.AdAccountId,
			BrandId:     request.BrandId,
			Token:       token.Token,
		},
	}

	jobId, err := h.scheduler.Execute(core.TaskArgs{
		Type:           core.TaskLoadAdGroupsInsights,
		CompositeKey:   map[string]interface{}{"ad_account_id": request.AdAccountId, "brand_id": request.BrandId},
		Data:           args.ToMap(),
		RelatedObjType: "sync_ad_group_criteria",
		RelatedObjId:   "1",
		IsDependent:    false,
	})
	identifier.Id = jobId
	return err
}

var _ pb.GoogleAdsHandler = (*tasksHandler)(nil)

func (h tasksHandler) UpdateTokenForBrandId(ctx context.Context, request *pb.UpdateTokenForBrandIdRequest, response *pb.UpdateTokenForBrandIdResponse) error {
	token, err := h.tokenRepo.GetToken(request.BrandId)
	if err != nil {
		return err
	}
	googleApi := google_api.NewGoogleApiService(google_api.CompleteServiceName, h.cl)
	newToken, err := googleApi.RefreshToken(ctx, &google_api.RefreshTockenRequest{
		RefreshToken: token.RefreshToken,
	})

	if err != nil {
		return err
	}

	return h.tokenRepo.Update(request.BrandId, newToken.Token)
}

func (h tasksHandler) SyncCampaigns(ctx context.Context, request *pb.SyncCampaignsRequest, response *pb.TaskIdentifier) error {
	token, err := h.tokenRepo.GetToken(request.BrandId)
	if err != nil {
		//h.tokenRepo.Delete(request.BrandId)
		return err
	}

	args := tasks.SyncAdCampaignsArgs{
		ContextArgs: tasks.ContextArgs{
			AdAccountId: request.AdAccountId,
			BrandId:     request.BrandId,
			Token:       token.Token,
		},
		IsIncremental: request.IsIcremental,
	}

	jobId, err := h.scheduler.Execute(core.TaskArgs{
		Type:           core.TaskSyncCampaigns,
		CompositeKey:   map[string]interface{}{"ad_account_id": request.AdAccountId, "brand_id": request.BrandId},
		Data:           args.ToMap(),
		RelatedObjType: "sync_campaigns",
		RelatedObjId:   "1",
		IsDependent:    false,
	})
	response.Id = jobId
	return err
}
