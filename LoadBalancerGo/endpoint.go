package LoadBalancerGo

import "sync"

type EndPointInterface interface {

  isAlive(ep *EndPoint) bool
}

type EndPoint struct {
  URL    string       `json:"url"`
  IsDead bool
  mu     sync.RWMutex
}

func (ep *EndPoint) SetAlive(b bool) {
  ep.mu.Lock()
  ep.IsDead = b
  ep.mu.Unlock()
}

func (ep *EndPoint) IsAlive() bool {
  ep.mu.RLock()
  isAlive := !ep.IsDead
  ep.mu.RUnlock()
  return isAlive
}