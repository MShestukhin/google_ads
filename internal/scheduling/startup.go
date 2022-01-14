package scheduling

import (
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/scheduling/utils"
	"gorm.io/gorm"
)

const (
	numOfThreads     = 1    // todo: increase for production
	defaultSecPeriod = 1800 // seconds
)

func NewTaskManager(pool *redis.Pool, db *gorm.DB, cl client.Client) *TaskManager {
	wp := work.NewWorkerPool(WorkContext{}, numOfThreads, core.AppNamespace, pool)
	enqueuer := work.NewEnqueuer(core.AppNamespace, pool)
	cli := work.NewClient(core.AppNamespace, pool)

	ctxStore := newContextStore()

	wp.Middleware(func(c *WorkContext, job *work.Job, next work.NextMiddlewareFunc) error {
		actionsCh := make(chan core.Action, 1)
		ctxStore.Add(job.Name, job.ID, actionsCh)
		c.acts = actionsCh
		c.db = db
		c.cl = cl
		c.pool = pool
		return next()
	})

	defaultOpts := work.JobOptions{MaxConcurrency: 5, MaxFails: 1, SkipDead: true}

	defaultOpts.MaxFails = 3

	wp.JobWithOptions(string(core.TaskSyncCampaigns), defaultOpts, (*WorkContext).SyncCampaignsInsight)
	wp.JobWithOptions(string(core.TaskSyncCampaigns), defaultOpts, (*WorkContext).SyncAdGroupsInsight)

	wp.JobWithOptions(string(core.TaskSyncCampaigns), defaultOpts, (*WorkContext).SyncCampaigns)
	wp.JobWithOptions(string(core.TaskLoadAdCampaigns), defaultOpts, (*WorkContext).LoadAdCampaigns)
	wp.JobWithOptions(string(core.TaskLoadAdGroups), defaultOpts, (*WorkContext).LoadAdGroups)
	wp.JobWithOptions(string(core.TaskLoadAds), defaultOpts, (*WorkContext).LoadAdGroupAds)

	wp.Start()

	return &TaskManager{enqueuer: enqueuer, wp: wp, ctxStore: ctxStore, db: db, cli: cli,
		log: logger.NewHelper(logger.NewLogger().Fields(map[string]interface{}{
			"type": "google_ads", "scope": "manager",
		})),
	}
}

// TaskManager is the component that stores all the schedulers for tasks...
type TaskManager struct {
	enqueuer *work.Enqueuer
	wp       *work.WorkerPool
	progress *work.ProgressClient
	cli      *work.Client

	ctxStore *contextStore
	db       *gorm.DB

	resulter core.TaskResulter
	coder    utils.UUIDCoder
	log      *logger.Helper
}

func (m *TaskManager) Execute(args core.TaskArgs) (string, error) {
	var job *work.Job
	var err error

	if args.IsRelative() {
		relatedObj := work.Relation{
			JobId: args.ParentTaskId,
			Obj: &work.RelatedObject{
				Id:          args.RelatedObjId,
				Type:        args.RelatedObjType,
				IsDependent: args.IsDependent,
			},
		}
		job, err = m.enqueuer.EnqueueRelatedUniqueByKey(string(args.Type), args.Data, args.CompositeKey, relatedObj)
	} else {
		job, err = m.enqueuer.EnqueueUniqueByKey(string(args.Type), args.Data, args.CompositeKey)
	}
	if err != nil {
		return "", err
	}

	return m.jobId(job, args)
}

func (m *TaskManager) jobId(job *work.Job, args core.TaskArgs) (string, error) {
	var id string
	var err error
	if job != nil {
		id = job.ID
	} else {
		id, err = m.resulter.FindTaskIdInProgress(args)
		if err != nil {
			m.log.WithError(err).Error("failed to find a job")
		}
		if err != nil || id == "" {
			return "", fmt.Errorf("failed to enqueue a job")
		}
	}
	encodedId := m.coder.EncodeTypeWithGuid(id, args.Type)
	return encodedId, nil
}

func (m *TaskManager) RepeatableExecute(args core.TaskArgs, periodSec int64) (string, error) {
	if periodSec < 1 {
		periodSec = int64(defaultSecPeriod)
	}

	job, err := m.enqueuer.EnqueueRepeatableUniqueByKey(string(args.Type), periodSec, args.Data, args.CompositeKey)
	if err != nil {
		return "", err
	}

	var id string
	if job != nil {
		id = job.ID
	} else {
		job, err = m.cli.ReadUniqJob(string(args.Type), args.CompositeKey)
		if err != nil {
			m.log.WithError(err).Error("failed to find a job")
		}
		id = job.ID
		if err != nil || id == "" {
			return "", fmt.Errorf("failed to enqueue a job")
		}
	}
	encodedId := m.coder.EncodeTypeWithGuid(id, args.Type)
	return encodedId, nil
}

func (m *TaskManager) DeleteRepeatedJob(id string) error {
	jobId, _, err := m.coder.DecodeTypeFromId(id)
	if err != nil {
		return err
	}
	return m.cli.DeleteRepeatedJob(jobId)
}

func (m *TaskManager) Cancel(name, id string) error {
	return m.sendAction(name, id, core.ActionCancel)
}

func (m *TaskManager) sendAction(name, id string, action core.Action) error {
	acts := m.ctxStore.Get(name, id)
	if acts == nil {
		return nil
	}
	select {
	case acts <- action:
	default:
		return fmt.Errorf("previous message hasn't been processed")
	}
	return nil
}
