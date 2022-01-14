package core

import (
	"context"
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/models"
	"time"
)

const (
	AppNamespace = "service:google_ads"
)

type Action int

const (
	ActionCancel = 1
)

type AdSetStatus int32

const (
	AdSetDeleted = AdSetStatus(1)
)

type TaskStatus string

const (
	StatusDone       = TaskStatus("done")
	StatusFailed     = TaskStatus("failed")
	StatusInProgress = TaskStatus("in_progress")
	StatusOnPause    = TaskStatus("on_pause")
	StatusWaiting    = TaskStatus("waiting")
)

type TaskArgs struct {
	Type          TaskType
	SchedulerSpec string
	StartAt       *time.Time
	CompositeKey  map[string]interface{}
	Data          map[string]interface{}

	RelatedObjType string
	RelatedObjId   string
	IsDependent    bool
	ParentTaskId   string
}

func (a *TaskArgs) IsRelative() bool {
	return a.RelatedObjType != "" && a.RelatedObjId != ""
}

type TaskResult struct {
	Type        TaskType
	Progress    float64
	Status      TaskStatus
	Description string
	Error       string
	Subtasks    []*TaskResult
	CompletedAt int64
	UpdatedAt   int64
}

type AdCampaignsRepository interface {
	Read(id string) (*models.AdCampaign, error)
	Save(campaigns []*models.AdCampaign) ([]*models.AdCampaign, error)
	AttachToBrand(brandId int32, adCampaignIds []string) error
	DeleteStrategy(adCampaignId string) error
	FindExcluded(adAccountId string, ids []string, before *time.Time) ([]*models.AdCampaign, error)
	MarkDeleted(ids []string) error
	UpdateById(id string, campaign *models.AdCampaign) error
}

type LockedObject string

const (
	LockStrategy = LockedObject("strategy")
)

// TaskLocker doesn't allow related tasks to be executed in the same time on the same object.
// e.g. create Delete task for a strategy before the previous task (Update) has finished.
type TaskLocker interface {
	Lock(obj LockedObject, idArgs map[string]interface{}) bool
	Unlock(obj LockedObject, idArgs map[string]interface{}) bool
}

// TaskScheduler describes the API for jobs execution.
type TaskScheduler interface {
	Execute(params TaskArgs) (string, error)
	RepeatableExecute(args TaskArgs, periodSec int64) (string, error)
	DeleteRepeatedJob(id string) error
	// Cancel is used for cancelling a job that is executing.
	// Note: never been used. May be removed if there are no use cases.
	Cancel(name, id string) error
}

type TaskResulter interface {
	FindTaskIdInProgress(args TaskArgs) (string, error)
	GetResult(encodedId string, withSubtasks bool) (*TaskResult, error)
}

// Run runs the job.
type Task interface {
	Run(actions <-chan Action, job *work.Job) error
}

//
type TaskMetadata interface {
	Metadata(ctx context.Context) interface{}
}
