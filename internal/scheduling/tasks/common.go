package tasks

import (
	"github.com/gocraft/work"
	"github.com/micro/go-micro/v2/logger"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	gapi "gitlab.somin.ai/analytics/platform/services/google_api/pb/google_api"
	"runtime/debug"
)

var (
	ErrTokenIsNotFoundOrInvalid = errors.New("token is not found or invalid")
)

type FbErrorType int32

//type IAifier core.Task

type ContextArgs struct {
	AdAccountId string `json:"adAccountId,omitempty" mapstructure:"adAccountId,omitempty"`
	BrandId     int32  `json:"brandId,omitempty" mapstructure:"brandId,omitempty"`

	Completed    bool   `json:"completed,omitempty" mapstructure:"completed,omitempty"`
	ParentTaskId string `json:"parentTaskId,omitempty" mapstructure:"parentTaskId,omitempty"`
	Token        string `json:",omitempty" mapstructure:"Token,omitempty"`
}

type M map[string]interface{}

func ToMap(v interface{}) M {
	m := make(M)
	_ = mapstructure.Decode(v, &m)
	return m
}

func ScanFromMap(m M, v interface{}) {
	_ = mapstructure.Decode(m, v)
}

// obsolete
type ParentArgs struct {
	ParentId string `json:",omitempty" mapstructure:"ParentId,omitempty"`
	Token    string `json:",omitempty" mapstructure:"Token,omitempty"`
}

// baseExecutor is a group of dependencies used for any task that calls google API.
type baseExecutor struct {
	api gapi.GoogleApiService
	log *logger.Helper
}

func percentage(base, max float64, stepsTotal, stepsLeft int) float64 {
	return base + (max-base)*(1-float64(stepsLeft)/float64(stepsTotal))
}

// JobFactory is the way produces *work.Job instances to use as for child tasks.
type JobFactory interface {
	New(taskType core.TaskType, args map[string]interface{}) *work.Job
}

func wrapPanic(p interface{}) error {
	return errors.Errorf("%v\n\n%s", p, debug.Stack())
}

func getAdCampaignById(repo core.AdCampaignsRepository, adCampaignId string) (adCampaign *models.AdCampaign, err error) {
	if adCampaign, err = repo.Read(adCampaignId); err != nil {
		err = errors.Wrap(err, "reading ad campaign failed")
		return
	}

	return
}
