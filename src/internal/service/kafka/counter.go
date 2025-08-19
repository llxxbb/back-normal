package kafka

import (
	"sync/atomic"
)

var (
	messageCounter = struct {
		sent     int64
		received int64
	}{}
)

// IncrementSentCount 增加发送计数
func IncrementSentCount() {
	atomic.AddInt64(&messageCounter.sent, 1)
}

// IncrementReceivedCount 增加接收计数
func IncrementReceivedCount() {
	atomic.AddInt64(&messageCounter.received, 1)
}

// GetMessageCounts 获取消息计数
func GetMessageCounts() (sent, received int64) {
	return atomic.LoadInt64(&messageCounter.sent), atomic.LoadInt64(&messageCounter.received)
}
