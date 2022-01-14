package scheduling

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/micro/go-micro/v2/logger"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"gitlab.somin.ai/analytics/platform/services/google_ads/pb/google_ads"
	"strings"
)

type taskLocker struct {
	pool *redis.Pool
	log  *logger.Helper
}

func NewTaskLocker(pool *redis.Pool) *taskLocker {
	return &taskLocker{
		pool: pool,
		log: logger.NewHelper(logger.NewLogger().Fields(map[string]interface{}{
			"type": google_ads.ServiceName,
		})),
	}
}

func (l *taskLocker) Lock(obj core.LockedObject, idArgs map[string]interface{}) bool {
	conn := l.pool.Get()
	defer conn.Close()

	key := l.makeKey(obj, idArgs)
	res, err := redis.String(conn.Do("SET", key, 1, "NX", "EX", 60))
	if err != nil {
		l.log.WithError(err)
		return false
	}
	return res == "OK"
}

func (l *taskLocker) Unlock(obj core.LockedObject, idArgs map[string]interface{}) bool {
	conn := l.pool.Get()
	defer conn.Close()

	key := l.makeKey(obj, idArgs)
	_, err := redis.Int(conn.Do("DEL", key))
	if err != nil {
		return false
	}
	return true
}

func (l *taskLocker) makeKey(objType core.LockedObject, args map[string]interface{}) string {
	var argPairs []string
	for k, v := range args {
		argPairs = append(argPairs, fmt.Sprintf("%s=%v", k, v))
	}
	id := strings.Join(argPairs, "&")
	return fmt.Sprintf("google_ads:locks:%s:%s", objType, id)
}
