package service

import (
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/pkg/errors"
	"gitlab.somin.ai/analytics/platform/pkg/app"
	"gitlab.somin.ai/analytics/platform/pkg/app/config/postgres/v2"
	"gitlab.somin.ai/analytics/platform/pkg/app/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/handlers"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/redis/sentinel"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/repositories"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/scheduling"
	"gitlab.somin.ai/analytics/platform/services/google_ads/pb/google_ads"
	pb "gitlab.somin.ai/analytics/platform/services/google_ads/pb/google_ads"
	"sync"
	"time"
)

type serviceConfig struct {
	Redis struct {
		MasterName string
		Addrs      []string
	}
}

var (
	_ core.Service    = (*svc)(nil)
	_ core.RpcService = (*svc)(nil)
)

func New(a app.App) *svc {
	return &svc{
		a: a,
		log: logger.NewHelper(a.Logger().Fields(map[string]interface{}{
			"type": "facebook_tasks", "scope": "services",
		})),
	}
}

type svc struct {
	a      app.App
	config serviceConfig
	log    *logger.Helper
}

func (s *svc) Name() string {
	return pb.ServiceName
}

func (s *svc) RegisterRpcService(srv server.Server) error {
	var (
		cfg = s.a.Micro().Options().Config
		cli = s.a.Micro().Client()
	)

	if err := cfg.Get(pb.ServiceName).Scan(&s.config); err != nil {
		return errors.Wrap(err, "reading config failed")
	}

	db, err := postgres.NewDatabase(cfg)
	if err != nil {
		s.log.WithError(err).Error()
		return err
	}

	var (
		pool      = s.newRedisPool(s.config.Redis.Addrs, s.config.Redis.MasterName)
		tokenRepo = repositories.NewTokenRepository(db)
		scheduler = scheduling.NewTaskManager(pool, db, cli)
		locker    = scheduling.NewTaskLocker(pool)
		handler   = handlers.NewHandler(
			locker,
			scheduler,
			tokenRepo,
			cli,
		)
	)

	if err = google_ads.RegisterGoogleAdsHandler(srv, handler); err != nil {
		s.log.WithError(err).Error("failed to register handler")
	}

	return nil
}

func (s *svc) newRedisPool(addrs []string, masterName string) *redis.Pool {
	var dialFn func() (redis.Conn, error)
	var testOnBorrowFn func(c redis.Conn, t time.Time) error

	if masterName == "" {
		idx := -1
		idxMx := sync.Mutex{}

		nextFn := func() string {
			idxMx.Lock()
			defer idxMx.Unlock()

			idx += 1
			if idx >= len(addrs) {
				idx = 0
			}

			return addrs[idx]
		}

		dialFn = func() (redis.Conn, error) {
			return redis.Dial("tcp", nextFn())
		}
	} else {
		// Using sentinel for detect master node
		sntnl := &sentinel.Sentinel{
			Addrs:      addrs,
			MasterName: masterName,
			Dial: func(addr string) (redis.Conn, error) {
				timeout := 500 * time.Millisecond
				c, err := redis.Dial("tcp", addr, redis.DialConnectTimeout(timeout),
					redis.DialReadTimeout(timeout), redis.DialWriteTimeout(timeout))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		}

		dialFn = func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		}
		testOnBorrowFn = func(c redis.Conn, t time.Time) error {
			if !sentinel.TestRole(c, "master") {
				return errors.New("Role check failed")
			} else {
				return nil
			}
		}
	}

	return &redis.Pool{
		MaxIdle:      3,
		MaxActive:    64,
		Wait:         true,
		IdleTimeout:  240 * time.Second,
		Dial:         dialFn,
		TestOnBorrow: testOnBorrowFn,
	}
}
