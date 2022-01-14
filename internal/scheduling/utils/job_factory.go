package utils

import (
	"github.com/gocraft/work"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
)

type IdProvider interface {
	MakeIdentifier() string
}

type testIdProvider struct{}

func (testIdProvider) MakeIdentifier() string {
	return "id"
}

func NewJobFactory() *JobFactory {
	return &JobFactory{clock: systemClock{}, idProvider: UUIDCoder{}}
}

func NewTestJobFactory() *JobFactory {
	return &JobFactory{clock: testClock{}, idProvider: testIdProvider{}}
}

type JobFactory struct {
	clock      Clock
	idProvider IdProvider
}

func (f *JobFactory) New(taskType core.TaskType, args map[string]interface{}) *work.Job {
	return &work.Job{
		ID:         f.idProvider.MakeIdentifier(),
		Name:       string(taskType),
		EnqueuedAt: f.clock.Now().Unix(),
		Args:       args,
	}
}
