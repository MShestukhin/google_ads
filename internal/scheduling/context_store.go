package scheduling

import (
	"fmt"
	"gitlab.somin.ai/analytics/platform/services/google_ads/internal/core"
	"sync"
)

func newContextStore() *contextStore {
	return &contextStore{actions: make(map[string]chan core.Action)}
}

type contextStore struct {
	mu      sync.RWMutex
	actions map[string]chan core.Action
}

func (c *contextStore) Get(name, id string) chan core.Action {
	c.mu.RLock()
	defer c.mu.RUnlock()
	k := makeKey(name, id)
	a, _ := c.actions[k]
	return a
}

func (c *contextStore) Add(name, id string, acts chan core.Action) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k := makeKey(name, id)
	c.actions[k] = acts
}

func (c *contextStore) Delete(name, id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k := makeKey(name, id)
	delete(c.actions, k)
}

func makeKey(name, id string) string {
	return fmt.Sprintf("%s:%s", name, id)
}
