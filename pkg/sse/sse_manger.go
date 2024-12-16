package sse

import (
	"sync"

	"go.uber.org/zap"
)

type SseEventType string

var (
	SSEIssueTrendEvent  SseEventType = "event_issue_trend_update"
	SSEScanStatsEvent   SseEventType = "event_scan_stats_update"
	SSEProjectListEvent SseEventType = "event_project_list_update"
	SSEScanTotalEvent   SseEventType = "event_scan_total_update"
	SSEUserInfoEvent    SseEventType = "event_user_info_update"
)

type EventManager struct {
	subscribers map[SseEventType][]chan struct{}
	lock        sync.Mutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		subscribers: make(map[SseEventType][]chan struct{}),
	}
}

// Subscribe 订阅给定事件类型的通知
func (em *EventManager) Subscribe(eventType SseEventType) (chan struct{}, func()) {
	ch := make(chan struct{}, 500)
	em.lock.Lock()
	em.subscribers[eventType] = append(em.subscribers[eventType], ch)
	em.lock.Unlock()

	return ch, func() {
		em.lock.Lock()
		defer em.lock.Unlock()
		for i, c := range em.subscribers[eventType] {
			if c == ch {
				em.subscribers[eventType] = append(em.subscribers[eventType][:i], em.subscribers[eventType][i+1:]...)
				break
			}
		}
		close(ch)
	}
}

// Notify 通知给定事件类型的所有订阅者
func (em *EventManager) Notify(eventType SseEventType) {
	em.lock.Lock()
	defer em.lock.Unlock()

	subscribers, ok := em.subscribers[eventType]
	if !ok {
		return
	}

	for _, ch := range subscribers {
		select {
		case ch <- struct{}{}:
		default:
			zap.L().Warn("SSE event channel is full, dropping event")
		}
	}
}

// GetChannel 返回给定事件类型的订阅通道
func (em *EventManager) GetChannel(eventType SseEventType) ([]chan struct{}, bool) {
	em.lock.Lock()
	defer em.lock.Unlock()

	chs, ok := em.subscribers[eventType]
	return chs, ok
}
