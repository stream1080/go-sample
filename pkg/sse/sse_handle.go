package sse

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	GlobalEventManager     *EventManager
	activeConnections      int
	activeConnectionsMutex sync.Mutex
)

const maxConnections = 100

func SetupGlobalEventManager() {
	GlobalEventManager = NewEventManager()
}

// HandleSSE 处理服务器发送事件
func HandleSSE(c *gin.Context, fetcher func(SseEventType, *gin.Context) (interface{}, error), eventNames ...SseEventType) {
	// 检查当前活动的连接数
	activeConnectionsMutex.Lock()
	if activeConnections >= maxConnections {
		zap.L().Error("sse connection limit reached")
		c.AbortWithError(http.StatusServiceUnavailable, errors.New("达到最大SSE连接数限制"))
		activeConnectionsMutex.Unlock()
		return
	}
	activeConnections++
	activeConnectionsMutex.Unlock()

	defer func() {
		activeConnectionsMutex.Lock()
		activeConnections--
		activeConnectionsMutex.Unlock()
	}()

	// 设置SSE相关的HTTP头部
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("X-Accel-Buffering", "no")
	c.Writer.Header().Set("Connection", "keep-alive")
	// 确保关闭响应体时能够发送所有缓冲区中的数据
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		zap.L().Error("streaming unsupported")
		return
	}

	// 发送连接成功的提示消息
	c.Writer.WriteString("event: message\ndata: create sse connection\n\n")
	flusher.Flush()

	// 心跳定时器
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	// 创建一个cancel通道来监听客户端断开连接
	cancel := c.Request.Context().Done()

	// 创建一个监听器map，每个事件类型一个监听器
	listeners := make(map[SseEventType][]chan struct{})
	for _, eventType := range eventNames {
		listeners[eventType], _ = GlobalEventManager.GetChannel(eventType)
	}

	// SSE事件循环
	for {
		select {
		case <-cancel:
			zap.L().Error("Connection closed")
			return
		case <-ticker.C:
			// 心跳
			c.Writer.WriteString("event: ping\n")
			flusher.Flush()
		default:
			// 遍历每个事件类型的所有监听器
			for _, eventType := range eventNames {
				chs, ok := GlobalEventManager.GetChannel(eventType)
				if !ok {
					continue
				}

				for _, ch := range chs {
					select {
					case <-ch:
						// 获取数据并写入SSE
						handleEvent(c, fetcher, eventType, flusher)
					case <-cancel:
						return
					default:
					}
				}
			}
		}
	}
}

func handleEvent(c *gin.Context, fetcher func(SseEventType, *gin.Context) (interface{}, error), eventType SseEventType, flusher http.Flusher) {
	data, err := fetcher(eventType, c)
	if err != nil {
		c.Writer.WriteString("event: error\n data: " + err.Error() + "\n\n")
		zap.L().Error("failed to fetch data for event", zap.String("event", string(eventType)), zap.Error(err))
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		zap.L().Error("failed to marshal data", zap.Error(err))
		return
	}

	// 将数据写入所有活跃的客户端
	c.Writer.WriteString("event: " + string(eventType) + "\ndata: " + string(jsonData) + "\n\n")
	flusher.Flush()
}
